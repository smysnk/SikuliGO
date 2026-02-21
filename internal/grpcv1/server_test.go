package grpcv1

import (
	"context"
	"testing"

	pb "github.com/sikulix/portgo/internal/grpcv1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFindAll(t *testing.T) {
	srv := NewServer()

	req := &pb.FindRequest{
		Source: grayImage("source", [][]uint8{
			{10, 10, 10, 10, 10, 10, 10, 10},
			{10, 0, 255, 10, 10, 10, 10, 10},
			{10, 255, 0, 10, 0, 255, 10, 10},
			{10, 10, 10, 10, 255, 0, 10, 10},
			{10, 10, 10, 10, 10, 10, 10, 10},
		}),
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
	}

	res, err := srv.FindAll(context.Background(), req)
	if err != nil {
		t.Fatalf("find all failed: %v", err)
	}
	if got := len(res.GetMatches()); got != 2 {
		t.Fatalf("expected 2 matches, got %d", got)
	}
	if m := res.GetMatches()[0].GetRect(); m.GetX() != 1 || m.GetY() != 1 {
		t.Fatalf("first match mismatch: %+v", res.GetMatches()[0])
	}
	if m := res.GetMatches()[1].GetRect(); m.GetX() != 4 || m.GetY() != 2 {
		t.Fatalf("second match mismatch: %+v", res.GetMatches()[1])
	}
}

func TestFindNotFoundMapsToNotFound(t *testing.T) {
	srv := NewServer()

	req := &pb.FindRequest{
		Source: grayImage("source", [][]uint8{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}),
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
	}

	_, err := srv.Find(context.Background(), req)
	if err == nil {
		t.Fatalf("expected not found error")
	}
	if code := status.Code(err); code != codes.NotFound {
		t.Fatalf("expected not found code, got %s", code)
	}
}

func TestFindInvalidImageMapsToInvalidArgument(t *testing.T) {
	srv := NewServer()

	req := &pb.FindRequest{
		Source: &pb.GrayImage{
			Name:   "bad",
			Width:  2,
			Height: 2,
			Pix:    []byte{0, 1, 2},
		},
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{1, 1},
				{1, 1},
			}),
		},
	}

	_, err := srv.Find(context.Background(), req)
	if err == nil {
		t.Fatalf("expected invalid argument error")
	}
	if code := status.Code(err); code != codes.InvalidArgument {
		t.Fatalf("expected invalid argument code, got %s", code)
	}
}

func TestFindTextEmptyQueryMapsToInvalidArgument(t *testing.T) {
	srv := NewServer()

	req := &pb.FindTextRequest{
		Source: grayImage("source", [][]uint8{
			{1, 1},
			{1, 1},
		}),
		Query: "   ",
	}

	_, err := srv.FindText(context.Background(), req)
	if err == nil {
		t.Fatalf("expected invalid argument error")
	}
	if code := status.Code(err); code != codes.InvalidArgument {
		t.Fatalf("expected invalid argument code, got %s", code)
	}
}

func grayImage(name string, rows [][]uint8) *pb.GrayImage {
	h := len(rows)
	w := len(rows[0])
	pix := make([]byte, 0, w*h)
	for _, row := range rows {
		pix = append(pix, row...)
	}
	return &pb.GrayImage{
		Name:   name,
		Width:  int32(w),
		Height: int32(h),
		Pix:    pix,
	}
}

func boolPtr(v bool) *bool {
	return &v
}
