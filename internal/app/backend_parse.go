package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sikulix/portgo/internal/core"
)

func parseBoolString(s string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true":
		return true, nil
	case "false", "":
		return false, nil
	default:
		return false, fmt.Errorf("unexpected bool value %q", strings.TrimSpace(s))
	}
}

func parseWindowsOutput(s string) ([]core.WindowInfo, error) {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	windows := make([]core.WindowInfo, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, "||")
		if len(parts) != 6 {
			return nil, fmt.Errorf("window parse failed for line %q", line)
		}
		x, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("window parse x failed: %w", err)
		}
		y, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			return nil, fmt.Errorf("window parse y failed: %w", err)
		}
		w, err := strconv.Atoi(strings.TrimSpace(parts[3]))
		if err != nil {
			return nil, fmt.Errorf("window parse w failed: %w", err)
		}
		h, err := strconv.Atoi(strings.TrimSpace(parts[4]))
		if err != nil {
			return nil, fmt.Errorf("window parse h failed: %w", err)
		}
		focused, err := parseBoolString(parts[5])
		if err != nil {
			return nil, fmt.Errorf("window parse focused failed: %w", err)
		}
		windows = append(windows, core.WindowInfo{
			Title:   parts[0],
			X:       x,
			Y:       y,
			W:       w,
			H:       h,
			Focused: focused,
		})
	}
	return windows, nil
}
