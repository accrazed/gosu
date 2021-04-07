package gosu

import "time"

type GradeCounts struct {
	A   int // Number of A ranked scores.
	S   int // Number of S ranked scores.
	SH  int // Number of Silver S ranked scores.
	SS  int // Number of SS ranked scores.
	SSH int // Number of Silver SS ranked scores.
}
type Group struct {
	ID             int
	Identifier     string // Unique string to identify the group.
	IsProbationary string // Whether members of this group are considered probationary.
	HasPlaymodes   bool   // If this group associates GameModes with a user's membership, e.g. BN/NAT members
	Name           string
	ShortName      string // Short Name of the group for display.
	Description    string
	Color          string
}
type ProfileBanner struct {
	ID           int
	TournamentID int
	Image        string
}

// FIXME: unspecified types
type ProfilePage struct {
	Me             string
	RecentActivity string
	Beatmaps       string
	Historical     string
	Kudosu         string
	TopRanks       string
	Medals         string
}
type User struct {
	UserCompact
	CoverURL        string // URL of profile cover
	Discord         string //
	HasSupported    bool   // whether or not ever being a supporter in the past
	Interests       string //
	JoinDate        time.Time
	AvailableKudosu int           //
	TotalKudosu     int           //
	Location        string        //
	MaxBlocks       int           // maximum int of users allowed to be blocked
	MaxFriends      int           // maximum int of friends allowed to be added
	Occupation      string        //
	Playmode        GameMode      //
	Playstyle       string        // Device choices of the user.
	PostCount       int           // int of forum posts
	ProfileOrder    []ProfilePage // ordered array of sections in user profile page
	Title           string        // user-specific title
	TitleURL        string        //
	Twitter         string        //
	Website         string        //
}
type UserAccountHistory struct {
	ID        int       //
	UserType  string    // .note, restriction, or silence.
	Timestamp time.Time //
	Length    int       // In seconds.
}
type UserBadge struct {
	AwardedAt   time.Time
	Description string
	ImageURL    string
	URL         string
}
type UserCompact struct {
	AvatarURL     string    // URL of user's avatar
	CountryCode   string    // two-letter code representing user's country
	DefaultGroup  string    // Identifier of the default Group the user belongs to.
	ID            int       // unique identifier for user
	IsActive      bool      // has this account been active in the last x months?
	IsBot         bool      // is this a bot account?
	IsDeleted     bool      // is the user currently online? (either on lazer or the new website)
	IsOnline      bool      //
	IsSupporter   bool      // does this user have supporter?
	LastVisit     time.Time // last access time. null if the user hides online presence
	PMFriendsOnly bool      // whether or not the user allows PM from other than friends
	ProfileColor  string    // color of username/profile highlight, hex code (e.g. #333333)
	Username      string    // user's display Name

	// Optional attributes
	AccountHistory                   []UserAccountHistory
	ActiveTournamentsBanner          ProfileBanner
	Badges                           []UserBadge
	BeatmapsPlaycountCount           int
	Blocks                           string // FIXME: Unspecified in docs
	Country                          string // FIXME: Unspecified in docs
	Cover                            Covers // FIXME: Unspecified in docs
	FavoriteBeatmapsetCount          int
	FollowerCount                    int
	Friends                          []string // FIXME: Unspecified in docs
	GraveyardBeatmapsetCount         int
	Groups                           []UserGroup
	IsAdmin                          bool
	IsBNG                            bool
	IsFullBN                         bool
	IsGMT                            bool
	IsLimitedBN                      bool
	IsModerator                      bool
	IsNAT                            bool
	IsRestricted                     bool
	IsSilenced                       bool
	LovedBeatmapsetCount             int
	MonthlyPlaycounts                []UserMonthlyPlaycount
	Page                             string   // FIXME: Unspecified in docs
	PreviousUsernames                []string // FIXME: Unspecified in docs
	RankedAndApprovedBeatmapsetCount int
	ReplaysWatchedCount              int
	ScoresBestCount                  int
	ScoresFirstCount                 int
	ScoresRecentCount                int
	Statistics                       *UserStatistics // FIXME: Unspecified in docs
	StatisticsRulesets               UserStatisticsRulesets
	SupportLevel                     int    // FIXME: Unspecified in docs
	UnrankedBeatmapsetCount          int    // FIXME: Unspecified in docs
	UnreadPMCount                    int    // FIXME: Unspecified in docs
	UserAchievements                 int    // FIXME: Unspecified in docs
	UserPreferences                  string // FIXME: Unspecified in docs
	RankHistory                      string // FIXME: Unspecified in docs
}

type UserGroup struct {
	ID             int      // ID (of Group)
	Identifier     string   // Unique string to identify the group.
	IsProbationary bool     // Whether members of this group are considered probationary.
	Name           string   //
	ShortName      string   // Short Name of the group for display.
	Description    string   //
	Color          string   //
	Playmodes      []string // GameModes which the member is responsible for, e.g. in the case of BN/NAT (only present when hasPlaymodes is set on Group)
}

type UserLevel struct {
	Current  int
	Progress int
}

type UserMonthlyPlaycount struct {
	// TODO: implement when specified in docs
}

type UserStatistics struct {
	GradeCounts            GradeCounts
	HitAccuracy            int          // Hit accuracy percentage
	IsRanked               bool         // Is actively ranked
	LevelCurrent           int          // Current level.
	LevelProgress          int          // Progress to next level.
	MaxCombo               int          // Highest maximum combo.
	PlayCount              int          // Number of maps played.
	PlayTime               int          // Cumulative time played.
	PP                     int          // Performance points
	GlobalRank             int          // Current rank according to pp.
	RankedScore            int          // Current ranked score.
	ReplaysWatchedByOthers int          // Number of replays watched by other users.
	TotalHits              int          // Total int of hits.
	TotalScore             int          // Total score.
	User                   *UserCompact // The associated user.
}

type UserStatisticsRulesets struct {
	// TODO: implement when specified in docs
}
