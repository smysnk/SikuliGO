package cv

import (
	"testing"
)

func TestParseMatcherEngine(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want MatcherEngine
		ok   bool
	}{
		{name: "default_empty", in: "", want: MatcherEngineTemplate, ok: true},
		{name: "template", in: "template", want: MatcherEngineTemplate, ok: true},
		{name: "template_alias_ncc", in: "ncc", want: MatcherEngineTemplate, ok: true},
		{name: "template_alias_opencv", in: "opencv", want: MatcherEngineTemplate, ok: true},
		{name: "orb", in: "orb", want: MatcherEngineORB, ok: true},
		{name: "hybrid", in: "hybrid", want: MatcherEngineHybrid, ok: true},
		{name: "invalid", in: "unknown", want: "", ok: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseMatcherEngine(tc.in)
			if tc.ok && err != nil {
				t.Fatalf("expected success, got err=%v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected parse error")
			}
			if tc.ok && got != tc.want {
				t.Fatalf("engine mismatch got=%q want=%q", got, tc.want)
			}
		})
	}
}

func TestNewMatcherForEngine(t *testing.T) {
	tests := []struct {
		name   string
		engine MatcherEngine
		ok     bool
	}{
		{name: "template", engine: MatcherEngineTemplate, ok: true},
		{name: "orb", engine: MatcherEngineORB, ok: true},
		{name: "hybrid", engine: MatcherEngineHybrid, ok: true},
		{name: "invalid", engine: MatcherEngine("invalid"), ok: false},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			m, err := NewMatcherForEngine(tc.engine)
			if tc.ok && err != nil {
				t.Fatalf("expected success, got err=%v", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected error for engine %q", tc.engine)
			}
			if tc.ok && m == nil {
				t.Fatalf("expected matcher instance")
			}
		})
	}
}
