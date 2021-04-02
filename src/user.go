package gosu

type GradeCounts struct {
	a   int // Number of A ranked scores.
	s   int // Number of S ranked scores.
	sh  int // Number of Silver S ranked scores.
	ss  int // Number of SS ranked scores.
	ssh int // Number of Silver SS ranked scores.
}
type Group struct {
	id             int
	identifier     string // Unique string to identify the group.
	isProbationary string // Whether members of this group are considered probationary.
	hasPlaymodes   bool   // If this group associates GameModes with a user's membership, e.g. BN/NAT members
	name           string
	shortName      string // Short name of the group for display.
	description    string
	color          string
}
type ProfileBanner struct {
	id           int
	tournamentID int
	image        string
}

type ProfilePage struct {
	me             string // FIXME: unspecified type
	recentActivity string // FIXME: unspecified type
	beatmaps       string // FIXME: unspecified type
	historical     string // FIXME: unspecified type
	kudosu         string // FIXME: unspecified type
	topRanks       string // FIXME: unspecified type
	medals         string // FIXME: unspecified type
}
type User struct {
	UserCompact
	coverURL        string // url of profile cover
	discord         string //
	hasSupported    bool   // whether or not ever being a supporter in the past
	interests       string //
	joinDate        Timestamp
	availableKudosu int           //
	totalKudosu     int           //
	location        string        //
	maxBlocks       int           // maximum int of users allowed to be blocked
	maxFriends      int           // maximum int of friends allowed to be added
	occupation      string        //
	playmode        GameMode      //
	playstyle       string        // Device choices of the user.
	postCount       int           // int of forum posts
	profileOrder    []ProfilePage // ordered array of sections in user profile page
	title           string        // user-specific title
	titleURL        string        //
	twitter         string        //
	website         string        //
}
type UserAccountHistory struct {
	id        int       //
	userType  string    // .note, restriction, or silence.
	timestamp Timestamp //
	length    int       // In seconds.
}
type UserBadge struct {
	awardedAt   Timestamp
	description string
	imageURL    string
	url         string
}
type UserCompact struct {
	avatarURL     string    // url of user's avatar
	countryCode   string    // two-letter code representing user's country
	defaultGroup  string    // Identifier of the default Group the user belongs to.
	id            int       // unique identifier for user
	isActive      bool      // has this account been active in the last x months?
	isBot         bool      // is this a bot account?
	isDeleted     bool      // is the user currently online? (either on lazer or the new website)
	isOnline      bool      //
	isSupporter   bool      // does this user have supporter?
	lastVisit     Timestamp // last access time. null if the user hides online presence
	pmFriendsOnly bool      // whether or not the user allows PM from other than friends
	profileColor  string    // color of username/profile highlight, hex code (e.g. #333333)
	username      string    // user's display name

	// Optional attributes
	accountHistory                   []UserAccountHistory
	activeTournamentsBanner          ProfileBanner
	badges                           []UserBadge
	BeatmapsPlaycountCount           int
	blocks                           string // FIXME: Unspecified in docs
	country                          string // FIXME: Unspecified in docs
	cover                            Covers // FIXME: Unspecified in docs
	favoriteBeatmapsetCount          int
	followerCount                    int
	friends                          []string // FIXME: Unspecified in docs
	graveyardBeatmapsetCount         int
	groups                           []UserGroup
	isAdmin                          bool
	isBNG                            bool
	isFullBN                         bool
	isGMT                            bool
	isLimitedBN                      bool
	isModerator                      bool
	isNAT                            bool
	isRestricted                     bool
	isSilenced                       bool
	lovedBeatmapsetCount             int
	monthlyPlaycounts                []UserMonthlyPlaycount
	page                             string   // FIXME: Unspecified in docs
	previousUsernames                []string // FIXME: Unspecified in docs
	rankedAndApprovedBeatmapsetCount int
	replaysWatchedCount              int
	scoresBestCount                  int
	scoresFirstCount                 int
	scoresRecentCount                int
	statistics                       *UserStatistics // FIXME: Unspecified in docs
	statisticsRulesets               UserStatisticsRulesets
	supportLevel                     int    // FIXME: Unspecified in docs
	unrankedBeatmapsetCount          int    // FIXME: Unspecified in docs
	unreadPMCount                    int    // FIXME: Unspecified in docs
	userAchievements                 int    // FIXME: Unspecified in docs
	userPreferences                  string // FIXME: Unspecified in docs
	rankHistory                      string // FIXME: Unspecified in docs
}

type UserGroup struct {
	id             int      // ID (of Group)
	identifier     string   // Unique string to identify the group.
	isProbationary bool     // Whether members of this group are considered probationary.
	name           string   //
	shortName      string   // Short name of the group for display.
	description    string   //
	color          string   //
	playmodes      []string // GameModes which the member is responsible for, e.g. in the case of BN/NAT (only present when hasPlaymodes is set on Group)
}

type UserLevel struct {
	current  int
	progress int
}

type UserMonthlyPlaycount struct {
	// TODO: implement when specified in docs
}

type UserStatistics struct {
	gradeCounts            GradeCounts
	hitAccuracy            int          // Hit accuracy percentage
	isRanked               bool         // Is actively ranked
	levelCurrent           int          // Current level.
	levelProgress          int          // Progress to next level.
	maxCombo               int          // Highest maximum combo.
	playCount              int          // Number of maps played.
	playTime               int          // Cumulative time played.
	pp                     int          // Performance points
	globalRanke            int          // Current rank according to pp.
	rankedScore            int          // Current ranked score.
	replaysWatchedByOthers int          // Number of replays watched by other users.
	totalHits              int          // Total int of hits.
	totalScore             int          // Total score.
	user                   *UserCompact // The associated user.
}

type UserStatisticsRulesets struct {
	// TODO: implement when specified in docs
}
