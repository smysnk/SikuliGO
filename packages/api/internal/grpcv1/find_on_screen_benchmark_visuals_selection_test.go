package grpcv1

import "testing"

func TestSelectMegaSummaryImagesPreferResolution(t *testing.T) {
	input := []findBenchScenarioSummaryImage{
		{ScenarioName: "vector_ui_baseline_800x600_i01_scale_a", Path: "/tmp/vector-800.png"},
		{ScenarioName: "vector_ui_baseline_1280x720_i01_scale_b", Path: "/tmp/vector-1280.png"},
		{ScenarioName: "photo_clutter_800x600_i02_rotate_c", Path: "/tmp/photo-800.png"},
		{ScenarioName: "photo_clutter_1280x720_i02_rotate_d", Path: "/tmp/photo-1280.png"},
		{ScenarioName: "orb_feature_rich_800x600_i07_rotate_e", Path: "/tmp/orb-800.png"},
	}

	selected := selectMegaSummaryImages(input, "1280x720")
	if got, want := len(selected), 3; got != want {
		t.Fatalf("selected count mismatch: got=%d want=%d", got, want)
	}

	byFamily := map[string]string{}
	for _, item := range selected {
		family, _ := splitScenarioFamilyAndResolution(item.ScenarioName)
		byFamily[family] = item.ScenarioName
	}
	if got := byFamily["vector_ui_baseline"]; got != "vector_ui_baseline_1280x720_i01_scale_b" {
		t.Fatalf("vector selection mismatch: got=%q", got)
	}
	if got := byFamily["photo_clutter"]; got != "photo_clutter_1280x720_i02_rotate_d" {
		t.Fatalf("photo selection mismatch: got=%q", got)
	}
	// Fallback to any available resolution when preferred is absent.
	if got := byFamily["orb_feature_rich"]; got != "orb_feature_rich_800x600_i07_rotate_e" {
		t.Fatalf("orb fallback mismatch: got=%q", got)
	}
}

func TestSplitScenarioFamilyAndResolution(t *testing.T) {
	family, res := splitScenarioFamilyAndResolution("noise_stress_random_1920x1080_i04_rotate_x")
	if family != "noise_stress_random" || res != "1920x1080" {
		t.Fatalf("unexpected parse result family=%q res=%q", family, res)
	}
}

