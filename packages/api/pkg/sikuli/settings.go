package sikuli

import "sync"

type RuntimeSettings struct {
	ImageCache       int
	ShowActions      bool
	WaitScanRate     float64
	ObserveScanRate  float64
	AutoWaitTimeout  float64
	MinSimilarity    float64
	FindFailedThrows bool
}

var (
	settingsMu sync.RWMutex
	settings   = RuntimeSettings{
		ImageCache:       64,
		ShowActions:      false,
		WaitScanRate:     DefaultWaitScanRate,
		ObserveScanRate:  DefaultObserveScanRate,
		AutoWaitTimeout:  DefaultAutoWaitTimeout,
		MinSimilarity:    DefaultSimilarity,
		FindFailedThrows: true,
	}
)

func GetSettings() RuntimeSettings {
	settingsMu.RLock()
	defer settingsMu.RUnlock()
	return settings
}

func UpdateSettings(apply func(*RuntimeSettings)) RuntimeSettings {
	settingsMu.Lock()
	defer settingsMu.Unlock()
	apply(&settings)
	return settings
}

func ResetSettings() RuntimeSettings {
	settingsMu.Lock()
	defer settingsMu.Unlock()
	settings = RuntimeSettings{
		ImageCache:       64,
		ShowActions:      false,
		WaitScanRate:     DefaultWaitScanRate,
		ObserveScanRate:  DefaultObserveScanRate,
		AutoWaitTimeout:  DefaultAutoWaitTimeout,
		MinSimilarity:    DefaultSimilarity,
		FindFailedThrows: true,
	}
	return settings
}

