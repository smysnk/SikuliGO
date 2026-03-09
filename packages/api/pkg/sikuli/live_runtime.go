package sikuli

import (
	"context"
	"fmt"
	"image"
	"net"
	"os"
	"strings"
	"time"

	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const defaultRuntimeAddress = "127.0.0.1:50051"
const defaultRuntimeTimeout = 5 * time.Second

type Runtime struct {
	address       string
	authToken     string
	rpcTimeout    time.Duration
	matcherEngine MatcherEngine
	conn          *grpc.ClientConn
	client        pb.SikuliServiceClient
	ownsConn      bool
}

type runtimeConfig struct {
	address       string
	authToken     string
	rpcTimeout    time.Duration
	dialTimeout   time.Duration
	matcherEngine MatcherEngine
	conn          *grpc.ClientConn
	dialer        func(context.Context, string) (net.Conn, error)
}

type RuntimeOption func(*runtimeConfig)

func WithRuntimeAuthToken(token string) RuntimeOption {
	return func(cfg *runtimeConfig) {
		cfg.authToken = strings.TrimSpace(token)
	}
}

func WithRuntimeRPCTimeout(timeout time.Duration) RuntimeOption {
	return func(cfg *runtimeConfig) {
		if timeout > 0 {
			cfg.rpcTimeout = timeout
		}
	}
}

func WithRuntimeDialTimeout(timeout time.Duration) RuntimeOption {
	return func(cfg *runtimeConfig) {
		if timeout > 0 {
			cfg.dialTimeout = timeout
		}
	}
}

func WithRuntimeMatcherEngine(engine MatcherEngine) RuntimeOption {
	return func(cfg *runtimeConfig) {
		cfg.matcherEngine = engine
	}
}

func WithRuntimeConn(conn *grpc.ClientConn) RuntimeOption {
	return func(cfg *runtimeConfig) {
		cfg.conn = conn
	}
}

func WithRuntimeContextDialer(dialer func(context.Context, string) (net.Conn, error)) RuntimeOption {
	return func(cfg *runtimeConfig) {
		cfg.dialer = dialer
	}
}

// NewRuntime connects to a running sikuli-go API runtime and exposes live screen operations.
func NewRuntime(address string, opts ...RuntimeOption) (*Runtime, error) {
	cfg := runtimeConfig{
		address:       strings.TrimSpace(address),
		authToken:     strings.TrimSpace(os.Getenv("SIKULI_GRPC_AUTH_TOKEN")),
		rpcTimeout:    defaultRuntimeTimeout,
		dialTimeout:   defaultRuntimeTimeout,
		matcherEngine: MatcherEngineHybrid,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&cfg)
		}
	}
	if cfg.address == "" {
		cfg.address = strings.TrimSpace(os.Getenv("SIKULI_GRPC_ADDR"))
	}
	if cfg.address == "" {
		cfg.address = defaultRuntimeAddress
	}
	if cfg.conn != nil {
		return &Runtime{
			address:       cfg.address,
			authToken:     cfg.authToken,
			rpcTimeout:    cfg.rpcTimeout,
			matcherEngine: cfg.matcherEngine,
			conn:          cfg.conn,
			client:        pb.NewSikuliServiceClient(cfg.conn),
			ownsConn:      false,
		}, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), cfg.dialTimeout)
	defer cancel()
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	if cfg.dialer != nil {
		grpcOpts = append(grpcOpts, grpc.WithContextDialer(cfg.dialer))
	}
	conn, err := grpc.DialContext(ctx, cfg.address, grpcOpts...)
	if err != nil {
		return nil, err
	}
	return &Runtime{
		address:       cfg.address,
		authToken:     cfg.authToken,
		rpcTimeout:    cfg.rpcTimeout,
		matcherEngine: cfg.matcherEngine,
		conn:          conn,
		client:        pb.NewSikuliServiceClient(conn),
		ownsConn:      true,
	}, nil
}

func (r *Runtime) Address() string {
	if r == nil {
		return ""
	}
	return r.address
}

func (r *Runtime) Close() error {
	if r == nil || r.conn == nil || !r.ownsConn {
		return nil
	}
	return r.conn.Close()
}

func (r *Runtime) Screens() ([]Screen, error) {
	if err := r.ensure(); err != nil {
		return nil, err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	res, err := r.client.ListScreens(r.withMetadata(ctx), &pb.ListScreensRequest{})
	if err != nil {
		return nil, r.mapError(err)
	}
	out := make([]Screen, 0, len(res.GetScreens()))
	for _, screen := range res.GetScreens() {
		out = append(out, r.screenFromProto(screen))
	}
	return out, nil
}

func (r *Runtime) PrimaryScreen() (Screen, error) {
	if err := r.ensure(); err != nil {
		return Screen{}, err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	res, err := r.client.GetPrimaryScreen(r.withMetadata(ctx), &pb.GetPrimaryScreenRequest{})
	if err != nil {
		return Screen{}, r.mapError(err)
	}
	if res.GetScreen() == nil {
		return Screen{}, fmt.Errorf("%w: primary screen missing", ErrBackendUnsupported)
	}
	return r.screenFromProto(res.GetScreen()), nil
}

func (r *Runtime) Screen(id int) (Screen, error) {
	screens, err := r.Screens()
	if err != nil {
		return Screen{}, err
	}
	for _, screen := range screens {
		if screen.ID == id {
			return screen, nil
		}
	}
	return Screen{}, fmt.Errorf("%w: screen id %d not found", ErrInvalidTarget, id)
}

func (r *Runtime) Capture() (*Image, error) {
	if err := r.ensure(); err != nil {
		return nil, err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	res, err := r.client.CaptureScreen(r.withMetadata(ctx), &pb.CaptureScreenRequest{})
	if err != nil {
		return nil, r.mapError(err)
	}
	return imageFromProtoGray(res.GetImage(), "capture")
}

func (r *Runtime) CaptureRegion(region Region) (*Image, error) {
	return r.Region(region).Capture()
}

func (r *Runtime) Region(region Region) LiveRegion {
	return LiveRegion{
		runtime:       r,
		region:        region,
		matcherEngine: r.defaultMatcherEngine(),
		waitScanRate:  region.WaitScanRate,
	}
}

func (r *Runtime) ensure() error {
	if r == nil || r.client == nil {
		return ErrRuntimeUnavailable
	}
	return nil
}

func (r *Runtime) defaultMatcherEngine() MatcherEngine {
	if r == nil || r.matcherEngine == "" {
		return MatcherEngineHybrid
	}
	return r.matcherEngine
}

func (r *Runtime) rpcContext() (context.Context, context.CancelFunc) {
	timeout := defaultRuntimeTimeout
	if r != nil && r.rpcTimeout > 0 {
		timeout = r.rpcTimeout
	}
	return context.WithTimeout(context.Background(), timeout)
}

func (r *Runtime) withMetadata(ctx context.Context) context.Context {
	if r == nil || strings.TrimSpace(r.authToken) == "" {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, "x-api-key", strings.TrimSpace(r.authToken))
}

func (r *Runtime) mapError(err error) error {
	if err == nil {
		return nil
	}
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.InvalidArgument:
		return fmt.Errorf("%w: %s", ErrInvalidTarget, st.Message())
	case codes.NotFound:
		return fmt.Errorf("%w: %s", ErrFindFailed, st.Message())
	case codes.DeadlineExceeded:
		return fmt.Errorf("%w: %s", ErrTimeout, st.Message())
	case codes.Unimplemented:
		return fmt.Errorf("%w: %s", ErrBackendUnsupported, st.Message())
	default:
		return err
	}
}

func (r *Runtime) screenFromProto(in *pb.ScreenDescriptor) Screen {
	if in == nil {
		return Screen{}
	}
	return Screen{
		ID:      int(in.GetId()),
		Name:    in.GetName(),
		Bounds:  rectFromProto(in.GetBounds()),
		Primary: in.GetPrimary(),
		runtime: r,
	}
}

type LiveRegion struct {
	runtime       *Runtime
	region        Region
	screenID      int
	hasScreenID   bool
	screenBounds  Rect
	matcherEngine MatcherEngine
	waitScanRate  float64
}

func (r LiveRegion) Bounds() Region {
	return r.region
}

func (r LiveRegion) Center() Point {
	return r.region.Center()
}

func (r LiveRegion) TargetPoint() Point {
	return r.Center()
}

func (r LiveRegion) Grow(dx, dy int) LiveRegion {
	r.region = r.region.Grow(dx, dy)
	return r
}

func (r LiveRegion) Offset(dx, dy int) LiveRegion {
	r.region = r.region.Offset(dx, dy)
	return r
}

func (r LiveRegion) MoveTo(x, y int) LiveRegion {
	r.region = r.region.MoveTo(x, y)
	return r
}

func (r LiveRegion) SetSize(w, h int) LiveRegion {
	r.region = r.region.SetSize(w, h)
	return r
}

func (r LiveRegion) WithMatcherEngine(engine MatcherEngine) LiveRegion {
	r.matcherEngine = engine
	return r
}

func (r LiveRegion) Capture() (*Image, error) {
	if err := r.ensure(); err != nil {
		return nil, err
	}
	ctx, cancel := r.runtime.rpcContext()
	defer cancel()
	res, err := r.runtime.client.CaptureScreen(r.runtime.withMetadata(ctx), &pb.CaptureScreenRequest{
		ScreenId: r.screenIDPtr(),
		Region:   r.captureRectProto(),
	})
	if err != nil {
		return nil, r.runtime.mapError(err)
	}
	return imageFromProtoGray(res.GetImage(), "capture")
}

func (r LiveRegion) captureOCRFinder() (*Finder, error) {
	if err := r.ensure(); err != nil {
		return nil, err
	}
	capture, err := r.Capture()
	if err != nil {
		return nil, err
	}
	return NewFinder(capture)
}

func (r LiveRegion) bindTextMatch(match TextMatch) TextMatch {
	match.Rect = NewRect(match.X+r.region.X, match.Y+r.region.Y, match.W, match.H)
	return match
}

func (r LiveRegion) bindOCRWord(word OCRWord) OCRWord {
	word.Rect = NewRect(word.X+r.region.X, word.Y+r.region.Y, word.W, word.H)
	return word
}

func (r LiveRegion) bindOCRLine(line OCRLine) OCRLine {
	line.Rect = NewRect(line.X+r.region.X, line.Y+r.region.Y, line.W, line.H)
	for i := range line.Words {
		line.Words[i] = r.bindOCRWord(line.Words[i])
	}
	return line
}

func (r LiveRegion) Find(pattern *Pattern) (Match, error) {
	if err := r.ensure(); err != nil {
		return Match{}, err
	}
	ctx, cancel := r.runtime.rpcContext()
	defer cancel()
	res, err := r.runtime.client.FindOnScreen(r.runtime.withMetadata(ctx), &pb.FindOnScreenRequest{
		Pattern: patternToProto(pattern),
		Opts:    r.screenQueryOptionsProto(nil),
	})
	if err != nil {
		return Match{}, r.runtime.mapError(err)
	}
	return r.bindMatch(matchFromProto(res.GetMatch())), nil
}

func (r LiveRegion) Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error) {
	if err := r.ensure(); err != nil {
		return Match{}, false, err
	}
	ctx, cancel := r.runtime.rpcContext()
	defer cancel()
	res, err := r.runtime.client.ExistsOnScreen(r.runtime.withMetadata(ctx), &pb.ExistsOnScreenRequest{
		Pattern: patternToProto(pattern),
		Opts:    r.screenQueryOptionsProto(&timeout),
	})
	if err != nil {
		return Match{}, false, r.runtime.mapError(err)
	}
	if !res.GetExists() || res.GetMatch() == nil {
		return Match{}, false, nil
	}
	return r.bindMatch(matchFromProto(res.GetMatch())), true, nil
}

func (r LiveRegion) Has(pattern *Pattern, timeout time.Duration) (bool, error) {
	_, ok, err := r.Exists(pattern, timeout)
	return ok, err
}

func (r LiveRegion) Wait(pattern *Pattern, timeout time.Duration) (Match, error) {
	if err := r.ensure(); err != nil {
		return Match{}, err
	}
	ctx, cancel := r.runtime.rpcContext()
	defer cancel()
	res, err := r.runtime.client.WaitOnScreen(r.runtime.withMetadata(ctx), &pb.WaitOnScreenRequest{
		Pattern: patternToProto(pattern),
		Opts:    r.screenQueryOptionsProto(&timeout),
	})
	if err != nil {
		return Match{}, r.runtime.mapError(err)
	}
	return r.bindMatch(matchFromProto(res.GetMatch())), nil
}

func (r LiveRegion) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error) {
	interval := r.waitInterval()
	return SearchWaitVanish(func() (Match, error) {
		return r.Find(pattern)
	}, timeout, interval)
}

func (r LiveRegion) ReadText(params OCRParams) (string, error) {
	finder, err := r.captureOCRFinder()
	if err != nil {
		return "", err
	}
	return finder.ReadText(params)
}

func (r LiveRegion) FindText(query string, params OCRParams) ([]TextMatch, error) {
	finder, err := r.captureOCRFinder()
	if err != nil {
		return nil, err
	}
	matches, err := finder.FindText(query, params)
	if err != nil {
		return nil, err
	}
	for i := range matches {
		matches[i] = r.bindTextMatch(matches[i])
	}
	return matches, nil
}

func (r LiveRegion) CollectWords(params OCRParams) ([]OCRWord, error) {
	finder, err := r.captureOCRFinder()
	if err != nil {
		return nil, err
	}
	words, err := finder.CollectWords(params)
	if err != nil {
		return nil, err
	}
	for i := range words {
		words[i] = r.bindOCRWord(words[i])
	}
	return words, nil
}

func (r LiveRegion) CollectLines(params OCRParams) ([]OCRLine, error) {
	finder, err := r.captureOCRFinder()
	if err != nil {
		return nil, err
	}
	lines, err := finder.CollectLines(params)
	if err != nil {
		return nil, err
	}
	for i := range lines {
		lines[i] = r.bindOCRLine(lines[i])
	}
	return lines, nil
}

func (r LiveRegion) Hover(opts InputOptions) error {
	return r.runtime.moveMouse(r.TargetPoint(), opts)
}

func (r LiveRegion) Click(opts InputOptions) error {
	return r.runtime.click(r.TargetPoint(), opts)
}

func (r LiveRegion) RightClick(opts InputOptions) error {
	return r.runtime.rightClick(r.TargetPoint(), opts)
}

func (r LiveRegion) DoubleClick(opts InputOptions) error {
	return r.runtime.doubleClick(r.TargetPoint(), opts)
}

func (r LiveRegion) MouseDown(opts InputOptions) error {
	return r.runtime.mouseDown(r.TargetPoint(), opts)
}

func (r LiveRegion) MouseUp(opts InputOptions) error {
	return r.runtime.mouseUp(r.TargetPoint(), opts)
}

func (r LiveRegion) TypeText(text string, opts InputOptions) error {
	if err := r.Click(opts); err != nil {
		return err
	}
	return r.runtime.typeText(text, opts)
}

func (r LiveRegion) Paste(text string, opts InputOptions) error {
	if err := r.Click(opts); err != nil {
		return err
	}
	return r.runtime.pasteText(text, opts)
}

func (r LiveRegion) DragDrop(target TargetPointProvider, opts InputOptions) error {
	return r.runtime.dragDrop(r.TargetPoint(), target, opts)
}

func (r LiveRegion) Wheel(direction WheelDirection, steps int, opts InputOptions) error {
	return r.runtime.wheel(r.TargetPoint(), direction, steps, opts)
}

func (r LiveRegion) KeyDown(keys ...string) error {
	return r.runtime.keyDown(keys...)
}

func (r LiveRegion) KeyUp(keys ...string) error {
	return r.runtime.keyUp(keys...)
}

func (r LiveRegion) ensure() error {
	if r.runtime == nil {
		return ErrRuntimeUnavailable
	}
	return r.runtime.ensure()
}

func (r LiveRegion) waitInterval() time.Duration {
	rate := r.waitScanRate
	if rate <= 0 {
		rate = DefaultWaitScanRate
	}
	interval := time.Duration(float64(time.Second) / rate)
	if interval < time.Millisecond {
		return time.Millisecond
	}
	return interval
}

func (r LiveRegion) captureRectProto() *pb.Rect {
	if r.hasScreenID {
		return rectToProto(NewRect(r.region.X-r.screenBounds.X, r.region.Y-r.screenBounds.Y, r.region.W, r.region.H))
	}
	if r.region.Empty() {
		return nil
	}
	return rectToProto(r.region.Rect)
}

func (r LiveRegion) screenQueryOptionsProto(timeout *time.Duration) *pb.ScreenQueryOptions {
	opts := &pb.ScreenQueryOptions{
		Region:        r.captureRectProto(),
		MatcherEngine: matcherEngineToProto(r.matcherEngine),
		ScreenId:      r.screenIDPtr(),
	}
	if timeout != nil {
		ms := timeout.Milliseconds()
		opts.TimeoutMillis = &ms
	}
	interval := r.waitInterval().Milliseconds()
	opts.IntervalMillis = &interval
	return opts
}

func (r LiveRegion) screenIDPtr() *int32 {
	if !r.hasScreenID {
		return nil
	}
	id := int32(r.screenID)
	return &id
}

func rectToProto(in Rect) *pb.Rect {
	return &pb.Rect{X: int32(in.X), Y: int32(in.Y), W: int32(in.W), H: int32(in.H)}
}

func rectFromProto(in *pb.Rect) Rect {
	if in == nil {
		return Rect{}
	}
	return NewRect(int(in.GetX()), int(in.GetY()), int(in.GetW()), int(in.GetH()))
}

func patternToProto(in *Pattern) *pb.Pattern {
	if in == nil || in.Image() == nil || in.Image().Gray() == nil {
		return nil
	}
	pbPattern := &pb.Pattern{
		Image:        grayImageFromImage(in.Image()),
		TargetOffset: &pb.Point{X: int32(in.Offset().X), Y: int32(in.Offset().Y)},
	}
	similarity := in.Similarity()
	pbPattern.Similarity = &similarity
	if in.ResizeFactor() != 1.0 {
		factor := in.ResizeFactor()
		pbPattern.ResizeFactor = &factor
	}
	if in.Similarity() >= ExactSimilarity {
		exact := true
		pbPattern.Exact = &exact
	}
	if mask := in.Mask(); mask != nil {
		pbPattern.Mask = grayImageFromGray("mask", mask)
	}
	return pbPattern
}

func grayImageFromImage(in *Image) *pb.GrayImage {
	if in == nil || in.Gray() == nil {
		return nil
	}
	return grayImageFromGray(in.Name(), in.Gray())
}

func grayImageFromGray(name string, gray *image.Gray) *pb.GrayImage {
	if gray == nil {
		return nil
	}
	pix := make([]byte, len(gray.Pix))
	copy(pix, gray.Pix)
	return &pb.GrayImage{
		Name:   name,
		Width:  int32(gray.Bounds().Dx()),
		Height: int32(gray.Bounds().Dy()),
		Pix:    pix,
	}
}

func imageFromProtoGray(in *pb.GrayImage, field string) (*Image, error) {
	if in == nil {
		return nil, fmt.Errorf("%w: %s is nil", ErrInvalidTarget, field)
	}
	if in.GetWidth() <= 0 || in.GetHeight() <= 0 {
		return nil, fmt.Errorf("%w: %s dimensions must be positive", ErrInvalidTarget, field)
	}
	gray := image.NewGray(image.Rect(0, 0, int(in.GetWidth()), int(in.GetHeight())))
	if got, want := len(in.GetPix()), len(gray.Pix); got != want {
		return nil, fmt.Errorf("%w: %s pix length mismatch got=%d want=%d", ErrInvalidTarget, field, got, want)
	}
	copy(gray.Pix, in.GetPix())
	name := strings.TrimSpace(in.GetName())
	if name == "" {
		name = field
	}
	return NewImageFromGray(name, gray)
}

func matchFromProto(in *pb.Match) Match {
	if in == nil {
		return Match{}
	}
	out := Match{
		Rect:  rectFromProto(in.GetRect()),
		Score: in.GetScore(),
		Index: int(in.GetIndex()),
	}
	if target := in.GetTarget(); target != nil {
		out.Target = NewPoint(int(target.GetX()), int(target.GetY()))
	} else {
		out.Target = NewPoint(out.X+out.W/2, out.Y+out.H/2)
	}
	return out
}

func (r LiveRegion) bindMatch(match Match) Match {
	if match == (Match{}) {
		return Match{}
	}
	match.runtime = r.runtime
	match.screenID = r.screenID
	match.hasScreenID = r.hasScreenID
	match.screenBounds = r.screenBounds
	match.matcherEngine = r.matcherEngine
	match.waitScanRate = r.waitScanRate
	return match
}

func matcherEngineToProto(in MatcherEngine) pb.MatcherEngine {
	switch in {
	case MatcherEngineTemplate:
		return pb.MatcherEngine_MATCHER_ENGINE_TEMPLATE
	case MatcherEngineORB:
		return pb.MatcherEngine_MATCHER_ENGINE_ORB
	case MatcherEngineAKAZE:
		return pb.MatcherEngine_MATCHER_ENGINE_AKAZE
	case MatcherEngineBRISK:
		return pb.MatcherEngine_MATCHER_ENGINE_BRISK
	case MatcherEngineKAZE:
		return pb.MatcherEngine_MATCHER_ENGINE_KAZE
	case MatcherEngineSIFT:
		return pb.MatcherEngine_MATCHER_ENGINE_SIFT
	case MatcherEngineHybrid, MatcherEngineDefault:
		return pb.MatcherEngine_MATCHER_ENGINE_HYBRID
	default:
		return pb.MatcherEngine_MATCHER_ENGINE_HYBRID
	}
}

func (s Screen) Live() bool {
	return s.runtime != nil
}

func (s Screen) TargetPoint() Point {
	return s.FullRegion().TargetPoint()
}

func (s Screen) FullRegion() LiveRegion {
	return LiveRegion{
		runtime:       s.runtime,
		region:        NewRegion(s.Bounds.X, s.Bounds.Y, s.Bounds.W, s.Bounds.H),
		screenID:      s.ID,
		hasScreenID:   true,
		screenBounds:  s.Bounds,
		matcherEngine: s.runtime.defaultMatcherEngine(),
		waitScanRate:  DefaultWaitScanRate,
	}
}

func (s Screen) Region(x, y, w, h int) LiveRegion {
	return s.RegionRect(NewRect(x, y, w, h))
}

func (s Screen) RegionRect(rect Rect) LiveRegion {
	live := s.FullRegion()
	live.region = NewRegion(s.Bounds.X+rect.X, s.Bounds.Y+rect.Y, rect.W, rect.H)
	live.matcherEngine = live.runtime.defaultMatcherEngine()
	return live
}

func (s Screen) Capture() (*Image, error) {
	return s.FullRegion().Capture()
}

func (s Screen) Find(pattern *Pattern) (Match, error) {
	return s.FullRegion().Find(pattern)
}

func (s Screen) Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error) {
	return s.FullRegion().Exists(pattern, timeout)
}

func (s Screen) Has(pattern *Pattern, timeout time.Duration) (bool, error) {
	return s.FullRegion().Has(pattern, timeout)
}

func (s Screen) Wait(pattern *Pattern, timeout time.Duration) (Match, error) {
	return s.FullRegion().Wait(pattern, timeout)
}

func (s Screen) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error) {
	return s.FullRegion().WaitVanish(pattern, timeout)
}

func (s Screen) ReadText(params OCRParams) (string, error) {
	return s.FullRegion().ReadText(params)
}

func (s Screen) FindText(query string, params OCRParams) ([]TextMatch, error) {
	return s.FullRegion().FindText(query, params)
}

func (s Screen) CollectWords(params OCRParams) ([]OCRWord, error) {
	return s.FullRegion().CollectWords(params)
}

func (s Screen) CollectLines(params OCRParams) ([]OCRLine, error) {
	return s.FullRegion().CollectLines(params)
}

func (s Screen) Hover(opts InputOptions) error {
	return s.FullRegion().Hover(opts)
}

func (s Screen) Click(opts InputOptions) error {
	return s.FullRegion().Click(opts)
}

func (s Screen) RightClick(opts InputOptions) error {
	return s.FullRegion().RightClick(opts)
}

func (s Screen) DoubleClick(opts InputOptions) error {
	return s.FullRegion().DoubleClick(opts)
}

func (s Screen) MouseDown(opts InputOptions) error {
	return s.FullRegion().MouseDown(opts)
}

func (s Screen) MouseUp(opts InputOptions) error {
	return s.FullRegion().MouseUp(opts)
}

func (s Screen) TypeText(text string, opts InputOptions) error {
	return s.FullRegion().TypeText(text, opts)
}

func (s Screen) Paste(text string, opts InputOptions) error {
	return s.FullRegion().Paste(text, opts)
}

func (s Screen) DragDrop(target TargetPointProvider, opts InputOptions) error {
	return s.FullRegion().DragDrop(target, opts)
}

func (s Screen) Wheel(direction WheelDirection, steps int, opts InputOptions) error {
	return s.FullRegion().Wheel(direction, steps, opts)
}

func (s Screen) KeyDown(keys ...string) error {
	return s.FullRegion().KeyDown(keys...)
}

func (s Screen) KeyUp(keys ...string) error {
	return s.FullRegion().KeyUp(keys...)
}

func (r *Runtime) moveMouse(target Point, opts InputOptions) error {
	if err := r.ensure(); err != nil {
		return err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	_, err := r.client.MoveMouse(r.withMetadata(ctx), &pb.MoveMouseRequest{
		X:    int32(target.X),
		Y:    int32(target.Y),
		Opts: inputOptionsToProto(opts),
	})
	return r.mapError(err)
}

func (r *Runtime) click(target Point, opts InputOptions) error {
	if err := r.ensure(); err != nil {
		return err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	_, err := r.client.Click(r.withMetadata(ctx), &pb.ClickRequest{
		X:    int32(target.X),
		Y:    int32(target.Y),
		Opts: inputOptionsToProto(opts),
	})
	return r.mapError(err)
}

func (r *Runtime) rightClick(target Point, opts InputOptions) error {
	opts.Button = MouseButtonRight
	return r.click(target, opts)
}

func (r *Runtime) doubleClick(target Point, opts InputOptions) error {
	if err := r.click(target, opts); err != nil {
		return err
	}
	return r.click(target, opts)
}

func (r *Runtime) mouseDown(target Point, opts InputOptions) error {
	if err := r.ensure(); err != nil {
		return err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	_, err := r.client.MouseDown(r.withMetadata(ctx), &pb.ClickRequest{
		X:    int32(target.X),
		Y:    int32(target.Y),
		Opts: inputOptionsToProto(opts),
	})
	return r.mapError(err)
}

func (r *Runtime) mouseUp(target Point, opts InputOptions) error {
	if err := r.ensure(); err != nil {
		return err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	_, err := r.client.MouseUp(r.withMetadata(ctx), &pb.ClickRequest{
		X:    int32(target.X),
		Y:    int32(target.Y),
		Opts: inputOptionsToProto(opts),
	})
	return r.mapError(err)
}

func (r *Runtime) typeText(text string, opts InputOptions) error {
	if err := r.ensure(); err != nil {
		return err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	_, err := r.client.TypeText(r.withMetadata(ctx), &pb.TypeTextRequest{
		Text: text,
		Opts: inputOptionsToProto(opts),
	})
	return r.mapError(err)
}

func (r *Runtime) pasteText(text string, opts InputOptions) error {
	if err := r.ensure(); err != nil {
		return err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	_, err := r.client.PasteText(r.withMetadata(ctx), &pb.TypeTextRequest{
		Text: text,
		Opts: inputOptionsToProto(opts),
	})
	return r.mapError(err)
}

func (r *Runtime) keyDown(keys ...string) error {
	if err := r.ensure(); err != nil {
		return err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	_, err := r.client.KeyDown(r.withMetadata(ctx), &pb.HotkeyRequest{Keys: keys})
	return r.mapError(err)
}

func (r *Runtime) keyUp(keys ...string) error {
	if err := r.ensure(); err != nil {
		return err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	_, err := r.client.KeyUp(r.withMetadata(ctx), &pb.HotkeyRequest{Keys: keys})
	return r.mapError(err)
}

func (r *Runtime) dragDrop(from Point, target TargetPointProvider, opts InputOptions) error {
	if err := r.ensure(); err != nil {
		return err
	}
	to, err := resolveTargetPoint(target)
	if err != nil {
		return err
	}
	if err := r.moveMouse(from, opts); err != nil {
		return err
	}
	if err := r.mouseDown(from, opts); err != nil {
		return err
	}
	if err := r.moveMouse(to, opts); err != nil {
		return err
	}
	return r.mouseUp(to, opts)
}

func (r *Runtime) wheel(target Point, direction WheelDirection, steps int, opts InputOptions) error {
	if err := r.ensure(); err != nil {
		return err
	}
	ctx, cancel := r.rpcContext()
	defer cancel()
	_, err := r.client.ScrollWheel(r.withMetadata(ctx), &pb.ScrollWheelRequest{
		X:         int32(target.X),
		Y:         int32(target.Y),
		Direction: string(direction),
		Steps:     int32(steps),
		Opts:      inputOptionsToProto(opts),
	})
	return r.mapError(err)
}

func inputOptionsToProto(in InputOptions) *pb.InputOptions {
	normalized := normalizeInputOptions(in)
	delay := normalized.Delay.Milliseconds()
	return &pb.InputOptions{
		DelayMillis: &delay,
		Button:      string(normalized.Button),
	}
}

func resolveTargetPoint(target TargetPointProvider) (Point, error) {
	if target == nil {
		return Point{}, fmt.Errorf("%w: target is nil", ErrInvalidTarget)
	}
	return target.TargetPoint(), nil
}
