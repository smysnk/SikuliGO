//go:build darwin

package grpcv1

/*
#cgo darwin LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>
#include <stdlib.h>

static int sikuli_go_active_display_count(uint32_t *count) {
	return CGGetActiveDisplayList(0, NULL, count);
}

static int sikuli_go_active_display_list(uint32_t max, CGDirectDisplayID *ids, uint32_t *count) {
	return CGGetActiveDisplayList(max, ids, count);
}

static uint32_t sikuli_go_main_display_id(void) {
	return CGMainDisplayID();
}

static CGRect sikuli_go_display_bounds(CGDirectDisplayID id) {
	return CGDisplayBounds(id);
}
*/
import "C"

import (
	"context"
	"fmt"

	"github.com/smysnk/sikuligo/pkg/sikuli"
)

func listScreens(context.Context) ([]sikuli.Screen, error) {
	var count C.uint32_t
	if code := C.sikuli_go_active_display_count(&count); code != 0 {
		return nil, fmt.Errorf("screen enumeration failed: cg error=%d", int(code))
	}
	if count == 0 {
		return nil, fmt.Errorf("no active displays found")
	}
	ids := make([]C.CGDirectDisplayID, int(count))
	if code := C.sikuli_go_active_display_list(count, &ids[0], &count); code != 0 {
		return nil, fmt.Errorf("screen enumeration failed: cg error=%d", int(code))
	}
	mainID := C.sikuli_go_main_display_id()
	out := make([]sikuli.Screen, 0, int(count))
	minX := 0
	minY := 0
	initialized := false
	for i := 0; i < int(count); i++ {
		bounds := C.sikuli_go_display_bounds(ids[i])
		x := int(bounds.origin.x)
		y := int(bounds.origin.y)
		w := int(bounds.size.width)
		h := int(bounds.size.height)
		if !initialized {
			minX = x
			minY = y
			initialized = true
		} else {
			if x < minX {
				minX = x
			}
			if y < minY {
				minY = y
			}
		}
		out = append(out, sikuli.Screen{
			ID:      int(ids[i]),
			Name:    fmt.Sprintf("display-%d", int(ids[i])),
			Bounds:  sikuli.NewRect(x, y, w, h),
			Primary: ids[i] == mainID,
		})
	}
	if minX != 0 || minY != 0 {
		for i := range out {
			out[i].Bounds = sikuli.NewRect(
				out[i].Bounds.X-minX,
				out[i].Bounds.Y-minY,
				out[i].Bounds.W,
				out[i].Bounds.H,
			)
		}
	}
	return out, nil
}
