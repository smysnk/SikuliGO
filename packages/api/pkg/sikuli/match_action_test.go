package sikuli

import (
	"context"
	"testing"
	"time"

	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"google.golang.org/grpc"
)

type recordingRuntimeClient struct {
	pb.SikuliServiceClient
	moveRequests       []*pb.MoveMouseRequest
	clickRequests      []*pb.ClickRequest
	typeRequests       []*pb.TypeTextRequest
	pasteRequests      []*pb.TypeTextRequest
	mouseDownRequests  []*pb.ClickRequest
	mouseUpRequests    []*pb.ClickRequest
	keyDownRequests    []*pb.HotkeyRequest
	keyUpRequests      []*pb.HotkeyRequest
	scrollWheelRequest []*pb.ScrollWheelRequest
}

func (c *recordingRuntimeClient) MoveMouse(_ context.Context, in *pb.MoveMouseRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	c.moveRequests = append(c.moveRequests, cloneMoveMouseRequest(in))
	return &pb.ActionResponse{}, nil
}

func (c *recordingRuntimeClient) Click(_ context.Context, in *pb.ClickRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	c.clickRequests = append(c.clickRequests, cloneClickRequest(in))
	return &pb.ActionResponse{}, nil
}

func (c *recordingRuntimeClient) TypeText(_ context.Context, in *pb.TypeTextRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	c.typeRequests = append(c.typeRequests, cloneTypeTextRequest(in))
	return &pb.ActionResponse{}, nil
}

func (c *recordingRuntimeClient) PasteText(_ context.Context, in *pb.TypeTextRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	c.pasteRequests = append(c.pasteRequests, cloneTypeTextRequest(in))
	return &pb.ActionResponse{}, nil
}

func (c *recordingRuntimeClient) MouseDown(_ context.Context, in *pb.ClickRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	c.mouseDownRequests = append(c.mouseDownRequests, cloneClickRequest(in))
	return &pb.ActionResponse{}, nil
}

func (c *recordingRuntimeClient) MouseUp(_ context.Context, in *pb.ClickRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	c.mouseUpRequests = append(c.mouseUpRequests, cloneClickRequest(in))
	return &pb.ActionResponse{}, nil
}

func (c *recordingRuntimeClient) KeyDown(_ context.Context, in *pb.HotkeyRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	c.keyDownRequests = append(c.keyDownRequests, cloneHotkeyRequest(in))
	return &pb.ActionResponse{}, nil
}

func (c *recordingRuntimeClient) KeyUp(_ context.Context, in *pb.HotkeyRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	c.keyUpRequests = append(c.keyUpRequests, cloneHotkeyRequest(in))
	return &pb.ActionResponse{}, nil
}

func (c *recordingRuntimeClient) ScrollWheel(_ context.Context, in *pb.ScrollWheelRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	c.scrollWheelRequest = append(c.scrollWheelRequest, cloneScrollWheelRequest(in))
	return &pb.ActionResponse{}, nil
}

func TestLiveMatchActionHelpersUseTargetPoint(t *testing.T) {
	client := &recordingRuntimeClient{}
	runtime := &Runtime{
		client:     client,
		rpcTimeout: time.Second,
	}
	match := Match{
		Rect:      NewRect(10, 20, 30, 40),
		Score:     0.99,
		Target:    NewPoint(42, 64),
		runtime:   runtime,
		screenID:  1,
		hasScreenID: true,
	}

	if err := match.Hover(InputOptions{}); err != nil {
		t.Fatalf("hover failed: %v", err)
	}
	if err := match.Click(InputOptions{}); err != nil {
		t.Fatalf("click failed: %v", err)
	}
	if err := match.RightClick(InputOptions{}); err != nil {
		t.Fatalf("right click failed: %v", err)
	}
	if err := match.DoubleClick(InputOptions{}); err != nil {
		t.Fatalf("double click failed: %v", err)
	}
	if err := match.MouseDown(InputOptions{}); err != nil {
		t.Fatalf("mouse down failed: %v", err)
	}
	if err := match.MouseUp(InputOptions{}); err != nil {
		t.Fatalf("mouse up failed: %v", err)
	}
	if err := match.TypeText("  hello  ", InputOptions{}); err != nil {
		t.Fatalf("type text failed: %v", err)
	}
	if err := match.Paste("  world  ", InputOptions{}); err != nil {
		t.Fatalf("paste failed: %v", err)
	}
	if err := match.DragDrop(NewPoint(80, 96), InputOptions{}); err != nil {
		t.Fatalf("drag drop failed: %v", err)
	}
	if err := match.Wheel(WheelDirectionDown, 3, InputOptions{}); err != nil {
		t.Fatalf("wheel failed: %v", err)
	}
	if err := match.KeyDown("cmd", "shift"); err != nil {
		t.Fatalf("key down failed: %v", err)
	}
	if err := match.KeyUp("cmd", "shift"); err != nil {
		t.Fatalf("key up failed: %v", err)
	}

	if len(client.moveRequests) != 3 {
		t.Fatalf("move request count mismatch got=%d want=3", len(client.moveRequests))
	}
	assertMoveRequest(t, client.moveRequests[0], 42, 64)
	assertMoveRequest(t, client.moveRequests[1], 42, 64)
	assertMoveRequest(t, client.moveRequests[2], 80, 96)

	if len(client.clickRequests) != 6 {
		t.Fatalf("click request count mismatch got=%d want=6", len(client.clickRequests))
	}
	assertClickRequest(t, client.clickRequests[0], 42, 64, string(MouseButtonLeft))
	assertClickRequest(t, client.clickRequests[1], 42, 64, string(MouseButtonRight))
	assertClickRequest(t, client.clickRequests[2], 42, 64, string(MouseButtonLeft))
	assertClickRequest(t, client.clickRequests[3], 42, 64, string(MouseButtonLeft))
	assertClickRequest(t, client.clickRequests[4], 42, 64, string(MouseButtonLeft))
	assertClickRequest(t, client.clickRequests[5], 42, 64, string(MouseButtonLeft))

	if len(client.mouseDownRequests) != 2 {
		t.Fatalf("mouse down count mismatch got=%d want=2", len(client.mouseDownRequests))
	}
	assertClickRequest(t, client.mouseDownRequests[0], 42, 64, string(MouseButtonLeft))
	assertClickRequest(t, client.mouseDownRequests[1], 42, 64, string(MouseButtonLeft))

	if len(client.mouseUpRequests) != 2 {
		t.Fatalf("mouse up count mismatch got=%d want=2", len(client.mouseUpRequests))
	}
	assertClickRequest(t, client.mouseUpRequests[0], 42, 64, string(MouseButtonLeft))
	assertClickRequest(t, client.mouseUpRequests[1], 80, 96, string(MouseButtonLeft))

	if len(client.typeRequests) != 1 || client.typeRequests[0].GetText() != "  hello  " {
		t.Fatalf("type text request mismatch: %+v", client.typeRequests)
	}
	if len(client.pasteRequests) != 1 || client.pasteRequests[0].GetText() != "  world  " {
		t.Fatalf("paste request mismatch: %+v", client.pasteRequests)
	}
	if len(client.scrollWheelRequest) != 1 {
		t.Fatalf("scroll wheel count mismatch got=%d want=1", len(client.scrollWheelRequest))
	}
	if req := client.scrollWheelRequest[0]; req.GetX() != 42 || req.GetY() != 64 || req.GetDirection() != string(WheelDirectionDown) || req.GetSteps() != 3 {
		t.Fatalf("scroll wheel request mismatch: %+v", req)
	}
	if len(client.keyDownRequests) != 1 || len(client.keyDownRequests[0].GetKeys()) != 2 {
		t.Fatalf("key down request mismatch: %+v", client.keyDownRequests)
	}
	if len(client.keyUpRequests) != 1 || len(client.keyUpRequests[0].GetKeys()) != 2 {
		t.Fatalf("key up request mismatch: %+v", client.keyUpRequests)
	}
}

func assertMoveRequest(t *testing.T, req *pb.MoveMouseRequest, x, y int32) {
	t.Helper()
	if req.GetX() != x || req.GetY() != y {
		t.Fatalf("move request mismatch: %+v", req)
	}
}

func assertClickRequest(t *testing.T, req *pb.ClickRequest, x, y int32, button string) {
	t.Helper()
	if req.GetX() != x || req.GetY() != y || req.GetOpts().GetButton() != button {
		t.Fatalf("click request mismatch: %+v", req)
	}
}

func cloneMoveMouseRequest(in *pb.MoveMouseRequest) *pb.MoveMouseRequest {
	if in == nil {
		return nil
	}
	return &pb.MoveMouseRequest{
		X:    in.GetX(),
		Y:    in.GetY(),
		Opts: cloneInputOptions(in.GetOpts()),
	}
}

func cloneClickRequest(in *pb.ClickRequest) *pb.ClickRequest {
	if in == nil {
		return nil
	}
	return &pb.ClickRequest{
		X:    in.GetX(),
		Y:    in.GetY(),
		Opts: cloneInputOptions(in.GetOpts()),
	}
}

func cloneTypeTextRequest(in *pb.TypeTextRequest) *pb.TypeTextRequest {
	if in == nil {
		return nil
	}
	return &pb.TypeTextRequest{
		Text: in.GetText(),
		Opts: cloneInputOptions(in.GetOpts()),
	}
}

func cloneHotkeyRequest(in *pb.HotkeyRequest) *pb.HotkeyRequest {
	if in == nil {
		return nil
	}
	return &pb.HotkeyRequest{Keys: append([]string(nil), in.GetKeys()...)}
}

func cloneScrollWheelRequest(in *pb.ScrollWheelRequest) *pb.ScrollWheelRequest {
	if in == nil {
		return nil
	}
	return &pb.ScrollWheelRequest{
		X:         in.GetX(),
		Y:         in.GetY(),
		Direction: in.GetDirection(),
		Steps:     in.GetSteps(),
		Opts:      cloneInputOptions(in.GetOpts()),
	}
}

func cloneInputOptions(in *pb.InputOptions) *pb.InputOptions {
	if in == nil {
		return nil
	}
	var delay *int64
	if in.DelayMillis != nil {
		v := in.GetDelayMillis()
		delay = &v
	}
	return &pb.InputOptions{
		DelayMillis: delay,
		Button:      in.GetButton(),
	}
}
