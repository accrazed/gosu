package gosu

type Beatmap struct {
	BeatmapCompact
	accuracy     float64
	ar           float64
	beatmapSetID int
	bpm          float64
	convert      bool
	countCircles int
	countSliders int
	cs           float64
	deletedAt    Timestamp
	drain        float64
	hitLength    int
	isScoreable  bool
	lastUpdated  Timestamp
	modeInt      int
	passCount    int
	playCount    int
	ranked       RankStatus
	url          string
}

type BeatmapCompact struct {
	difficultyRating float64
	id               int
	mode             GameMode
	status           RankStatus
	totalLength      int
	version          string

	// Optional attributes
	beatmapset    Beatmapset
	checksum      string
	failtimesExit [100]int
	failtimesFail [100]int
}

type BeatmapScores struct {
	scores    []Score
	userScore BeatmapUserScore
}

type BeatmapUserScore struct {
	position int
	score    Score
}
