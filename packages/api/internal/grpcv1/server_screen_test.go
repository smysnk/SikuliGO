package grpcv1

import (
	"context"
	"testing"

	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"github.com/smysnk/sikuligo/pkg/sikuli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestListScreensAndPrimaryScreen(t *testing.T) {
	screens := []sikuli.Screen{
		{ID: 1, Name: "left", Bounds: sikuli.NewRect(0, 0, 4, 4), Primary: false},
		{ID: 2, Name: "right", Bounds: sikuli.NewRect(4, 0, 4, 4), Primary: true},
	}
	srv := NewServer(WithScreenLister(func(context.Context) ([]sikuli.Screen, error) {
		return screens, nil
	}))

	listRes, err := srv.ListScreens(context.Background(), &pb.ListScreensRequest{})
	if err != nil {
		t.Fatalf("list screens failed: %v", err)
	}
	if got := len(listRes.GetScreens()); got != 2 {
		t.Fatalf("screen count mismatch got=%d want=2", got)
	}
	if got := listRes.GetScreens()[1]; got.GetId() != 2 || !got.GetPrimary() || got.GetBounds().GetX() != 4 {
		t.Fatalf("unexpected second screen descriptor: %+v", got)
	}

	primaryRes, err := srv.GetPrimaryScreen(context.Background(), &pb.GetPrimaryScreenRequest{})
	if err != nil {
		t.Fatalf("get primary screen failed: %v", err)
	}
	if primaryRes.GetScreen() == nil || primaryRes.GetScreen().GetId() != 2 {
		t.Fatalf("primary screen mismatch: %+v", primaryRes.GetScreen())
	}
}

func TestCaptureScreenUsesSelectedScreenAndLocalRegion(t *testing.T) {
	screen := sikuliImageFromRows(t, "screen", [][]uint8{
		{10, 10, 10, 10, 10, 10, 10, 10},
		{10, 0, 255, 10, 10, 0, 255, 10},
		{10, 255, 0, 10, 10, 255, 0, 10},
		{10, 10, 10, 10, 10, 10, 10, 10},
	})
	screens := []sikuli.Screen{
		{ID: 1, Name: "left", Bounds: sikuli.NewRect(0, 0, 4, 4), Primary: true},
		{ID: 2, Name: "right", Bounds: sikuli.NewRect(4, 0, 4, 4), Primary: false},
	}
	srv := NewServer(
		WithCaptureScreen(func(_ context.Context, _ string) (*sikuli.Image, error) { return screen, nil }),
		WithScreenLister(func(context.Context) ([]sikuli.Screen, error) { return screens, nil }),
	)

	screenID := int32(2)
	res, err := srv.CaptureScreen(context.Background(), &pb.CaptureScreenRequest{
		ScreenId: &screenID,
		Region:   &pb.Rect{X: 1, Y: 1, W: 2, H: 2},
	})
	if err != nil {
		t.Fatalf("capture screen failed: %v", err)
	}
	if res.GetScreen() == nil || res.GetScreen().GetId() != 2 {
		t.Fatalf("expected selected screen descriptor, got %+v", res.GetScreen())
	}
	img := res.GetImage()
	if img == nil || img.GetWidth() != 2 || img.GetHeight() != 2 {
		t.Fatalf("capture dimensions mismatch: %+v", img)
	}
	wantPix := []byte{0, 255, 255, 0}
	if got := img.GetPix(); len(got) != len(wantPix) || got[0] != wantPix[0] || got[1] != wantPix[1] || got[2] != wantPix[2] || got[3] != wantPix[3] {
		t.Fatalf("capture pix mismatch got=%v want=%v", got, wantPix)
	}
}

func TestCaptureScreenRejectsRegionOutsideSelectedScreen(t *testing.T) {
	screen := sikuliImageFromRows(t, "screen", [][]uint8{
		{10, 10, 10, 10, 10, 10, 10, 10},
		{10, 0, 255, 10, 10, 0, 255, 10},
		{10, 255, 0, 10, 10, 255, 0, 10},
		{10, 10, 10, 10, 10, 10, 10, 10},
	})
	screens := []sikuli.Screen{
		{ID: 1, Name: "left", Bounds: sikuli.NewRect(0, 0, 4, 4), Primary: true},
		{ID: 2, Name: "right", Bounds: sikuli.NewRect(4, 0, 4, 4), Primary: false},
	}
	srv := NewServer(
		WithCaptureScreen(func(_ context.Context, _ string) (*sikuli.Image, error) { return screen, nil }),
		WithScreenLister(func(context.Context) ([]sikuli.Screen, error) { return screens, nil }),
	)

	screenID := int32(2)
	_, err := srv.CaptureScreen(context.Background(), &pb.CaptureScreenRequest{
		ScreenId: &screenID,
		Region:   &pb.Rect{X: 10, Y: 10, W: 2, H: 2},
	})
	if err == nil {
		t.Fatalf("expected invalid argument error")
	}
	if code := status.Code(err); code != codes.InvalidArgument {
		t.Fatalf("expected invalid argument, got %s", code)
	}
}

func TestFindOnScreenResolvesScreenIDAndLocalRegion(t *testing.T) {
	screen := sikuliImageFromRows(t, "screen", [][]uint8{
		{10, 10, 10, 10, 10, 10, 10, 10},
		{10, 0, 255, 10, 10, 0, 255, 10},
		{10, 255, 0, 10, 10, 255, 0, 10},
		{10, 10, 10, 10, 10, 10, 10, 10},
	})
	screens := []sikuli.Screen{
		{ID: 1, Name: "left", Bounds: sikuli.NewRect(0, 0, 4, 4), Primary: true},
		{ID: 2, Name: "right", Bounds: sikuli.NewRect(4, 0, 4, 4), Primary: false},
	}
	srv := NewServer(
		WithCaptureScreen(func(_ context.Context, _ string) (*sikuli.Image, error) { return screen, nil }),
		WithScreenLister(func(context.Context) ([]sikuli.Screen, error) { return screens, nil }),
	)

	screenID := int32(2)
	res, err := srv.FindOnScreen(context.Background(), &pb.FindOnScreenRequest{
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
		Opts: &pb.ScreenQueryOptions{
			ScreenId: &screenID,
			Region:   &pb.Rect{X: 1, Y: 1, W: 2, H: 2},
		},
	})
	if err != nil {
		t.Fatalf("find_on_screen failed: %v", err)
	}
	match := res.GetMatch()
	if match == nil {
		t.Fatalf("expected match in response")
	}
	if got := match.GetRect(); got.GetX() != 5 || got.GetY() != 1 {
		t.Fatalf("screen-scoped match mismatch: %+v", got)
	}
}
