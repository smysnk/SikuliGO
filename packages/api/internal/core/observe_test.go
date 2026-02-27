package core

import (
	"image"
	"testing"
)

func TestObserveRequestValidate(t *testing.T) {
	base := ObserveRequest{
		Source:  image.NewGray(image.Rect(0, 0, 10, 10)),
		Region:  image.Rect(0, 0, 10, 10),
		Pattern: image.NewGray(image.Rect(0, 0, 2, 2)),
		Event:   ObserveEventAppear,
	}
	if err := base.Validate(); err != nil {
		t.Fatalf("base request should validate: %v", err)
	}

	cases := []ObserveRequest{
		{
			Source:  nil,
			Region:  image.Rect(0, 0, 10, 10),
			Pattern: image.NewGray(image.Rect(0, 0, 2, 2)),
			Event:   ObserveEventAppear,
		},
		{
			Source:  image.NewGray(image.Rect(0, 0, 10, 10)),
			Region:  image.Rect(0, 0, 0, 0),
			Pattern: image.NewGray(image.Rect(0, 0, 2, 2)),
			Event:   ObserveEventAppear,
		},
		{
			Source: image.NewGray(image.Rect(0, 0, 10, 10)),
			Region: image.Rect(0, 0, 10, 10),
			Event:  ObserveEventAppear,
		},
		{
			Source:  image.NewGray(image.Rect(0, 0, 10, 10)),
			Region:  image.Rect(0, 0, 10, 10),
			Pattern: image.NewGray(image.Rect(0, 0, 2, 2)),
			Event:   "bad-event",
		},
	}

	for i := range cases {
		if err := cases[i].Validate(); err == nil {
			t.Fatalf("case %d should fail validation", i)
		}
	}

	change := ObserveRequest{
		Source: image.NewGray(image.Rect(0, 0, 10, 10)),
		Region: image.Rect(0, 0, 10, 10),
		Event:  ObserveEventChange,
	}
	if err := change.Validate(); err != nil {
		t.Fatalf("change event should allow nil pattern: %v", err)
	}
}
