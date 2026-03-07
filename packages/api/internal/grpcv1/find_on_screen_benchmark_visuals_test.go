package grpcv1

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type findBenchVisualCollector struct {
	outDir             string
	maxAttempts        int
	rpcTimeout         time.Duration
	attemptOverlay     bool
	summaryNative      bool
	summaryShowPattern bool

	mu        sync.Mutex
	scenarios map[string]*findBenchScenarioVisual
}

type findBenchScenarioVisual struct {
	scenario findBenchScenario
	engines  map[string][]findBenchAttemptVisual
	patterns []findBenchPatternVisual
}

type findBenchPatternVisual struct {
	Label string
	Image *image.Gray
}

type findBenchScenarioSummaryImage struct {
	ScenarioName string
	Path         string
}

var (
	scenarioResolutionPattern = regexp.MustCompile(`^(.*)_([0-9]+x[0-9]+)_.+$`)
	scenarioStablePathPattern = regexp.MustCompile(`^(.+_[0-9]+x[0-9]+_i[0-9]+)(?:_[a-z0-9]+)?$`)
	scenarioSeedSuffixPattern = regexp.MustCompile(`_(?:s)?[0-9a-fA-F]{8,}$`)
)

type findBenchVisualQuery struct {
	ID         string
	Request    *pb.FindOnScreenRequest
	Expected   *pb.Rect
	Alternates []*pb.Rect
	Label      string
}

type findBenchAttemptQueryVisual struct {
	ID         string
	Label      string
	Status     string
	Error      string
	Explain    string
	Overlap    float64
	AreaRatio  float64
	Expected   *pb.Rect
	Alternates []*pb.Rect
	Found      *pb.Rect
	Template   *image.Gray
}

type findBenchAttemptVisual struct {
	Attempt  int
	Duration time.Duration
	Retried  bool
	Status   string
	Error    string
	Queries  []findBenchAttemptQueryVisual
	File     string
}

func newFindBenchVisualCollectorFromEnv(t testing.TB) *findBenchVisualCollector {
	t.Helper()
	if !envFlagTrue(os.Getenv("FIND_BENCH_VISUAL")) {
		return nil
	}

	outDir := strings.TrimSpace(os.Getenv("FIND_BENCH_VISUAL_DIR"))
	if outDir == "" {
		outDir = ".test-results/bench/visuals"
	}
	maxAttempts := parseEnvInt("FIND_BENCH_VISUAL_MAX_ATTEMPTS", 2)
	if maxAttempts < 1 {
		maxAttempts = 1
	}
	rpcTimeout := parseEnvDuration("FIND_BENCH_VISUAL_TIMEOUT", 5*time.Second)
	attemptOverlay := true
	if raw := strings.TrimSpace(os.Getenv("FIND_BENCH_VISUAL_ATTEMPT_OVERLAY")); raw != "" {
		attemptOverlay = envFlagTrue(raw)
	}
	summaryNative := envFlagTrue(os.Getenv("FIND_BENCH_VISUAL_SUMMARY_NATIVE"))
	summaryShowPattern := envFlagTrue(os.Getenv("FIND_BENCH_VISUAL_SUMMARY_SHOW_PATTERN"))

	attemptsDir := filepath.Join(outDir, "attempts")
	summariesDir := filepath.Join(outDir, "summaries")
	if err := os.RemoveAll(attemptsDir); err != nil {
		t.Fatalf("reset benchmark visual attempts directory: %v", err)
	}
	if err := os.RemoveAll(summariesDir); err != nil {
		t.Fatalf("reset benchmark visual summary directory: %v", err)
	}
	if err := os.MkdirAll(attemptsDir, 0o755); err != nil {
		t.Fatalf("create benchmark visual attempts directory: %v", err)
	}
	if err := os.MkdirAll(summariesDir, 0o755); err != nil {
		t.Fatalf("create benchmark visual summary directory: %v", err)
	}
	t.Logf("benchmark visuals enabled: dir=%s max_attempts=%d timeout=%s attempt_overlay=%v summary_native=%v show_pattern=%v", outDir, maxAttempts, rpcTimeout, attemptOverlay, summaryNative, summaryShowPattern)

	return &findBenchVisualCollector{
		outDir:             outDir,
		maxAttempts:        maxAttempts,
		rpcTimeout:         rpcTimeout,
		attemptOverlay:     attemptOverlay,
		summaryNative:      summaryNative,
		summaryShowPattern: summaryShowPattern,
		scenarios:          map[string]*findBenchScenarioVisual{},
	}
}

func (c *findBenchVisualCollector) CaptureAttempts(
	t testing.TB,
	client pb.SikuliServiceClient,
	queries []findBenchVisualQuery,
	source *image.Gray,
	engine findBenchEngine,
	scenario findBenchScenario,
) {
	t.Helper()
	if c == nil || client == nil || source == nil || len(queries) == 0 {
		return
	}
	patternSet := collectFindBenchPatterns(queries)

	attempts := make([]findBenchAttemptVisual, 0, c.maxAttempts)
	expectedRects := make([]*pb.Rect, 0, len(queries))
	for _, query := range queries {
		expectedRects = appendExpectedCandidates(expectedRects, query.Expected, query.Alternates)
	}
	for attempt := 1; attempt <= c.maxAttempts; attempt++ {
		rec := findBenchAttemptVisual{
			Attempt: attempt,
			Retried: attempt > 1,
			Status:  "error",
			Queries: make([]findBenchAttemptQueryVisual, 0, len(queries)),
		}

		start := time.Now()
		for queryIdx, query := range queries {
			var queryTemplate *image.Gray
			if qimg := patternGrayFromRequest(query.Request); qimg != nil {
				queryTemplate = cloneGray(qimg)
			}
			qrec := findBenchAttemptQueryVisual{
				ID:         strings.TrimSpace(query.ID),
				Label:      strings.TrimSpace(query.Label),
				Status:     "error",
				Expected:   cloneRect(query.Expected),
				Alternates: cloneRectSlice(query.Alternates),
				Template:   queryTemplate,
			}
			if qrec.Label == "" {
				qrec.Label = fmt.Sprintf("target-%02d", queryIdx+1)
			}
			if query.Request == nil || query.Expected == nil {
				qrec.Error = "missing query request or expected target"
				rec.Queries = append(rec.Queries, qrec)
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), c.rpcTimeout)
			res, err := client.FindOnScreen(ctx, query.Request)
			cancel()
			if err != nil {
				qrec.Status = visualStatusFromError(err)
				qrec.Error = err.Error()
				qrec.Explain = fmt.Sprintf("rpc %s", strings.ReplaceAll(qrec.Status, "_", " "))
			} else {
				rect := res.GetMatch().GetRect()
				if rect == nil {
					qrec.Status = "error"
					qrec.Error = "missing match rect"
					qrec.Explain = "rpc returned missing match rect"
				} else {
					qrec.Found = cloneRect(rect)
					qrec.Overlap = rectOverlapRatio(rect, query.Expected)
					qrec.AreaRatio = rectAreaRatio(rect, query.Expected)
					pattern := query.Request.GetPattern().GetImage()
					matchClass := classifyFindBenchPositiveMatch(
						rect,
						query.Expected,
						query.Alternates,
						expectedRects,
						pattern,
						scenario.tolerance,
						scenario.maxAreaRatio,
						scenario.allowPartial,
					)
					if matchClass == findBenchMatchClassOK {
						qrec.Status = "ok"
					} else if matchClass == findBenchMatchClassWrongRegion {
						qrec.Status = "wrong_region"
						qrec.Error = "matched another configured region"
					} else {
						qrec.Status = "overlap_miss"
					}
					qrec.Explain = explainFindBenchMatchOutcome(
						matchClass,
						rect,
						query.Expected,
						query.Alternates,
						pattern,
						scenario.tolerance,
						scenario.maxAreaRatio,
						scenario.allowPartial,
					)
				}
			}
			if strings.TrimSpace(qrec.Explain) == "" {
				qrec.Explain = strings.ReplaceAll(qrec.Status, "_", " ")
			}
			rec.Queries = append(rec.Queries, qrec)
		}
		rec.Duration = time.Since(start)
		rec.Status = aggregateAttemptStatus(rec.Queries)
		rec.Error = summarizeAttemptErrors(rec.Queries)

		path, writeErr := c.writeAttemptImage(source, scenario, engine, rec)
		if writeErr != nil {
			t.Logf("write benchmark attempt image scenario=%s engine=%s attempt=%d: %v", scenario.name, engine.name, attempt, writeErr)
		} else {
			rec.File = path
		}
		attempts = append(attempts, rec)

		if rec.Status == "ok" {
			break
		}
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.scenarios[scenario.name]
	if !ok {
		entry = &findBenchScenarioVisual{
			scenario: scenario,
			engines:  map[string][]findBenchAttemptVisual{},
			patterns: cloneFindBenchPatternSet(patternSet),
		}
		c.scenarios[scenario.name] = entry
	}
	if len(entry.patterns) == 0 && len(patternSet) > 0 {
		entry.patterns = cloneFindBenchPatternSet(patternSet)
	}
	entry.engines[engine.name] = append([]findBenchAttemptVisual(nil), attempts...)
}

func (c *findBenchVisualCollector) WriteScenarioSummaries() error {
	if c == nil {
		return nil
	}
	c.mu.Lock()
	snapshots := make([]*findBenchScenarioVisual, 0, len(c.scenarios))
	for _, v := range c.scenarios {
		cp := &findBenchScenarioVisual{
			scenario: v.scenario,
			engines:  map[string][]findBenchAttemptVisual{},
		}
		if len(v.patterns) > 0 {
			cp.patterns = cloneFindBenchPatternSet(v.patterns)
		}
		for engine, rows := range v.engines {
			cp.engines[engine] = append([]findBenchAttemptVisual(nil), rows...)
		}
		snapshots = append(snapshots, cp)
	}
	c.mu.Unlock()

	sort.Slice(snapshots, func(i, j int) bool {
		return snapshots[i].scenario.name < snapshots[j].scenario.name
	})
	summaryImages := make([]findBenchScenarioSummaryImage, 0, len(snapshots))
	for _, scenario := range snapshots {
		if err := c.writeScenarioSummary(scenario); err != nil {
			return err
		}
		summaryImages = append(summaryImages, findBenchScenarioSummaryImage{
			ScenarioName: scenario.scenario.name,
			Path:         c.scenarioSummaryPath(scenario.scenario.name),
		})
	}
	preferredMegaResolution := strings.TrimSpace(os.Getenv("FIND_BENCH_VISUAL_MEGA_RES"))
	if preferredMegaResolution == "" {
		preferredMegaResolution = "1280x720"
	}
	selected := selectMegaSummaryImages(summaryImages, preferredMegaResolution)
	if len(selected) > 0 {
		summaryImages = selected
	}
	if err := c.writeRunMegaSummary(summaryImages, preferredMegaResolution); err != nil {
		return err
	}
	return nil
}

func (c *findBenchVisualCollector) writeAttemptImage(source *image.Gray, scenario findBenchScenario, engine findBenchEngine, rec findBenchAttemptVisual) (string, error) {
	canvas := grayToRGBA(source)
	globalZoneGuideAngle, globalZoneGuideAngleOK := 0.0, false
	if shouldRotateExpectedGuideZone(scenario) {
		globalZoneGuideAngle, globalZoneGuideAngleOK = inferGlobalZoneGuideAngle(source, rec.Queries)
	}
	for _, query := range rec.Queries {
		expectedTarget := selectVisualExpectedTarget(query, scenario)
		expectedGuide := resolveExpectedGuideQuad(source, expectedTarget, query.Template, scenario, query.Found, globalZoneGuideAngle, globalZoneGuideAngleOK)
		if len(expectedGuide) > 0 {
			drawDashedQuadOutline(canvas, expectedGuide, color.RGBA{R: 148, G: 158, B: 172, A: 255}, 2, 10, 7)
			drawQuadCross(canvas, expectedGuide, color.RGBA{R: 148, G: 158, B: 172, A: 26}, 2)
		}
		if query.Found == nil {
			continue
		}
		boxColor := color.RGBA{R: 230, G: 80, B: 80, A: 255} // false positive/wrong place
		if query.Status == "ok" {
			boxColor = color.RGBA{R: 40, G: 210, B: 95, A: 255} // correct match
		}
		drawRectCross(canvas, query.Found, colorWithAlpha(boxColor, 26), 2)
		if expectedTarget != nil {
			ex, ey := quadCenter(expectedGuide)
			if len(expectedGuide) == 0 {
				ex, ey = rectCenter(expectedTarget)
			}
			fx, fy := rectCenter(query.Found)
			drawBidirectionalArrow(canvas, benchPointF{X: ex, Y: ey}, benchPointF{X: fx, Y: fy}, colorWithAlpha(boxColor, 224), 2)
		}
		drawRectOutline(canvas, query.Found, boxColor, 3)
		drawMatchDetailLabel(canvas, query, scenario, boxColor)
	}

	if c.attemptOverlay {
		overlayHeight := 74
		if canvas.Bounds().Dy() < overlayHeight+8 {
			overlayHeight = maxInt(24, canvas.Bounds().Dy()/2)
		}
		draw.Draw(
			canvas,
			image.Rect(0, 0, canvas.Bounds().Dx(), overlayHeight),
			&image.Uniform{C: color.RGBA{R: 14, G: 19, B: 28, A: 220}},
			image.Point{},
			draw.Over,
		)

		line1 := fmt.Sprintf("ENGINE %s ATTEMPT %d", engine.name, rec.Attempt)
		line2 := fmt.Sprintf("STATUS %s DURATION %.3fMS RETRY %s", rec.Status, float64(rec.Duration.Microseconds())/1000.0, boolWord(rec.Retried))
		okCount, notFoundCount, overlapCount, errorCount := summarizeAttemptQueryCounts(rec.Queries)
		line3 := fmt.Sprintf("MATCH OK %d/%d NOT_FOUND %d OVERLAP %d ERROR %d", okCount, len(rec.Queries), notFoundCount, overlapCount, errorCount)
		line4 := attemptQueryStatusLine(rec.Queries)
		if rec.Error != "" {
			line4 = truncateUpper(fmt.Sprintf("ERROR %s", rec.Error), 88)
		}

		textScale := 2
		y := 7
		y += drawTinyText(canvas, 8, y, line1, color.RGBA{R: 255, G: 255, B: 255, A: 255}, textScale)
		y += 3
		y += drawTinyText(canvas, 8, y, line2, color.RGBA{R: 220, G: 231, B: 248, A: 255}, textScale)
		y += 3
		y += drawTinyText(canvas, 8, y, line3, color.RGBA{R: 160, G: 245, B: 183, A: 255}, textScale)
		y += 3
		_ = drawTinyText(canvas, 8, y, line4, color.RGBA{R: 250, G: 189, B: 189, A: 255}, textScale)
	}

	scenarioDir := filepath.Join(c.outDir, "attempts", stableScenarioImageToken(scenario.name))
	if err := os.MkdirAll(scenarioDir, 0o755); err != nil {
		return "", err
	}
	filename := fmt.Sprintf("engine-%s-attempt-%d.png", sanitizeFileToken(engine.name), rec.Attempt)
	path := filepath.Join(scenarioDir, filename)
	if err := writePNG(path, canvas); err != nil {
		return "", err
	}
	return path, nil
}

func (c *findBenchVisualCollector) writeScenarioSummary(scenario *findBenchScenarioVisual) error {
	if scenario == nil {
		return nil
	}
	engines := make([]string, 0, len(scenario.engines))
	for name := range scenario.engines {
		engines = append(engines, name)
	}
	sort.Strings(engines)
	if len(engines) == 0 {
		return nil
	}

	rows := 1
	for _, engine := range engines {
		if n := len(scenario.engines[engine]); n > rows {
			rows = n
		}
	}

	srcW := maxInt(1, scenario.scenario.screenW)
	srcH := maxInt(1, scenario.scenario.screenH)
	summaryScale := 1
	if c.summaryNative {
		summaryScale = 1
	}
	screenW := 420 * summaryScale
	screenH := maxInt(120*summaryScale, int(float64(srcH)*float64(screenW)/float64(srcW)))
	if c.summaryNative {
		screenW = srcW
		screenH = srcH
	}
	margin := 12 * summaryScale
	headerH := 78 * summaryScale
	rowLabelH := 22 * summaryScale
	patternGap := 8 * summaryScale
	patternPanelW := 0
	if c.summaryShowPattern && len(scenario.patterns) > 0 {
		if c.summaryNative {
			patternPanelW = maxInt(96, minInt(180, screenW/8))
		} else {
			patternPanelW = 160 * summaryScale
		}
	}
	cellW := screenW
	if patternPanelW > 0 {
		cellW += patternPanelW + patternGap
	}
	cellH := rowLabelH + screenH
	cols := len(engines)

	width := margin + cols*(cellW+margin)
	height := headerH + margin + rows*(cellH+margin)
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{C: color.RGBA{R: 242, G: 246, B: 252, A: 255}}, image.Point{}, draw.Src)
	draw.Draw(canvas, image.Rect(0, 0, width, headerH), &image.Uniform{C: color.RGBA{R: 16, G: 26, B: 40, A: 255}}, image.Point{}, draw.Src)

	title := fmt.Sprintf("SCENARIO %s RES %dx%d ROT %d VARIANT %s", scenario.scenario.name, scenario.scenario.screenW, scenario.scenario.screenH, scenario.scenario.rotation, scenario.scenario.variant)
	subtitle := "GREEN CORRECT MATCH   RED FALSE POSITIVE   NO BOX NO MATCH   LABELS ABOVE SCREEN"
	drawTinyText(canvas, 10*summaryScale, 8*summaryScale, truncateUpper(title, 108), color.RGBA{R: 255, G: 255, B: 255, A: 255}, 2*summaryScale)
	drawTinyText(canvas, 10*summaryScale, 42*summaryScale, subtitle, color.RGBA{R: 190, G: 208, B: 236, A: 255}, 2*summaryScale)

	for col, engine := range engines {
		records := scenario.engines[engine]
		for row := 0; row < rows; row++ {
			x := margin + col*(cellW+margin)
			y := headerH + margin + row*(cellH+margin)
			bodyY := y + rowLabelH
			screenX := x
			if patternPanelW > 0 {
				screenX += patternPanelW + patternGap
			}
			screenRect := image.Rect(screenX, bodyY, screenX+screenW, bodyY+screenH)
			draw.Draw(canvas, screenRect, &image.Uniform{C: color.RGBA{R: 222, G: 228, B: 237, A: 255}}, image.Point{}, draw.Src)

			label := fmt.Sprintf("ENGINE %s ATTEMPT %d NO RETRY", engine, row+1)
			if row < len(records) {
				rec := records[row]
				label = fmt.Sprintf("ENGINE %s ATTEMPT %d %s %.3fMS RETRY %s", engine, rec.Attempt, rec.Status, float64(rec.Duration.Microseconds())/1000.0, boolWord(rec.Retried))
				img, err := readImage(rec.File)
				if err == nil {
					iw := img.Bounds().Dx()
					ih := img.Bounds().Dy()
					tw, th := fitWithinNoUpscale(screenW, screenH, iw, ih)
					if tw > 0 && th > 0 {
						thumb := resizeNearest(img, tw, th)
						px := screenRect.Min.X + (screenRect.Dx()-tw)/2
						py := screenRect.Min.Y + (screenRect.Dy()-th)/2
						draw.Draw(canvas, image.Rect(px, py, px+tw, py+th), thumb, image.Point{}, draw.Src)
					}
				}
			}

			if patternPanelW > 0 {
				patternRect := image.Rect(x, bodyY, x+patternPanelW, bodyY+screenH)
				draw.Draw(canvas, patternRect, &image.Uniform{C: color.RGBA{R: 228, G: 234, B: 243, A: 255}}, image.Point{}, draw.Src)
				drawPatternStackPanel(canvas, patternRect, scenario.patterns, summaryScale)
				drawTinyText(canvas, x+4, y+4, "PATTERNS", color.RGBA{R: 35, G: 46, B: 64, A: 255}, summaryScale)
				drawRectOutline(canvas, &pb.Rect{X: int32(patternRect.Min.X), Y: int32(patternRect.Min.Y), W: int32(patternRect.Dx()), H: int32(patternRect.Dy())}, color.RGBA{R: 104, G: 116, B: 137, A: 255}, maxInt(1, summaryScale))
			}

			drawTinyText(canvas, screenX+4, y+4, truncateUpper(label, 84), color.RGBA{R: 35, G: 46, B: 64, A: 255}, summaryScale)
			drawRectOutline(canvas, &pb.Rect{X: int32(screenRect.Min.X), Y: int32(screenRect.Min.Y), W: int32(screenRect.Dx()), H: int32(screenRect.Dy())}, color.RGBA{R: 82, G: 95, B: 116, A: 255}, maxInt(1, summaryScale))
		}
	}

	path := c.scenarioSummaryPath(scenario.scenario.name)
	return writePNG(path, canvas)
}

func collectFindBenchPatterns(queries []findBenchVisualQuery) []findBenchPatternVisual {
	out := make([]findBenchPatternVisual, 0, len(queries))
	for idx, query := range queries {
		img := patternGrayFromRequest(query.Request)
		if img == nil {
			continue
		}
		label := strings.TrimSpace(query.Label)
		if label == "" {
			label = fmt.Sprintf("target-%02d", idx+1)
		}
		out = append(out, findBenchPatternVisual{
			Label: label,
			Image: cloneGray(img),
		})
	}
	return out
}

func cloneFindBenchPatternSet(in []findBenchPatternVisual) []findBenchPatternVisual {
	if len(in) == 0 {
		return nil
	}
	out := make([]findBenchPatternVisual, 0, len(in))
	for _, item := range in {
		if item.Image == nil {
			continue
		}
		out = append(out, findBenchPatternVisual{
			Label: item.Label,
			Image: cloneGray(item.Image),
		})
	}
	return out
}

func drawPatternStackPanel(canvas *image.RGBA, panel image.Rectangle, patterns []findBenchPatternVisual, summaryScale int) {
	if canvas == nil || panel.Empty() || len(patterns) == 0 {
		return
	}
	inner := panel.Inset(maxInt(2, summaryScale))
	if inner.Empty() {
		return
	}

	slotGap := maxInt(2, summaryScale)
	slotCount := len(patterns)
	totalGap := slotGap * (slotCount - 1)
	slotH := (inner.Dy() - totalGap) / maxInt(1, slotCount)
	if slotH < 16 {
		slotH = 16
	}

	for idx, pattern := range patterns {
		slotY := inner.Min.Y + idx*(slotH+slotGap)
		if slotY >= inner.Max.Y {
			break
		}
		slotRect := image.Rect(inner.Min.X, slotY, inner.Max.X, minInt(slotY+slotH, inner.Max.Y))
		if slotRect.Dy() < 8 {
			continue
		}
		draw.Draw(canvas, slotRect, &image.Uniform{C: color.RGBA{R: 214, G: 223, B: 236, A: 255}}, image.Point{}, draw.Src)
		drawRectOutline(
			canvas,
			&pb.Rect{X: int32(slotRect.Min.X), Y: int32(slotRect.Min.Y), W: int32(slotRect.Dx()), H: int32(slotRect.Dy())},
			color.RGBA{R: 135, G: 149, B: 170, A: 255},
			maxInt(1, summaryScale),
		)

		if pattern.Image == nil {
			continue
		}
		label := strings.TrimSpace(pattern.Label)
		if label == "" {
			label = fmt.Sprintf("target-%02d", idx+1)
		}
		drawTinyText(canvas, slotRect.Min.X+2, slotRect.Min.Y+2, truncateUpper(label, 16), color.RGBA{R: 35, G: 46, B: 64, A: 255}, summaryScale)

		topPad := 12 * summaryScale
		availW := maxInt(1, slotRect.Dx()-6)
		availH := maxInt(1, slotRect.Dy()-topPad-3)
		pw, ph := fitWithinNoUpscale(availW, availH, pattern.Image.Bounds().Dx(), pattern.Image.Bounds().Dy())
		if pw < 1 || ph < 1 {
			continue
		}
		px := slotRect.Min.X + (slotRect.Dx()-pw)/2
		py := slotRect.Min.Y + topPad + (availH-ph)/2
		patternImg := resizeNearest(grayToRGBA(pattern.Image), pw, ph)
		draw.Draw(canvas, image.Rect(px, py, px+pw, py+ph), patternImg, image.Point{}, draw.Src)
	}
}

func (c *findBenchVisualCollector) scenarioSummaryPath(scenarioName string) string {
	return filepath.Join(c.outDir, "summaries", fmt.Sprintf("summary-%s.png", stableScenarioImageToken(scenarioName)))
}

func (c *findBenchVisualCollector) writeRunMegaSummary(images []findBenchScenarioSummaryImage, preferredResolution string) error {
	if len(images) == 0 {
		return nil
	}
	sort.Slice(images, func(i, j int) bool {
		return images[i].ScenarioName < images[j].ScenarioName
	})

	const (
		megaScale = 2
		margin    = 16 * megaScale
		headerH   = 96 * megaScale
		tileW     = 680 * megaScale
		tileH     = 330 * megaScale
		captionH  = 24 * megaScale
	)
	cols := int(math.Ceil(math.Sqrt(float64(len(images)))))
	if cols < 1 {
		cols = 1
	}
	rows := int(math.Ceil(float64(len(images)) / float64(cols)))

	width := margin + cols*(tileW+margin)
	height := headerH + margin + rows*(tileH+captionH+margin)
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{C: color.RGBA{R: 238, G: 243, B: 250, A: 255}}, image.Point{}, draw.Src)
	draw.Draw(canvas, image.Rect(0, 0, width, headerH), &image.Uniform{C: color.RGBA{R: 14, G: 24, B: 39, A: 255}}, image.Point{}, draw.Src)

	title := "RUN-LEVEL MEGA SUMMARY   FINDONSCREEN BENCHMARK"
	resolutionLabel := strings.TrimSpace(preferredResolution)
	if resolutionLabel == "" {
		resolutionLabel = "auto"
	}
	meta := fmt.Sprintf("SCENARIOS %d   RES %s   GRID %dX%d   GENERATED %s", len(images), resolutionLabel, cols, rows, time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	drawTinyText(canvas, 10*megaScale, 12*megaScale, title, color.RGBA{R: 255, G: 255, B: 255, A: 255}, 2*megaScale)
	drawTinyText(canvas, 10*megaScale, 46*megaScale, truncateUpper(meta, 120), color.RGBA{R: 184, G: 202, B: 232, A: 255}, 2*megaScale)

	for i, summary := range images {
		col := i % cols
		row := i / cols
		x := margin + col*(tileW+margin)
		y := headerH + margin + row*(tileH+captionH+margin)

		tileRect := image.Rect(x, y, x+tileW, y+tileH)
		draw.Draw(canvas, tileRect, &image.Uniform{C: color.RGBA{R: 219, G: 226, B: 236, A: 255}}, image.Point{}, draw.Src)
		if img, err := readImage(summary.Path); err == nil {
			iw := img.Bounds().Dx()
			ih := img.Bounds().Dy()
			tw, th := fitWithinNoUpscale(tileW, tileH, iw, ih)
			if tw > 0 && th > 0 {
				thumb := resizeNearest(img, tw, th)
				px := tileRect.Min.X + (tileRect.Dx()-tw)/2
				py := tileRect.Min.Y + (tileRect.Dy()-th)/2
				draw.Draw(canvas, image.Rect(px, py, px+tw, py+th), thumb, image.Point{}, draw.Src)
			}
		}
		drawRectOutline(canvas, &pb.Rect{X: int32(x), Y: int32(y), W: int32(tileW), H: int32(tileH)}, color.RGBA{R: 77, G: 90, B: 112, A: 255}, 2*megaScale)
		drawTinyText(
			canvas,
			x+8,
			y+tileH+10,
			truncateUpper(fmt.Sprintf("SCENARIO %s", summary.ScenarioName), 72),
			color.RGBA{R: 33, G: 43, B: 60, A: 255},
			megaScale,
		)
	}

	return writeJPEG(filepath.Join(c.outDir, "summaries", "summary-run-mega.jpg"), canvas, 80)
}

func selectMegaSummaryImages(images []findBenchScenarioSummaryImage, preferredResolution string) []findBenchScenarioSummaryImage {
	if len(images) == 0 {
		return nil
	}
	type keyed struct {
		family     string
		resolution string
		entry      findBenchScenarioSummaryImage
	}
	byFamily := map[string][]keyed{}
	for _, img := range images {
		family, resolution := splitScenarioFamilyAndResolution(img.ScenarioName)
		if family == "" {
			family = img.ScenarioName
		}
		byFamily[family] = append(byFamily[family], keyed{
			family:     family,
			resolution: resolution,
			entry:      img,
		})
	}

	families := make([]string, 0, len(byFamily))
	for family := range byFamily {
		families = append(families, family)
	}
	sort.Strings(families)

	out := make([]findBenchScenarioSummaryImage, 0, len(families))
	for _, family := range families {
		rows := byFamily[family]
		sort.Slice(rows, func(i, j int) bool {
			return rows[i].entry.ScenarioName < rows[j].entry.ScenarioName
		})
		chosen := rows[0]
		if preferredResolution != "" {
			for _, row := range rows {
				if row.resolution == preferredResolution {
					chosen = row
					break
				}
			}
		}
		out = append(out, chosen.entry)
	}
	return out
}

func splitScenarioFamilyAndResolution(name string) (family string, resolution string) {
	m := scenarioResolutionPattern.FindStringSubmatch(strings.TrimSpace(name))
	if len(m) != 3 {
		return strings.TrimSpace(name), ""
	}
	return strings.TrimSpace(m[1]), strings.TrimSpace(m[2])
}

func writePNG(path string, in image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	return png.Encode(f, in)
}

func writeJPEG(path string, in image.Image, quality int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	if quality < 1 {
		quality = 1
	}
	if quality > 100 {
		quality = 100
	}
	return jpeg.Encode(f, in, &jpeg.Options{Quality: quality})
}

func readImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func grayToRGBA(in *image.Gray) *image.RGBA {
	b := in.Bounds()
	out := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			v := in.Pix[in.PixOffset(b.Min.X+x, b.Min.Y+y)]
			out.SetRGBA(x, y, color.RGBA{R: v, G: v, B: v, A: 255})
		}
	}
	return out
}

func resizeNearest(src image.Image, dstW, dstH int) *image.RGBA {
	if dstW < 1 {
		dstW = 1
	}
	if dstH < 1 {
		dstH = 1
	}
	sb := src.Bounds()
	sw := maxInt(1, sb.Dx())
	sh := maxInt(1, sb.Dy())
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))
	for y := 0; y < dstH; y++ {
		sy := sb.Min.Y + (y*sh)/dstH
		for x := 0; x < dstW; x++ {
			sx := sb.Min.X + (x*sw)/dstW
			dst.Set(x, y, src.At(sx, sy))
		}
	}
	return dst
}

func resizeGrayNearest(src *image.Gray, dstW, dstH int) *image.Gray {
	if src == nil || dstW < 1 || dstH < 1 {
		return nil
	}
	sb := src.Bounds()
	sw := maxInt(1, sb.Dx())
	sh := maxInt(1, sb.Dy())
	dst := image.NewGray(image.Rect(0, 0, dstW, dstH))
	for y := 0; y < dstH; y++ {
		sy := sb.Min.Y + (y*sh)/dstH
		for x := 0; x < dstW; x++ {
			sx := sb.Min.X + (x*sw)/dstW
			dst.Pix[dst.PixOffset(x, y)] = src.Pix[src.PixOffset(sx, sy)]
		}
	}
	return dst
}

func colorWithAlpha(c color.RGBA, a uint8) color.RGBA {
	c.A = a
	return c
}

func drawRectOutline(img *image.RGBA, rect *pb.Rect, c color.RGBA, thickness int) {
	if img == nil || rect == nil || thickness < 1 {
		return
	}
	b := img.Bounds()
	x0 := int(rect.GetX())
	y0 := int(rect.GetY())
	x1 := x0 + int(rect.GetW()) - 1
	y1 := y0 + int(rect.GetH()) - 1
	if x1 < x0 || y1 < y0 {
		return
	}
	for t := 0; t < thickness; t++ {
		yt := y0 + t
		yb := y1 - t
		for x := x0; x <= x1; x++ {
			setRGBAIfInBounds(img, b, x, yt, c)
			setRGBAIfInBounds(img, b, x, yb, c)
		}
		xl := x0 + t
		xr := x1 - t
		for y := y0; y <= y1; y++ {
			setRGBAIfInBounds(img, b, xl, y, c)
			setRGBAIfInBounds(img, b, xr, y, c)
		}
	}
}

func drawRectCross(img *image.RGBA, rect *pb.Rect, c color.RGBA, thickness int) {
	if rect == nil {
		return
	}
	drawLineSegment(img, benchPointF{X: float64(rect.GetX()), Y: float64(rect.GetY())}, benchPointF{X: float64(rect.GetX() + rect.GetW()), Y: float64(rect.GetY() + rect.GetH())}, c, thickness, 0, 0, true)
	drawLineSegment(img, benchPointF{X: float64(rect.GetX() + rect.GetW()), Y: float64(rect.GetY())}, benchPointF{X: float64(rect.GetX()), Y: float64(rect.GetY() + rect.GetH())}, c, thickness, 0, 0, true)
}

func drawQuadCross(img *image.RGBA, quad []benchPointF, c color.RGBA, thickness int) {
	if len(quad) != 4 {
		return
	}
	drawLineSegment(img, quad[0], quad[2], c, thickness, 0, 0, true)
	drawLineSegment(img, quad[1], quad[3], c, thickness, 0, 0, true)
}

func drawDashedQuadOutline(img *image.RGBA, quad []benchPointF, c color.RGBA, thickness int, dashLen float64, gapLen float64) {
	if len(quad) != 4 {
		return
	}
	for i := 0; i < 4; i++ {
		a := quad[i]
		b := quad[(i+1)%4]
		drawLineSegment(img, a, b, c, thickness, dashLen, gapLen, false)
	}
}

func drawBidirectionalArrow(img *image.RGBA, a benchPointF, b benchPointF, c color.RGBA, thickness int) {
	drawLineSegment(img, a, b, c, thickness, 0, 0, false)
	drawArrowHead(img, a, b, c, thickness)
	drawArrowHead(img, b, a, c, thickness)
}

func drawArrowHead(img *image.RGBA, tip benchPointF, other benchPointF, c color.RGBA, thickness int) {
	dx := tip.X - other.X
	dy := tip.Y - other.Y
	dist := math.Hypot(dx, dy)
	if dist < 1 {
		return
	}
	ux := dx / dist
	uy := dy / dist
	arrowLen := math.Max(8.0, float64(thickness)*5.0)
	arrowWidth := arrowLen * 0.45
	baseX := tip.X - ux*arrowLen
	baseY := tip.Y - uy*arrowLen
	px := -uy
	py := ux
	left := benchPointF{X: baseX + px*arrowWidth, Y: baseY + py*arrowWidth}
	right := benchPointF{X: baseX - px*arrowWidth, Y: baseY - py*arrowWidth}
	drawLineSegment(img, tip, left, c, thickness, 0, 0, false)
	drawLineSegment(img, tip, right, c, thickness, 0, 0, false)
}

func drawLineSegment(img *image.RGBA, a benchPointF, b benchPointF, c color.RGBA, thickness int, dashLen float64, gapLen float64, blend bool) {
	if img == nil || thickness < 1 {
		return
	}
	dx := b.X - a.X
	dy := b.Y - a.Y
	dist := math.Hypot(dx, dy)
	if dist < 0.5 {
		drawLineBrush(img, int(math.Round(a.X)), int(math.Round(a.Y)), c, thickness, blend)
		return
	}
	steps := int(math.Ceil(dist))
	if steps < 1 {
		steps = 1
	}
	for step := 0; step <= steps; step++ {
		t := float64(step) / float64(steps)
		if dashLen > 0 {
			mod := math.Mod(t*dist, dashLen+gapLen)
			if mod > dashLen {
				continue
			}
		}
		x := a.X + dx*t
		y := a.Y + dy*t
		drawLineBrush(img, int(math.Round(x)), int(math.Round(y)), c, thickness, blend)
	}
}

func drawLineBrush(img *image.RGBA, x int, y int, c color.RGBA, thickness int, blend bool) {
	b := img.Bounds()
	radius := maxInt(0, thickness/2)
	for yy := y - radius; yy <= y+radius; yy++ {
		for xx := x - radius; xx <= x+radius; xx++ {
			if blend {
				blendRGBAIfInBounds(img, b, xx, yy, c)
			} else {
				setRGBAIfInBounds(img, b, xx, yy, c)
			}
		}
	}
}

func blendRGBAIfInBounds(img *image.RGBA, bounds image.Rectangle, x, y int, c color.RGBA) {
	if x < bounds.Min.X || y < bounds.Min.Y || x >= bounds.Max.X || y >= bounds.Max.Y {
		return
	}
	if c.A == 255 {
		img.SetRGBA(x, y, c)
		return
	}
	dst := img.RGBAAt(x, y)
	alpha := float64(c.A) / 255.0
	inv := 1.0 - alpha
	img.SetRGBA(x, y, color.RGBA{
		R: uint8(math.Round(float64(c.R)*alpha + float64(dst.R)*inv)),
		G: uint8(math.Round(float64(c.G)*alpha + float64(dst.G)*inv)),
		B: uint8(math.Round(float64(c.B)*alpha + float64(dst.B)*inv)),
		A: 255,
	})
}

func rectToQuad(rect *pb.Rect) []benchPointF {
	if rect == nil {
		return nil
	}
	return []benchPointF{
		{X: float64(rect.GetX()), Y: float64(rect.GetY())},
		{X: float64(rect.GetX() + rect.GetW()), Y: float64(rect.GetY())},
		{X: float64(rect.GetX() + rect.GetW()), Y: float64(rect.GetY() + rect.GetH())},
		{X: float64(rect.GetX()), Y: float64(rect.GetY() + rect.GetH())},
	}
}

func quadCenter(quad []benchPointF) (float64, float64) {
	if len(quad) == 0 {
		return 0, 0
	}
	var sumX float64
	var sumY float64
	for _, p := range quad {
		sumX += p.X
		sumY += p.Y
	}
	return sumX / float64(len(quad)), sumY / float64(len(quad))
}

func rotatedRectQuad(cx, cy, w, h, angleDeg float64) []benchPointF {
	hw := w / 2.0
	hh := h / 2.0
	theta := angleDeg * math.Pi / 180.0
	cosT := math.Cos(theta)
	sinT := math.Sin(theta)
	points := []benchPointF{
		{X: -hw, Y: -hh},
		{X: hw, Y: -hh},
		{X: hw, Y: hh},
		{X: -hw, Y: hh},
	}
	out := make([]benchPointF, 0, 4)
	for _, p := range points {
		out = append(out, benchPointF{
			X: cx + p.X*cosT - p.Y*sinT,
			Y: cy + p.X*sinT + p.Y*cosT,
		})
	}
	return out
}

func setRGBAIfInBounds(img *image.RGBA, bounds image.Rectangle, x, y int, c color.RGBA) {
	if x < bounds.Min.X || y < bounds.Min.Y || x >= bounds.Max.X || y >= bounds.Max.Y {
		return
	}
	img.SetRGBA(x, y, c)
}

func cloneRect(in *pb.Rect) *pb.Rect {
	if in == nil {
		return nil
	}
	return &pb.Rect{X: in.GetX(), Y: in.GetY(), W: in.GetW(), H: in.GetH()}
}

func patternGrayFromRequest(req *pb.FindOnScreenRequest) *image.Gray {
	if req == nil || req.GetPattern() == nil || req.GetPattern().GetImage() == nil {
		return nil
	}
	gi := req.GetPattern().GetImage()
	w := int(gi.GetWidth())
	h := int(gi.GetHeight())
	if w <= 0 || h <= 0 {
		return nil
	}
	pix := gi.GetPix()
	if len(pix) < w*h {
		return nil
	}
	out := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		start := y * w
		copy(out.Pix[y*out.Stride:y*out.Stride+w], pix[start:start+w])
	}
	return out
}

func fitWithin(maxW, maxH, srcW, srcH int) (int, int) {
	if maxW <= 0 || maxH <= 0 || srcW <= 0 || srcH <= 0 {
		return 0, 0
	}
	scale := math.Min(float64(maxW)/float64(srcW), float64(maxH)/float64(srcH))
	if scale <= 0 {
		return 0, 0
	}
	w := maxInt(1, int(math.Round(float64(srcW)*scale)))
	h := maxInt(1, int(math.Round(float64(srcH)*scale)))
	return w, h
}

func fitWithinNoUpscale(maxW, maxH, srcW, srcH int) (int, int) {
	if maxW <= 0 || maxH <= 0 || srcW <= 0 || srcH <= 0 {
		return 0, 0
	}
	if srcW <= maxW && srcH <= maxH {
		return srcW, srcH
	}
	return fitWithin(maxW, maxH, srcW, srcH)
}

func visualStatusFromError(err error) string {
	code := status.Code(err)
	switch code {
	case codes.NotFound:
		return "not_found"
	case codes.Unimplemented:
		return "unsupported"
	case codes.DeadlineExceeded:
		return "timeout"
	default:
		return "error"
	}
}

func aggregateAttemptStatus(queries []findBenchAttemptQueryVisual) string {
	if len(queries) == 0 {
		return "error"
	}
	allOK := true
	hasNotFound := false
	hasOverlap := false
	hasTimeout := false
	hasUnsupported := false
	hasError := false
	for _, query := range queries {
		switch query.Status {
		case "ok":
			// no-op
		case "not_found":
			allOK = false
			hasNotFound = true
		case "overlap_miss":
			allOK = false
			hasOverlap = true
		case "wrong_region":
			allOK = false
			hasOverlap = true
		case "timeout":
			allOK = false
			hasTimeout = true
		case "unsupported":
			allOK = false
			hasUnsupported = true
		default:
			allOK = false
			hasError = true
		}
	}
	if allOK {
		return "ok"
	}
	if hasError {
		return "error"
	}
	if hasTimeout {
		return "timeout"
	}
	if hasUnsupported {
		return "unsupported"
	}
	if hasOverlap {
		return "overlap_miss"
	}
	if hasNotFound {
		return "not_found"
	}
	return "error"
}

func summarizeAttemptErrors(queries []findBenchAttemptQueryVisual) string {
	parts := make([]string, 0, 2)
	for _, query := range queries {
		if strings.TrimSpace(query.Error) == "" {
			continue
		}
		label := strings.TrimSpace(query.Label)
		if label == "" {
			label = "query"
		}
		parts = append(parts, fmt.Sprintf("%s: %s", label, query.Error))
		if len(parts) >= 2 {
			break
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, " | ")
}

func summarizeAttemptQueryCounts(queries []findBenchAttemptQueryVisual) (okCount int, notFoundCount int, overlapCount int, errorCount int) {
	for _, query := range queries {
		switch query.Status {
		case "ok":
			okCount++
		case "not_found":
			notFoundCount++
		case "overlap_miss", "wrong_region":
			overlapCount++
		default:
			errorCount++
		}
	}
	return okCount, notFoundCount, overlapCount, errorCount
}

func attemptQueryStatusLine(queries []findBenchAttemptQueryVisual) string {
	if len(queries) == 0 {
		return "TARGETS NONE"
	}
	parts := make([]string, 0, len(queries))
	for idx, query := range queries {
		label := strings.TrimSpace(query.Label)
		if label == "" {
			label = fmt.Sprintf("target-%02d", idx+1)
		}
		parts = append(parts, fmt.Sprintf("%s=%s", label, query.Status))
	}
	return truncateUpper(fmt.Sprintf("TARGETS %s", strings.Join(parts, " ")), 88)
}

func grayImageDimensionsProto(img *image.Gray) *pb.GrayImage {
	if img == nil {
		return nil
	}
	b := img.Bounds()
	if b.Dx() < 1 || b.Dy() < 1 {
		return nil
	}
	return &pb.GrayImage{Width: int32(b.Dx()), Height: int32(b.Dy())}
}

func bestExpectedCandidateMetrics(
	found *pb.Rect,
	expected *pb.Rect,
	expectedAlternates []*pb.Rect,
	pattern *pb.GrayImage,
	tolerance float64,
	maxAreaRatio float64,
	allowPartial bool,
) (target *pb.Rect, dx float64, dy float64, dist float64, limX float64, limY float64, ok bool) {
	target = expected
	if expected == nil {
		return nil, 0, 0, 0, 0, 0, false
	}
	dx, dy, dist, limX, limY, ok = centerDeltaMetrics(found, expected, pattern, tolerance, maxAreaRatio, allowPartial)
	for _, alt := range expectedAlternates {
		if alt == nil {
			continue
		}
		adx, ady, adist, alimX, alimY, aok := centerDeltaMetrics(found, alt, pattern, tolerance, maxAreaRatio, allowPartial)
		if adist < dist {
			target = alt
			dx, dy, dist, limX, limY, ok = adx, ady, adist, alimX, alimY, aok
		}
	}
	return target, dx, dy, dist, limX, limY, ok
}

func selectVisualExpectedTarget(query findBenchAttemptQueryVisual, scenario findBenchScenario) *pb.Rect {
	if query.Found != nil {
		templateDims := grayImageDimensionsProto(query.Template)
		target, _, _, _, _, _, _ := bestExpectedCandidateMetrics(
			query.Found,
			query.Expected,
			query.Alternates,
			templateDims,
			scenario.tolerance,
			scenario.maxAreaRatio,
			scenario.allowPartial,
		)
		if target != nil {
			return target
		}
	}
	return query.Expected
}

func explainFindBenchMatchOutcome(
	matchClass findBenchMatchClass,
	found *pb.Rect,
	expected *pb.Rect,
	expectedAlternates []*pb.Rect,
	pattern *pb.GrayImage,
	tolerance float64,
	maxAreaRatio float64,
	allowPartial bool,
) string {
	if found == nil || expected == nil {
		return "missing match or expected rect"
	}
	target, dx, dy, dist, limX, _, _ := bestExpectedCandidateMetrics(found, expected, expectedAlternates, pattern, tolerance, maxAreaRatio, allowPartial)
	if !areaRatioWithinLimit(found, pattern, maxAreaRatio) {
		ratio := 0.0
		if pattern != nil {
			patternArea := float64(max32(1, pattern.GetWidth()*pattern.GetHeight()))
			if patternArea > 0 {
				ratio = rectAreaFloat(found) / patternArea
			}
		}
		return fmt.Sprintf("area miss ratio=%.2f max=%.2f", ratio, maxAreaRatio)
	}

	switch matchClass {
	case findBenchMatchClassOK:
		if regionActsAsZone(target, pattern) {
			return fmt.Sprintf("zone ok d=%.0f dx=%.0f dy=%.0f lim=%.0f", dist, dx, dy, limX)
		}
		return fmt.Sprintf("center ok d=%.0f dx=%.0f dy=%.0f lim=%.0f", dist, dx, dy, limX)
	case findBenchMatchClassWrongRegion:
		return fmt.Sprintf("center in peer d=%.0f dx=%.0f dy=%.0f", dist, dx, dy)
	default:
		if regionActsAsZone(target, pattern) {
			return fmt.Sprintf("zone miss d=%.0f dx=%.0f dy=%.0f lim=%.0f", dist, dx, dy, limX)
		}
		return fmt.Sprintf("center miss d=%.0f dx=%.0f dy=%.0f lim=%.0f", dist, dx, dy, limX)
	}
}

func resolveExpectedGuideQuad(source *image.Gray, expected *pb.Rect, template *image.Gray, scenario findBenchScenario, found *pb.Rect, fallbackAngle float64, fallbackAngleOK bool) []benchPointF {
	if expected == nil {
		return nil
	}
	base := rectToQuad(expected)
	if template == nil || source == nil {
		return base
	}
	templateDims := grayImageDimensionsProto(template)
	if regionActsAsZone(expected, templateDims) {
		if !shouldRotateExpectedGuideZone(scenario) {
			return base
		}
		angle, ok := 0.0, false
		if found != nil {
			angle, _, ok = inferGuideAngleAndScaleFromRect(source, found, template)
		}
		if !ok && fallbackAngleOK {
			angle, ok = fallbackAngle, true
		}
		if !ok || math.Abs(angle) < 2.0 {
			return base
		}
		cx, cy := rectCenter(expected)
		return rotatedRectQuad(cx, cy, float64(expected.GetW()), float64(expected.GetH()), angle)
	}

	angle, scale, err := estimateGuideAngleAndScale(float64(template.Bounds().Dx()), float64(template.Bounds().Dy()), float64(expected.GetW()), float64(expected.GetH()))
	if err > math.Max(6.0, 0.12*(float64(expected.GetW())+float64(expected.GetH()))) {
		angle = expectedGuideAngleHint(scenario)
		scale = 1.0
	}
	if math.Abs(angle) < 0.5 {
		return base
	}

	if math.Abs(expectedGuideAngleHint(scenario)) < 0.5 {
		angle = chooseGuideAngleSign(source, expected, template, scale, angle)
	} else if expectedGuideAngleHint(scenario) < 0 {
		angle = -math.Abs(angle)
	} else {
		angle = math.Abs(angle)
	}

	cx, cy := rectCenter(expected)
	guideW := float64(template.Bounds().Dx()) * scale
	guideH := float64(template.Bounds().Dy()) * scale
	if guideW < 2 || guideH < 2 {
		return base
	}
	return rotatedRectQuad(cx, cy, guideW, guideH, angle)
}

func shouldRotateExpectedGuideZone(scenario findBenchScenario) bool {
	return strings.EqualFold(strings.TrimSpace(scenario.scenarioTypeID), "scale_rotate_sweep")
}

func expectedGuideAngleHint(scenario findBenchScenario) float64 {
	if scenario.transformKind == "rotate" && math.Abs(scenario.transformA) > 0.5 {
		return scenario.transformA
	}
	switch ((scenario.rotation % 360) + 360) % 360 {
	case 90:
		return 90
	case 270:
		return -90
	default:
		return 0
	}
}

func estimateGuideAngleAndScale(templateW, templateH, observedW, observedH float64) (angleDeg float64, scale float64, err float64) {
	if templateW < 1 || templateH < 1 || observedW < 1 || observedH < 1 {
		return 0, 1, math.MaxFloat64
	}
	bestErr := math.MaxFloat64
	bestAngle := 0.0
	bestScale := 1.0
	for angle := 0.0; angle <= 89.5; angle += 0.5 {
		theta := angle * math.Pi / 180.0
		cosT := math.Abs(math.Cos(theta))
		sinT := math.Abs(math.Sin(theta))
		bboxW := templateW*cosT + templateH*sinT
		bboxH := templateW*sinT + templateH*cosT
		if bboxW < 1 || bboxH < 1 {
			continue
		}
		scaleX := observedW / bboxW
		scaleY := observedH / bboxH
		candidateScale := (scaleX + scaleY) / 2.0
		predW := bboxW * candidateScale
		predH := bboxH * candidateScale
		candidateErr := math.Abs(predW-observedW) + math.Abs(predH-observedH)
		if candidateErr < bestErr {
			bestErr = candidateErr
			bestAngle = angle
			bestScale = candidateScale
		}
	}
	return bestAngle, bestScale, bestErr
}

func chooseGuideAngleSign(source *image.Gray, expected *pb.Rect, template *image.Gray, scale float64, angleMag float64) float64 {
	if source == nil || expected == nil || template == nil || angleMag < 0.5 {
		return angleMag
	}
	posScore := guideTemplateCropScore(source, expected, template, scale, angleMag)
	negScore := guideTemplateCropScore(source, expected, template, scale, -angleMag)
	if negScore < posScore {
		return -angleMag
	}
	return angleMag
}

func guideTemplateCropScore(source *image.Gray, expected *pb.Rect, template *image.Gray, scale float64, angleDeg float64) float64 {
	if source == nil || expected == nil || template == nil {
		return math.MaxFloat64
	}
	scaled := scaleGrayNearestBench(template, scale)
	rotated := rotateGrayBilinearBench(scaled, angleDeg, 128)
	target := cropGray(source, image.Rect(int(expected.GetX()), int(expected.GetY()), int(expected.GetX()+expected.GetW()), int(expected.GetY()+expected.GetH())))
	if target.Bounds().Dx() < 1 || target.Bounds().Dy() < 1 {
		return math.MaxFloat64
	}
	candidate := resizeGrayNearest(rotated, target.Bounds().Dx(), target.Bounds().Dy())
	if candidate == nil {
		return math.MaxFloat64
	}
	var total float64
	pixels := target.Bounds().Dx() * target.Bounds().Dy()
	if pixels <= 0 {
		return math.MaxFloat64
	}
	for y := 0; y < target.Bounds().Dy(); y++ {
		for x := 0; x < target.Bounds().Dx(); x++ {
			tv := target.Pix[target.PixOffset(x, y)]
			cv := candidate.Pix[candidate.PixOffset(x, y)]
			total += math.Abs(float64(tv) - float64(cv))
		}
	}
	return total / float64(pixels)
}

func inferGlobalZoneGuideAngle(source *image.Gray, queries []findBenchAttemptQueryVisual) (float64, bool) {
	for _, query := range queries {
		if query.Template == nil || query.Found == nil {
			continue
		}
		angle, _, ok := inferGuideAngleAndScaleFromRect(source, query.Found, query.Template)
		if ok {
			return angle, true
		}
	}
	return 0, false
}

func inferGuideAngleAndScaleFromRect(source *image.Gray, rect *pb.Rect, template *image.Gray) (float64, float64, bool) {
	if source == nil || rect == nil || template == nil {
		return 0, 1, false
	}
	target := cropGray(source, image.Rect(int(rect.GetX()), int(rect.GetY()), int(rect.GetX()+rect.GetW()), int(rect.GetY()+rect.GetH())))
	if target.Bounds().Dx() < 4 || target.Bounds().Dy() < 4 {
		return 0, 1, false
	}

	type guideCandidate struct {
		angle float64
		scale float64
		score float64
	}
	scoreAt := func(angle float64) guideCandidate {
		candScale := estimateGuideScaleForAngle(
			float64(template.Bounds().Dx()),
			float64(template.Bounds().Dy()),
			float64(rect.GetW()),
			float64(rect.GetH()),
			angle,
		)
		return guideCandidate{
			angle: angle,
			scale: candScale,
			score: guideTemplateCropScore(source, rect, template, candScale, angle),
		}
	}

	best := scoreAt(0)
	zero := best
	for angle := -180.0; angle <= 180.0; angle += 10.0 {
		candidate := scoreAt(angle)
		if candidate.score < best.score {
			best = candidate
		}
	}
	refined := best
	for angle := best.angle - 10.0; angle <= best.angle+10.0; angle += 1.0 {
		candidate := scoreAt(angle)
		if candidate.score < refined.score {
			refined = candidate
		}
	}
	improvement := 0.0
	if zero.score > 0 {
		improvement = (zero.score - refined.score) / zero.score
	}
	if math.Abs(refined.angle) < 3.0 || improvement < 0.05 {
		return 0, 1, false
	}
	return normalizeGuideAngle(refined.angle), refined.scale, true
}

func estimateGuideScaleForAngle(templateW, templateH, observedW, observedH, angleDeg float64) float64 {
	theta := math.Abs(angleDeg) * math.Pi / 180.0
	cosT := math.Abs(math.Cos(theta))
	sinT := math.Abs(math.Sin(theta))
	bboxW := templateW*cosT + templateH*sinT
	bboxH := templateW*sinT + templateH*cosT
	if bboxW < 1 || bboxH < 1 {
		return 1.0
	}
	scaleX := observedW / bboxW
	scaleY := observedH / bboxH
	candidateScale := (scaleX + scaleY) / 2.0
	if candidateScale <= 0 {
		return 1.0
	}
	return candidateScale
}

func normalizeGuideAngle(angle float64) float64 {
	for angle <= -180 {
		angle += 360
	}
	for angle > 180 {
		angle -= 360
	}
	if angle > 90 {
		angle -= 180
	}
	if angle <= -90 {
		angle += 180
	}
	return angle
}

func drawMatchDetailLabel(canvas *image.RGBA, query findBenchAttemptQueryVisual, scenario findBenchScenario, border color.RGBA) {
	if canvas == nil || query.Found == nil {
		return
	}
	b := canvas.Bounds()
	label := strings.ToLower(strings.TrimSpace(query.Label))
	if label == "" {
		label = "target"
	}
	templateDims := grayImageDimensionsProto(query.Template)
	target, dx, dy, _, lim, _, _ := bestExpectedCandidateMetrics(
		query.Found,
		query.Expected,
		query.Alternates,
		templateDims,
		scenario.tolerance,
		scenario.maxAreaRatio,
		scenario.allowPartial,
	)
	matchType := "region"
	if regionActsAsZone(target, templateDims) {
		matchType = "zone"
	}
	textScale := 1
	labelH := 8*textScale + 2
	pad := 4
	maxPanelW := maxInt(96, minInt(240, b.Dx()-8))
	if maxPanelW <= pad*2 {
		return
	}
	maxTextChars := maxInt(10, (maxPanelW-pad*2)/(6*textScale)-1)
	titleText := truncateUpper(fmt.Sprintf("%s %s", label, matchType), maxTextChars)
	detailText := truncateUpper(fmt.Sprintf("dx=%.0f dy=%.0f lim=%.0f", dx, dy, lim), maxTextChars)
	panelW := maxInt(len(strings.ToUpper(titleText))*6*textScale, len(strings.ToUpper(detailText))*6*textScale)
	panelW += pad * 2
	panelW = minInt(panelW, maxPanelW)
	panelH := (labelH * 2) + pad*3

	x := int(query.Found.GetX())
	y := int(query.Found.GetY()) - panelH - 6
	if x+panelW > b.Max.X {
		x = b.Max.X - panelW
	}
	if x < b.Min.X {
		x = b.Min.X
	}
	if y < b.Min.Y {
		y = minInt(b.Max.Y-panelH, int(query.Found.GetY()+query.Found.GetH())+6)
	}
	if y < b.Min.Y {
		y = b.Min.Y
	}
	panel := image.Rect(x, y, x+panelW, y+panelH)
	draw.Draw(canvas, panel, &image.Uniform{C: color.RGBA{R: 12, G: 18, B: 30, A: 224}}, image.Point{}, draw.Over)

	drawTinyText(canvas, x+pad, y+pad, titleText, color.RGBA{R: 232, G: 244, B: 255, A: 255}, textScale)
	drawTinyText(canvas, x+pad, y+pad+labelH, detailText, color.RGBA{R: 174, G: 219, B: 255, A: 255}, textScale)
	drawRectOutline(
		canvas,
		&pb.Rect{X: int32(panel.Min.X), Y: int32(panel.Min.Y), W: int32(panel.Dx()), H: int32(panel.Dy())},
		border,
		2,
	)
}

func boolWord(v bool) string {
	if v {
		return "YES"
	}
	return "NO"
}

func parseEnvInt(name string, fallback int) int {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return fallback
	}
	v, err := strconv.Atoi(raw)
	if err != nil {
		return fallback
	}
	return v
}

func parseEnvDuration(name string, fallback time.Duration) time.Duration {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return fallback
	}
	v, err := time.ParseDuration(raw)
	if err != nil || v <= 0 {
		return fallback
	}
	return v
}

func envFlagTrue(v string) bool {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

func sanitizeFileToken(in string) string {
	if in == "" {
		return "untitled"
	}
	var b strings.Builder
	lastDash := false
	for _, r := range in {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'):
			b.WriteRune(r)
			lastDash = false
		case r == '-' || r == '_' || r == '.':
			b.WriteRune(r)
			lastDash = false
		default:
			if !lastDash {
				b.WriteRune('-')
				lastDash = true
			}
		}
	}
	out := strings.Trim(b.String(), "-")
	if out == "" {
		return "untitled"
	}
	return out
}

func stableScenarioImageToken(in string) string {
	name := strings.TrimSpace(in)
	if name == "" {
		return "scenario"
	}
	name = scenarioSeedSuffixPattern.ReplaceAllString(name, "")
	if m := scenarioStablePathPattern.FindStringSubmatch(name); len(m) == 2 {
		return sanitizeFileToken(strings.TrimSpace(m[1]))
	}
	return sanitizeFileToken(name)
}

func truncateUpper(in string, maxLen int) string {
	upper := strings.ToUpper(in)
	if len(upper) <= maxLen || maxLen < 4 {
		return upper
	}
	return upper[:maxLen-3] + "..."
}

func drawTinyText(img *image.RGBA, x, y int, text string, c color.RGBA, scale int) int {
	if scale < 1 {
		scale = 1
	}
	penX := x
	upper := strings.ToUpper(text)
	for _, r := range upper {
		g, ok := tinyFont5x7[r]
		if !ok {
			g = tinyFont5x7['?']
		}
		drawTinyGlyph(img, penX, y, g, c, scale)
		penX += 6 * scale
	}
	return 8 * scale
}

func drawTinyGlyph(img *image.RGBA, x, y int, rows [7]uint8, c color.RGBA, scale int) {
	b := img.Bounds()
	for row := 0; row < 7; row++ {
		mask := rows[row]
		for col := 0; col < 5; col++ {
			if mask&(1<<(4-col)) == 0 {
				continue
			}
			for dy := 0; dy < scale; dy++ {
				for dx := 0; dx < scale; dx++ {
					setRGBAIfInBounds(img, b, x+col*scale+dx, y+row*scale+dy, c)
				}
			}
		}
	}
}

var tinyFont5x7 = map[rune][7]uint8{
	' ': {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	'.': {0x00, 0x00, 0x00, 0x00, 0x00, 0x0C, 0x0C},
	',': {0x00, 0x00, 0x00, 0x00, 0x0C, 0x0C, 0x08},
	':': {0x00, 0x0C, 0x0C, 0x00, 0x0C, 0x0C, 0x00},
	'-': {0x00, 0x00, 0x00, 0x1F, 0x00, 0x00, 0x00},
	'_': {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1F},
	'/': {0x01, 0x02, 0x04, 0x08, 0x10, 0x00, 0x00},
	'=': {0x00, 0x00, 0x1F, 0x00, 0x1F, 0x00, 0x00},
	'%': {0x19, 0x19, 0x02, 0x04, 0x08, 0x13, 0x13},
	'(': {0x02, 0x04, 0x08, 0x08, 0x08, 0x04, 0x02},
	')': {0x08, 0x04, 0x02, 0x02, 0x02, 0x04, 0x08},
	'?': {0x0E, 0x11, 0x01, 0x02, 0x04, 0x00, 0x04},

	'0': {0x0E, 0x11, 0x13, 0x15, 0x19, 0x11, 0x0E},
	'1': {0x04, 0x0C, 0x04, 0x04, 0x04, 0x04, 0x0E},
	'2': {0x0E, 0x11, 0x01, 0x06, 0x08, 0x10, 0x1F},
	'3': {0x0E, 0x11, 0x01, 0x06, 0x01, 0x11, 0x0E},
	'4': {0x02, 0x06, 0x0A, 0x12, 0x1F, 0x02, 0x02},
	'5': {0x1F, 0x10, 0x1E, 0x01, 0x01, 0x11, 0x0E},
	'6': {0x06, 0x08, 0x10, 0x1E, 0x11, 0x11, 0x0E},
	'7': {0x1F, 0x11, 0x01, 0x02, 0x04, 0x04, 0x04},
	'8': {0x0E, 0x11, 0x11, 0x0E, 0x11, 0x11, 0x0E},
	'9': {0x0E, 0x11, 0x11, 0x0F, 0x01, 0x02, 0x0C},

	'A': {0x0E, 0x11, 0x11, 0x1F, 0x11, 0x11, 0x11},
	'B': {0x1E, 0x11, 0x11, 0x1E, 0x11, 0x11, 0x1E},
	'C': {0x0E, 0x11, 0x10, 0x10, 0x10, 0x11, 0x0E},
	'D': {0x1C, 0x12, 0x11, 0x11, 0x11, 0x12, 0x1C},
	'E': {0x1F, 0x10, 0x10, 0x1E, 0x10, 0x10, 0x1F},
	'F': {0x1F, 0x10, 0x10, 0x1E, 0x10, 0x10, 0x10},
	'G': {0x0E, 0x11, 0x10, 0x17, 0x11, 0x11, 0x0E},
	'H': {0x11, 0x11, 0x11, 0x1F, 0x11, 0x11, 0x11},
	'I': {0x0E, 0x04, 0x04, 0x04, 0x04, 0x04, 0x0E},
	'J': {0x07, 0x02, 0x02, 0x02, 0x02, 0x12, 0x0C},
	'K': {0x11, 0x12, 0x14, 0x18, 0x14, 0x12, 0x11},
	'L': {0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x1F},
	'M': {0x11, 0x1B, 0x15, 0x15, 0x11, 0x11, 0x11},
	'N': {0x11, 0x19, 0x19, 0x15, 0x13, 0x13, 0x11},
	'O': {0x0E, 0x11, 0x11, 0x11, 0x11, 0x11, 0x0E},
	'P': {0x1E, 0x11, 0x11, 0x1E, 0x10, 0x10, 0x10},
	'Q': {0x0E, 0x11, 0x11, 0x11, 0x15, 0x12, 0x0D},
	'R': {0x1E, 0x11, 0x11, 0x1E, 0x14, 0x12, 0x11},
	'S': {0x0F, 0x10, 0x10, 0x0E, 0x01, 0x01, 0x1E},
	'T': {0x1F, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04},
	'U': {0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x0E},
	'V': {0x11, 0x11, 0x11, 0x11, 0x0A, 0x0A, 0x04},
	'W': {0x11, 0x11, 0x11, 0x15, 0x15, 0x1B, 0x11},
	'X': {0x11, 0x11, 0x0A, 0x04, 0x0A, 0x11, 0x11},
	'Y': {0x11, 0x11, 0x0A, 0x04, 0x04, 0x04, 0x04},
	'Z': {0x1F, 0x01, 0x02, 0x04, 0x08, 0x10, 0x1F},
}

func TestSelectVisualExpectedTarget_PrefersPrimaryExpectedRect(t *testing.T) {
	query := findBenchAttemptQueryVisual{
		Expected: &pb.Rect{X: 120, Y: 80, W: 92, H: 88},
		Alternates: []*pb.Rect{
			{X: 96, Y: 54, W: 228, H: 154},
		},
		Found:    &pb.Rect{X: 124, Y: 84, W: 90, H: 86},
		Template: buildBenchPattern("orbtex", 64),
	}

	got := selectVisualExpectedTarget(query, findBenchScenario{scenarioTypeID: "orb_feature_rich", tolerance: 0.22, maxAreaRatio: 2.0})
	if got == nil {
		t.Fatalf("expected visual target")
	}
	if got.GetX() != 120 || got.GetY() != 80 || got.GetW() != 92 || got.GetH() != 88 {
		t.Fatalf("expected visual target to stay aligned with primary expected rect, got=%+v", got)
	}
}

func TestResolveExpectedGuideQuad_ZoneStaysAxisAlignedForOrbFeatureRich(t *testing.T) {
	template := buildBenchPattern("orbtex", 64)
	source := image.NewGray(image.Rect(0, 0, 360, 240))
	setRect(source, 0, 0, 360, 240, 180)
	blitGray(source, template, 120, 48)

	expected := &pb.Rect{X: 96, Y: 28, W: 228, H: 154}
	found := &pb.Rect{X: 120, Y: 48, W: int32(template.Bounds().Dx()), H: int32(template.Bounds().Dy())}
	quad := resolveExpectedGuideQuad(source, expected, template, findBenchScenario{scenarioTypeID: "orb_feature_rich"}, found, 0, false)
	if len(quad) != 4 {
		t.Fatalf("expected 4-point guide quad, got=%d", len(quad))
	}
	if angle := math.Abs(quadEdgeAngleDegrees(quad)); angle > 1.0 {
		t.Fatalf("expected orb feature guide to stay axis-aligned, got angle=%.2f", angle)
	}
}

func TestResolveExpectedGuideQuad_RotatesZoneForScaleRotateSweep(t *testing.T) {
	template := buildBenchPattern("photo", 56)
	rotated := rotateGrayBilinearBench(template, 28, 128)
	source := image.NewGray(image.Rect(0, 0, 420, 300))
	setRect(source, 0, 0, 420, 300, 200)
	blitGray(source, rotated, 126, 84)

	found := &pb.Rect{
		X: 126,
		Y: 84,
		W: int32(rotated.Bounds().Dx()),
		H: int32(rotated.Bounds().Dy()),
	}
	expected := &pb.Rect{X: 102, Y: 60, W: int32(rotated.Bounds().Dx() + 72), H: int32(rotated.Bounds().Dy() + 56)}
	quad := resolveExpectedGuideQuad(source, expected, template, findBenchScenario{scenarioTypeID: "scale_rotate_sweep"}, found, 0, false)
	if len(quad) != 4 {
		t.Fatalf("expected 4-point guide quad, got=%d", len(quad))
	}
	angle := math.Abs(quadEdgeAngleDegrees(quad))
	if angle < 8.0 {
		t.Fatalf("expected rotated guide for scale_rotate_sweep zone, got angle=%.2f", angle)
	}
}

func quadEdgeAngleDegrees(quad []benchPointF) float64 {
	if len(quad) < 2 {
		return 0
	}
	angle := math.Atan2(quad[1].Y-quad[0].Y, quad[1].X-quad[0].X) * 180 / math.Pi
	for angle <= -90 {
		angle += 180
	}
	for angle > 90 {
		angle -= 180
	}
	return angle
}
