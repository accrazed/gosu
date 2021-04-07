package gosu

import "time"

// FIXME: Score struct has no type specifications
type Score struct {
	ID         int
	BestID     int
	UserID     int
	Accuracy   float64
	Mods       GameMod
	Score      int
	MaxCombo   int
	Perfect    bool
	Statistics ScoreCount
	PP         float64
	Rank       RankStatus
	CreatedAt  time.Time
	Mode       string
	ModeInt    int
	Replay     string

	// Optional attributes
	Beatmap     Beatmap
	Beatmapset  Beatmapset
	RankCountry int
	RankGlobal  int
	Weight      float64
	User        User
	Match       string
}

type ScoreCount struct {
	Count50   int
	Count100  int
	Count300  int
	CountGeki int
	CountKatu int
	CountMiss int
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
