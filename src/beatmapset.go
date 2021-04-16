package gosu

import "time"

type Beatmapset struct {
	BeatmapsetCompact `json:"beatmapset_compact"`
	Availability
	BPM               float64   `json:"bpm"`
	CanBeHyped        bool      `json:"can_be_hyped"`
	Creator           string    `json:"creator"`
	DiscussionEnabled bool      `json:"discussion_enabled"`
	DiscussionLocked  bool      `json:"discussion_locked"`
	Hype              Hype      `json:"hype"`
	IsScoreable       bool      `json:"is_scoreable"`
	LastUpdated       time.Time `json:"last_updated"`
	LegacyThreadURL   string    `json:"legacy_thread_url"`
	Ranked            int       `json:"ranked"`
	RankedDate        time.Time `json:"ranked_date"`
	Source            string    `json:"source"`
	Storyboard        bool      `json:"storyboard"`
	SubmittedDate     time.Time `json:"submitted_date"`
	Tags              string    `json:"tags"`
}

type BeatmapsetCompact struct {
	Artist        string `json:"artist"`
	ArtistUnicode string `json:"artist_unicode"`
	Covers        Covers `json:"covers"`
	Creator       string `json:"creators"`
	FavoriteCount int    `json:"favourite_count"`
	ID            int    `json:"id"`
	NSFW          bool   `json:"nsfw"`
	PlayCount     int    `json:"play_count"`
	PreviewURL    string `json:"preview_url"`
	Source        string `json:"source"`
	Status        string `json:"status"`
	Title         string `json:"title"`
	TitleUnicode  string `json:"title_unicode"`
	UserID        int    `json:"user_id"`
	Video         bool   `json:"video"`

	// Optional attributes
	// FIXME: Actual Types for most of these fields are unspecified
	Beatmaps              []Beatmap   `json:"beatmaps"`
	Converts              string      `json:"converts"`
	CurrentUserAttributes string      `json:"current_user_attributes"`
	Description           string      `json:"description"`
	Discussions           string      `json:"discussions"`
	Events                string      `json:"events"`
	Genre                 string      `json:"genre"`
	HasFavorited          bool        `json:"has_favourited"` // always included
	Language              string      `json:"language"`
	Nominations           Nominations `json:"nominations_summary"`
	Ratings               [100]int    `json:"ratings"`
	RecentFavorites       string      `json:"recent_favourites"`
	RelatedUsers          string      `json:"related_users"`
	User                  string      `json:"user"`
}

type Availability struct {
	DownloadDisabled bool   `json:"download_disabled"`
	MoreInformation  string `json:"more_information"`
}

type Hype struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}

type Nominations struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}
