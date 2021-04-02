package gosu

type BeatmapCompact struct {
	difficultyRating float64
	id               int
	mode             Gamemode
	status           RankStatus // same as ranked - originally a string
	totalLength      int
	version          string
}

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
	deletedAt    string // TODO: may become a Timestamp object
	drain        float64
	hitLength    int
	isScoreable  bool
	lastUpdated  string // TODO: may become a Timestamp object
	modeInt      int
	passCount    int
	playCount    int
	ranked       RankStatus // same as status - originally an int
	url          string
}
