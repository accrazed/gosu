package gosu

type MultiplayerScore struct {
	ID             int
	UserID         int
	RoomID         int
	PlaylistItemID int
	BeatmapID      int
	Rank           RankStatus
	TotalScore     int
	Accuracy       int
	MaxCombo       int
	Mods           []GameMod
	Statistics     UserStatistics
	Passed         bool
	Position       int
	ScoresAround   MultiplayerScoresAround // Scores around the specified score.
	User           User
}

type MultiplayerScores struct {
	Cursor     MultiplayerScoresCursor
	Params     []string // not specified
	Scores     []MultiplayerScore
	Total      int
	UserScores MultiplayerScore
}

type MultiplayerScoresAround struct {
	Higher *MultiplayerScores
	Lower  *MultiplayerScores
}

type MultiplayerScoresCursor struct {
	ScoreID    int // Last score ID of current result (score_asc, score_desc)
	TotalScore int // Last score's total score of current result (score_asc, score_desc)
}

type MultiplayerScoresSort struct {
	ScoreAsc  string // TODO: Figure out what the fuck this means
	ScoreDesc string
}
