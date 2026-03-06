package grpcv1

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/gen2brain/avif"
	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"github.com/smysnk/sikuligo/pkg/sikuli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type findBenchEngine struct {
	name string
	enum pb.MatcherEngine
}

type findBenchScenario struct {
	name               string
	kind               string
	variant            string
	scenarioTypeID     string
	targetSource       string
	targetAsset        string
	sourceImagePath    string
	benchmarkImagePath string
	sourceTargets      []findBenchTargetRegion
	benchmarkTargets   []findBenchTargetRegion
	size               int
	rotation           int
	screenW            int
	screenH            int
	tolerance          float64
	maxAreaRatio       float64
	transformKind      string
	transformA         float64
	transformB         float64
	transformC         float64
	queryFromBase      bool
	seed               uint64

	// Expected matching behavior.
	expectedPositive bool
	allowPartial     bool
	strictBBox       bool

	// Runtime knobs from manifest defaults.
	retryAttempts int
	rpcTimeout    time.Duration

	// Background controls.
	backgroundContinuousCanvas bool
	backgroundClutterDensity   float64
	backgroundPalette          string
	backgroundTextureSeed      int

	// Photometric perturbations.
	brightnessDelta  float64
	contrastFactor   float64
	gammaFactor      float64
	blurSigma        float64
	jpegQuality      int
	noiseGaussian    float64
	noisePoisson     float64
	noiseSaltPepper  float64
	noiseBanding     float64
	noiseCompression float64

	// Occlusion controls.
	occlusionEnabled  bool
	occlusionCoverage float64

	// Decoy controls.
	decoyEnabled    bool
	decoyCount      int
	decoySimilarity float64
	decoyPlacement  string

	// Monitor profile simulation.
	monitorID         string
	monitorMode       string
	monitorGamma      float64
	monitorSharpness  float64
	monitorColorShift int

	// Hybrid policy metadata (influences scenario transform planning).
	hybridMustConsiderAll bool
	hybridSelectBy        string
	hybridFallbackOrder   []string
}

type findBenchTargetRegion struct {
	ID    string
	Label string
	X     int
	Y     int
	W     int
	H     int
}

type findBenchFixtureQuery struct {
	Pattern            *pb.GrayImage
	Expected           *pb.Rect
	ExpectedAlternates []*pb.Rect
	Label              string
}

type findBenchMatchClass string

const (
	findBenchMatchClassOK          findBenchMatchClass = "ok"
	findBenchMatchClassWrongRegion findBenchMatchClass = "wrong_region"
	findBenchMatchClassOverlapMiss findBenchMatchClass = "overlap_miss"
)

type findBenchResolutionPack struct {
	screenW      int
	screenH      int
	baseSize     int
	tolerance    float64
	maxAreaRatio float64
	rotateDeg    float64
	resizeDown   float64
	resizeUp     float64
}

func BenchmarkFindOnScreenE2E(b *testing.B) {
	b.ReportAllocs()
	visuals := newFindBenchVisualCollectorFromEnv(b)

	engines := []findBenchEngine{
		{name: "template", enum: pb.MatcherEngine_MATCHER_ENGINE_TEMPLATE},
		{name: "orb", enum: pb.MatcherEngine_MATCHER_ENGINE_ORB},
		{name: "akaze", enum: pb.MatcherEngine_MATCHER_ENGINE_AKAZE},
		{name: "brisk", enum: pb.MatcherEngine_MATCHER_ENGINE_BRISK},
		{name: "kaze", enum: pb.MatcherEngine_MATCHER_ENGINE_KAZE},
		{name: "sift", enum: pb.MatcherEngine_MATCHER_ENGINE_SIFT},
		{name: "hybrid", enum: pb.MatcherEngine_MATCHER_ENGINE_HYBRID},
	}

	highResEnabled := true
	ultraResEnabled := true
	scenarios := []findBenchScenario{}
	if strings.TrimSpace(os.Getenv("FIND_BENCH_SCENARIO_MANIFEST")) != "" {
		loaded, sourcePath, err := loadFindBenchScenariosFromManifest(highResEnabled, ultraResEnabled)
		if err != nil {
			b.Fatalf("load benchmark scenarios from manifest: %v", err)
		}
		scenarios = loaded
		b.Logf(
			"find bench scenario config count=%d source=manifest path=%s high_res=%v ultra_res=%v",
			len(scenarios),
			sourcePath,
			highResEnabled,
			ultraResEnabled,
		)
	} else {
		scenarios = defaultFindBenchScenarios(highResEnabled, ultraResEnabled)
		b.Logf(
			"find bench scenario config count=%d source=default high_res=%v ultra_res=%v",
			len(scenarios),
			highResEnabled,
			ultraResEnabled,
		)
	}
	if len(scenarios) == 0 {
		b.Fatalf("no benchmark scenarios configured")
	}

	for _, engine := range engines {
		engine := engine
		b.Run("engine="+engine.name, func(b *testing.B) {
			for _, scenario := range scenarios {
				scenario := scenario
				b.Run(scenario.name, func(b *testing.B) {
					runFindOnScreenScenarioBenchmark(b, engine, scenario, visuals)
				})
			}
		})
	}

	if visuals != nil {
		if err := visuals.WriteScenarioSummaries(); err != nil {
			b.Fatalf("write benchmark visual summaries: %v", err)
		}
	}
}

func defaultFindBenchScenarios(highResEnabled, ultraResEnabled bool) []findBenchScenario {
	scenarios := []findBenchScenario{}
	basePacks := []findBenchResolutionPack{
		{screenW: 480, screenH: 270, baseSize: 32, tolerance: 0.30, maxAreaRatio: 1.50, rotateDeg: 8, resizeDown: 0.78, resizeUp: 1.25},
		{screenW: 640, screenH: 360, baseSize: 48, tolerance: 0.30, maxAreaRatio: 1.50, rotateDeg: 10, resizeDown: 0.76, resizeUp: 1.28},
		{screenW: 800, screenH: 450, baseSize: 64, tolerance: 0.28, maxAreaRatio: 1.60, rotateDeg: 12, resizeDown: 0.74, resizeUp: 1.30},
		{screenW: 960, screenH: 540, baseSize: 80, tolerance: 0.26, maxAreaRatio: 1.70, rotateDeg: 14, resizeDown: 0.72, resizeUp: 1.35},
		{screenW: 1024, screenH: 576, baseSize: 96, tolerance: 0.24, maxAreaRatio: 1.80, rotateDeg: 16, resizeDown: 0.70, resizeUp: 1.40},
	}
	for _, pack := range basePacks {
		scenarios = append(scenarios, resolutionScenarioPack(pack)...)
	}
	if highResEnabled {
		highResPacks := []findBenchResolutionPack{
			{screenW: 1280, screenH: 720, baseSize: 112, tolerance: 0.22, maxAreaRatio: 2.00, rotateDeg: 18, resizeDown: 0.68, resizeUp: 1.45},
			{screenW: 1920, screenH: 1080, baseSize: 144, tolerance: 0.18, maxAreaRatio: 2.80, rotateDeg: 22, resizeDown: 0.64, resizeUp: 1.55},
		}
		for _, pack := range highResPacks {
			scenarios = append(scenarios, resolutionScenarioPack(pack)...)
		}
	}
	if ultraResEnabled {
		scenarios = append(
			scenarios,
			resolutionScenarioPack(
				findBenchResolutionPack{
					screenW:      2560,
					screenH:      1440,
					baseSize:     176,
					tolerance:    0.16,
					maxAreaRatio: 3.00,
					rotateDeg:    26,
					resizeDown:   0.60,
					resizeUp:     1.65,
				},
			)...,
		)
	}
	return scenarios
}

func resolutionScenarioPack(cfg findBenchResolutionPack) []findBenchScenario {
	res := fmt.Sprintf("%dx%d", cfg.screenW, cfg.screenH)
	styleTol := maxFloat(0.14, cfg.tolerance-0.02)
	transformTol := maxFloat(0.12, cfg.tolerance-0.08)
	styleSize := cfg.baseSize + 16
	transformSize := cfg.baseSize + 48

	base := func(kind, name, variant string, size, rotation int, tol, area float64) findBenchScenario {
		return findBenchScenario{
			name:                       name,
			kind:                       kind,
			variant:                    variant,
			size:                       size,
			rotation:                   rotation,
			screenW:                    cfg.screenW,
			screenH:                    cfg.screenH,
			tolerance:                  tol,
			maxAreaRatio:               area,
			queryFromBase:              true,
			expectedPositive:           true,
			allowPartial:               false,
			strictBBox:                 true,
			retryAttempts:              1,
			rpcTimeout:                 5 * time.Second,
			backgroundContinuousCanvas: true,
			backgroundClutterDensity:   0.84,
			backgroundPalette:          "mixed",
			brightnessDelta:            0.0,
			contrastFactor:             1.0,
			gammaFactor:                1.0,
			decoyEnabled:               true,
			decoyCount:                 24,
			decoySimilarity:            0.90,
			decoyPlacement:             "grid",
		}
	}

	vector := base("vector_ui", "vector_r0_"+res, "vector", styleSize, 0, styleTol, cfg.maxAreaRatio)
	vector.backgroundPalette = "ui_light"
	vector.decoyCount = 18

	photo := base("photographic", "photo_r90_"+res, "photo", styleSize, 90, styleTol, cfg.maxAreaRatio)
	photo.decoyPlacement = "mixed"
	photo.decoyCount = 30
	photo.noiseGaussian = 0.06

	ui := base("template_control", "ui_r180_"+res, "ui", styleSize, 180, styleTol, cfg.maxAreaRatio)
	ui.backgroundPalette = "ui_dark"
	ui.decoySimilarity = 0.92

	resizeDown := base("scale_rotate", fmt.Sprintf("mix_resize_%03d_%s", int(cfg.resizeDown*100), res), "photo", transformSize, 0, transformTol, cfg.maxAreaRatio+1.30)
	resizeDown.transformKind = "scale"
	resizeDown.transformA = cfg.resizeDown
	resizeDown.allowPartial = true
	resizeDown.decoyPlacement = "mixed"
	resizeDown.noiseCompression = 0.06

	resizeUp := base("scale_rotate", fmt.Sprintf("mix_resize_%03d_%s", int(cfg.resizeUp*100), res), "photo", transformSize, 0, transformTol, cfg.maxAreaRatio+1.30)
	resizeUp.transformKind = "scale"
	resizeUp.transformA = cfg.resizeUp
	resizeUp.allowPartial = true
	resizeUp.decoyPlacement = "mixed"
	resizeUp.noiseCompression = 0.06

	rotate := base("scale_rotate", fmt.Sprintf("mix_rotate_%ddeg_%s", int(cfg.rotateDeg), res), "ui", transformSize, 0, transformTol, cfg.maxAreaRatio+1.10)
	rotate.transformKind = "rotate"
	rotate.transformA = cfg.rotateDeg
	rotate.allowPartial = true
	rotate.decoyPlacement = "clustered"
	rotate.occlusionEnabled = true
	rotate.occlusionCoverage = 0.12

	return []findBenchScenario{vector, photo, ui, resizeDown, resizeUp, rotate}
}

func runFindOnScreenScenarioBenchmark(b *testing.B, engine findBenchEngine, scenario findBenchScenario, visuals *findBenchVisualCollector) {
	source, queries := buildFindBenchFixture(b, scenario)
	if len(queries) == 0 {
		b.Fatalf("empty benchmark query set scenario=%s", scenario.name)
	}

	svc := NewServer(WithCaptureScreen(func(_ context.Context, _ string) (*sikuli.Image, error) {
		return source, nil
	}))

	lis := bufconn.Listen(1024 * 1024)
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(UnaryInterceptors("", nil, NewMetricsRegistry(), nil)...),
	)
	pb.RegisterSikuliServiceServer(srv, svc)

	go func() {
		_ = srv.Serve(lis)
	}()
	b.Cleanup(func() {
		srv.Stop()
		_ = lis.Close()
	})

	dialCtx, dialCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer dialCancel()
	conn, err := grpc.DialContext(
		dialCtx,
		"bufnet",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return lis.DialContext(ctx)
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		b.Fatalf("dial bufconn: %v", err)
	}
	b.Cleanup(func() { _ = conn.Close() })

	client := pb.NewSikuliServiceClient(conn)
	requests := make([]*pb.FindOnScreenRequest, 0, len(queries))
	visualQueries := make([]findBenchVisualQuery, 0, len(queries))
	expectedRects := make([]*pb.Rect, 0, len(queries))
	for _, q := range queries {
		request := &pb.FindOnScreenRequest{
			Pattern: &pb.Pattern{
				Image:      q.Pattern,
				Similarity: benchSimilarityPtr(engine.name),
			},
			Opts: &pb.ScreenQueryOptions{
				MatcherEngine: engine.enum,
			},
		}
		requests = append(requests, request)
		visualQueries = append(visualQueries, findBenchVisualQuery{
			Request:    request,
			Expected:   q.Expected,
			Alternates: cloneRectSlice(q.ExpectedAlternates),
			Label:      q.Label,
		})
		expectedRects = appendExpectedCandidates(expectedRects, q.Expected, q.ExpectedAlternates)
	}

	if visuals != nil && len(visualQueries) > 0 {
		visuals.CaptureAttempts(b, client, visualQueries, source.Gray(), engine, scenario)
	}

	for _, request := range requests {
		_, _ = benchFindOnScreenCall(client, request, scenario)
	}

	successCount := 0
	notFoundCount := 0
	unsupportedCount := 0
	errorCount := 0
	overlapMissCount := 0
	queryTotalCount := 0
	queryOKCount := 0
	queryNotFoundCount := 0
	queryUnsupportedCount := 0
	queryErrorCount := 0
	queryOverlapMissCount := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterationOK := true

		if !scenario.expectedPositive {
			for _, request := range requests {
				queryTotalCount++
				_, err := benchFindOnScreenCall(client, request, scenario)
				if err == nil {
					queryOverlapMissCount++
					if iterationOK {
						overlapMissCount++
					}
					iterationOK = false
					continue
				}
				switch status.Code(err) {
				case codes.NotFound:
					queryOKCount++
					// expected for negative scenarios.
				case codes.Unimplemented:
					queryUnsupportedCount++
					if iterationOK {
						unsupportedCount++
					}
					iterationOK = false
				default:
					queryErrorCount++
					if iterationOK {
						errorCount++
					}
					iterationOK = false
				}
			}
			if iterationOK {
				successCount++
			}
			continue
		}

		for queryIdx, request := range requests {
			queryTotalCount++
			res, err := benchFindOnScreenCall(client, request, scenario)
			if err != nil {
				switch status.Code(err) {
				case codes.NotFound:
					queryNotFoundCount++
					if iterationOK {
						notFoundCount++
					}
				case codes.Unimplemented:
					queryUnsupportedCount++
					if iterationOK {
						unsupportedCount++
					}
				default:
					queryErrorCount++
					if iterationOK {
						errorCount++
					}
				}
				iterationOK = false
				continue
			}
			if res.GetMatch() == nil || res.GetMatch().GetRect() == nil {
				queryErrorCount++
				if iterationOK {
					errorCount++
				}
				iterationOK = false
				continue
			}
			expectedRect := queries[queryIdx].Expected
			if expectedRect == nil {
				queryErrorCount++
				if iterationOK {
					errorCount++
				}
				iterationOK = false
				continue
			}
			matchClass := classifyFindBenchPositiveMatch(
				res.GetMatch().GetRect(),
				expectedRect,
				queries[queryIdx].ExpectedAlternates,
				expectedRects,
				request.GetPattern().GetImage(),
				scenario.tolerance,
				scenario.maxAreaRatio,
				scenario.allowPartial,
			)
			if matchClass != findBenchMatchClassOK {
				queryOverlapMissCount++
				if iterationOK {
					overlapMissCount++
				}
				iterationOK = false
				continue
			}
			queryOKCount++
		}
		if iterationOK {
			successCount++
		}
	}
	if b.N > 0 {
		denom := float64(b.N)
		b.ReportMetric(float64(successCount)/denom, "success/op")
		b.ReportMetric(float64(notFoundCount)/denom, "not_found/op")
		b.ReportMetric(float64(unsupportedCount)/denom, "unsupported/op")
		b.ReportMetric(float64(errorCount)/denom, "error/op")
		b.ReportMetric(float64(overlapMissCount)/denom, "overlap_miss/op")
		b.ReportMetric(float64(queryTotalCount)/denom, "query_total/op")
		b.ReportMetric(float64(queryOKCount)/denom, "query_ok/op")
		b.ReportMetric(float64(queryNotFoundCount)/denom, "query_not_found/op")
		b.ReportMetric(float64(queryUnsupportedCount)/denom, "query_unsupported/op")
		b.ReportMetric(float64(queryErrorCount)/denom, "query_error/op")
		b.ReportMetric(float64(queryOverlapMissCount)/denom, "query_overlap_miss/op")
	}
}

func classifyFindBenchPositiveMatch(
	found *pb.Rect,
	expected *pb.Rect,
	expectedAlternates []*pb.Rect,
	allExpected []*pb.Rect,
	pattern *pb.GrayImage,
	tolerance float64,
	maxAreaRatio float64,
	allowPartial bool,
) findBenchMatchClass {
	if found == nil || expected == nil {
		return findBenchMatchClassOverlapMiss
	}

	ownExpected := expectedCandidateSet(expected, expectedAlternates)
	for _, own := range ownExpected {
		if matchRectAgainstRegion(found, own, pattern, tolerance, maxAreaRatio, allowPartial) {
			return findBenchMatchClassOK
		}
	}
	for _, peer := range allExpected {
		if peer == nil || rectExistsInSet(peer, ownExpected) {
			continue
		}
		if matchRectAgainstRegion(found, peer, pattern, tolerance, maxAreaRatio, allowPartial) {
			return findBenchMatchClassWrongRegion
		}
	}
	return findBenchMatchClassOverlapMiss
}

func expectedCandidateSet(primary *pb.Rect, alternates []*pb.Rect) []*pb.Rect {
	out := make([]*pb.Rect, 0, 1+len(alternates))
	if primary != nil {
		out = append(out, primary)
	}
	for _, alt := range alternates {
		if alt == nil {
			continue
		}
		if rectExistsInSet(alt, out) {
			continue
		}
		out = append(out, alt)
	}
	return out
}

func rectExistsInSet(target *pb.Rect, set []*pb.Rect) bool {
	if target == nil {
		return false
	}
	for _, item := range set {
		if rectsEquivalent(item, target) {
			return true
		}
	}
	return false
}

func appendExpectedCandidates(dst []*pb.Rect, primary *pb.Rect, alternates []*pb.Rect) []*pb.Rect {
	for _, candidate := range expectedCandidateSet(primary, alternates) {
		if rectExistsInSet(candidate, dst) {
			continue
		}
		dst = append(dst, candidate)
	}
	return dst
}

func cloneRectSlice(in []*pb.Rect) []*pb.Rect {
	if len(in) == 0 {
		return nil
	}
	out := make([]*pb.Rect, 0, len(in))
	for _, rect := range in {
		if rect == nil {
			continue
		}
		out = append(out, cloneRect(rect))
	}
	return out
}

func matchRectAgainstRegion(
	found *pb.Rect,
	expected *pb.Rect,
	pattern *pb.GrayImage,
	tolerance float64,
	maxAreaRatio float64,
	allowPartial bool,
) bool {
	if found == nil || expected == nil {
		return false
	}
	_, _, _, _, _, ok := centerDeltaMetrics(found, expected, pattern, tolerance, maxAreaRatio, allowPartial)
	return ok
}

func rectsEquivalent(a, b *pb.Rect) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.GetX() == b.GetX() &&
		a.GetY() == b.GetY() &&
		a.GetW() == b.GetW() &&
		a.GetH() == b.GetH()
}

func centerDeltaMetrics(
	found *pb.Rect,
	expected *pb.Rect,
	pattern *pb.GrayImage,
	tolerance float64,
	maxAreaRatio float64,
	allowPartial bool,
) (dxAbs float64, dyAbs float64, distance float64, limitX float64, limitY float64, ok bool) {
	if found == nil || expected == nil {
		return 0, 0, 0, 0, 0, false
	}
	if !areaRatioWithinLimit(found, pattern, maxAreaRatio) {
		limitX, limitY = centerDeltaThreshold(expected, tolerance, allowPartial)
		return 0, 0, 0, limitX, limitY, false
	}
	if regionActsAsZone(expected, pattern) {
		dxAbs, dyAbs, distance = centerDeltaToRect(found, expected)
	} else {
		fx, fy := rectCenter(found)
		ex, ey := rectCenter(expected)
		dxAbs = math.Abs(fx - ex)
		dyAbs = math.Abs(fy - ey)
		distance = math.Hypot(dxAbs, dyAbs)
	}
	limitX, limitY = centerDeltaThreshold(expected, tolerance, allowPartial)
	ok = dxAbs <= limitX && dyAbs <= limitY
	return dxAbs, dyAbs, distance, limitX, limitY, ok
}

func areaRatioWithinLimit(found *pb.Rect, pattern *pb.GrayImage, maxAreaRatio float64) bool {
	if found == nil || pattern == nil {
		return true
	}
	if maxAreaRatio <= 0 {
		maxAreaRatio = 1.50
	}
	patternArea := float64(max32(1, pattern.GetWidth()*pattern.GetHeight()))
	foundArea := rectAreaFloat(found)
	if patternArea <= 0 || foundArea <= 0 {
		return true
	}
	return (foundArea / patternArea) <= maxAreaRatio
}

func regionActsAsZone(expected *pb.Rect, pattern *pb.GrayImage) bool {
	if expected == nil || pattern == nil {
		return false
	}
	zoneArea := rectAreaFloat(expected)
	patternArea := float64(max32(1, pattern.GetWidth()*pattern.GetHeight()))
	return zoneArea > patternArea*1.35
}

func centerDeltaToRect(found *pb.Rect, expected *pb.Rect) (dxAbs float64, dyAbs float64, distance float64) {
	if found == nil || expected == nil {
		return 0, 0, 0
	}
	fx, fy := rectCenter(found)
	left := float64(expected.GetX())
	top := float64(expected.GetY())
	right := left + float64(expected.GetW())
	bottom := top + float64(expected.GetH())

	if fx < left {
		dxAbs = left - fx
	} else if fx > right {
		dxAbs = fx - right
	}
	if fy < top {
		dyAbs = top - fy
	} else if fy > bottom {
		dyAbs = fy - bottom
	}
	distance = math.Hypot(dxAbs, dyAbs)
	return dxAbs, dyAbs, distance
}

func centerDeltaThreshold(expected *pb.Rect, tolerance float64, allowPartial bool) (float64, float64) {
	tol := math.Max(0.0, math.Min(1.0, tolerance))
	factor := 0.18 + tol*0.45
	if allowPartial {
		factor *= 1.08
	}
	factor = math.Max(0.16, math.Min(0.50, factor))
	limitX := math.Max(8.0, float64(max32(1, expected.GetW()))*factor)
	limitY := math.Max(8.0, float64(max32(1, expected.GetH()))*factor)
	return limitX, limitY
}

func rectCenter(r *pb.Rect) (float64, float64) {
	if r == nil {
		return 0, 0
	}
	return float64(r.GetX()) + float64(r.GetW())/2.0, float64(r.GetY()) + float64(r.GetH())/2.0
}

func benchFindOnScreenCall(client pb.SikuliServiceClient, request *pb.FindOnScreenRequest, scenario findBenchScenario) (*pb.FindResponse, error) {
	retries := scenario.retryAttempts
	if retries < 0 {
		retries = 0
	}
	timeout := scenario.rpcTimeout
	if timeout <= 0 {
		timeout = 5 * time.Second
	}

	var lastResp *pb.FindResponse
	var lastErr error
	for attempt := 0; attempt <= retries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		resp, err := client.FindOnScreen(ctx, request)
		cancel()
		lastResp, lastErr = resp, err
		if err == nil {
			return resp, nil
		}
		code := status.Code(err)
		if code != codes.NotFound && code != codes.DeadlineExceeded {
			return resp, err
		}
	}
	return lastResp, lastErr
}

func buildFindBenchFixture(t testing.TB, scenario findBenchScenario) (*sikuli.Image, []findBenchFixtureQuery) {
	t.Helper()

	if source, queries, ok := buildFindBenchFixtureFromRegionSpec(t, scenario); ok {
		return source, queries
	}

	if source, queries, ok := buildFindBenchFixtureFromPhotoAsset(t, scenario); ok {
		return source, queries
	}

	basePattern := buildBenchPattern(scenario.variant, scenario.size)
	queryPattern := rotateGrayByQuarterTurns(basePattern, scenario.rotation)
	targetPattern := applyBenchTransform(queryPattern, scenario)
	haystack := buildBenchHaystack(scenario.screenW, scenario.screenH, scenario)

	// Apply monitor profile simulation first to the capture plane.
	applyMonitorProfile(haystack, scenario)

	pbounds := targetPattern.Bounds()
	insertX := (scenario.screenW - pbounds.Dx()) / 3
	insertY := (scenario.screenH - pbounds.Dy()) / 2
	if insertX < 0 || insertY < 0 {
		t.Fatalf("pattern does not fit haystack scenario=%s pattern=%dx%d haystack=%dx%d", scenario.name, pbounds.Dx(), pbounds.Dy(), scenario.screenW, scenario.screenH)
	}
	// Fill the full frame with near-match decoys first so the target is camouflaged.
	populateNearMatchDecoys(haystack, targetPattern, insertX, insertY, scenario)
	// Blend target edges slightly into local context, then use the inserted patch as the query pattern.
	blitGrayFeather(haystack, targetPattern, insertX, insertY, maxInt(2, scenario.size/20))
	applyTargetOcclusion(haystack, insertX, insertY, pbounds.Dx(), pbounds.Dy(), scenario)
	applyPhotometricProfile(haystack, scenario, true)
	applyPhotometricProfile(targetPattern, scenario, false)
	applySeamSmoothing(haystack, scenario.variant)

	source, err := sikuli.NewImageFromGray(fmt.Sprintf("bench-%s-source", scenario.name), haystack)
	if err != nil {
		t.Fatalf("new source image: %v", err)
	}

	// Benchmark query always comes from source image-1, then is searched in modified image-2.
	patternGray := queryPattern
	pattern := grayProtoFromGray(fmt.Sprintf("bench-%s-pattern", scenario.name), patternGray)
	expected := &pb.Rect{X: int32(insertX), Y: int32(insertY), W: int32(pbounds.Dx()), H: int32(pbounds.Dy())}
	return source, []findBenchFixtureQuery{
		{
			Pattern:  pattern,
			Expected: expected,
			Label:    "primary",
		},
	}
}

func buildFindBenchFixtureFromPhotoAsset(t testing.TB, scenario findBenchScenario) (*sikuli.Image, []findBenchFixtureQuery, bool) {
	t.Helper()

	assetPath := strings.TrimSpace(scenario.targetAsset)
	if assetPath == "" {
		return nil, nil, false
	}
	if !(strings.EqualFold(scenario.kind, "photographic") || strings.EqualFold(scenario.targetSource, "asset") || strings.EqualFold(scenario.targetSource, "mixed")) {
		return nil, nil, false
	}

	assetGray, err := loadGrayFromFile(assetPath)
	if err != nil {
		t.Logf("find bench photo asset load failed path=%s err=%v fallback=synthetic", assetPath, err)
		return nil, nil, false
	}

	haystack := buildPhotoHaystackFromAsset(assetGray, scenario.screenW, scenario.screenH, scenario.seed)
	if haystack == nil {
		t.Logf("find bench photo asset haystack failed path=%s fallback=synthetic", assetPath)
		return nil, nil, false
	}

	hb := haystack.Bounds()
	assetSize := minInt(minInt(hb.Dx(), hb.Dy()), maxInt(16, scenario.size))
	if assetSize < 16 {
		t.Logf("find bench photo asset too small path=%s fallback=synthetic", assetPath)
		return nil, nil, false
	}

	cropX, cropY := chooseBenchCropOrigin(hb.Dx(), hb.Dy(), assetSize, scenario.seed)
	basePattern := cropGray(haystack, image.Rect(cropX, cropY, cropX+assetSize, cropY+assetSize))
	queryPattern := rotateGrayByQuarterTurns(basePattern, scenario.rotation)
	targetPattern := applyBenchTransform(queryPattern, scenario)
	pbounds := targetPattern.Bounds()

	insertX := (scenario.screenW - pbounds.Dx()) / 3
	insertY := (scenario.screenH - pbounds.Dy()) / 2
	if insertX < 0 || insertY < 0 {
		t.Logf("find bench photo asset transformed target does not fit scenario=%s path=%s fallback=synthetic", scenario.name, assetPath)
		return nil, nil, false
	}

	populateNearMatchDecoys(haystack, targetPattern, insertX, insertY, scenario)
	blitGrayFeather(haystack, targetPattern, insertX, insertY, maxInt(2, scenario.size/20))
	applyTargetOcclusion(haystack, insertX, insertY, pbounds.Dx(), pbounds.Dy(), scenario)
	applyPhotometricProfile(haystack, scenario, true)
	applyPhotometricProfile(targetPattern, scenario, false)
	applySeamSmoothing(haystack, scenario.variant)

	source, err := sikuli.NewImageFromGray(fmt.Sprintf("bench-%s-source", scenario.name), haystack)
	if err != nil {
		t.Logf("find bench photo asset source image failed scenario=%s path=%s err=%v fallback=synthetic", scenario.name, assetPath, err)
		return nil, nil, false
	}

	// Benchmark query always comes from source image-1, then is searched in modified image-2.
	patternGray := queryPattern
	pattern := grayProtoFromGray(fmt.Sprintf("bench-%s-pattern", scenario.name), patternGray)
	expected := &pb.Rect{X: int32(insertX), Y: int32(insertY), W: int32(pbounds.Dx()), H: int32(pbounds.Dy())}
	return source, []findBenchFixtureQuery{
		{
			Pattern:  pattern,
			Expected: expected,
			Label:    "primary",
		},
	}, true
}

type benchPointF struct {
	X float64
	Y float64
}

func buildFindBenchFixtureFromRegionSpec(t testing.TB, scenario findBenchScenario) (*sikuli.Image, []findBenchFixtureQuery, bool) {
	t.Helper()

	sourcePath := strings.TrimSpace(scenario.sourceImagePath)
	if sourcePath == "" || len(scenario.sourceTargets) == 0 {
		return nil, nil, false
	}

	raw, err := loadGrayFromFile(sourcePath)
	if err != nil {
		t.Logf("find bench region-spec source load failed path=%s err=%v fallback=synthetic", sourcePath, err)
		return nil, nil, false
	}
	sourceNativeW := raw.Bounds().Dx()
	sourceNativeH := raw.Bounds().Dy()

	sourceScene, sourceRegions := normalizeSceneAndRegions(raw, scenario.screenW, scenario.screenH, scenario.sourceTargets)
	if sourceScene == nil || len(sourceRegions) == 0 {
		t.Logf("find bench region-spec normalize failed scenario=%s path=%s fallback=synthetic", scenario.name, sourcePath)
		return nil, nil, false
	}
	if projected := projectTargetRegionsToScreen(scenario.sourceTargets, sourceNativeW, sourceNativeH, scenario.screenW, scenario.screenH); len(projected) > 0 {
		sourceRegions = projected
	}

	regionsForBenchmark := scenario.sourceTargets
	if len(scenario.benchmarkTargets) > 0 {
		regionsForBenchmark = scenario.benchmarkTargets
	}

	benchScene := sourceScene
	benchRegions := append([]findBenchTargetRegion(nil), sourceRegions...)
	benchNativeW := sourceNativeW
	benchNativeH := sourceNativeH
	if projected := projectTargetRegionsToScreen(regionsForBenchmark, benchNativeW, benchNativeH, scenario.screenW, scenario.screenH); len(projected) > 0 {
		benchRegions = projected
	}
	benchPath := strings.TrimSpace(scenario.benchmarkImagePath)
	if benchPath != "" && benchPath != sourcePath {
		benchRaw, benchErr := loadGrayFromFile(benchPath)
		if benchErr != nil {
			t.Logf("find bench region-spec benchmark load failed path=%s err=%v fallback=source", benchPath, benchErr)
		} else {
			benchNativeW = benchRaw.Bounds().Dx()
			benchNativeH = benchRaw.Bounds().Dy()
			candidateScene, candidateRegions := normalizeSceneAndRegions(benchRaw, scenario.screenW, scenario.screenH, regionsForBenchmark)
			if candidateScene == nil || len(candidateRegions) == 0 {
				t.Logf("find bench region-spec benchmark normalize failed scenario=%s path=%s fallback=source", scenario.name, benchPath)
			} else {
				benchScene = candidateScene
				benchRegions = candidateRegions
			}
			if projected := projectTargetRegionsToScreen(regionsForBenchmark, benchNativeW, benchNativeH, scenario.screenW, scenario.screenH); len(projected) > 0 {
				benchRegions = projected
			}
		}
	}

	source, err := sikuli.NewImageFromGray(fmt.Sprintf("bench-%s-source", scenario.name), benchScene)
	if err != nil {
		t.Logf("find bench region-spec source image failed scenario=%s err=%v fallback=synthetic", scenario.name, err)
		return nil, nil, false
	}

	queries := make([]findBenchFixtureQuery, 0, len(sourceRegions))
	for idx, sourceRegion := range sourceRegions {
		if sourceRegion.W <= 0 || sourceRegion.H <= 0 {
			continue
		}
		sourcePattern := cropGray(sourceScene, image.Rect(sourceRegion.X, sourceRegion.Y, sourceRegion.X+sourceRegion.W, sourceRegion.Y+sourceRegion.H))
		if sourcePattern.Bounds().Dx() < 4 || sourcePattern.Bounds().Dy() < 4 {
			continue
		}
		benchRegion, ok := findScenarioTargetRegionByID(benchRegions, sourceRegion.ID)
		if !ok && idx < len(benchRegions) {
			benchRegion = benchRegions[idx]
			ok = true
		}
		if !ok && len(benchRegions) > 0 {
			benchRegion = chooseScenarioTargetRegion(benchRegions, scenario.seed+uint64(idx))
			ok = true
		}
		if !ok || benchRegion.W <= 0 || benchRegion.H <= 0 {
			continue
		}
		expectedRegion, alternates := chooseExpectedRegionsForScenario(scenario, sourceRegion, benchRegion)
		label := strings.TrimSpace(sourceRegion.Label)
		if label == "" {
			label = strings.TrimSpace(sourceRegion.ID)
		}
		if label == "" {
			label = fmt.Sprintf("target-%02d", idx+1)
		}
		queries = append(queries, findBenchFixtureQuery{
			Pattern:            grayProtoFromGray(fmt.Sprintf("bench-%s-pattern-%02d", scenario.name, idx+1), sourcePattern),
			Expected:           benchRectFromTargetRegion(expectedRegion),
			ExpectedAlternates: alternates,
			Label:              label,
		})
	}
	if len(queries) == 0 {
		t.Logf("find bench region-spec no valid mapped queries scenario=%s fallback=synthetic", scenario.name)
		return nil, nil, false
	}
	return source, queries, true
}

func chooseScenarioTargetRegionPair(regions []findBenchTargetRegion, seed uint64) (findBenchTargetRegion, findBenchTargetRegion) {
	if len(regions) == 0 {
		return findBenchTargetRegion{}, findBenchTargetRegion{}
	}
	firstIdx := int(seed % uint64(len(regions)))
	if firstIdx < 0 || firstIdx >= len(regions) {
		firstIdx = 0
	}
	secondIdx := firstIdx
	if len(regions) > 1 {
		offset := 1 + int((seed>>17)%uint64(len(regions)-1))
		secondIdx = (firstIdx + offset) % len(regions)
	}
	return regions[firstIdx], regions[secondIdx]
}

func benchRectFromTargetRegion(region findBenchTargetRegion) *pb.Rect {
	if region.W <= 0 || region.H <= 0 {
		return &pb.Rect{}
	}
	return &pb.Rect{
		X: int32(region.X),
		Y: int32(region.Y),
		W: int32(region.W),
		H: int32(region.H),
	}
}

func composeMultiMonitorDPISceneFromSource(
	sourceScene *image.Gray,
	sourceRegions []findBenchTargetRegion,
	scenario findBenchScenario,
) (*image.Gray, []findBenchTargetRegion, []findBenchTargetRegion) {
	if sourceScene == nil || scenario.screenW < 4 || scenario.screenH < 4 {
		return nil, nil, nil
	}

	out := image.NewGray(image.Rect(0, 0, scenario.screenW, scenario.screenH))
	setRect(out, 0, 0, scenario.screenW, scenario.screenH, 214)

	gap := maxInt(8, scenario.screenW/120)
	leftW := maxInt(1, (scenario.screenW-gap)/2)
	rightW := maxInt(1, scenario.screenW-gap-leftW)
	leftRect := image.Rect(0, 0, leftW, scenario.screenH)
	rightRect := image.Rect(leftW+gap, 0, leftW+gap+rightW, scenario.screenH)
	setRect(out, leftRect.Min.X, leftRect.Min.Y, leftRect.Dx(), leftRect.Dy(), 226)
	setRect(out, rightRect.Min.X, rightRect.Min.Y, rightRect.Dx(), rightRect.Dy(), 226)
	setRect(out, leftW, 0, gap, scenario.screenH, 168)

	leftScale := 0.98 + (float64(int((scenario.seed>>3)%5)-2) * 0.01)
	rightScale := 1.22 + (float64(int((scenario.seed>>9)%5)-2) * 0.02)
	leftMapped := blitScaledSceneIntoMonitorRect(out, leftRect, sourceScene, sourceRegions, leftScale)
	rightMapped := blitScaledSceneIntoMonitorRect(out, rightRect, sourceScene, sourceRegions, rightScale)
	return out, leftMapped, rightMapped
}

func blitScaledSceneIntoMonitorRect(
	dst *image.Gray,
	dstRect image.Rectangle,
	src *image.Gray,
	srcRegions []findBenchTargetRegion,
	scaleFactor float64,
) []findBenchTargetRegion {
	if dst == nil || src == nil || dstRect.Dx() < 1 || dstRect.Dy() < 1 {
		return nil
	}
	sb := src.Bounds()
	sw := maxInt(1, sb.Dx())
	sh := maxInt(1, sb.Dy())
	fit := math.Min(float64(dstRect.Dx())/float64(sw), float64(dstRect.Dy())/float64(sh))
	if fit <= 0 {
		fit = 1
	}
	scale := fit * scaleFactor
	if scale <= 0 {
		scale = fit
	}
	scaled := scaleGrayNearestBench(src, scale)
	swb := scaled.Bounds()
	offX := dstRect.Min.X + (dstRect.Dx()-swb.Dx())/2
	offY := dstRect.Min.Y + (dstRect.Dy()-swb.Dy())/2
	blitGray(dst, scaled, offX, offY)

	mapped := make([]findBenchTargetRegion, 0, len(srcRegions))
	for _, region := range srcRegions {
		if region.W <= 0 || region.H <= 0 {
			continue
		}
		x := offX + int(math.Round(float64(region.X)*scale))
		y := offY + int(math.Round(float64(region.Y)*scale))
		w := maxInt(1, int(math.Round(float64(region.W)*scale)))
		h := maxInt(1, int(math.Round(float64(region.H)*scale)))
		rr := image.Rect(x, y, x+w, y+h).Intersect(dstRect).Intersect(dst.Bounds())
		if rr.Empty() || rr.Dx() < 2 || rr.Dy() < 2 {
			continue
		}
		mapped = append(mapped, findBenchTargetRegion{
			ID:    region.ID,
			Label: region.Label,
			X:     rr.Min.X,
			Y:     rr.Min.Y,
			W:     rr.Dx(),
			H:     rr.Dy(),
		})
	}
	return mapped
}

func chooseScenarioTargetRegion(regions []findBenchTargetRegion, seed uint64) findBenchTargetRegion {
	if len(regions) == 0 {
		return findBenchTargetRegion{}
	}
	idx := int(seed % uint64(len(regions)))
	if idx < 0 || idx >= len(regions) {
		idx = 0
	}
	return regions[idx]
}

func chooseExpectedRegionsForScenario(
	scenario findBenchScenario,
	sourceRegion findBenchTargetRegion,
	benchmarkRegion findBenchTargetRegion,
) (findBenchTargetRegion, []*pb.Rect) {
	primary := sourceRegion
	if shouldPreferBenchmarkExpectedRegion(scenario) {
		primary = benchmarkRegion
	}

	alternates := make([]*pb.Rect, 0, 1)
	if !targetRegionsEquivalent(sourceRegion, benchmarkRegion) {
		if shouldPreferBenchmarkExpectedRegion(scenario) {
			alternates = append(alternates, benchRectFromTargetRegion(sourceRegion))
		} else {
			alternates = append(alternates, benchRectFromTargetRegion(benchmarkRegion))
		}
	}
	return primary, alternates
}

func shouldPreferBenchmarkExpectedRegion(scenario findBenchScenario) bool {
	kind := strings.ToLower(strings.TrimSpace(scenario.kind))
	switch kind {
	case "multi_monitor_dpi", "multi_monitor":
		return true
	default:
		return false
	}
}

func targetRegionsEquivalent(a, b findBenchTargetRegion) bool {
	return a.X == b.X && a.Y == b.Y && a.W == b.W && a.H == b.H
}

func findScenarioTargetRegionByID(regions []findBenchTargetRegion, id string) (findBenchTargetRegion, bool) {
	id = strings.TrimSpace(id)
	if id == "" {
		return findBenchTargetRegion{}, false
	}
	for _, region := range regions {
		if strings.TrimSpace(region.ID) == id {
			return region, true
		}
	}
	return findBenchTargetRegion{}, false
}

func normalizeSceneAndRegions(src *image.Gray, screenW, screenH int, regions []findBenchTargetRegion) (*image.Gray, []findBenchTargetRegion) {
	if src == nil {
		return nil, nil
	}
	sb := src.Bounds()
	sw := maxInt(1, sb.Dx())
	sh := maxInt(1, sb.Dy())
	if screenW < 1 || screenH < 1 {
		return nil, nil
	}
	scale := math.Max(float64(screenW)/float64(sw), float64(screenH)/float64(sh))
	if scale <= 0 {
		scale = 1
	}
	scaled := scaleGrayNearestBench(src, scale)
	scaledW := maxInt(1, scaled.Bounds().Dx())
	scaledH := maxInt(1, scaled.Bounds().Dy())
	cropX := maxInt(0, (scaledW-screenW)/2)
	cropY := maxInt(0, (scaledH-screenH)/2)
	scene := cropGray(scaled, image.Rect(cropX, cropY, cropX+screenW, cropY+screenH))
	out := projectTargetRegionsToScreen(regions, sw, sh, screenW, screenH)
	return scene, out
}

func projectTargetRegionsToScreen(regions []findBenchTargetRegion, sourceW, sourceH, screenW, screenH int) []findBenchTargetRegion {
	if len(regions) == 0 || sourceW < 1 || sourceH < 1 || screenW < 1 || screenH < 1 {
		return nil
	}
	scale := math.Max(float64(screenW)/float64(sourceW), float64(screenH)/float64(sourceH))
	if scale <= 0 {
		scale = 1
	}
	scaledW := maxInt(1, int(math.Round(float64(sourceW)*scale)))
	scaledH := maxInt(1, int(math.Round(float64(sourceH)*scale)))
	cropX := maxInt(0, (scaledW-screenW)/2)
	cropY := maxInt(0, (scaledH-screenH)/2)
	screenRect := image.Rect(0, 0, screenW, screenH)

	out := make([]findBenchTargetRegion, 0, len(regions))
	for _, r := range regions {
		if r.W <= 0 || r.H <= 0 {
			continue
		}
		x := int(math.Round(float64(r.X)*scale)) - cropX
		y := int(math.Round(float64(r.Y)*scale)) - cropY
		w := maxInt(1, int(math.Round(float64(r.W)*scale)))
		h := maxInt(1, int(math.Round(float64(r.H)*scale)))
		rr := image.Rect(x, y, x+w, y+h).Intersect(screenRect)
		if rr.Empty() || rr.Dx() < 2 || rr.Dy() < 2 {
			continue
		}
		out = append(out, findBenchTargetRegion{
			ID:    r.ID,
			Label: r.Label,
			X:     rr.Min.X,
			Y:     rr.Min.Y,
			W:     rr.Dx(),
			H:     rr.Dy(),
		})
	}
	return out
}

func applyGlobalTransformToSceneAndRegion(scene *image.Gray, region findBenchTargetRegion, scenario findBenchScenario) (*image.Gray, *pb.Rect, bool) {
	if scene == nil || region.W <= 0 || region.H <= 0 {
		return nil, nil, false
	}
	current := cloneGray(scene)
	corners := regionToCorners(region)

	turns := ((scenario.rotation % 360) + 360) % 360 / 90
	for i := 0; i < turns; i++ {
		current, corners = rotateSceneAndCorners90(current, corners)
	}

	switch scenario.transformKind {
	case "":
		// no-op
	case "scale":
		f := scenario.transformA
		if f <= 0 {
			f = 1
		}
		current = scaleGrayNearestBench(current, f)
		for i := range corners {
			corners[i].X *= f
			corners[i].Y *= f
		}
	case "rotate":
		current, corners = rotateSceneAndCornersArbitrary(current, corners, scenario.transformA)
	case "skewx":
		current, corners = skewSceneAndCorners(current, corners, scenario.transformA)
	case "perspective":
		topScale := scenario.transformA
		bottomScale := scenario.transformB
		shift := scenario.transformC
		if topScale <= 0 {
			topScale = 0.90
		}
		if bottomScale <= 0 {
			bottomScale = 1.08
		}
		current, corners = perspectiveSceneAndCorners(current, corners, topScale, bottomScale, shift)
	default:
		current = applyBenchTransform(current, scenario)
	}

	finalScene, finalCorners := fitSceneAndCornersToScreen(current, corners, scenario.screenW, scenario.screenH)
	rect := cornersToRect(finalCorners, finalScene.Bounds())
	if rect.GetW() < 2 || rect.GetH() < 2 {
		return nil, nil, false
	}
	return finalScene, rect, true
}

func rotateSceneAndCorners90(src *image.Gray, corners []benchPointF) (*image.Gray, []benchPointF) {
	sb := src.Bounds()
	sw := sb.Dx()
	sh := sb.Dy()
	dst := rotate90Gray(src)
	out := make([]benchPointF, 0, len(corners))
	for _, p := range corners {
		x := float64(sh-1) - p.Y
		y := p.X
		out = append(out, benchPointF{X: x, Y: y})
	}
	if sw <= 0 || sh <= 0 {
		return dst, corners
	}
	return dst, out
}

func rotateSceneAndCornersArbitrary(src *image.Gray, corners []benchPointF, degrees float64) (*image.Gray, []benchPointF) {
	sb := src.Bounds()
	sw, sh := sb.Dx(), sb.Dy()
	if sw <= 0 || sh <= 0 {
		return rotateGrayBilinearBench(src, degrees, 128), corners
	}
	theta := degrees * math.Pi / 180.0
	cosT := math.Cos(theta)
	sinT := math.Sin(theta)
	cx := (float64(sw) - 1) / 2.0
	cy := (float64(sh) - 1) / 2.0
	imageCorners := [][2]float64{
		{-cx, -cy},
		{float64(sw-1) - cx, -cy},
		{float64(sw-1) - cx, float64(sh-1) - cy},
		{-cx, float64(sh-1) - cy},
	}
	minX, maxX := math.MaxFloat64, -math.MaxFloat64
	minY, maxY := math.MaxFloat64, -math.MaxFloat64
	for _, c := range imageCorners {
		x := cosT*c[0] - sinT*c[1]
		y := sinT*c[0] + cosT*c[1]
		minX = math.Min(minX, x)
		maxX = math.Max(maxX, x)
		minY = math.Min(minY, y)
		maxY = math.Max(maxY, y)
	}
	dst := rotateGrayBilinearBench(src, degrees, 128)
	out := make([]benchPointF, 0, len(corners))
	for _, p := range corners {
		xr := p.X - cx
		yr := p.Y - cy
		rx := cosT*xr - sinT*yr
		ry := sinT*xr + cosT*yr
		out = append(out, benchPointF{X: rx - minX, Y: ry - minY})
	}
	_ = maxX
	_ = maxY
	return dst, out
}

func skewSceneAndCorners(src *image.Gray, corners []benchPointF, skew float64) (*image.Gray, []benchPointF) {
	sb := src.Bounds()
	sw, sh := sb.Dx(), sb.Dy()
	cx := (float64(sw) - 1) / 2.0
	cy := (float64(sh) - 1) / 2.0
	cornersImg := [][2]float64{
		{-cx + skew*(-cy), -cy},
		{float64(sw-1) - cx + skew*(-cy), -cy},
		{float64(sw-1) - cx + skew*(float64(sh-1)-cy), float64(sh-1) - cy},
		{-cx + skew*(float64(sh-1)-cy), float64(sh-1) - cy},
	}
	minX := math.MaxFloat64
	for _, c := range cornersImg {
		minX = math.Min(minX, c[0])
	}
	dst := skewGrayXBench(src, skew, 128)
	out := make([]benchPointF, 0, len(corners))
	for _, p := range corners {
		yr := p.Y - cy
		xr := (p.X - cx) + skew*yr
		out = append(out, benchPointF{
			X: xr - minX,
			Y: p.Y,
		})
	}
	return dst, out
}

func perspectiveSceneAndCorners(src *image.Gray, corners []benchPointF, topScale, bottomScale, shift float64) (*image.Gray, []benchPointF) {
	sb := src.Bounds()
	sw, sh := sb.Dx(), sb.Dy()
	cx := (float64(sw) - 1) / 2.0
	dst := perspectiveKeystoneBench(src, topScale, bottomScale, shift, 128)
	out := make([]benchPointF, 0, len(corners))
	for _, p := range corners {
		t := 0.0
		if sh > 1 {
			t = p.Y / float64(sh-1)
		}
		scale := topScale*(1-t) + bottomScale*t
		shiftX := shift * (0.5 - t) * float64(sw)
		x := (p.X-cx)*scale + cx + shiftX
		out = append(out, benchPointF{X: x, Y: p.Y})
	}
	return dst, out
}

func fitSceneAndCornersToScreen(src *image.Gray, corners []benchPointF, screenW, screenH int) (*image.Gray, []benchPointF) {
	sb := src.Bounds()
	sw, sh := sb.Dx(), sb.Dy()
	out := image.NewGray(image.Rect(0, 0, screenW, screenH))
	for y := 0; y < screenH; y++ {
		for x := 0; x < screenW; x++ {
			out.Pix[out.PixOffset(x, y)] = 128
		}
	}

	cropX := maxInt(0, (sw-screenW)/2)
	cropY := maxInt(0, (sh-screenH)/2)
	padX := maxInt(0, (screenW-sw)/2)
	padY := maxInt(0, (screenH-sh)/2)

	for y := 0; y < screenH; y++ {
		sy := y + cropY - padY
		if sy < 0 || sy >= sh {
			continue
		}
		for x := 0; x < screenW; x++ {
			sx := x + cropX - padX
			if sx < 0 || sx >= sw {
				continue
			}
			out.Pix[out.PixOffset(x, y)] = src.Pix[src.PixOffset(sx, sy)]
		}
	}

	shiftX := float64(-cropX + padX)
	shiftY := float64(-cropY + padY)
	adjusted := make([]benchPointF, 0, len(corners))
	for _, p := range corners {
		adjusted = append(adjusted, benchPointF{X: p.X + shiftX, Y: p.Y + shiftY})
	}
	return out, adjusted
}

func regionToCorners(r findBenchTargetRegion) []benchPointF {
	return []benchPointF{
		{X: float64(r.X), Y: float64(r.Y)},
		{X: float64(r.X + r.W), Y: float64(r.Y)},
		{X: float64(r.X + r.W), Y: float64(r.Y + r.H)},
		{X: float64(r.X), Y: float64(r.Y + r.H)},
	}
}

func cornersToRect(corners []benchPointF, bounds image.Rectangle) *pb.Rect {
	if len(corners) == 0 {
		return &pb.Rect{}
	}
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := -math.MaxFloat64, -math.MaxFloat64
	for _, p := range corners {
		minX = math.Min(minX, p.X)
		minY = math.Min(minY, p.Y)
		maxX = math.Max(maxX, p.X)
		maxY = math.Max(maxY, p.Y)
	}
	x0 := maxInt(bounds.Min.X, int(math.Floor(minX)))
	y0 := maxInt(bounds.Min.Y, int(math.Floor(minY)))
	x1 := minInt(bounds.Max.X, int(math.Ceil(maxX)))
	y1 := minInt(bounds.Max.Y, int(math.Ceil(maxY)))
	if x1 <= x0 {
		x1 = minInt(bounds.Max.X, x0+1)
	}
	if y1 <= y0 {
		y1 = minInt(bounds.Max.Y, y0+1)
	}
	return &pb.Rect{
		X: int32(x0),
		Y: int32(y0),
		W: int32(maxInt(1, x1-x0)),
		H: int32(maxInt(1, y1-y0)),
	}
}

func loadGrayFromFile(path string) (*image.Gray, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	b := img.Bounds()
	out := image.NewGray(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(out, out.Bounds(), img, b.Min, draw.Src)
	return out, nil
}

func buildPhotoHaystackFromAsset(asset *image.Gray, screenW, screenH int, seed uint64) *image.Gray {
	if asset == nil {
		return nil
	}
	working := asset
	b := working.Bounds()
	if b.Dx() <= 0 || b.Dy() <= 0 {
		return nil
	}

	if b.Dx() < screenW || b.Dy() < screenH {
		fx := float64(screenW) / float64(maxInt(1, b.Dx()))
		fy := float64(screenH) / float64(maxInt(1, b.Dy()))
		working = scaleGrayNearestBench(working, math.Max(fx, fy))
		b = working.Bounds()
	}
	if b.Dx() < screenW || b.Dy() < screenH {
		return nil
	}

	x, y := chooseBenchCropOrigin(b.Dx(), b.Dy(), minInt(screenW, screenH), seed^0x9e3779b97f4a7c15)
	if x+screenW > b.Dx() {
		x = maxInt(0, b.Dx()-screenW)
	}
	if y+screenH > b.Dy() {
		y = maxInt(0, b.Dy()-screenH)
	}
	return cropGray(working, image.Rect(x, y, x+screenW, y+screenH))
}

func chooseBenchCropOrigin(width, height, size int, seed uint64) (int, int) {
	maxX := maxInt(0, width-size)
	maxY := maxInt(0, height-size)
	x := 0
	y := 0
	if maxX > 0 {
		x = int((seed ^ (seed >> 11)) % uint64(maxX+1))
	}
	if maxY > 0 {
		y = int(((seed >> 23) ^ (seed << 7)) % uint64(maxY+1))
	}
	return x, y
}

func buildBenchPattern(variant string, size int) *image.Gray {
	if size < 16 {
		size = 16
	}
	img := image.NewGray(image.Rect(0, 0, size, size))

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			img.Pix[img.PixOffset(x, y)] = 235
		}
	}

	for i := 0; i < size; i++ {
		img.Pix[img.PixOffset(i, i)] = uint8(70 + (i*3)%70)
		img.Pix[img.PixOffset(size-1-i, i)] = uint8(180 - (i*5)%70)
	}

	switch variant {
	case "vector":
		// Hard edges / flat fills similar to vector iconography.
		setRect(img, size/8, size/8, size*3/4, size/6, 28)
		setRect(img, size/7, size/2, size*5/7, size/7, 210)
		for i := 0; i < size; i += maxInt(3, size/14) {
			setRect(img, i, (i*3)%size, maxInt(1, size/28), maxInt(1, size/16), 245)
			setRect(img, size-1-i, (i*5)%size, maxInt(1, size/30), maxInt(1, size/20), 15)
		}
		setRect(img, size/3, size/3, size/5, size/5, 128)
	case "photo":
		// Smooth gradients and textured patches approximating photographic content.
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				gx := (x * 255) / maxInt(1, size-1)
				gy := (y * 255) / maxInt(1, size-1)
				v := (gx*3 + gy*2) / 5
				noise := (x*41 + y*29 + (x*y)%97 + size*13) & 31
				img.Pix[img.PixOffset(x, y)] = uint8(minInt(255, maxInt(0, v+noise-16)))
			}
		}
		setRect(img, size/6, size/6, size/4, size/3, 180)
		setRect(img, size/2, size/3, size/3, size/4, 70)
	case "ui":
		// Panels, controls, and text-like bars similar to UI screenshots.
		setRect(img, size/12, size/12, size*10/12, size*10/12, 235)
		setRect(img, size/8, size/6, size*3/4, size/8, 190)
		setRect(img, size/8, size*2/5, size/3, size/10, 120)
		setRect(img, size/2, size*2/5, size*3/8, size/10, 140)
		for i := 0; i < 5; i++ {
			y := size*3/5 + i*maxInt(2, size/16)
			setRect(img, size/6, y, size*2/3, maxInt(1, size/40), uint8(70+i*12))
		}
	case "glyph":
		step := maxInt(6, size/6)
		for y := step; y < size-step; y += step {
			for x := step; x < size-step; x += step {
				if (x+y)/step%2 == 0 {
					setRect(img, x-1, y-1, 3, 3, 20)
				} else {
					setRect(img, x-1, y-1, 3, 3, 245)
				}
			}
		}
		setRect(img, size/4, size/4, size/2, maxInt(3, size/10), 30)
		setRect(img, size/3, size/2, size/3, maxInt(3, size/10), 210)
	case "noise":
		for y := 1; y < size-1; y++ {
			for x := 1; x < size-1; x++ {
				v := (x*73 + y*151 + (x*y)%251 + size*19) & 0xFF
				img.Pix[img.PixOffset(x, y)] = uint8(v)
			}
		}
		setRect(img, size/8, size/8, size/3, size/20+2, 0)
		setRect(img, size/2, size/3, size/4, size/18+2, 255)
		setRect(img, size/3, size/2, size/20+2, size/3, 0)
	case "orbtex":
		for y := 1; y < size-1; y++ {
			for x := 1; x < size-1; x++ {
				v := (x*97 + y*193 + (x*y)%239 + size*41) & 0xFF
				if ((x+y)%7 == 0) || ((x*3+y*5)%11 == 0) {
					v = 255 - v
				}
				img.Pix[img.PixOffset(x, y)] = uint8(v)
			}
		}
		ringStep := maxInt(10, size/8)
		for r := ringStep; r < size-ringStep; r += ringStep {
			setRect(img, r, r, size-r*2, 2, uint8((r*37)&0xFF))
			setRect(img, r, r, 2, size-r*2, uint8((r*59)&0xFF))
			setRect(img, size-r-2, r, 2, size-r*2, uint8((r*83)&0xFF))
			setRect(img, r, size-r-2, size-r*2, 2, uint8((r*97)&0xFF))
		}
		for i := 0; i < size; i += 3 {
			img.Pix[img.PixOffset(i, (i*5)%size)] = 0
			img.Pix[img.PixOffset((i*7)%size, i)] = 255
		}
	default:
		for y := 2; y < size-2; y++ {
			for x := 2; x < size-2; x++ {
				if (x/4+y/4)%2 == 0 {
					img.Pix[img.PixOffset(x, y)] = 55
				}
			}
		}
		setRect(img, size/5, size/5, size/2, size/2, 180)
		setRect(img, size/2, size/3, size/4, size/4, 25)
	}

	seed := 43
	if variant == "vector" {
		seed = 67
	} else if variant == "photo" {
		seed = 131
	} else if variant == "ui" {
		seed = 173
	} else if variant == "glyph" {
		seed = 97
	} else if variant == "noise" {
		seed = 157
	} else if variant == "orbtex" {
		seed = 211
	}
	applyFeatureTexture(img, seed)

	return img
}

func buildBenchHaystack(w, h int, scenario findBenchScenario) *image.Gray {
	if w < 64 {
		w = 64
	}
	if h < 64 {
		h = 64
	}
	img := image.NewGray(image.Rect(0, 0, w, h))
	variant := scenario.variant
	seed := scenario.backgroundTextureSeed
	if seed == 0 {
		seed = 37
		if variant == "vector" {
			seed = 61
		} else if variant == "photo" {
			seed = 119
		} else if variant == "ui" {
			seed = 181
		} else if variant == "glyph" {
			seed = 91
		} else if variant == "noise" {
			seed = 149
		} else if variant == "orbtex" {
			seed = 227
		}
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			v := uint8(90 + ((x*13 + y*7 + seed) % 41))
			if scenario.backgroundPalette == "ui_dark" {
				v = uint8(maxInt(20, int(v)-55))
			} else if scenario.backgroundPalette == "ui_light" {
				v = uint8(minInt(245, int(v)+40))
			}
			img.Pix[img.PixOffset(x, y)] = v
		}
	}
	applyDenseArtifactField(img, seed, variant, scenario.backgroundClutterDensity, scenario.backgroundContinuousCanvas)
	return img
}

func populateNearMatchDecoys(haystack *image.Gray, target *image.Gray, targetX, targetY int, scenario findBenchScenario) {
	if haystack == nil || target == nil {
		return
	}
	if !scenario.decoyEnabled {
		return
	}
	hb := haystack.Bounds()
	tb := target.Bounds()
	tw := tb.Dx()
	th := tb.Dy()
	if tw <= 0 || th <= 0 {
		return
	}

	decoyCount := scenario.decoyCount
	if decoyCount <= 0 {
		decoyCount = maxInt(8, (scenario.screenW*scenario.screenH)/maxInt(1, tw*th*2))
	}
	if decoyCount > 400 {
		decoyCount = 400
	}
	similarity := scenario.decoySimilarity
	if similarity <= 0 || similarity > 1 {
		similarity = 0.90
	}
	placement := strings.ToLower(strings.TrimSpace(scenario.decoyPlacement))
	if placement == "" {
		placement = "grid"
	}

	seed := hashKey(fmt.Sprintf("%d:%s:%s:%dx%d", scenario.seed, scenario.name, placement, tw, th))
	next := func(max int) int {
		seed = seed*6364136223846793005 + 1442695040888963407
		if max <= 0 {
			return 0
		}
		return int(seed % uint64(max))
	}

	gridCols := maxInt(1, int(math.Ceil(math.Sqrt(float64(decoyCount)))))
	gridRows := maxInt(1, int(math.Ceil(float64(decoyCount)/float64(gridCols))))
	stepX := maxInt(4, (hb.Dx()-tw)/maxInt(1, gridCols))
	stepY := maxInt(4, (hb.Dy()-th)/maxInt(1, gridRows))

	clusterCount := maxInt(2, minInt(6, decoyCount/8))
	clusterCenters := make([]image.Point, 0, clusterCount)
	for i := 0; i < clusterCount; i++ {
		cx := hb.Min.X + next(maxInt(1, hb.Dx()-tw))
		cy := hb.Min.Y + next(maxInt(1, hb.Dy()-th))
		clusterCenters = append(clusterCenters, image.Point{X: cx, Y: cy})
	}

	placed := 0
	maxAttempts := decoyCount * 8
	for attempt := 0; attempt < maxAttempts && placed < decoyCount; attempt++ {
		var px, py int
		switch placement {
		case "random":
			px = hb.Min.X + next(maxInt(1, hb.Dx()-tw))
			py = hb.Min.Y + next(maxInt(1, hb.Dy()-th))
		case "clustered":
			center := clusterCenters[next(len(clusterCenters))]
			px = center.X + (next(maxInt(2, tw)) - tw/2)
			py = center.Y + (next(maxInt(2, th)) - th/2)
		case "mixed":
			switch attempt % 3 {
			case 0:
				px = hb.Min.X + next(maxInt(1, hb.Dx()-tw))
				py = hb.Min.Y + next(maxInt(1, hb.Dy()-th))
			case 1:
				col := placed % gridCols
				row := placed / gridCols
				px = hb.Min.X + col*stepX + (next(maxInt(2, tw/3)) - tw/6)
				py = hb.Min.Y + row*stepY + (next(maxInt(2, th/3)) - th/6)
			default:
				center := clusterCenters[next(len(clusterCenters))]
				px = center.X + (next(maxInt(2, tw)) - tw/2)
				py = center.Y + (next(maxInt(2, th)) - th/2)
			}
		default: // grid
			col := placed % gridCols
			row := placed / gridCols
			px = hb.Min.X + col*stepX + (next(maxInt(2, tw/4)) - tw/8)
			py = hb.Min.Y + row*stepY + (next(maxInt(2, th/4)) - th/8)
		}

		if px < hb.Min.X || py < hb.Min.Y || px+tw >= hb.Max.X || py+th >= hb.Max.Y {
			continue
		}
		if rectsOverlapInt(px, py, tw, th, targetX, targetY, tw, th) {
			continue
		}

		decoy := makeNearMatchVariant(target, int(seed&0x7fffffff), similarity)
		blitGrayFeather(haystack, decoy, px, py, maxInt(2, minInt(tw, th)/10))
		placed++
	}
}

func makeNearMatchVariant(src *image.Gray, seed int, similarity float64) *image.Gray {
	dst := cloneGray(src)
	b := dst.Bounds()
	w := b.Dx()
	h := b.Dy()
	total := maxInt(1, w*h)
	similarity = clampFloat(similarity, 0.50, 0.999)
	mutationScale := 1.0 - similarity

	// Spread subtle brightness deltas based on similarity (higher similarity => fewer mutations).
	mutationCount := maxInt(1, int(float64(total)*(0.03+0.22*mutationScale)))
	mutationCount += seed % maxInt(1, total/25+1)
	for i := 0; i < mutationCount; i++ {
		x := (i*73 + seed*19 + (i*i)%137) % w
		y := (i*97 + seed*11 + (i*i)%149) % h
		maxDelta := maxInt(2, int(18*mutationScale))
		delta := ((i*17 + seed*7) % (maxDelta*2 + 1)) - maxDelta
		off := dst.PixOffset(b.Min.X+x, b.Min.Y+y)
		v := int(dst.Pix[off]) + delta
		if v < 0 {
			v = 0
		} else if v > 255 {
			v = 255
		}
		dst.Pix[off] = uint8(v)
	}

	// Add faint line cuts that remain subtle but reduce exact identity.
	if w > 6 && h > 6 {
		lineY := 1 + ((seed*13 + w) % maxInt(1, h-2))
		lineX := 1 + ((seed*17 + h) % maxInt(1, w-2))
		for x := 1; x < w-1; x++ {
			off := dst.PixOffset(b.Min.X+x, b.Min.Y+lineY)
			delta := int(float64(((x+seed)%7)-3) * (0.5 + mutationScale))
			v := int(dst.Pix[off]) + delta
			if v < 0 {
				v = 0
			} else if v > 255 {
				v = 255
			}
			dst.Pix[off] = uint8(v)
		}
		for y := 1; y < h-1; y++ {
			off := dst.PixOffset(b.Min.X+lineX, b.Min.Y+y)
			delta := int(float64(((y+seed)%7)-3) * (0.5 + mutationScale))
			v := int(dst.Pix[off]) + delta
			if v < 0 {
				v = 0
			} else if v > 255 {
				v = 255
			}
			dst.Pix[off] = uint8(v)
		}
	}

	return dst
}

func rectsOverlapInt(ax, ay, aw, ah, bx, by, bw, bh int) bool {
	ax1, ay1 := ax+aw, ay+ah
	bx1, by1 := bx+bw, by+bh
	return ax < bx1 && ax1 > bx && ay < by1 && ay1 > by
}

func applyDenseArtifactField(img *image.Gray, seed int, variant string, clutterDensity float64, continuousCanvas bool) {
	if img == nil {
		return
	}
	if clutterDensity <= 0 {
		clutterDensity = 0.75
	}
	if clutterDensity > 1 {
		clutterDensity = 1
	}
	artifactGate := int((1.0 - clutterDensity) * 7.0)
	if artifactGate < 0 {
		artifactGate = 0
	}

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			off := img.PixOffset(x, y)
			v := int(img.Pix[off])
			hash := (x*29 + y*41 + seed*17 + (x*y)%233) & 0xFF
			if artifactGate > 0 && (hash%8) < artifactGate {
				continue
			}
			delta := (hash % 9) - 4
			if variant == "noise" || variant == "orbtex" || variant == "photo" {
				delta += (hash % 7) - 3
			}
			if variant == "ui" || variant == "vector" {
				delta = (hash % 5) - 2
			}
			v += delta
			if v < 0 {
				v = 0
			} else if v > 255 {
				v = 255
			}
			img.Pix[off] = uint8(v)
		}
	}

	spacing := 8
	if variant == "noise" || variant == "orbtex" {
		spacing = 6
	} else if variant == "photo" {
		spacing = 10
	} else if variant == "ui" || variant == "vector" {
		spacing = 12
	}
	for y := b.Min.Y + 2; y < b.Max.Y-2; y += spacing {
		for x := b.Min.X + 2; x < b.Max.X-2; x += spacing {
			base := uint8((x*11 + y*13 + seed*23) & 0xFF)
			if !continuousCanvas {
				if (x+y+seed)%2 == 0 {
					continue
				}
			}
			img.Pix[img.PixOffset(x, y)] = base
			img.Pix[img.PixOffset(x+1, y)] = 255 - base/2
			img.Pix[img.PixOffset(x, y+1)] = base / 2
			if (x+y+seed)%3 == 0 {
				img.Pix[img.PixOffset(x-1, y)] = 255 - base
			}
		}
	}
}

func blitGray(dst *image.Gray, src *image.Gray, atX, atY int) {
	db := dst.Bounds()
	sb := src.Bounds()
	for y := 0; y < sb.Dy(); y++ {
		for x := 0; x < sb.Dx(); x++ {
			dx := atX + x
			dy := atY + y
			if dx < db.Min.X || dy < db.Min.Y || dx >= db.Max.X || dy >= db.Max.Y {
				continue
			}
			dst.Pix[dst.PixOffset(dx, dy)] = src.Pix[src.PixOffset(sb.Min.X+x, sb.Min.Y+y)]
		}
	}
}

func blitGrayFeather(dst *image.Gray, src *image.Gray, atX, atY, feather int) {
	if feather < 1 {
		blitGray(dst, src, atX, atY)
		return
	}
	db := dst.Bounds()
	sb := src.Bounds()
	sw := sb.Dx()
	sh := sb.Dy()
	for y := 0; y < sh; y++ {
		for x := 0; x < sw; x++ {
			dx := atX + x
			dy := atY + y
			if dx < db.Min.X || dy < db.Min.Y || dx >= db.Max.X || dy >= db.Max.Y {
				continue
			}
			srcV := int(src.Pix[src.PixOffset(sb.Min.X+x, sb.Min.Y+y)])
			dstOff := dst.PixOffset(dx, dy)
			dstV := int(dst.Pix[dstOff])

			distL := x
			distR := sw - 1 - x
			distT := y
			distB := sh - 1 - y
			edgeDist := minInt(minInt(distL, distR), minInt(distT, distB))
			if edgeDist >= feather {
				dst.Pix[dstOff] = uint8(srcV)
				continue
			}
			alphaNum := edgeDist + 1
			alphaDen := feather + 1
			blended := (dstV*(alphaDen-alphaNum) + srcV*alphaNum) / alphaDen
			if blended < 0 {
				blended = 0
			} else if blended > 255 {
				blended = 255
			}
			dst.Pix[dstOff] = uint8(blended)
		}
	}
}

func applySeamSmoothing(img *image.Gray, variant string) {
	if img == nil {
		return
	}
	passes := 1
	if variant == "noise" || variant == "orbtex" {
		passes = 2
	}
	for pass := 0; pass < passes; pass++ {
		src := cloneGray(img)
		b := src.Bounds()
		for y := b.Min.Y + 1; y < b.Max.Y-1; y++ {
			for x := b.Min.X + 1; x < b.Max.X-1; x++ {
				c := int(src.Pix[src.PixOffset(x, y)])
				u := int(src.Pix[src.PixOffset(x, y-1)])
				d := int(src.Pix[src.PixOffset(x, y+1)])
				l := int(src.Pix[src.PixOffset(x-1, y)])
				r := int(src.Pix[src.PixOffset(x+1, y)])
				mix := (c*5 + u + d + l + r) / 9
				img.Pix[img.PixOffset(x, y)] = uint8(mix)
			}
		}
	}
}

func applyTargetOcclusion(img *image.Gray, x, y, w, h int, scenario findBenchScenario) {
	if img == nil || !scenario.occlusionEnabled || scenario.occlusionCoverage <= 0 {
		return
	}
	coverage := clampFloat(scenario.occlusionCoverage, 0.0, 0.95)
	targetArea := maxInt(1, w*h)
	occArea := maxInt(1, int(float64(targetArea)*coverage))

	seed := hashKey(fmt.Sprintf("occ:%d:%s:%d:%d:%d:%d", scenario.seed, scenario.name, x, y, w, h))
	next := func(max int) int {
		seed = seed*6364136223846793005 + 1442695040888963407
		if max <= 0 {
			return 0
		}
		return int(seed % uint64(max))
	}

	remaining := occArea
	for remaining > 0 {
		blockW := maxInt(2, minInt(w, int(math.Sqrt(float64(remaining)))))
		blockH := maxInt(2, minInt(h, maxInt(2, remaining/maxInt(1, blockW))))
		blockX := x + next(maxInt(1, w-blockW+1))
		blockY := y + next(maxInt(1, h-blockH+1))
		fill := uint8(64 + next(128))
		setRect(img, blockX, blockY, blockW, blockH, fill)
		remaining -= blockW * blockH
	}
}

func applyPhotometricProfile(img *image.Gray, scenario findBenchScenario, isCapture bool) {
	if img == nil {
		return
	}
	brightness := scenario.brightnessDelta
	contrast := scenario.contrastFactor
	gamma := scenario.gammaFactor
	if contrast <= 0 {
		contrast = 1.0
	}
	if gamma <= 0 {
		gamma = 1.0
	}

	b := img.Bounds()
	seed := hashKey(fmt.Sprintf("photo:%d:%s:%t", scenario.seed, scenario.name, isCapture))
	nextUnit := func() float64 {
		seed = seed*6364136223846793005 + 1442695040888963407
		return float64(seed%10000) / 9999.0
	}
	nextSigned := func() float64 {
		return nextUnit()*2.0 - 1.0
	}

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			off := img.PixOffset(x, y)
			v := float64(img.Pix[off])

			v = (v-128.0)*contrast + 128.0 + brightness*255.0
			if gamma != 1.0 {
				n := clampFloat(v/255.0, 0.0, 1.0)
				v = math.Pow(n, 1.0/gamma) * 255.0
			}

			if scenario.noiseGaussian > 0 {
				v += nextSigned() * 255.0 * scenario.noiseGaussian * 0.25
			}
			if scenario.noisePoisson > 0 {
				v += (math.Sqrt(clampFloat(v, 0.0, 255.0)) / 16.0) * nextSigned() * 255.0 * scenario.noisePoisson * 0.15
			}
			if scenario.noiseSaltPepper > 0 && nextUnit() < scenario.noiseSaltPepper*0.02 {
				if nextUnit() < 0.5 {
					v = 0
				} else {
					v = 255
				}
			}
			if scenario.noiseBanding > 0 {
				band := math.Sin(float64(y)*0.2+float64(x)*0.03) * 20.0 * scenario.noiseBanding
				v += band
			}
			if scenario.noiseCompression > 0 {
				q := maxInt(2, int(2.0+scenario.noiseCompression*18.0))
				v = math.Round(v/float64(q)) * float64(q)
			}

			if v < 0 {
				v = 0
			} else if v > 255 {
				v = 255
			}
			img.Pix[off] = uint8(v)
		}
	}

	if scenario.blurSigma > 0 {
		passes := maxInt(1, int(math.Round(scenario.blurSigma)))
		for i := 0; i < passes; i++ {
			applySeamSmoothing(img, scenario.variant)
		}
	}
}

func applyMonitorProfile(img *image.Gray, scenario findBenchScenario) {
	if img == nil {
		return
	}
	if scenario.monitorGamma == 0 && scenario.monitorSharpness == 0 && scenario.monitorColorShift == 0 {
		return
	}
	gamma := scenario.monitorGamma
	if gamma <= 0 {
		gamma = 1.0
	}
	shift := clampFloat(float64(scenario.monitorColorShift)/2000.0, -1.0, 1.0) * 18.0
	sharpness := scenario.monitorSharpness
	if sharpness <= 0 {
		sharpness = 1.0
	}

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			off := img.PixOffset(x, y)
			v := float64(img.Pix[off]) + shift
			n := clampFloat(v/255.0, 0.0, 1.0)
			v = math.Pow(n, 1.0/gamma) * 255.0
			if v < 0 {
				v = 0
			} else if v > 255 {
				v = 255
			}
			img.Pix[off] = uint8(v)
		}
	}
	if sharpness < 1.0 {
		passes := maxInt(1, int(math.Round((1.0-sharpness)*3.0)))
		for i := 0; i < passes; i++ {
			applySeamSmoothing(img, scenario.variant)
		}
	}
}

func cropGray(src *image.Gray, r image.Rectangle) *image.Gray {
	b := src.Bounds()
	c := r.Intersect(b)
	if c.Empty() {
		return image.NewGray(image.Rect(0, 0, 1, 1))
	}
	out := image.NewGray(image.Rect(0, 0, c.Dx(), c.Dy()))
	for y := 0; y < c.Dy(); y++ {
		srcStart := src.PixOffset(c.Min.X, c.Min.Y+y)
		srcEnd := srcStart + c.Dx()
		dstStart := y * out.Stride
		copy(out.Pix[dstStart:dstStart+c.Dx()], src.Pix[srcStart:srcEnd])
	}
	return out
}

func rotateGrayByQuarterTurns(src *image.Gray, degrees int) *image.Gray {
	turns := ((degrees % 360) + 360) % 360 / 90
	switch turns {
	case 0:
		return cloneGray(src)
	case 1:
		return rotate90Gray(src)
	case 2:
		return rotate180Gray(src)
	case 3:
		return rotate270Gray(src)
	default:
		return cloneGray(src)
	}
}

func applyBenchTransform(src *image.Gray, scenario findBenchScenario) *image.Gray {
	switch scenario.transformKind {
	case "":
		return cloneGray(src)
	case "scale":
		factor := scenario.transformA
		if factor <= 0 {
			factor = 1.0
		}
		return scaleGrayNearestBench(src, factor)
	case "rotate":
		return rotateGrayBilinearBench(src, scenario.transformA, 128)
	case "perspective":
		topScale := scenario.transformA
		bottomScale := scenario.transformB
		shift := scenario.transformC
		if topScale <= 0 {
			topScale = 0.90
		}
		if bottomScale <= 0 {
			bottomScale = 1.08
		}
		return perspectiveKeystoneBench(src, topScale, bottomScale, shift, 128)
	case "skewx":
		return skewGrayXBench(src, scenario.transformA, 128)
	default:
		return cloneGray(src)
	}
}

func scaleGrayNearestBench(src *image.Gray, factor float64) *image.Gray {
	if factor <= 0 {
		factor = 1.0
	}
	sb := src.Bounds()
	sw, sh := sb.Dx(), sb.Dy()
	dw := maxInt(1, int(math.Round(float64(sw)*factor)))
	dh := maxInt(1, int(math.Round(float64(sh)*factor)))
	dst := image.NewGray(image.Rect(0, 0, dw, dh))
	for y := 0; y < dh; y++ {
		sy := sb.Min.Y + minInt(sh-1, int(float64(y)/factor))
		for x := 0; x < dw; x++ {
			sx := sb.Min.X + minInt(sw-1, int(float64(x)/factor))
			dst.SetGray(x, y, src.GrayAt(sx, sy))
		}
	}
	return dst
}

func rotateGrayBilinearBench(src *image.Gray, degrees float64, bg uint8) *image.Gray {
	sb := src.Bounds()
	sw, sh := sb.Dx(), sb.Dy()
	if sw <= 0 || sh <= 0 {
		return image.NewGray(image.Rect(0, 0, 1, 1))
	}
	theta := degrees * math.Pi / 180.0
	cosT := math.Cos(theta)
	sinT := math.Sin(theta)
	cx := (float64(sw) - 1) / 2.0
	cy := (float64(sh) - 1) / 2.0

	corners := [][2]float64{
		{-cx, -cy},
		{float64(sw-1) - cx, -cy},
		{float64(sw-1) - cx, float64(sh-1) - cy},
		{-cx, float64(sh-1) - cy},
	}
	minX, maxX := math.MaxFloat64, -math.MaxFloat64
	minY, maxY := math.MaxFloat64, -math.MaxFloat64
	for _, c := range corners {
		x := cosT*c[0] - sinT*c[1]
		y := sinT*c[0] + cosT*c[1]
		minX = math.Min(minX, x)
		maxX = math.Max(maxX, x)
		minY = math.Min(minY, y)
		maxY = math.Max(maxY, y)
	}
	dw := maxInt(1, int(math.Ceil(maxX-minX))+1)
	dh := maxInt(1, int(math.Ceil(maxY-minY))+1)
	dst := image.NewGray(image.Rect(0, 0, dw, dh))
	for y := 0; y < dh; y++ {
		for x := 0; x < dw; x++ {
			dxr := float64(x) + minX
			dyr := float64(y) + minY
			sxr := cosT*dxr + sinT*dyr
			syr := -sinT*dxr + cosT*dyr
			sx := sxr + cx
			sy := syr + cy
			dst.SetGray(x, y, color.Gray{Y: sampleGrayBilinearBench(src, sx, sy, bg)})
		}
	}
	return dst
}

func skewGrayXBench(src *image.Gray, skew float64, bg uint8) *image.Gray {
	sb := src.Bounds()
	sw, sh := sb.Dx(), sb.Dy()
	cx := (float64(sw) - 1) / 2.0
	cy := (float64(sh) - 1) / 2.0

	corners := [][2]float64{
		{-cx + skew*(-cy), -cy},
		{float64(sw-1) - cx + skew*(-cy), -cy},
		{float64(sw-1) - cx + skew*(float64(sh-1)-cy), float64(sh-1) - cy},
		{-cx + skew*(float64(sh-1)-cy), float64(sh-1) - cy},
	}
	minX, maxX := math.MaxFloat64, -math.MaxFloat64
	for _, c := range corners {
		minX = math.Min(minX, c[0])
		maxX = math.Max(maxX, c[0])
	}
	dw := maxInt(1, int(math.Ceil(maxX-minX))+1)
	dh := sh
	dst := image.NewGray(image.Rect(0, 0, dw, dh))
	for y := 0; y < dh; y++ {
		yr := float64(y) - cy
		for x := 0; x < dw; x++ {
			xr := float64(x) + minX
			sxr := xr - skew*yr
			sx := sxr + cx
			sy := yr + cy
			dst.SetGray(x, y, color.Gray{Y: sampleGrayBilinearBench(src, sx, sy, bg)})
		}
	}
	return dst
}

func perspectiveKeystoneBench(src *image.Gray, topScale, bottomScale, shift float64, bg uint8) *image.Gray {
	sb := src.Bounds()
	sw, sh := sb.Dx(), sb.Dy()
	dst := image.NewGray(image.Rect(0, 0, sw, sh))
	cx := (float64(sw) - 1) / 2.0
	for y := 0; y < sh; y++ {
		t := 0.0
		if sh > 1 {
			t = float64(y) / float64(sh-1)
		}
		scale := topScale*(1-t) + bottomScale*t
		shiftX := shift * (0.5 - t) * float64(sw)
		for x := 0; x < sw; x++ {
			sx := (float64(x)-cx-shiftX)/scale + cx
			sy := float64(y)
			dst.SetGray(x, y, color.Gray{Y: sampleGrayBilinearBench(src, sx, sy, bg)})
		}
	}
	return dst
}

func sampleGrayBilinearBench(src *image.Gray, fx, fy float64, bg uint8) uint8 {
	sb := src.Bounds()
	if fx < float64(sb.Min.X) || fy < float64(sb.Min.Y) || fx > float64(sb.Max.X-1) || fy > float64(sb.Max.Y-1) {
		return bg
	}
	x0 := int(math.Floor(fx))
	y0 := int(math.Floor(fy))
	x1 := minInt(sb.Max.X-1, x0+1)
	y1 := minInt(sb.Max.Y-1, y0+1)
	dx := fx - float64(x0)
	dy := fy - float64(y0)
	p00 := float64(src.GrayAt(x0, y0).Y)
	p10 := float64(src.GrayAt(x1, y0).Y)
	p01 := float64(src.GrayAt(x0, y1).Y)
	p11 := float64(src.GrayAt(x1, y1).Y)
	top := p00*(1-dx) + p10*dx
	bot := p01*(1-dx) + p11*dx
	v := int(math.Round(top*(1-dy) + bot*dy))
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return uint8(v)
}

func rotate90Gray(src *image.Gray) *image.Gray {
	sb := src.Bounds()
	dst := image.NewGray(image.Rect(0, 0, sb.Dy(), sb.Dx()))
	for y := 0; y < sb.Dy(); y++ {
		for x := 0; x < sb.Dx(); x++ {
			dx := sb.Dy() - 1 - y
			dy := x
			dst.Pix[dst.PixOffset(dx, dy)] = src.Pix[src.PixOffset(sb.Min.X+x, sb.Min.Y+y)]
		}
	}
	return dst
}

func rotate180Gray(src *image.Gray) *image.Gray {
	sb := src.Bounds()
	dst := image.NewGray(image.Rect(0, 0, sb.Dx(), sb.Dy()))
	for y := 0; y < sb.Dy(); y++ {
		for x := 0; x < sb.Dx(); x++ {
			dx := sb.Dx() - 1 - x
			dy := sb.Dy() - 1 - y
			dst.Pix[dst.PixOffset(dx, dy)] = src.Pix[src.PixOffset(sb.Min.X+x, sb.Min.Y+y)]
		}
	}
	return dst
}

func rotate270Gray(src *image.Gray) *image.Gray {
	sb := src.Bounds()
	dst := image.NewGray(image.Rect(0, 0, sb.Dy(), sb.Dx()))
	for y := 0; y < sb.Dy(); y++ {
		for x := 0; x < sb.Dx(); x++ {
			dx := y
			dy := sb.Dx() - 1 - x
			dst.Pix[dst.PixOffset(dx, dy)] = src.Pix[src.PixOffset(sb.Min.X+x, sb.Min.Y+y)]
		}
	}
	return dst
}

func cloneGray(src *image.Gray) *image.Gray {
	sb := src.Bounds()
	dst := image.NewGray(image.Rect(0, 0, sb.Dx(), sb.Dy()))
	for y := 0; y < sb.Dy(); y++ {
		copy(dst.Pix[y*dst.Stride:y*dst.Stride+sb.Dx()], src.Pix[(sb.Min.Y+y)*src.Stride+sb.Min.X:(sb.Min.Y+y)*src.Stride+sb.Min.X+sb.Dx()])
	}
	return dst
}

func grayProtoFromGray(name string, in *image.Gray) *pb.GrayImage {
	b := in.Bounds()
	pix := make([]byte, 0, b.Dx()*b.Dy())
	for y := 0; y < b.Dy(); y++ {
		rowStart := (b.Min.Y+y)*in.Stride + b.Min.X
		pix = append(pix, in.Pix[rowStart:rowStart+b.Dx()]...)
	}
	return &pb.GrayImage{
		Name:   name,
		Width:  int32(b.Dx()),
		Height: int32(b.Dy()),
		Pix:    pix,
	}
}

func rectMatchSatisfies(got *pb.Rect, want *pb.Rect, minOverlap float64, maxAreaRatio float64) bool {
	if maxAreaRatio <= 0 {
		maxAreaRatio = 1.50
	}
	if rectAreaRatio(got, want) > maxAreaRatio {
		return false
	}
	return rectOverlapRatio(got, want) >= math.Max(0.0, math.Min(1.0, minOverlap))
}

func rectAreaRatio(got *pb.Rect, want *pb.Rect) float64 {
	gotArea := float64(max32(1, got.GetW()*got.GetH()))
	wantArea := float64(max32(1, want.GetW()*want.GetH()))
	return gotArea / wantArea
}

func rectOverlapRatio(got *pb.Rect, want *pb.Rect) float64 {
	interArea := rectIntersectionArea(got, want)
	wantArea := float64(max32(1, want.GetW()*want.GetH()))
	return interArea / wantArea
}

func rectIntersectionArea(a *pb.Rect, b *pb.Rect) float64 {
	ax0, ay0 := a.GetX(), a.GetY()
	ax1, ay1 := ax0+a.GetW(), ay0+a.GetH()
	bx0, by0 := b.GetX(), b.GetY()
	bx1, by1 := bx0+b.GetW(), by0+b.GetH()

	ix0 := max32(ax0, bx0)
	iy0 := max32(ay0, by0)
	ix1 := min32(ax1, bx1)
	iy1 := min32(ay1, by1)
	if ix1 <= ix0 || iy1 <= iy0 {
		return 0
	}
	return float64((ix1 - ix0) * (iy1 - iy0))
}

func rectAreaFloat(r *pb.Rect) float64 {
	if r == nil {
		return 0
	}
	return float64(max32(1, r.GetW()*r.GetH()))
}

func rectContainsPoint(r *pb.Rect, x, y float64) bool {
	if r == nil {
		return false
	}
	left := float64(r.GetX())
	top := float64(r.GetY())
	right := left + float64(r.GetW())
	bottom := top + float64(r.GetH())
	return x >= left && x <= right && y >= top && y <= bottom
}

func setRect(img *image.Gray, x, y, w, h int, value uint8) {
	b := img.Bounds()
	for yy := y; yy < y+h; yy++ {
		if yy < b.Min.Y || yy >= b.Max.Y {
			continue
		}
		for xx := x; xx < x+w; xx++ {
			if xx < b.Min.X || xx >= b.Max.X {
				continue
			}
			img.Pix[img.PixOffset(xx, yy)] = value
		}
	}
}

func applyFeatureTexture(img *image.Gray, seed int) {
	b := img.Bounds()
	for y := 2; y < b.Dy()-2; y++ {
		for x := 2; x < b.Dx()-2; x++ {
			base := int(img.Pix[img.PixOffset(x, y)])
			hash := (x*131 + y*197 + seed*53 + (x*y)%251) & 0xFF
			mix := (base*3 + hash) / 4
			if ((x+y+seed)%5 == 0) || ((x*7+y*11+seed)%13 == 0) {
				mix = (base + hash) / 2
			}
			img.Pix[img.PixOffset(x, y)] = uint8(mix)
		}
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func min32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func benchSimilarityPtr(engine string) *float64 {
	v := 0.99
	switch engine {
	case "orb", "akaze", "brisk", "kaze", "sift":
		v = 0.10
	}
	return &v
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func benchEnvEnabled(name string, defaultValue bool) bool {
	raw := strings.TrimSpace(strings.ToLower(os.Getenv(name)))
	if raw == "" {
		return defaultValue
	}
	switch raw {
	case "1", "true", "yes", "on":
		return true
	case "0", "false", "no", "off":
		return false
	default:
		return defaultValue
	}
}
