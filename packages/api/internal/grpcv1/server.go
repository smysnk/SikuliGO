package grpcv1

import (
	"context"
	"errors"
	"fmt"
	"image"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/smysnk/sikuligo/internal/cv"
	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"github.com/smysnk/sikuligo/pkg/sikuli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedSikuliServiceServer
	captureScreen       func(context.Context, string) (*sikuli.Image, error)
	clickOnScreen       func(int, int, sikuli.InputOptions) error
	newFinder           func(*sikuli.Image) (*sikuli.Finder, error)
	newFinderWithEngine func(*sikuli.Image, cv.MatcherEngine) (*sikuli.Finder, error)
	listScreens         func(context.Context) ([]sikuli.Screen, error)
}

type ServerOption func(*Server)

var debugEnabled = func() bool {
	v := strings.TrimSpace(os.Getenv("SIKULI_DEBUG"))
	switch strings.ToLower(v) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}()

func debugLogf(format string, args ...any) {
	if !debugEnabled {
		return
	}
	log.Printf("[debug] "+format, args...)
}

func infoLogf(format string, args ...any) {
	log.Printf("[info] "+format, args...)
}

func errorLogf(format string, args ...any) {
	log.Printf("[error] "+format, args...)
}

var captureScreenFn = captureScreenImage

var clickOnScreenFn = func(x, y int, opts sikuli.InputOptions) error {
	c := sikuli.NewInputController()
	return c.Click(x, y, opts)
}

var newFinderFn = sikuli.NewFinder
var newFinderWithEngineFn = newFinderWithEngine

func WithCaptureScreen(fn func(context.Context, string) (*sikuli.Image, error)) ServerOption {
	return func(s *Server) {
		s.captureScreen = fn
	}
}

func WithClickOnScreen(fn func(int, int, sikuli.InputOptions) error) ServerOption {
	return func(s *Server) {
		s.clickOnScreen = fn
	}
}

func WithFinderFactory(fn func(*sikuli.Image) (*sikuli.Finder, error)) ServerOption {
	return func(s *Server) {
		s.newFinder = fn
	}
}

func WithFinderWithEngineFactory(fn func(*sikuli.Image, cv.MatcherEngine) (*sikuli.Finder, error)) ServerOption {
	return func(s *Server) {
		s.newFinderWithEngine = fn
	}
}

func NewServer(opts ...ServerOption) *Server {
	s := &Server{}
	for _, opt := range opts {
		if opt != nil {
			opt(s)
		}
	}
	return s
}

func (s *Server) finder(source *sikuli.Image) (*sikuli.Finder, error) {
	if s != nil && s.newFinder != nil {
		return s.newFinder(source)
	}
	return newFinderFn(source)
}

func (s *Server) finderWithEngine(source *sikuli.Image, engine cv.MatcherEngine) (*sikuli.Finder, error) {
	if s != nil && s.newFinderWithEngine != nil {
		return s.newFinderWithEngine(source, engine)
	}
	return newFinderWithEngineFn(source, engine)
}

func (s *Server) capture(ctx context.Context, name string) (*sikuli.Image, error) {
	if s != nil && s.captureScreen != nil {
		return s.captureScreen(ctx, name)
	}
	return captureScreenFn(ctx, name)
}

func (s *Server) clickAt(x, y int, opts sikuli.InputOptions) error {
	if s != nil && s.clickOnScreen != nil {
		return s.clickOnScreen(x, y, opts)
	}
	return clickOnScreenFn(x, y, opts)
}

func (s *Server) ListScreens(ctx context.Context, _ *pb.ListScreensRequest) (*pb.ListScreensResponse, error) {
	screens, err := s.screens(ctx)
	if err != nil {
		return nil, mapStatusError(err)
	}
	out := make([]*pb.ScreenDescriptor, 0, len(screens))
	for _, screen := range screens {
		out = append(out, toProtoScreenDescriptor(screen))
	}
	return &pb.ListScreensResponse{Screens: out}, nil
}

func (s *Server) GetPrimaryScreen(ctx context.Context, _ *pb.GetPrimaryScreenRequest) (*pb.GetPrimaryScreenResponse, error) {
	screen, err := s.primaryScreen(ctx)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.GetPrimaryScreenResponse{Screen: toProtoScreenDescriptor(screen)}, nil
}

func (s *Server) CaptureScreen(ctx context.Context, req *pb.CaptureScreenRequest) (*pb.CaptureScreenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	full, err := s.capture(ctx, "screen")
	if err != nil {
		return nil, mapStatusError(err)
	}
	screen, region, hasRegion, err := s.resolveCaptureTarget(ctx, req.ScreenId != nil, int(req.GetScreenId()), req.GetRegion())
	if err != nil {
		return nil, mapStatusError(err)
	}
	if hasRegion {
		full, err = full.Crop(region.Rect)
		if err != nil {
			return nil, mapStatusError(err)
		}
	}
	var screenProto *pb.ScreenDescriptor
	if req.ScreenId != nil {
		screenProto = toProtoScreenDescriptor(screen)
	}
	return &pb.CaptureScreenResponse{
		Image:  toProtoGrayImage(full),
		Screen: screenProto,
	}, nil
}

func (s *Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	source, pattern, err := findRequestParts(req)
	if err != nil {
		return nil, mapStatusError(err)
	}
	engine, err := matcherEngineFromFindRequest(req)
	if err != nil {
		return nil, mapStatusError(err)
	}
	f, err := s.finderWithEngine(source, engine)
	if err != nil {
		return nil, mapStatusError(err)
	}
	match, err := f.Find(pattern)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.FindResponse{Match: toProtoMatch(match)}, nil
}

func (s *Server) FindAll(ctx context.Context, req *pb.FindRequest) (*pb.FindAllResponse, error) {
	source, pattern, err := findRequestParts(req)
	if err != nil {
		return nil, mapStatusError(err)
	}
	engine, err := matcherEngineFromFindRequest(req)
	if err != nil {
		return nil, mapStatusError(err)
	}
	f, err := s.finderWithEngine(source, engine)
	if err != nil {
		return nil, mapStatusError(err)
	}
	matches, err := f.FindAll(pattern)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.FindAllResponse{Matches: toProtoMatches(matches)}, nil
}

func (s *Server) FindOnScreen(ctx context.Context, req *pb.FindOnScreenRequest) (*pb.FindResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	engine, err := matcherEngineFromScreenOptions(req.GetOpts())
	if err != nil {
		return nil, mapStatusError(err)
	}
	match, err := s.findOnScreenOnce(ctx, req.GetPattern(), req.GetOpts(), engine)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.FindResponse{Match: toProtoMatch(match)}, nil
}

func (s *Server) ExistsOnScreen(ctx context.Context, req *pb.ExistsOnScreenRequest) (*pb.ExistsOnScreenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	engine, err := matcherEngineFromScreenOptions(req.GetOpts())
	if err != nil {
		return nil, mapStatusError(err)
	}
	opts := screenQueryFromProto(req.GetOpts())
	match, ok, err := sikuli.SearchExists(func() (sikuli.Match, error) {
		return s.findOnScreenOnce(ctx, req.GetPattern(), req.GetOpts(), engine)
	}, opts.Timeout, opts.Interval)
	if err != nil {
		return nil, mapStatusError(err)
	}
	if !ok {
		return &pb.ExistsOnScreenResponse{Exists: false}, nil
	}
	return &pb.ExistsOnScreenResponse{Exists: true, Match: toProtoMatch(match)}, nil
}

func (s *Server) WaitOnScreen(ctx context.Context, req *pb.WaitOnScreenRequest) (*pb.FindResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	engine, err := matcherEngineFromScreenOptions(req.GetOpts())
	if err != nil {
		return nil, mapStatusError(err)
	}
	opts := screenQueryFromProto(req.GetOpts())
	if opts.Timeout <= 0 {
		opts.Timeout = time.Duration(sikuli.DefaultAutoWaitTimeout * float64(time.Second))
	}
	match, err := sikuli.SearchWait(func() (sikuli.Match, error) {
		return s.findOnScreenOnce(ctx, req.GetPattern(), req.GetOpts(), engine)
	}, opts.Timeout, opts.Interval)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.FindResponse{Match: toProtoMatch(match)}, nil
}

func (s *Server) ClickOnScreen(ctx context.Context, req *pb.ClickOnScreenRequest) (*pb.FindResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	waitReq := &pb.WaitOnScreenRequest{
		Pattern: req.GetPattern(),
		Opts:    req.GetOpts(),
	}
	debugLogf("click_on_screen.start")
	found, err := s.WaitOnScreen(ctx, waitReq)
	if err != nil {
		errorLogf("click_on_screen.wait.error err=%v", err)
		return nil, err
	}
	match := found.GetMatch()
	if match == nil || match.GetTarget() == nil {
		debugLogf("click_on_screen.match_missing_target")
		return nil, status.Error(codes.Internal, "match target missing")
	}
	clickStart := time.Now()
	if err := s.clickAt(int(match.GetTarget().GetX()), int(match.GetTarget().GetY()), inputOptionsFromProto(req.GetClickOpts())); err != nil {
		debugLogf("click_on_screen.click.error duration=%s err=%v", time.Since(clickStart), err)
		return nil, mapStatusError(err)
	}
	debugLogf("click_on_screen.click.ok duration=%s target=(%d,%d)", time.Since(clickStart), match.GetTarget().GetX(), match.GetTarget().GetY())
	return found, nil
}

func (s *Server) ReadText(_ context.Context, req *pb.ReadTextRequest) (*pb.ReadTextResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	source, err := imageFromProto(req.GetSource(), "source")
	if err != nil {
		return nil, mapStatusError(err)
	}
	f, err := s.finder(source)
	if err != nil {
		return nil, mapStatusError(err)
	}
	text, err := f.ReadText(ocrParamsFromProto(req.GetParams()))
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ReadTextResponse{Text: text}, nil
}

type screenQuery struct {
	Region      sikuli.Region
	Timeout     time.Duration
	Interval    time.Duration
	ScreenID    int
	HasScreenID bool
}

func screenQueryFromProto(in *pb.ScreenQueryOptions) screenQuery {
	out := screenQuery{
		Region:   sikuli.NewRegion(0, 0, 0, 0),
		Timeout:  0,
		Interval: time.Millisecond * 100,
	}
	if in == nil {
		return out
	}
	out.Region = regionFromProto(in.GetRegion())
	if in.TimeoutMillis != nil {
		out.Timeout = durationMillis(in.GetTimeoutMillis())
		if out.Timeout < 0 {
			out.Timeout = 0
		}
	}
	if in.IntervalMillis != nil {
		out.Interval = durationMillis(in.GetIntervalMillis())
		if out.Interval <= 0 {
			out.Interval = time.Millisecond * 100
		}
	}
	if in.ScreenId != nil {
		out.ScreenID = int(in.GetScreenId())
		out.HasScreenID = true
	}
	return out
}

func matcherEngineFromFindRequest(req *pb.FindRequest) (cv.MatcherEngine, error) {
	if req == nil {
		return cv.MatcherEngineHybrid, nil
	}
	return matcherEngineFromProto(req.GetMatcherEngine())
}

func matcherEngineFromScreenOptions(opts *pb.ScreenQueryOptions) (cv.MatcherEngine, error) {
	if opts == nil {
		return cv.MatcherEngineHybrid, nil
	}
	return matcherEngineFromProto(opts.GetMatcherEngine())
}

func matcherEngineFromProto(in pb.MatcherEngine) (cv.MatcherEngine, error) {
	switch in {
	case pb.MatcherEngine_MATCHER_ENGINE_UNSPECIFIED:
		return cv.MatcherEngineHybrid, nil
	case pb.MatcherEngine_MATCHER_ENGINE_TEMPLATE:
		return cv.MatcherEngineTemplate, nil
	case pb.MatcherEngine_MATCHER_ENGINE_ORB:
		return cv.MatcherEngineORB, nil
	case pb.MatcherEngine_MATCHER_ENGINE_AKAZE:
		return cv.MatcherEngineAKAZE, nil
	case pb.MatcherEngine_MATCHER_ENGINE_BRISK:
		return cv.MatcherEngineBRISK, nil
	case pb.MatcherEngine_MATCHER_ENGINE_KAZE:
		return cv.MatcherEngineKAZE, nil
	case pb.MatcherEngine_MATCHER_ENGINE_SIFT:
		return cv.MatcherEngineSIFT, nil
	case pb.MatcherEngine_MATCHER_ENGINE_HYBRID:
		return cv.MatcherEngineHybrid, nil
	default:
		return "", fmt.Errorf("%w: unsupported matcher engine enum value=%d", sikuli.ErrInvalidTarget, in)
	}
}

func newFinderWithEngine(source *sikuli.Image, engine cv.MatcherEngine) (*sikuli.Finder, error) {
	f, err := sikuli.NewFinder(source)
	if err != nil {
		return nil, err
	}
	matcher, err := cv.NewMatcherForEngine(engine)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", sikuli.ErrInvalidTarget, err)
	}
	f.SetMatcher(matcher)
	return f, nil
}

func (s *Server) findOnScreenOnce(ctx context.Context, patternReq *pb.Pattern, optsReq *pb.ScreenQueryOptions, engine cv.MatcherEngine) (sikuli.Match, error) {
	start := time.Now()
	traceID := traceIDFromContext(ctx)
	infoLogf("find_on_screen.start engine=%s opencv=%t trace_id=%s", engine, cv.OpenCVEnabled(), traceID)
	pattern, err := patternFromProto(patternReq)
	if err != nil {
		debugLogf("find_on_screen.pattern.error trace_id=%s err=%v", traceID, err)
		return sikuli.Match{}, err
	}
	captureStart := time.Now()
	debugLogf("find_on_screen.capture.start trace_id=%s", traceID)
	source, err := s.capture(ctx, "screen")
	if err != nil {
		debugLogf("find_on_screen.capture.error trace_id=%s duration=%s total=%s err=%v", traceID, time.Since(captureStart), time.Since(start), err)
		return sikuli.Match{}, err
	}
	debugLogf(
		"find_on_screen.capture.ok trace_id=%s duration=%s width=%d height=%d",
		traceID,
		time.Since(captureStart),
		source.Width(),
		source.Height(),
	)

	opts := screenQueryFromProto(optsReq)
	if opts.HasScreenID {
		resolved, err := s.resolveSearchRegion(ctx, opts)
		if err != nil {
			debugLogf("find_on_screen.screen_region.error trace_id=%s err=%v", traceID, err)
			return sikuli.Match{}, err
		}
		opts.Region = resolved
	}
	matchStart := time.Now()
	patternW := pattern.Image().Width()
	patternH := pattern.Image().Height()
	deadlineRemaining := deadlineRemainingMillis(ctx)
	if opts.Region.Empty() {
		f, err := s.finderWithEngine(source, engine)
		if err != nil {
			debugLogf("find_on_screen.finder.error trace_id=%s err=%v", traceID, err)
			return sikuli.Match{}, err
		}
		sourceW := source.Width()
		sourceH := source.Height()
		positions := searchPositions(sourceW, sourceH, patternW, patternH)
		debugLogf(
			"find_on_screen.match.start trace_id=%s engine=%s source=%dx%d pattern=%dx%d positions=%d similarity=%.3f deadline_remaining_ms=%d region=full_screen",
			traceID,
			engine,
			sourceW,
			sourceH,
			patternW,
			patternH,
			positions,
			pattern.Similarity(),
			deadlineRemaining,
		)
		stopProgress := startMatchProgressLogger(ctx, traceID, engine, matchStart)
		m, err := f.Find(pattern)
		stopProgress()
		if err != nil {
			errorLogf("find_on_screen.match.error trace_id=%s duration=%s total=%s err=%v", traceID, time.Since(matchStart), time.Since(start), err)
			return sikuli.Match{}, err
		}
		debugLogf(
			"find_on_screen.match.ok trace_id=%s duration=%s total=%s rect=(%d,%d %dx%d) score=%.3f",
			traceID,
			time.Since(matchStart),
			time.Since(start),
			m.Rect.X,
			m.Rect.Y,
			m.Rect.W,
			m.Rect.H,
			m.Score,
		)
		return m, nil
	}
	regionSource, err := source.Crop(opts.Region.Rect)
	if err != nil {
		debugLogf("find_on_screen.region_crop.error trace_id=%s duration=%s total=%s err=%v", traceID, time.Since(matchStart), time.Since(start), err)
		return sikuli.Match{}, err
	}
	f, err := s.finderWithEngine(regionSource, engine)
	if err != nil {
		debugLogf("find_on_screen.region_finder.error trace_id=%s duration=%s total=%s err=%v", traceID, time.Since(matchStart), time.Since(start), err)
		return sikuli.Match{}, err
	}
	regionW := regionSource.Width()
	regionH := regionSource.Height()
	positions := searchPositions(regionW, regionH, patternW, patternH)
	debugLogf(
		"find_on_screen.region_match.start trace_id=%s engine=%s region_source=%dx%d pattern=%dx%d positions=%d similarity=%.3f deadline_remaining_ms=%d region=(%d,%d %dx%d)",
		traceID,
		engine,
		regionW,
		regionH,
		patternW,
		patternH,
		positions,
		pattern.Similarity(),
		deadlineRemaining,
		opts.Region.X,
		opts.Region.Y,
		opts.Region.W,
		opts.Region.H,
	)
	stopProgress := startMatchProgressLogger(ctx, traceID, engine, matchStart)
	m, err := f.Find(pattern)
	stopProgress()
	if err != nil {
		debugLogf("find_on_screen.region_match.error trace_id=%s duration=%s total=%s err=%v", traceID, time.Since(matchStart), time.Since(start), err)
		return sikuli.Match{}, err
	}
	debugLogf(
		"find_on_screen.region_match.ok trace_id=%s duration=%s total=%s rect=(%d,%d %dx%d) score=%.3f",
		traceID,
		time.Since(matchStart),
		time.Since(start),
		m.Rect.X,
		m.Rect.Y,
		m.Rect.W,
		m.Rect.H,
		m.Score,
	)
	return m, nil
}

func (s *Server) resolveSearchRegion(ctx context.Context, opts screenQuery) (sikuli.Region, error) {
	if !opts.HasScreenID {
		return opts.Region, nil
	}
	screen, err := s.screenByID(ctx, opts.ScreenID)
	if err != nil {
		return sikuli.Region{}, err
	}
	screenRegion := sikuli.NewRegion(screen.Bounds.X, screen.Bounds.Y, screen.Bounds.W, screen.Bounds.H)
	if opts.Region.Empty() {
		return screenRegion, nil
	}
	translated := opts.Region.Offset(screen.Bounds.X, screen.Bounds.Y)
	resolved := screenRegion.Intersection(translated)
	if resolved.Empty() {
		return sikuli.Region{}, fmt.Errorf("%w: screen region outside selected screen", sikuli.ErrInvalidTarget)
	}
	return resolved, nil
}

func (s *Server) resolveCaptureTarget(ctx context.Context, hasScreenID bool, screenID int, region *pb.Rect) (sikuli.Screen, sikuli.Region, bool, error) {
	if hasScreenID {
		screen, err := s.screenByID(ctx, screenID)
		if err != nil {
			return sikuli.Screen{}, sikuli.Region{}, false, err
		}
		screenRegion := sikuli.NewRegion(screen.Bounds.X, screen.Bounds.Y, screen.Bounds.W, screen.Bounds.H)
		if region == nil {
			return screen, screenRegion, true, nil
		}
		local := regionFromProto(region)
		if local.Empty() {
			return screen, screenRegion, true, nil
		}
		translated := local.Offset(screen.Bounds.X, screen.Bounds.Y)
		resolved := screenRegion.Intersection(translated)
		if resolved.Empty() {
			return sikuli.Screen{}, sikuli.Region{}, false, fmt.Errorf("%w: capture region outside selected screen", sikuli.ErrInvalidTarget)
		}
		return screen, resolved, true, nil
	}
	global := regionFromProto(region)
	if global.Empty() {
		return sikuli.Screen{}, sikuli.Region{}, false, nil
	}
	return sikuli.Screen{}, global, true, nil
}

func searchPositions(sourceW, sourceH, patternW, patternH int) int64 {
	x := sourceW - patternW + 1
	y := sourceH - patternH + 1
	if x < 0 || y < 0 {
		return 0
	}
	return int64(x) * int64(y)
}

func deadlineRemainingMillis(ctx context.Context) int64 {
	if ctx == nil {
		return -1
	}
	deadline, ok := ctx.Deadline()
	if !ok {
		return -1
	}
	return maxInt64(0, time.Until(deadline).Milliseconds())
}

func startMatchProgressLogger(ctx context.Context, traceID string, engine cv.MatcherEngine, started time.Time) func() {
	if !debugEnabled {
		return func() {}
	}
	done := make(chan struct{})
	var once sync.Once
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				debugLogf(
					"find_on_screen.match.progress trace_id=%s engine=%s elapsed=%s deadline_remaining_ms=%d",
					traceID,
					engine,
					time.Since(started),
					deadlineRemainingMillis(ctx),
				)
			case <-ctx.Done():
				debugLogf(
					"find_on_screen.match.context_done trace_id=%s engine=%s elapsed=%s err=%v",
					traceID,
					engine,
					time.Since(started),
					ctx.Err(),
				)
				return
			}
		}
	}()
	return func() {
		once.Do(func() {
			close(done)
		})
	}
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (s *Server) FindText(_ context.Context, req *pb.FindTextRequest) (*pb.FindTextResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	source, err := imageFromProto(req.GetSource(), "source")
	if err != nil {
		return nil, mapStatusError(err)
	}
	f, err := s.finder(source)
	if err != nil {
		return nil, mapStatusError(err)
	}
	matches, err := f.FindText(req.GetQuery(), ocrParamsFromProto(req.GetParams()))
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.FindTextResponse{Matches: toProtoTextMatches(matches)}, nil
}

func (s *Server) MoveMouse(_ context.Context, req *pb.MoveMouseRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.MoveMouse(int(req.GetX()), int(req.GetY()), inputOptionsFromProto(req.GetOpts())); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) Click(_ context.Context, req *pb.ClickRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.Click(int(req.GetX()), int(req.GetY()), inputOptionsFromProto(req.GetOpts())); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) MouseDown(_ context.Context, req *pb.ClickRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.MouseDown(int(req.GetX()), int(req.GetY()), inputOptionsFromProto(req.GetOpts())); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) MouseUp(_ context.Context, req *pb.ClickRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.MouseUp(int(req.GetX()), int(req.GetY()), inputOptionsFromProto(req.GetOpts())); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) TypeText(_ context.Context, req *pb.TypeTextRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.TypeText(req.GetText(), inputOptionsFromProto(req.GetOpts())); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) PasteText(_ context.Context, req *pb.TypeTextRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.Paste(req.GetText(), inputOptionsFromProto(req.GetOpts())); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) Hotkey(_ context.Context, req *pb.HotkeyRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.Hotkey(req.GetKeys()...); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) KeyDown(_ context.Context, req *pb.HotkeyRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.KeyDown(req.GetKeys()...); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) KeyUp(_ context.Context, req *pb.HotkeyRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.KeyUp(req.GetKeys()...); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) ScrollWheel(_ context.Context, req *pb.ScrollWheelRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewInputController()
	if err := c.Wheel(
		int(req.GetX()),
		int(req.GetY()),
		sikuli.WheelDirection(strings.ToLower(strings.TrimSpace(req.GetDirection()))),
		int(req.GetSteps()),
		inputOptionsFromProto(req.GetOpts()),
	); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) ObserveAppear(_ context.Context, req *pb.ObserveRequest) (*pb.ObserveResponse, error) {
	source, region, pattern, opts, err := observeRequestParts(req, true)
	if err != nil {
		return nil, mapStatusError(err)
	}
	c := sikuli.NewObserverController()
	events, err := c.ObserveAppear(source, region, pattern, opts)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ObserveResponse{Events: toProtoObserveEvents(events)}, nil
}

func (s *Server) ObserveVanish(_ context.Context, req *pb.ObserveRequest) (*pb.ObserveResponse, error) {
	source, region, pattern, opts, err := observeRequestParts(req, true)
	if err != nil {
		return nil, mapStatusError(err)
	}
	c := sikuli.NewObserverController()
	events, err := c.ObserveVanish(source, region, pattern, opts)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ObserveResponse{Events: toProtoObserveEvents(events)}, nil
}

func (s *Server) ObserveChange(_ context.Context, req *pb.ObserveChangeRequest) (*pb.ObserveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	source, err := imageFromProto(req.GetSource(), "source")
	if err != nil {
		return nil, mapStatusError(err)
	}
	region := regionFromProto(req.GetRegion())
	opts := observeOptionsFromProto(req.GetOpts())

	c := sikuli.NewObserverController()
	events, err := c.ObserveChange(source, region, opts)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ObserveResponse{Events: toProtoObserveEvents(events)}, nil
}

func (s *Server) OpenApp(_ context.Context, req *pb.AppActionRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewAppController()
	if err := c.Open(req.GetName(), req.GetArgs(), appOptionsFromProto(req.GetOpts())); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) FocusApp(_ context.Context, req *pb.AppActionRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewAppController()
	if err := c.Focus(req.GetName(), appOptionsFromProto(req.GetOpts())); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) CloseApp(_ context.Context, req *pb.AppActionRequest) (*pb.ActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewAppController()
	if err := c.Close(req.GetName(), appOptionsFromProto(req.GetOpts())); err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ActionResponse{}, nil
}

func (s *Server) IsAppRunning(_ context.Context, req *pb.AppActionRequest) (*pb.IsAppRunningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewAppController()
	running, err := c.IsRunning(req.GetName(), appOptionsFromProto(req.GetOpts()))
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.IsAppRunningResponse{Running: running}, nil
}

func (s *Server) ListWindows(_ context.Context, req *pb.AppActionRequest) (*pb.ListWindowsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewAppController()
	windows, err := c.ListWindows(req.GetName(), appOptionsFromProto(req.GetOpts()))
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ListWindowsResponse{Windows: windowsToProto(windows)}, nil
}

func (s *Server) FindWindows(_ context.Context, req *pb.WindowQueryRequest) (*pb.ListWindowsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewAppController()
	windows, err := c.FindWindows(req.GetName(), windowQueryFromProto(req.GetQuery()), appOptionsFromProto(req.GetOpts()))
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ListWindowsResponse{Windows: windowsToProto(windows)}, nil
}

func (s *Server) GetWindow(_ context.Context, req *pb.WindowQueryRequest) (*pb.GetWindowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewAppController()
	window, found, err := c.GetWindow(req.GetName(), windowQueryFromProto(req.GetQuery()), appOptionsFromProto(req.GetOpts()))
	if err != nil {
		return nil, mapStatusError(err)
	}
	if !found {
		return &pb.GetWindowResponse{Found: false}, nil
	}
	return &pb.GetWindowResponse{Found: true, Window: windowToProto(window)}, nil
}

func (s *Server) GetFocusedWindow(_ context.Context, req *pb.AppActionRequest) (*pb.GetWindowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	c := sikuli.NewAppController()
	window, found, err := c.FocusedWindow(req.GetName(), appOptionsFromProto(req.GetOpts()))
	if err != nil {
		return nil, mapStatusError(err)
	}
	if !found {
		return &pb.GetWindowResponse{Found: false}, nil
	}
	return &pb.GetWindowResponse{Found: true, Window: windowToProto(window)}, nil
}

func windowQueryFromProto(in *pb.WindowQuery) sikuli.WindowQuery {
	if in == nil {
		return sikuli.WindowQuery{}
	}
	index := 0
	if in.Index != nil {
		index = int(in.GetIndex())
	}
	return sikuli.WindowQuery{
		ID:            in.GetId(),
		TitleExact:    in.GetTitleExact(),
		TitleContains: in.GetTitleContains(),
		FocusedOnly:   in.GetFocusedOnly(),
		Index:         index,
	}
}

func windowToProto(in sikuli.Window) *pb.Window {
	return &pb.Window{
		Id:      in.ID,
		App:     in.App,
		Pid:     int32(in.PID),
		Title:   in.Title,
		Bounds:  &pb.Rect{X: int32(in.Bounds.X), Y: int32(in.Bounds.Y), W: int32(in.Bounds.W), H: int32(in.Bounds.H)},
		Focused: in.Focused,
	}
}

func windowsToProto(in []sikuli.Window) []*pb.Window {
	out := make([]*pb.Window, 0, len(in))
	for _, window := range in {
		out = append(out, windowToProto(window))
	}
	return out
}

func findRequestParts(req *pb.FindRequest) (*sikuli.Image, *sikuli.Pattern, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request is nil", sikuli.ErrInvalidTarget)
	}
	source, err := imageFromProto(req.GetSource(), "source")
	if err != nil {
		return nil, nil, err
	}
	pattern, err := patternFromProto(req.GetPattern())
	if err != nil {
		return nil, nil, err
	}
	return source, pattern, nil
}

func observeRequestParts(req *pb.ObserveRequest, patternRequired bool) (*sikuli.Image, sikuli.Region, *sikuli.Pattern, sikuli.ObserveOptions, error) {
	if req == nil {
		return nil, sikuli.Region{}, nil, sikuli.ObserveOptions{}, fmt.Errorf("%w: request is nil", sikuli.ErrInvalidTarget)
	}
	source, err := imageFromProto(req.GetSource(), "source")
	if err != nil {
		return nil, sikuli.Region{}, nil, sikuli.ObserveOptions{}, err
	}
	region := regionFromProto(req.GetRegion())
	opts := observeOptionsFromProto(req.GetOpts())
	if !patternRequired {
		return source, region, nil, opts, nil
	}
	pattern, err := patternFromProto(req.GetPattern())
	if err != nil {
		return nil, sikuli.Region{}, nil, sikuli.ObserveOptions{}, err
	}
	return source, region, pattern, opts, nil
}

func imageFromProto(in *pb.GrayImage, field string) (*sikuli.Image, error) {
	if in == nil {
		return nil, fmt.Errorf("%w: %s is nil", sikuli.ErrInvalidTarget, field)
	}
	w := int(in.GetWidth())
	h := int(in.GetHeight())
	if w <= 0 || h <= 0 {
		return nil, fmt.Errorf("%w: %s dimensions must be positive", sikuli.ErrInvalidTarget, field)
	}
	if got, want := len(in.GetPix()), w*h; got != want {
		return nil, fmt.Errorf("%w: %s pix length mismatch got=%d want=%d", sikuli.ErrInvalidTarget, field, got, want)
	}
	gray := image.NewGray(image.Rect(0, 0, w, h))
	copy(gray.Pix, in.GetPix())
	name := strings.TrimSpace(in.GetName())
	if name == "" {
		name = field
	}
	img, err := sikuli.NewImageFromGray(name, gray)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func grayFromProto(in *pb.GrayImage, field string) (*image.Gray, error) {
	img, err := imageFromProto(in, field)
	if err != nil {
		return nil, err
	}
	return img.Gray(), nil
}

func patternFromProto(in *pb.Pattern) (*sikuli.Pattern, error) {
	if in == nil {
		return nil, fmt.Errorf("%w: pattern is nil", sikuli.ErrInvalidTarget)
	}
	img, err := imageFromProto(in.GetImage(), "pattern.image")
	if err != nil {
		return nil, err
	}
	pattern, err := sikuli.NewPattern(img)
	if err != nil {
		return nil, err
	}
	if in.Exact != nil && in.GetExact() {
		pattern.Exact()
	} else if in.Similarity != nil {
		pattern.Similar(in.GetSimilarity())
	}
	if in.ResizeFactor != nil {
		pattern.Resize(in.GetResizeFactor())
	}
	if off := in.GetTargetOffset(); off != nil {
		pattern.TargetOffset(int(off.GetX()), int(off.GetY()))
	}
	if in.GetMask() != nil {
		mask, err := grayFromProto(in.GetMask(), "pattern.mask")
		if err != nil {
			return nil, err
		}
		if _, err := pattern.WithMask(mask); err != nil {
			return nil, err
		}
	}
	return pattern, nil
}

func ocrParamsFromProto(in *pb.OCRParams) sikuli.OCRParams {
	out := sikuli.OCRParams{}
	if in == nil {
		return out
	}
	out.Language = in.GetLanguage()
	out.TrainingDataPath = in.GetTrainingDataPath()
	if in.MinConfidence != nil {
		out.MinConfidence = in.GetMinConfidence()
	}
	if in.TimeoutMillis != nil {
		out.Timeout = durationMillis(in.GetTimeoutMillis())
	}
	if in.CaseSensitive != nil {
		out.CaseSensitive = in.GetCaseSensitive()
	}
	return out
}

func inputOptionsFromProto(in *pb.InputOptions) sikuli.InputOptions {
	out := sikuli.InputOptions{}
	if in == nil {
		return out
	}
	if in.DelayMillis != nil {
		out.Delay = durationMillis(in.GetDelayMillis())
	}
	if btn := strings.TrimSpace(in.GetButton()); btn != "" {
		out.Button = sikuli.MouseButton(strings.ToLower(btn))
	}
	return out
}

func observeOptionsFromProto(in *pb.ObserveOptions) sikuli.ObserveOptions {
	out := sikuli.ObserveOptions{}
	if in == nil {
		return out
	}
	if in.IntervalMillis != nil {
		out.Interval = durationMillis(in.GetIntervalMillis())
	}
	if in.TimeoutMillis != nil {
		out.Timeout = durationMillis(in.GetTimeoutMillis())
	}
	return out
}

func appOptionsFromProto(in *pb.AppOptions) sikuli.AppOptions {
	out := sikuli.AppOptions{}
	if in == nil {
		return out
	}
	if in.TimeoutMillis != nil {
		out.Timeout = durationMillis(in.GetTimeoutMillis())
	}
	return out
}

func regionFromProto(in *pb.Rect) sikuli.Region {
	if in == nil {
		return sikuli.NewRegion(0, 0, 0, 0)
	}
	return sikuli.NewRegion(int(in.GetX()), int(in.GetY()), int(in.GetW()), int(in.GetH()))
}

func toProtoMatches(in []sikuli.Match) []*pb.Match {
	out := make([]*pb.Match, 0, len(in))
	for _, m := range in {
		out = append(out, toProtoMatch(m))
	}
	return out
}

func toProtoMatch(in sikuli.Match) *pb.Match {
	return &pb.Match{
		Rect: &pb.Rect{
			X: int32(in.X),
			Y: int32(in.Y),
			W: int32(in.W),
			H: int32(in.H),
		},
		Score: in.Score,
		Target: &pb.Point{
			X: int32(in.Target.X),
			Y: int32(in.Target.Y),
		},
		Index: int32(in.Index),
	}
}

func toProtoGrayImage(in *sikuli.Image) *pb.GrayImage {
	if in == nil || in.Gray() == nil {
		return nil
	}
	gray := in.Gray()
	pix := make([]byte, len(gray.Pix))
	copy(pix, gray.Pix)
	return &pb.GrayImage{
		Name:   in.Name(),
		Width:  int32(in.Width()),
		Height: int32(in.Height()),
		Pix:    pix,
	}
}

func toProtoScreenDescriptor(in sikuli.Screen) *pb.ScreenDescriptor {
	return &pb.ScreenDescriptor{
		Id:      int32(in.ID),
		Name:    in.Name,
		Bounds:  &pb.Rect{X: int32(in.Bounds.X), Y: int32(in.Bounds.Y), W: int32(in.Bounds.W), H: int32(in.Bounds.H)},
		Primary: in.Primary,
	}
}

func toProtoTextMatches(in []sikuli.TextMatch) []*pb.TextMatch {
	out := make([]*pb.TextMatch, 0, len(in))
	for _, m := range in {
		out = append(out, &pb.TextMatch{
			Rect: &pb.Rect{
				X: int32(m.X),
				Y: int32(m.Y),
				W: int32(m.W),
				H: int32(m.H),
			},
			Text:       m.Text,
			Confidence: m.Confidence,
			Index:      int32(m.Index),
		})
	}
	return out
}

func toProtoObserveEvents(in []sikuli.ObserveEvent) []*pb.ObserveEvent {
	out := make([]*pb.ObserveEvent, 0, len(in))
	for _, e := range in {
		out = append(out, &pb.ObserveEvent{
			Type:                string(e.Type),
			Match:               toProtoMatch(e.Match),
			TimestampUnixMillis: e.Timestamp.UnixMilli(),
		})
	}
	return out
}

func durationMillis(ms int64) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

func mapStatusError(err error) error {
	if err == nil {
		return nil
	}
	if st, ok := status.FromError(err); ok {
		return st.Err()
	}
	switch {
	case errors.Is(err, sikuli.ErrInvalidTarget):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, sikuli.ErrFindFailed):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, sikuli.ErrTimeout):
		return status.Error(codes.DeadlineExceeded, err.Error())
	case errors.Is(err, sikuli.ErrBackendUnsupported):
		return status.Error(codes.Unimplemented, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
