package grpcv1

import (
	"fmt"
	"image"
	"path/filepath"
	"strings"
	"testing"

	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
)

func TestBuildFindBenchFixtureFromRegionSpecUsesAllTargets(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	manifestPath := filepath.Join(root, "docs", "bench", "find-on-screen-scenarios.example.json")
	t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", manifestPath)

	scenarios, _, err := loadFindBenchScenariosFromManifest(true, false)
	if err != nil {
		t.Fatalf("load manifest scenarios: %v", err)
	}
	if len(scenarios) == 0 {
		t.Fatalf("expected at least one scenario from manifest")
	}

	source, queries, ok := buildFindBenchFixtureFromRegionSpec(t, scenarios[0])
	if !ok {
		t.Fatalf("expected region-spec fixture to be available")
	}
	if source == nil {
		t.Fatalf("expected source image")
	}
	if got, wantMin := len(queries), 3; got < wantMin {
		t.Fatalf("expected at least %d region queries, got %d", wantMin, got)
	}
	seenExpected := map[string]struct{}{}
	for i, query := range queries {
		if query.Pattern == nil {
			t.Fatalf("query[%d] missing pattern", i)
		}
		if query.Expected == nil || query.Expected.GetW() <= 0 || query.Expected.GetH() <= 0 {
			t.Fatalf("query[%d] missing expected rect", i)
		}
		key := rectKey(query.Expected)
		seenExpected[key] = struct{}{}
	}
	if got, wantMin := len(seenExpected), 3; got < wantMin {
		t.Fatalf("expected at least %d unique expected rects, got %d", wantMin, got)
	}
}

func TestProjectTargetRegionsToScreen_ScalesExpectedForResolution(t *testing.T) {
	in := []findBenchTargetRegion{
		{ID: "target-01", Label: "primary", X: 1508, Y: 900, W: 460, H: 278},
	}
	got := projectTargetRegionsToScreen(in, 2560, 1440, 1280, 720)
	if len(got) != 1 {
		t.Fatalf("expected one projected target, got %d", len(got))
	}
	if got[0].X != 754 || got[0].Y != 450 || got[0].W != 230 || got[0].H != 139 {
		t.Fatalf(
			"unexpected projected target got=(x=%d y=%d w=%d h=%d) want=(x=%d y=%d w=%d h=%d)",
			got[0].X, got[0].Y, got[0].W, got[0].H,
			754, 450, 230, 139,
		)
	}
}

func TestBuildFindBenchFixtureFromRegionSpecExpectedWithinScenarioBounds(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	manifestPath := filepath.Join(root, "docs", "bench", "find-on-screen-scenarios.example.json")
	t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", manifestPath)

	scenarios, _, err := loadFindBenchScenariosFromManifest(true, true)
	if err != nil {
		t.Fatalf("load manifest scenarios: %v", err)
	}
	if len(scenarios) == 0 {
		t.Fatalf("expected at least one scenario from manifest")
	}

	var selected *findBenchScenario
	for i := range scenarios {
		s := scenarios[i]
		if strings.TrimSpace(s.scenarioTypeID) == "hybrid_gate_conflicts" && s.screenW == 1280 && s.screenH == 720 {
			selected = &s
			break
		}
	}
	if selected == nil {
		t.Fatalf("missing hybrid_gate_conflicts 1280x720 scenario in manifest materialization")
	}

	_, queries, ok := buildFindBenchFixtureFromRegionSpec(t, *selected)
	if !ok {
		t.Fatalf("expected region-spec fixture to be available")
	}
	if len(queries) == 0 {
		t.Fatalf("expected at least one query")
	}

	for i, q := range queries {
		if q.Expected == nil {
			t.Fatalf("query[%d] missing expected rect", i)
		}
		if q.Expected.GetX() < 0 || q.Expected.GetY() < 0 {
			t.Fatalf("query[%d] has negative expected origin: %+v", i, q.Expected)
		}
		if q.Expected.GetW() <= 0 || q.Expected.GetH() <= 0 {
			t.Fatalf("query[%d] has non-positive expected size: %+v", i, q.Expected)
		}
		if q.Expected.GetX()+q.Expected.GetW() > int32(selected.screenW) || q.Expected.GetY()+q.Expected.GetH() > int32(selected.screenH) {
			t.Fatalf(
				"query[%d] expected rect exceeds scenario bounds rect=%+v bounds=%dx%d",
				i, q.Expected, selected.screenW, selected.screenH,
			)
		}

		// For downscaled scenarios, expected benchmark regions must be downscaled too.
		if selected.screenW < 2560 || selected.screenH < 1440 {
			matchRaw := findBenchTargetRegion{}
			foundRaw := false
			for _, rawTarget := range selected.benchmarkTargets {
				if strings.EqualFold(strings.TrimSpace(rawTarget.Label), strings.TrimSpace(q.Label)) {
					matchRaw = rawTarget
					foundRaw = true
					break
				}
			}
			if foundRaw && matchRaw.W > 0 && matchRaw.H > 0 {
				if q.Expected.GetW() >= int32(matchRaw.W) || q.Expected.GetH() >= int32(matchRaw.H) {
					t.Fatalf(
						"query[%d] expected rect was not downscaled rect=%+v raw_target=%+v",
						i, q.Expected, matchRaw,
					)
				}
			}
		}
	}
}

func TestBuildFindBenchFixtureFromRegionSpec_PrefersSourceExpectedForNonMultiMonitor(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	manifestPath := filepath.Join(root, "docs", "bench", "find-on-screen-scenarios.example.json")
	t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", manifestPath)

	scenarios, _, err := loadFindBenchScenariosFromManifest(true, true)
	if err != nil {
		t.Fatalf("load manifest scenarios: %v", err)
	}

	var selected *findBenchScenario
	for i := range scenarios {
		s := scenarios[i]
		if strings.TrimSpace(s.scenarioTypeID) == "vector_ui_baseline" && s.screenW == 1280 && s.screenH == 720 {
			selected = &s
			break
		}
	}
	if selected == nil {
		t.Fatalf("missing vector_ui_baseline 1280x720 scenario in manifest materialization")
	}

	_, queries, ok := buildFindBenchFixtureFromRegionSpec(t, *selected)
	if !ok {
		t.Fatalf("expected region-spec fixture to be available")
	}
	if len(queries) == 0 {
		t.Fatalf("expected at least one query")
	}

	var tertiary *findBenchFixtureQuery
	for i := range queries {
		if strings.EqualFold(strings.TrimSpace(queries[i].Label), "tertiary") {
			tertiary = &queries[i]
			break
		}
	}
	if tertiary == nil {
		t.Fatalf("missing tertiary query")
	}
	if tertiary.Expected == nil {
		t.Fatalf("tertiary expected is nil")
	}
	// Source target for vector_ui_baseline tertiary projects near y=564 at 1280x720.
	// Benchmark target projects near y=278, which should be treated as alternate.
	if tertiary.Expected.GetY() < 500 {
		t.Fatalf("expected source-aligned tertiary region, got y=%d", tertiary.Expected.GetY())
	}
	if len(tertiary.ExpectedAlternates) == 0 {
		t.Fatalf("expected alternate benchmark region for tertiary")
	}
}

func TestBuildFindBenchFixtureFromRegionSpec_PrefersPreciseBenchmarkExpectedForMultiMonitor(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	manifestPath := filepath.Join(root, "docs", "bench", "find-on-screen-scenarios.example.json")
	t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", manifestPath)

	scenarios, _, err := loadFindBenchScenariosFromManifest(true, true)
	if err != nil {
		t.Fatalf("load manifest scenarios: %v", err)
	}

	var selected *findBenchScenario
	for i := range scenarios {
		s := scenarios[i]
		if strings.TrimSpace(s.scenarioTypeID) == "multi_monitor_dpi_shift" && s.screenW == 1280 && s.screenH == 720 {
			selected = &s
			break
		}
	}
	if selected == nil {
		t.Fatalf("missing multi_monitor_dpi_shift 1280x720 scenario in manifest materialization")
	}

	_, queries, ok := buildFindBenchFixtureFromRegionSpec(t, *selected)
	if !ok {
		t.Fatalf("expected region-spec fixture to be available")
	}
	if len(queries) == 0 {
		t.Fatalf("expected at least one query")
	}

	var primary *findBenchFixtureQuery
	for i := range queries {
		if strings.EqualFold(strings.TrimSpace(queries[i].Label), "primary") {
			primary = &queries[i]
			break
		}
	}
	if primary == nil {
		t.Fatalf("missing primary query")
	}
	if primary.Expected == nil {
		t.Fatalf("primary expected is nil")
	}

	sourceRaw, err := loadGrayFromFile(selected.sourceImagePath)
	if err != nil {
		t.Fatalf("load source image: %v", err)
	}
	broadProjected := projectTargetRegionsToScreen(
		selected.benchmarkTargets,
		sourceRaw.Bounds().Dx(),
		sourceRaw.Bounds().Dy(),
		selected.screenW,
		selected.screenH,
	)
	broadPrimary, ok := findScenarioTargetRegionByID(broadProjected, "target-01")
	if !ok {
		t.Fatalf("missing projected broad benchmark primary region")
	}
	candidates := findBenchPreciseMultiMonitorBenchmarkCandidates(*selected, sourceRaw.Bounds().Dx(), sourceRaw.Bounds().Dy())
	precisePrimary, ok := choosePreciseMultiMonitorExpectedRegion(
		findBenchTargetRegion{ID: "target-01", Label: "primary"},
		broadPrimary,
		candidates,
	)
	if !ok {
		t.Fatalf("missing precise benchmark candidate for primary")
	}

	if got, want := rectKey(primary.Expected), rectKey(benchRectFromTargetRegion(precisePrimary)); got != want {
		t.Fatalf("expected primary rect to prefer precise benchmark target, got=%s want=%s", got, want)
	}
	if len(primary.ExpectedAlternates) != 0 {
		t.Fatalf("expected no alternates for multi-monitor primary, got=%d", len(primary.ExpectedAlternates))
	}
}

func TestMultiMonitorRegionSpecBenchmarkTargetsMatchComposedGeometry(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	specPath := filepath.Join(root, "packages", "api", "internal", "grpcv1", "testdata", "find-bench-assets", "regions.json")
	doc, resolvedSpecPath, err := loadFindBenchRegionSpec(specPath, filepath.Dir(specPath))
	if err != nil {
		t.Fatalf("load region spec: %v", err)
	}
	assignments, err := buildFindBenchRegionAssignments(doc, filepath.Dir(resolvedSpecPath))
	if err != nil {
		t.Fatalf("build region assignments: %v", err)
	}
	selected, ok := assignments["multi_monitor_dpi_shift"]
	if !ok {
		t.Fatalf("missing multi_monitor_dpi_shift assignment")
	}

	gotByID := map[string]findBenchTargetRegion{}
	for _, region := range selected.BenchmarkTargets {
		gotByID[strings.TrimSpace(region.ID)] = region
	}

	benchmarkRaw, err := loadGrayFromFile(selected.BenchmarkImagePath)
	if err != nil {
		t.Fatalf("load benchmark image: %v", err)
	}
	benchmarkNativeW := benchmarkRaw.Bounds().Dx()
	benchmarkNativeH := benchmarkRaw.Bounds().Dy()
	panelW := int(float64(benchmarkNativeW) * 0.68)
	panelH := int(float64(benchmarkNativeH) * 0.86)
	panelX := (benchmarkNativeW - panelW) / 2
	panelY := (benchmarkNativeH - panelH) / 2
	gap := maxInt(8, benchmarkNativeW/120)
	leftW := (benchmarkNativeW - gap) / 2
	rightW := benchmarkNativeW - gap - leftW
	leftRect := image.Rect(0, 0, leftW, benchmarkNativeH)
	rightRect := image.Rect(leftW+gap, 0, leftW+gap+rightW, benchmarkNativeH)

	leftMapped := mapRegionsWithPanelGeometryBench(selected.Targets, panelX, panelY, panelW, panelH, leftRect, 0.98)
	rightMapped := mapRegionsWithPanelGeometryBench(selected.Targets, panelX, panelY, panelW, panelH, rightRect, 1.22)

	want := map[string]findBenchTargetRegion{
		"target-01": leftMapped[0],
		"target-02": rightMapped[1],
		"target-03": rightMapped[2],
	}
	for id, expected := range want {
		got, ok := gotByID[id]
		if !ok {
			t.Fatalf("missing benchmark target %s", id)
		}
		if got != expected {
			t.Fatalf("benchmark target %s mismatch got=%+v want=%+v", id, got, expected)
		}
	}
}

func TestBuildFindBenchFixtureFromRegionSpec_UsesTransformedExpectedForScaleRotateSweep(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	manifestPath := filepath.Join(root, "docs", "bench", "find-on-screen-scenarios.example.json")
	t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", manifestPath)

	scenarios, _, err := loadFindBenchScenariosFromManifest(true, true)
	if err != nil {
		t.Fatalf("load manifest scenarios: %v", err)
	}

	var selected *findBenchScenario
	for i := range scenarios {
		s := scenarios[i]
		if strings.TrimSpace(s.scenarioTypeID) == "scale_rotate_sweep" && s.screenW == 1280 && s.screenH == 720 {
			selected = &s
			break
		}
	}
	if selected == nil {
		t.Fatalf("missing scale_rotate_sweep 1280x720 scenario in manifest materialization")
	}

	_, queries, ok := buildFindBenchFixtureFromRegionSpec(t, *selected)
	if !ok {
		t.Fatalf("expected region-spec fixture to be available")
	}
	if len(queries) == 0 {
		t.Fatalf("expected at least one query")
	}

	var tertiary *findBenchFixtureQuery
	for i := range queries {
		if strings.EqualFold(strings.TrimSpace(queries[i].Label), "tertiary") {
			tertiary = &queries[i]
			break
		}
	}
	if tertiary == nil {
		t.Fatalf("missing tertiary query")
	}
	if tertiary.Expected == nil {
		t.Fatalf("tertiary expected is nil")
	}
	if len(tertiary.ExpectedAlternates) == 0 {
		t.Fatalf("expected alternate source region for scale_rotate_sweep tertiary")
	}

	sourceRaw, err := loadGrayFromFile(selected.sourceImagePath)
	if err != nil {
		t.Fatalf("load source image: %v", err)
	}
	sourceScene, sourceProjected := normalizeSceneAndRegions(sourceRaw, selected.screenW, selected.screenH, selected.sourceTargets)
	if sourceScene == nil || len(sourceProjected) == 0 {
		t.Fatalf("normalize source scene failed")
	}
	if projected := projectTargetRegionsToScreen(
		selected.sourceTargets,
		sourceRaw.Bounds().Dx(),
		sourceRaw.Bounds().Dy(),
		selected.screenW,
		selected.screenH,
	); len(projected) > 0 {
		sourceProjected = projected
	}
	sourceRegion, ok := findScenarioTargetRegionByID(sourceProjected, "target-03")
	if !ok {
		t.Fatalf("missing projected source tertiary region")
	}
	derivedRegion, ok := deriveBenchmarkExpectedRegionFromTransform(sourceScene, sourceRegion, *selected)
	if !ok {
		t.Fatalf("missing transformed tertiary region")
	}

	if got, want := rectKey(tertiary.Expected), rectKey(benchRectFromTargetRegion(derivedRegion)); got != want {
		t.Fatalf("expected tertiary primary rect to prefer transformed benchmark target, got=%s want=%s", got, want)
	}
	if got, want := rectKey(tertiary.ExpectedAlternates[0]), rectKey(benchRectFromTargetRegion(sourceRegion)); got != want {
		t.Fatalf("expected tertiary alternate rect to preserve source target, got=%s want=%s", got, want)
	}
}

func TestBuildFindBenchFixtureFromRegionSpec_UsesTransformedExpectedForOrbFeatureRich(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	manifestPath := filepath.Join(root, "docs", "bench", "find-on-screen-scenarios.example.json")
	t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", manifestPath)

	scenarios, _, err := loadFindBenchScenariosFromManifest(true, true)
	if err != nil {
		t.Fatalf("load manifest scenarios: %v", err)
	}

	var selected *findBenchScenario
	for i := range scenarios {
		s := scenarios[i]
		if strings.TrimSpace(s.scenarioTypeID) == "orb_feature_rich" && s.screenW == 1280 && s.screenH == 720 {
			selected = &s
			break
		}
	}
	if selected == nil {
		t.Fatalf("missing orb_feature_rich 1280x720 scenario in manifest materialization")
	}

	_, queries, ok := buildFindBenchFixtureFromRegionSpec(t, *selected)
	if !ok {
		t.Fatalf("expected region-spec fixture to be available")
	}
	if len(queries) == 0 {
		t.Fatalf("expected at least one query")
	}

	var primary *findBenchFixtureQuery
	for i := range queries {
		if strings.EqualFold(strings.TrimSpace(queries[i].Label), "primary") {
			primary = &queries[i]
			break
		}
	}
	if primary == nil {
		t.Fatalf("missing primary query")
	}
	if primary.Expected == nil {
		t.Fatalf("primary expected is nil")
	}

	sourceRaw, err := loadGrayFromFile(selected.sourceImagePath)
	if err != nil {
		t.Fatalf("load source image: %v", err)
	}
	sourceScene, sourceProjected := normalizeSceneAndRegions(sourceRaw, selected.screenW, selected.screenH, selected.sourceTargets)
	if sourceScene == nil || len(sourceProjected) == 0 {
		t.Fatalf("normalize source scene failed")
	}
	if projected := projectTargetRegionsToScreen(
		selected.sourceTargets,
		sourceRaw.Bounds().Dx(),
		sourceRaw.Bounds().Dy(),
		selected.screenW,
		selected.screenH,
	); len(projected) > 0 {
		sourceProjected = projected
	}
	sourceRegion, ok := findScenarioTargetRegionByID(sourceProjected, "target-01")
	if !ok {
		t.Fatalf("missing projected source primary region")
	}
	derivedRegion, ok := deriveBenchmarkExpectedRegionFromTransform(sourceScene, sourceRegion, *selected)
	if !ok {
		t.Fatalf("missing transformed primary region")
	}
	if got, want := rectKey(primary.Expected), rectKey(benchRectFromTargetRegion(derivedRegion)); got != want {
		t.Fatalf("expected primary rect to prefer transformed benchmark target, got=%s want=%s", got, want)
	}
}

func TestBuildFindBenchFixtureFromRegionSpec_SkipsQueriesWithoutBenchmarkRegionMatch(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	manifestPath := filepath.Join(root, "docs", "bench", "find-on-screen-scenarios.example.json")
	t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", manifestPath)

	scenarios, _, err := loadFindBenchScenariosFromManifest(true, true)
	if err != nil {
		t.Fatalf("load manifest scenarios: %v", err)
	}

	var selected *findBenchScenario
	for i := range scenarios {
		s := scenarios[i]
		if strings.TrimSpace(s.scenarioTypeID) == "vector_ui_baseline" && s.screenW == 1280 && s.screenH == 720 {
			selected = &s
			break
		}
	}
	if selected == nil {
		t.Fatalf("missing vector_ui_baseline 1280x720 scenario in manifest materialization")
	}

	mutated := *selected
	mutated.benchmarkTargets = append([]findBenchTargetRegion(nil), selected.benchmarkTargets...)
	if len(mutated.benchmarkTargets) < 3 {
		t.Fatalf("expected at least three benchmark targets")
	}
	mutated.benchmarkTargets[0].ID = "non-matching-id"
	mutated.benchmarkTargets[0].Label = "non-matching-label"

	_, queries, ok := buildFindBenchFixtureFromRegionSpec(t, mutated)
	if !ok {
		t.Fatalf("expected region-spec fixture to be available")
	}
	if got, want := len(queries), len(selected.sourceTargets)-1; got != want {
		t.Fatalf("expected one query to be skipped when benchmark target id is missing, got=%d want=%d", got, want)
	}
	for _, query := range queries {
		if strings.TrimSpace(query.ID) == "target-01" {
			t.Fatalf("expected target-01 query to be skipped when benchmark target mapping is missing")
		}
	}
}

func TestClassifyFindBenchPositiveMatch_RegionAware(t *testing.T) {
	primary := &pb.Rect{X: 100, Y: 120, W: 80, H: 60}
	secondary := &pb.Rect{X: 340, Y: 120, W: 80, H: 60}
	all := []*pb.Rect{primary, secondary}
	pattern := &pb.GrayImage{Width: 80, Height: 60}

	if got := classifyFindBenchPositiveMatch(
		&pb.Rect{X: 102, Y: 122, W: 80, H: 60},
		primary,
		nil,
		all,
		pattern,
		0.6,
		1.5,
		false,
	); got != findBenchMatchClassOK {
		t.Fatalf("expected ok classification, got %q", got)
	}

	if got := classifyFindBenchPositiveMatch(
		&pb.Rect{X: 340, Y: 120, W: 80, H: 60},
		primary,
		nil,
		all,
		pattern,
		0.6,
		1.5,
		false,
	); got != findBenchMatchClassWrongRegion {
		t.Fatalf("expected wrong_region classification, got %q", got)
	}

	if got := classifyFindBenchPositiveMatch(
		&pb.Rect{X: 10, Y: 10, W: 80, H: 60},
		primary,
		nil,
		all,
		pattern,
		0.6,
		1.5,
		false,
	); got != findBenchMatchClassOverlapMiss {
		t.Fatalf("expected overlap_miss classification, got %q", got)
	}
}

func TestClassifyFindBenchPositiveMatch_ZoneExpected(t *testing.T) {
	zone := &pb.Rect{X: 500, Y: 320, W: 460, H: 278}
	peer := &pb.Rect{X: 1000, Y: 320, W: 460, H: 278}
	all := []*pb.Rect{zone, peer}
	pattern := &pb.GrayImage{Width: 160, Height: 98}

	if got := classifyFindBenchPositiveMatch(
		&pb.Rect{X: 602, Y: 410, W: 170, H: 110},
		zone,
		nil,
		all,
		pattern,
		0.22,
		2.0,
		false,
	); got != findBenchMatchClassOK {
		t.Fatalf("expected zone match to classify as ok, got %q", got)
	}

	if got := classifyFindBenchPositiveMatch(
		&pb.Rect{X: 560, Y: 360, W: 520, H: 350},
		zone,
		nil,
		all,
		pattern,
		0.22,
		2.0,
		false,
	); got != findBenchMatchClassOverlapMiss {
		t.Fatalf("expected oversized in-zone match to classify as overlap_miss, got %q", got)
	}

	if got := classifyFindBenchPositiveMatch(
		&pb.Rect{X: 120, Y: 80, W: 170, H: 110},
		zone,
		nil,
		all,
		pattern,
		0.22,
		2.0,
		false,
	); got != findBenchMatchClassOverlapMiss {
		t.Fatalf("expected center-outside-zone match to classify as overlap_miss, got %q", got)
	}
}

func TestClassifyFindBenchPositiveMatch_AlternateExpected(t *testing.T) {
	primary := &pb.Rect{X: 700, Y: 300, W: 180, H: 110}
	alt := &pb.Rect{X: 340, Y: 280, W: 180, H: 110}
	peer := &pb.Rect{X: 980, Y: 280, W: 180, H: 110}
	all := []*pb.Rect{primary, alt, peer}
	pattern := &pb.GrayImage{Width: 180, Height: 110}

	if got := classifyFindBenchPositiveMatch(
		&pb.Rect{X: 338, Y: 282, W: 180, H: 110},
		primary,
		[]*pb.Rect{alt},
		all,
		pattern,
		0.22,
		2.0,
		false,
	); got != findBenchMatchClassOK {
		t.Fatalf("expected alternate region match to classify as ok, got %q", got)
	}
}

func TestCenterDeltaMetrics_UsesRadialDistanceTolerance(t *testing.T) {
	expected := &pb.Rect{X: 100, Y: 100, W: 100, H: 100}
	pattern := &pb.GrayImage{Width: 100, Height: 100}

	dx, dy, dist, limX, limY, ok := centerDeltaMetrics(
		&pb.Rect{X: 135, Y: 135, W: 100, H: 100},
		expected,
		pattern,
		0.22,
		1.5,
		false,
	)
	if ok {
		t.Fatalf("expected diagonal center offset to fail radial tolerance, got ok dx=%.2f dy=%.2f dist=%.2f lim=%.2f,%.2f", dx, dy, dist, limX, limY)
	}
	if dx <= 0 || dy <= 0 || dist <= 0 {
		t.Fatalf("expected positive center deltas, got dx=%.2f dy=%.2f dist=%.2f", dx, dy, dist)
	}
	if limX != limY {
		t.Fatalf("expected radial tolerance to use symmetric limit, got lim=%.2f,%.2f", limX, limY)
	}
}

func rectKey(r *pb.Rect) string {
	if r == nil {
		return "nil"
	}
	return fmt.Sprintf("%d:%d:%d:%d", r.GetX(), r.GetY(), r.GetW(), r.GetH())
}
