package sikuli

// MatcherEngine selects the server-side matcher implementation used for live screen queries.
type MatcherEngine string

const (
	MatcherEngineDefault  MatcherEngine = ""
	MatcherEngineTemplate MatcherEngine = "template"
	MatcherEngineORB      MatcherEngine = "orb"
	MatcherEngineHybrid   MatcherEngine = "hybrid"
	MatcherEngineAKAZE    MatcherEngine = "akaze"
	MatcherEngineBRISK    MatcherEngine = "brisk"
	MatcherEngineKAZE     MatcherEngine = "kaze"
	MatcherEngineSIFT     MatcherEngine = "sift"
)
