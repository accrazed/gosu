package gosu

import (
	"encoding/json"
	"fmt"
	"time"
)

// FIXME: Score struct has no type specifications
type Score struct {
	ID         int        `json:"id"`
	BestID     int        `json:"best_id"`
	UserID     int        `json:"user_id"`
	Accuracy   float64    `json:"accuracy"`
	Mods       []string   `json:"mods"` // TODO: add enum support
	Score      int        `json:"score"`
	MaxCombo   int        `json:"max_combo"`
	Perfect    bool       `json:"perfect"`
	Statistics ScoreCount `json:"statistics"`
	PP         float64    `json:"pp"`
	Rank       string     `json:"rank"` // TODO: add enum support
	CreatedAt  time.Time  `json:"created_at"`
	Mode       string     `json:"mode"`
	ModeInt    int        `json:"mode_int"`
	Replay     bool       `json:"replay"`

	// Optional attributes
	Beatmap     Beatmap    `json:"beatmap"`
	Beatmapset  Beatmapset `json:"beatmapset"`
	RankCountry int        `json:"rank_country"`
	RankGlobal  int        `json:"rank_global"`
	Weight      float64    `json:"weight"`
	User        User       `json:"user"`
	Match       string     `json:"match"`
}

type ScoreCount struct {
	Count50   int `json:"count_50"`
	Count100  int `json:"count_100"`
	Count300  int `json:"count_300"`
	CountGeki int `json:"count_geki"`
	CountKatu int `json:"count_katu"`
	CountMiss int `json:"count_miss"`
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

/*
Returns a *BeatmapScores, given:
	`beatmapID`: the ID of the beatmap
	`params`:
		`mode`: the gamemode to get scores for (string)
		`mods`: An array of matching mods
		`type`: Beatmap score ranking type
*/
func (c *GosuClient) GetBeatmapScores(beatmapID string, params map[string]interface{}) (*BeatmapScores, error) {
	var ret *BeatmapScores
	requestURL := "beatmaps/" + beatmapID + "/scores"

	if resp, err := c.DoRequest("GET", requestURL, params); err != nil {
		fmt.Println("Error in GET request for beatmaps/{beatmap}/scores:", err)
		return nil, err
	} else {
		if err = json.Unmarshal(resp, &ret); err != nil {
			fmt.Println("Error in unmarshaling GET request for beatmaps/{beatmap}/scores:", err)
			return nil, err
		}
	}

	return ret, nil
}
