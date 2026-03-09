package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/smysnk/sikuligo/internal/core"
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
		if len(parts) != 6 && len(parts) != 9 {
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
		id := ""
		appName := ""
		pid := 0
		if len(parts) == 9 {
			id = strings.TrimSpace(parts[6])
			appName = strings.TrimSpace(parts[7])
			pid, err = strconv.Atoi(strings.TrimSpace(parts[8]))
			if err != nil {
				return nil, fmt.Errorf("window parse pid failed: %w", err)
			}
		}
		windows = append(windows, core.WindowInfo{
			ID:      id,
			App:     appName,
			PID:     pid,
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
