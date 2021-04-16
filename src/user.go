package gosu

import "time"

type GradeCounts struct {
	A   int `json:"a"`   // Number of A ranked scores.
	S   int `json:"s"`   // Number of S ranked scores.
	SH  int `json:"sh"`  // Number of Silver S ranked scores.
	SS  int `json:"ss"`  // Number of SS ranked scores.
	SSH int `json:"ssh"` // Number of Silver SS ranked scores.
}
type Group struct {
	ID             int    `json:"id"`
	Identifier     string `json:"identifier"`      // Unique string to identify the group.
	IsProbationary string `json:"is_probationary"` // Whether members of this group are considered probationary.
	HasPlaymodes   bool   `json:"has_playmodes"`   // If this group associates GameModes with a user's membership, e.g. BN/NAT members
	Name           string `json:"name"`
	ShortName      string `json:"short_name"` // Short Name of the group for display.
	Description    string `json:"description"`
	Color          string `json:"colour"`
}
type ProfileBanner struct {
	ID           int    `json:"id"`
	TournamentID int    `json:"tournament_id"`
	Image        string `json:"image"`
}

// FIXME: unspecified types
type ProfilePage struct {
	Me             string `json:"me"`
	RecentActivity string `json:"recent_activity"`
	Beatmaps       string `json:"beatmaps"`
	Historical     string `json:"historical"`
	Kudosu         Kudosu `json:"kudosu"`
	TopRanks       string `json:"top_ranks"`
	Medals         string `json:"medals"`
}
type User struct {
	UserCompact
	CoverURL     string        `json:"cover_url"`     // URL of profile cover
	Discord      string        `json:"discord"`       //
	HasSupported bool          `json:"has_supported"` // whether or not ever being a supporter in the past
	Interests    string        `json:"interests"`     //
	JoinDate     time.Time     `json:"join_date"`     //
	Kudosu       Kudosu        `json:"Kudosu"`        //
	Location     string        `json:"location"`      //
	MaxBlocks    int           `json:"max_blocks"`    // maximum int of users allowed to be blocked
	MaxFriends   int           `json:"max_friends"`   // maximum int of friends allowed to be added
	Occupation   string        `json:"occupation"`    //
	Playmode     GameMode      `json:"playmode"`      //
	Playstyle    string        `json:"playstyle"`     // Device choices of the user.
	PostCount    int           `json:"post_count"`    // int of forum posts
	ProfileOrder []ProfilePage `json:"profile_order"` // ordered array of sections in user profile page
	Title        string        `json:"title"`         // user-specific title
	TitleURL     string        `json:"title_url"`     //
	Twitter      string        `json:"twitter"`       //
	Website      string        `json:"website"`       //
}
type UserAccountHistory struct {
	ID        int       `json:"id"`        //
	UserType  string    `json:"type"`      // .note, restriction, or silence.
	Timestamp time.Time `json:"timestamp"` //
	Length    int       `json:"length"`    // In seconds.
}
type UserBadge struct {
	AwardedAt   time.Time `json:"awarded_at"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	URL         string    `json:"url"`
}
type UserCompact struct {
	AvatarURL     string    `json:"avatar_url"`      // URL of user's avatar
	CountryCode   string    `json:"country_code"`    // two-letter code representing user's country
	DefaultGroup  string    `json:"default_group"`   // Identifier of the default Group the user belongs to.
	ID            int       `json:"id"`              // unique identifier for user
	IsActive      bool      `json:"is_active"`       // has this account been active in the last x months?
	IsBot         bool      `json:"is_bot"`          // is this a bot account?
	IsDeleted     bool      `json:"is_deleted"`      // is the user currently online? (either on lazer or the new website)
	IsOnline      bool      `json:"is_online"`       //
	IsSupporter   bool      `json:"is_supported"`    // does this user have supporter?
	LastVisit     time.Time `json:"last_visit"`      // last access time. null if the user hides online presence
	PMFriendsOnly bool      `json:"pm_friends_only"` // whether or not the user allows PM from other than friends
	ProfileColor  string    `json:"profile_colour"`  // color of username/profile highlight, hex code (e.g. #333333)
	Username      string    `json:"username"`        // user's display Name

	// Optional attributes
	AccountHistory                   []UserAccountHistory   `json:"account_history"`
	ActiveTournamentBanner           ProfileBanner          `json:"active_tournament_banner"`
	Badges                           []UserBadge            `json:"badges"`
	BeatmapsPlaycountCount           int                    `json:"beatmap_playcounts_count"`
	Blocks                           string                 `json:"blocks"`  // FIXME: Unspecified in docs
	Country                          Country                `json:"country"` // FIXME: Unspecified in docs
	Cover                            UserCover              `json:"cover"`   // FIXME: Unspecified in docs
	FavoriteBeatmapsetCount          int                    `json:"favourite_beatmapset_count"`
	FollowerCount                    int                    `json:"follower_count"`
	Friends                          []string               `json:"friends"` // FIXME: Unspecified in docs
	GraveyardBeatmapsetCount         int                    `json:"graveyard_beatmapset_count"`
	Groups                           []UserGroup            `json:"groups"`
	IsAdmin                          bool                   `json:"is_admin"`
	IsBNG                            bool                   `json:"is_bng"`
	IsFullBN                         bool                   `json:"is_full_bn"`
	IsGMT                            bool                   `json:"is_gmt"`
	IsLimitedBN                      bool                   `json:"is_limited_bn"`
	IsModerator                      bool                   `json:"is_moderator"`
	IsNAT                            bool                   `json:"is_nat"`
	IsRestricted                     bool                   `json:"is_restricted"`
	IsSilenced                       bool                   `json:"is_silenced"`
	LovedBeatmapsetCount             int                    `json:"loved_beatmapset_count"`
	MonthlyPlaycounts                []UserMonthlyPlaycount `json:"monthly_playcounts"`
	Page                             string                 `json:"page "`              // FIXME: Unspecified in docs
	PreviousUsernames                []string               `json:"previous_usernames"` // FIXME: Unspecified in docs
	RankedAndApprovedBeatmapsetCount int                    `json:"ranked_and_approved_beatmapset_count"`
	ReplaysWatchedCount              int                    `json:"replays_watched_counts"`
	ScoresBestCount                  int                    `json:"scores_best_count"`
	ScoresFirstCount                 int                    `json:"scores_first_count"`
	ScoresRecentCount                int                    `json:"scores_recent_count"`
	Statistics                       *UserStatistics        `json:"statistics"` // FIXME: Unspecified in docs
	StatisticsRulesets               UserStatisticsRulesets `json:"statistics_rulesets"`
	SupportLevel                     int                    `json:"support_level"`             // FIXME: Unspecified in docs
	UnrankedBeatmapsetCount          int                    `json:"unranked_beatmapset_count"` // FIXME: Unspecified in docs
	UnreadPMCount                    int                    `json:"unread_pm_count"`           // FIXME: Unspecified in docs
	UserAchievements                 int                    `json:"user_achievements"`         // FIXME: Unspecified in docs
	UserPreferences                  string                 `json:"user_preferences"`          // FIXME: Unspecified in docs
	RankHistory                      string                 `json:"rank_history"`              // FIXME: Unspecified in docs
}

type UserGroup struct {
	ID             int      `json:"id"`              // ID (of Group)
	Identifier     string   `json:"identifier"`      // Unique string to identify the group.
	IsProbationary bool     `json:"is_probationary"` // Whether members of this group are considered probationary.
	Name           string   `json:"name"`            //
	ShortName      string   `json:"short_name"`      // Short Name of the group for display.
	Description    string   `json:"description"`     //
	Color          string   `json:"colour"`          //
	Playmodes      []string `json:"playmodes"`       // GameModes which the member is responsible for, e.g. in the case of BN/NAT (only present when hasPlaymodes is set on Group)
}

type UserLevel struct {
	Current  int `json:"current"`
	Progress int `json:"progress"`
}

type UserMonthlyPlaycount struct {
	// TODO: implement when specified in docs
}

type UserStatistics struct {
	GradeCounts            GradeCounts
	HitAccuracy            int          `json:"hit_accuracy"`              // Hit accuracy percentage
	IsRanked               bool         `json:"is_ranked"`                 // Is actively ranked
	Level                  UserLevel    `json:"level"`                     //
	MaxCombo               int          `json:"maximum_combo"`             // Highest maximum combo.
	PlayCount              int          `json:"play_count"`                // Number of maps played.
	PlayTime               int          `json:"play_time"`                 // Cumulative time played.
	PP                     int          `json:"pp"`                        // Performance points
	GlobalRank             int          `json:"global_rank"`               // Current rank according to pp.
	RankedScore            int          `json:"ranked_score"`              // Current ranked score.
	ReplaysWatchedByOthers int          `json:"replays_watched_by_others"` // Number of replays watched by other users.
	TotalHits              int          `json:"total_hits"`                // Total int of hits.
	TotalScore             int          `json:"total_score"`               // Total score.
	User                   *UserCompact `json:"user"`                      // The associated user.
}

type UserStatisticsRulesets struct {
	// TODO: implement when specified in docs
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Kudosu struct {
	Available int `json:"available"`
	Total     int `json:"total"`
}
