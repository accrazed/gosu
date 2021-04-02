package gosu

// FIXME: Score struct has no type specifications
type Score struct {
	id         int
	bestID     int
	userID     int
	accuracy   float64
	mods       GameMod
	score      int
	maxCombo   int
	perfect    bool
	statistics ScoreCount
	pp         float64
	rank       RankStatus
	createdAt  Timestamp
	mode       string
	modeInt    int
	replay     string

	// Optional attributes
	beatmap     Beatmap
	beatmapset  Beatmapset
	rankCountry int
	rankGlobal  int
	weight      float64
	user        User
	match       string
}
type ScoreCount struct {
	count50   int
	count100  int
	count300  int
	countGeki int
	countKatu int
	countMiss int
}

type GameMod int

const (
	NoFail GameMod = 1 << iota
	Easy
	TouchDevice
	Hidden
	HardRock
	SuddenDeath
	DoubleTime
	Relax
	HalfTime
	Nightcore // Only set along with DoubleTime. i.e: NC only gives 576
	Flashlight
	Autoplay
	SpunOut
	Relax2  // Autopilot
	Perfect // Only set along with SuddenDeath. i.e: PF only gives 16416
	Key4
	Key5
	Key6
	Key7
	Key8
	FadeIn
	Random
	Cinema
	Target
	Key9
	KeyCoop
	Key1
	Key3
	Key2
	ScoreV2
	Mirror
	None              = 0
	KeyMod            = Key1 | Key2 | Key3 | Key4 | Key5 | Key6 | Key7 | Key8 | Key9 | KeyCoop
	FreeModAllowed    = NoFail | Easy | Hidden | HardRock | SuddenDeath | Flashlight | FadeIn | Relax | Relax2 | SpunOut | KeyMod
	ScoreIncreaseMods = Hidden | HardRock | DoubleTime | Flashlight | FadeIn
)
