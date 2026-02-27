//go:build !opencv

package grpcv1

import (
	"context"
	"testing"

	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFindORBWithoutOpenCVMapsToUnimplemented(t *testing.T) {
	srv := NewServer()
	req := &pb.FindRequest{
		Source: grayImage("source", [][]uint8{
			{10, 10, 10, 10},
			{10, 0, 255, 10},
			{10, 255, 0, 10},
			{10, 10, 10, 10},
		}),
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
		MatcherEngine: pb.MatcherEngine_MATCHER_ENGINE_ORB,
	}
	_, err := srv.Find(context.Background(), req)
	if err == nil {
		t.Fatalf("expected unimplemented error")
	}
	if code := status.Code(err); code != codes.Unimplemented {
		t.Fatalf("expected unimplemented code, got %s", code)
	}
}
