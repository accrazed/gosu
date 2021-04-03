package gosu

type Beatmap struct {
	BeatmapCompact
	Accuracy     float64
	AR           float64
	BeatmapSetID int
	BPM          float64
	Convert      bool
	CountCircles int
	CountSliders int
	CS           float64
	DeletedAt    Timestamp
	Drain        float64
	HitLength    int
	IsScoreable  bool
	LastUpdated  Timestamp
	ModeInt      int
	PassCount    int
	PlayCount    int
	Ranked       RankStatus
	URL          string
}

type BeatmapCompact struct {
	DifficultyRating float64
	ID               int
	Mode             GameMode
	Status           RankStatus
	TotalLength      int
	Version          string

	// Optional attributes
	Beatmapset    Beatmapset
	Checksum      string
	FailtimesExit [100]int
	FailtimesFail [100]int
}

type BeatmapScores struct {
	Scores    []Score
	UserScore BeatmapUserScore
}

type BeatmapUserScore struct {
	Position int
	Score    Score
}
