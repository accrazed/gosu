package gosu

type Beatmapset struct {
	BeatmapsetCompact
	DownloadDisabled    bool
	MoreInformation     string
	BPM                 float64
	CanBeHyped          bool
	Creator             string
	DiscussionEnabled   bool
	DiscussionLocked    bool
	CurrentHype         int
	RequiredHype        int
	IsScoreable         bool
	LastUpdated         Timestamp
	LegacyThreadURL     string
	CurrentNominations  int
	RequiredNominations int
	Ranked              int
	RankedDate          Timestamp
	Storyboard          bool
	SubmittedDate       Timestamp
	Tags                string
}

type BeatmapsetCompact struct {
	Artist        string
	ArtistUnicode string
	Covers        Covers
	Creator       string
	FavoriteCount int
	ID            int
	NSFW          bool
	PlayCount     int
	PreviewURL    string
	Source        string
	Status        RankStatus
	Title         string
	TitleUnicode  string
	UserID        int
	Video         string

	// Optional attributes
	// FIXME: Actual Types for most of these fields are unspecified
	Beatmaps              []Beatmap
	Converts              string
	CurrentUserAttributes string
	Description           string
	Discussions           string
	Events                string
	Genre                 string
	HasFavorited          bool // always included
	Language              string
	Nominations           string
	Ratings               string
	RecentFavorites       string
	RelatedUsers          string
	User                  string
}
