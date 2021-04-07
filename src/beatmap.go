package gosu

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Beatmap struct {
	BeatmapCompact
	Accuracy      float64    `json:"accuracy"`
	AR            float64    `json:"ar"`
	BeatmapSetID  int        `json:"beatmapset_id"`
	BPM           float64    `json:"bpm"`
	Convert       bool       `json:"convert"`
	CountCircles  int        `json:"count_circles"`
	CountSliders  int        `json:"count_sliders"`
	CountSpinners int        `json:"count_spinners"`
	CS            float64    `json:"cs"`
	DeletedAt     time.Time  `json:"deleted_at"`
	Drain         float64    `json:"drain"`
	HitLength     int        `json:"hit_length"`
	IsScoreable   bool       `json:"is_scoreable"`
	LastUpdated   time.Time  `json:"last_updated"`
	ModeInt       int        `json:"mode_int"`
	PassCount     int        `json:"passcount"`
	PlayCount     int        `json:"playcount"`
	Ranked        RankStatus `json:"ranked"`
	URL           string     `json:"url"`
}

type BeatmapCompact struct {
	DifficultyRating float64 `json:"difficulty_rating"`
	ID               int     `json:"id"`
	Mode             string  `json:"mode"` // TODO: make this compatible with enum (pain in the butt)
	Status           string  `json:"status"`
	TotalLength      int     `json:"total_length"`
	Version          string  `json:"version"`

	// Optional attributes
	Beatmapset Beatmapset `json:"beatmapset"`
	Checksum   string     `json:"checksum"`
	Failtimes  Failtimes  `json:"failtimes"`
	MaxCombo   int        `json:"max_combo"`
}

type BeatmapScores struct {
	Scores    []Score
	UserScore BeatmapUserScore
}

type BeatmapUserScore struct {
	Position int
	Score    Score
}

type Failtimes struct {
	Fail [100]int `json:"fail"`
	Exit [100]int `json:"exit"`
}

// Returns a *Beatmap, given data and query type (checksum, filename, or id)
func (c *GosuClient) LookupBeatmap(data string, queryType string) (*Beatmap, error) {
	queryType = strings.Trim(queryType, " ")

	// Validate the token
	if err := c.ValidateToken(); err != nil {
		fmt.Println("Error in validating token: %w", err)
		return nil, err
	}

	var ret *Beatmap

	if resp, err := c.DoRequest("GET", "beatmaps/lookup", map[string]interface{}{queryType: data}); err != nil {
		fmt.Println("Error in GET request for beatmaps/lookup: %w", err)
		return nil, err
	} else {
		if err = json.Unmarshal(resp, &ret); err != nil {
			fmt.Println("Error in unmarshaling GET request for beatmaps/lookup: %w", err)
			return nil, err
		}
	}

	return ret, nil
}
