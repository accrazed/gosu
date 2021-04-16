package gosu

import (
	"encoding/json"
	"fmt"
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
	Beatmapset BeatmapsetCompact `json:"beatmapset"`
	Checksum   string            `json:"checksum"`
	Failtimes  Failtimes         `json:"failtimes"`
	MaxCombo   int               `json:"max_combo"`
}

type BeatmapScores struct {
	Scores    []Score
	UserScore BeatmapUserScore
}

type BeatmapUserScore struct {
	Position int   `json:"position"`
	Score    Score `json:"score"`
}

type Failtimes struct {
	Fail [100]int `json:"fail"`
	Exit [100]int `json:"exit"`
}

// Returns a *Beatmap, given data and query type (checksum, filename, or id)
func (c *GosuClient) LookupBeatmap(data string, queryType string) (*Beatmap, error) {
	var ret *Beatmap

	if resp, err := c.DoRequest("GET", "beatmaps/lookup", map[string]interface{}{queryType: data}); err != nil {
		fmt.Println("Error in GET request for beatmaps/lookup:", err)
		return nil, err
	} else {
		if err = json.Unmarshal(resp, &ret); err != nil {
			fmt.Println("Error in unmarshaling GET request for beatmaps/lookup:", err)
			return nil, err
		}
	}

	return ret, nil
}

/*
Returns a *BeatmapUserScore, given:
	`beatmapID`: the ID of the beatmap
	`userID`: the user you want to get the score of
	`params`:
		`mode`: the gamemode to get scores for (string)
		`mods`: An array of matching mods
*/
func (c *GosuClient) GetUserBeatmapScore(beatmapID string, userID string, params map[string]interface{}) (*BeatmapUserScore, error) {
	var ret *BeatmapUserScore
	requestURL := "beatmaps/" + beatmapID + "/scores/users/" + userID

	if resp, err := c.DoRequest("GET", requestURL, params); err != nil {
		fmt.Println("Error in GET request for beatmaps/{beatmap}/scores/users/{user}:", err)
		return nil, err
	} else {
		if err = json.Unmarshal(resp, &ret); err != nil {
			fmt.Println("Error in unmarshaling GET request for beatmaps/{beatmap}/scores/users/{user}:", err)
			return nil, err
		}
	}

	return ret, nil
}

/*
Returns a *Beatmap, given:
	`beatmapID`: the ID of the beatmap
*/
func (c *GosuClient) GetBeatmap(beatmapID string) (*Beatmap, error) {
	var ret *Beatmap
	requestURL := "beatmaps/" + beatmapID

	if resp, err := c.DoRequest("GET", requestURL, nil); err != nil {
		fmt.Println("Error in GET request for beatmaps/{beatmap}:", err)
		return nil, err
	} else {
		if err = json.Unmarshal(resp, &ret); err != nil {
			fmt.Println("Error in unmarshaling GET request for beatmaps/{beatmap}:", err)
			return nil, err
		}
	}

	return ret, nil
}
