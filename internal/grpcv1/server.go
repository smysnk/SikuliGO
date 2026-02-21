package grpcv1

import (
	"context"
	"errors"
	"fmt"
	"image"
	"strings"
	"time"

	pb "github.com/sikulix/portgo/internal/grpcv1/pb"
	"github.com/sikulix/portgo/pkg/sikuli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedSikuliServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Find(_ context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	source, pattern, err := findRequestParts(req)
	if err != nil {
		return nil, mapStatusError(err)
	}
	f, err := sikuli.NewFinder(source)
	if err != nil {
		return nil, mapStatusError(err)
	}
	match, err := f.Find(pattern)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.FindResponse{Match: toProtoMatch(match)}, nil
}

func (s *Server) FindAll(_ context.Context, req *pb.FindRequest) (*pb.FindAllResponse, error) {
	source, pattern, err := findRequestParts(req)
	if err != nil {
		return nil, mapStatusError(err)
	}
	f, err := sikuli.NewFinder(source)
	if err != nil {
		return nil, mapStatusError(err)
	}
	matches, err := f.FindAll(pattern)
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.FindAllResponse{Matches: toProtoMatches(matches)}, nil
}

func (s *Server) ReadText(_ context.Context, req *pb.ReadTextRequest) (*pb.ReadTextResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	source, err := imageFromProto(req.GetSource(), "source")
	if err != nil {
		return nil, mapStatusError(err)
	}
	f, err := sikuli.NewFinder(source)
	if err != nil {
		return nil, mapStatusError(err)
	}
	text, err := f.ReadText(ocrParamsFromProto(req.GetParams()))
	if err != nil {
		return nil, mapStatusError(err)
	}
	return &pb.ReadTextResponse{Text: text}, nil
}

func (s *Server) FindText(_ context.Context, req *pb.FindTextRequest) (*pb.FindTextResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	source, err := imageFromProto(req.GetSource(), "source")
	if err != nil {
		return nil, mapStatusError(err)
	}
	f, err := sikuli.NewFinder(source)
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
	out := make([]*pb.Window, 0, len(windows))
	for _, w := range windows {
		out = append(out, &pb.Window{
			Title:   w.Title,
			Bounds:  &pb.Rect{X: int32(w.Bounds.X), Y: int32(w.Bounds.Y), W: int32(w.Bounds.W), H: int32(w.Bounds.H)},
			Focused: w.Focused,
		})
	}
	return &pb.ListWindowsResponse{Windows: out}, nil
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
