package grpcv1

import (
	"context"
	"log"
	"net"
	"strings"
	"testing"
	"time"

	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"github.com/smysnk/sikuligo/pkg/sikuli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const rpcSurfaceBufSize = 1024 * 1024

func TestRPCSurfaceIntegrationViaBufconn(t *testing.T) {
	t.Parallel()

	lis := bufconn.Listen(rpcSurfaceBufSize)
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(UnaryInterceptors("", log.Default(), NewMetricsRegistry(), nil)...),
	)
	pb.RegisterSikuliServiceServer(srv, NewServer(
		WithCaptureScreen(func(_ context.Context, _ string) (*sikuli.Image, error) {
			return sikuliImageFromRows(t, "screen", [][]uint8{
				{10, 10},
				{10, 10},
			}), nil
		}),
		WithScreenLister(func(context.Context) ([]sikuli.Screen, error) {
			return []sikuli.Screen{
				{ID: 1, Name: "primary", Bounds: sikuli.NewRect(0, 0, 2, 2), Primary: true},
			}, nil
		}),
	))

	go func() {
		_ = srv.Serve(lis)
	}()
	t.Cleanup(func() {
		srv.Stop()
		_ = lis.Close()
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return lis.DialContext(ctx)
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("dial bufconn: %v", err)
	}
	t.Cleanup(func() { _ = conn.Close() })

	client := pb.NewSikuliServiceClient(conn)
	img := &pb.GrayImage{
		Name:   "tiny",
		Width:  1,
		Height: 1,
		Pix:    []byte{0},
	}
	pattern := &pb.Pattern{Image: img}
	region := &pb.Rect{X: 0, Y: 0, W: 1, H: 1}
	screenOpts := &pb.ScreenQueryOptions{Region: region, TimeoutMillis: int64PtrRPCSurface(5), IntervalMillis: int64PtrRPCSurface(1)}
	obsOpts := &pb.ObserveOptions{TimeoutMillis: int64PtrRPCSurface(5), IntervalMillis: int64PtrRPCSurface(1)}

	rpcTimeout := 800 * time.Millisecond
	check := func(name string, call func(context.Context) error) {
		t.Helper()
		callCtx, callCancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer callCancel()
		err := call(callCtx)
		assertNoContractRegression(t, name, err)
	}

	check("Find", func(c context.Context) error {
		_, err := client.Find(c, &pb.FindRequest{Source: img, Pattern: pattern})
		return err
	})
	check("FindAll", func(c context.Context) error {
		_, err := client.FindAll(c, &pb.FindRequest{Source: img, Pattern: pattern})
		return err
	})
	check("FindOnScreen", func(c context.Context) error {
		_, err := client.FindOnScreen(c, &pb.FindOnScreenRequest{Pattern: pattern, Opts: screenOpts})
		return err
	})
	check("ListScreens", func(c context.Context) error {
		_, err := client.ListScreens(c, &pb.ListScreensRequest{})
		return err
	})
	check("GetPrimaryScreen", func(c context.Context) error {
		_, err := client.GetPrimaryScreen(c, &pb.GetPrimaryScreenRequest{})
		return err
	})
	check("CaptureScreen", func(c context.Context) error {
		_, err := client.CaptureScreen(c, &pb.CaptureScreenRequest{})
		return err
	})
	check("ExistsOnScreen", func(c context.Context) error {
		_, err := client.ExistsOnScreen(c, &pb.ExistsOnScreenRequest{Pattern: pattern, Opts: screenOpts})
		return err
	})
	check("WaitOnScreen", func(c context.Context) error {
		_, err := client.WaitOnScreen(c, &pb.WaitOnScreenRequest{Pattern: pattern, Opts: screenOpts})
		return err
	})
	check("ClickOnScreen", func(c context.Context) error {
		_, err := client.ClickOnScreen(c, &pb.ClickOnScreenRequest{Pattern: pattern, Opts: screenOpts, ClickOpts: &pb.InputOptions{}})
		return err
	})
	check("ReadText", func(c context.Context) error {
		_, err := client.ReadText(c, &pb.ReadTextRequest{Source: img})
		return err
	})
	check("FindText", func(c context.Context) error {
		_, err := client.FindText(c, &pb.FindTextRequest{Source: img, Query: "x"})
		return err
	})
	check("MoveMouse", func(c context.Context) error {
		_, err := client.MoveMouse(c, &pb.MoveMouseRequest{X: 0, Y: 0, Opts: &pb.InputOptions{}})
		return err
	})
	check("Click", func(c context.Context) error {
		_, err := client.Click(c, &pb.ClickRequest{X: 0, Y: 0, Opts: &pb.InputOptions{}})
		return err
	})
	check("TypeText", func(c context.Context) error {
		_, err := client.TypeText(c, &pb.TypeTextRequest{Text: "x", Opts: &pb.InputOptions{}})
		return err
	})
	check("PasteText", func(c context.Context) error {
		_, err := client.PasteText(c, &pb.TypeTextRequest{Text: "x", Opts: &pb.InputOptions{}})
		return err
	})
	check("Hotkey", func(c context.Context) error {
		_, err := client.Hotkey(c, &pb.HotkeyRequest{Keys: []string{"cmd", "c"}})
		return err
	})
	check("MouseDown", func(c context.Context) error {
		_, err := client.MouseDown(c, &pb.ClickRequest{X: 0, Y: 0, Opts: &pb.InputOptions{}})
		return err
	})
	check("MouseUp", func(c context.Context) error {
		_, err := client.MouseUp(c, &pb.ClickRequest{X: 0, Y: 0, Opts: &pb.InputOptions{}})
		return err
	})
	check("KeyDown", func(c context.Context) error {
		_, err := client.KeyDown(c, &pb.HotkeyRequest{Keys: []string{"cmd"}})
		return err
	})
	check("KeyUp", func(c context.Context) error {
		_, err := client.KeyUp(c, &pb.HotkeyRequest{Keys: []string{"cmd"}})
		return err
	})
	check("ScrollWheel", func(c context.Context) error {
		_, err := client.ScrollWheel(c, &pb.ScrollWheelRequest{X: 0, Y: 0, Direction: "down", Steps: 1, Opts: &pb.InputOptions{}})
		return err
	})
	check("ObserveAppear", func(c context.Context) error {
		_, err := client.ObserveAppear(c, &pb.ObserveRequest{Source: img, Region: region, Pattern: pattern, Opts: obsOpts})
		return err
	})
	check("ObserveVanish", func(c context.Context) error {
		_, err := client.ObserveVanish(c, &pb.ObserveRequest{Source: img, Region: region, Pattern: pattern, Opts: obsOpts})
		return err
	})
	check("ObserveChange", func(c context.Context) error {
		_, err := client.ObserveChange(c, &pb.ObserveChangeRequest{Source: img, Region: region, Opts: obsOpts})
		return err
	})
	check("OpenApp", func(c context.Context) error {
		_, err := client.OpenApp(c, &pb.AppActionRequest{Name: "TestApp", Args: []string{}})
		return err
	})
	check("FocusApp", func(c context.Context) error {
		_, err := client.FocusApp(c, &pb.AppActionRequest{Name: "TestApp"})
		return err
	})
	check("CloseApp", func(c context.Context) error {
		_, err := client.CloseApp(c, &pb.AppActionRequest{Name: "TestApp"})
		return err
	})
	check("IsAppRunning", func(c context.Context) error {
		_, err := client.IsAppRunning(c, &pb.AppActionRequest{Name: "TestApp"})
		return err
	})
	check("ListWindows", func(c context.Context) error {
		_, err := client.ListWindows(c, &pb.AppActionRequest{Name: "TestApp"})
		return err
	})
	check("FindWindows", func(c context.Context) error {
		_, err := client.FindWindows(c, &pb.WindowQueryRequest{Name: "TestApp", Query: &pb.WindowQuery{TitleContains: "main"}})
		return err
	})
	check("GetWindow", func(c context.Context) error {
		idx := int32(0)
		_, err := client.GetWindow(c, &pb.WindowQueryRequest{Name: "TestApp", Query: &pb.WindowQuery{TitleContains: "main", Index: &idx}})
		return err
	})
	check("GetFocusedWindow", func(c context.Context) error {
		_, err := client.GetFocusedWindow(c, &pb.AppActionRequest{Name: "TestApp"})
		return err
	})
}

func assertNoContractRegression(t *testing.T, method string, err error) {
	t.Helper()
	if err == nil {
		return
	}
	st, ok := status.FromError(err)
	if !ok {
		msg := strings.ToLower(err.Error())
		if strings.Contains(msg, "unknown method") || strings.Contains(msg, "serialization failure") {
			t.Fatalf("%s contract regression: %v", method, err)
		}
		return
	}
	msg := strings.ToLower(st.Message())
	if strings.Contains(msg, "unknown method") || strings.Contains(msg, "serialization failure") {
		t.Fatalf("%s contract regression: %v", method, err)
	}
}

func int64PtrRPCSurface(v int64) *int64 {
	return &v
}
