package gosu

// See https://osu.ppy.sh/docs/index.html?javascript#chatchannel
type Channel string

const (
	PUBLIC      Channel = "PUBLIC"
	PRIVATE     Channel = "PRIVATE"
	MULTIPLAYER Channel = "MULTIPLAYER"
	SPECTATOR   Channel = "SPECTATOR"
	PM          Channel = "PM"
	GROUP       Channel = "GROUP"
)

type BeatmapsetDiscussion struct {
	Beatmap               BeatmapCompact
	BeatmapID             int
	Beatmapset            BeatmapsetCompact
	BeatmapsetID          int
	CanBeResolved         bool
	CanGrandKudosu        bool
	CreatedAt             string
	CurrentUserAttributes CurrentUserAttributes
	DeletedAt             Timestamp
	DeletedByID           int
	ID                    int
	KudosuDenied          bool
	LastPostAt            Timestamp
	MessageType           string // FIXME: originally a "MessageType" object, not in docs
	ParentID              int
	Posts                 []BeatmapsetDiscussionPost
	Resolved              bool
	StartingPost          BeatmapsetDiscussionPost
	Timestamp             Timestamp
	UpdatedAt             Timestamp
	UserID                int
	Votes                 []BeatmapsetDiscussionVote // FIXME: structure will change
}

type BeatmapsetDiscussionPost struct {
	BeatmapsetDiscussionID int
	CreatedAt              Timestamp
	DeletedAt              Timestamp
	DeletedByID            int
	ID                     int
	LastEditorID           int
	Message                string
	System                 bool
	UpdatedAt              Timestamp
	UserID                 int
}
type BeatmapsetDiscussionVote struct {
	BeatmapsetDiscussionID int
	CreatedAt              Timestamp
	ID                     int
	Score                  int
	UpdatedAt              Timestamp
	UserID                 int
}

type ChatChannel struct {
	ChannelID      int
	Name           string
	Description    string
	Icon           string        // display icon for the channel
	ChannelType    Channel       // see channel types below
	FirstMessageID int           // messageID of first message (only returned in presence responses)
	LastReadID     int           // messageID of last message read (only returned in presence responses)
	LastMessageID  int           // messageID of last known message (only returned in presence responses)
	RecentMessages []ChatMessage // up to 50 most recent messages
	Moderated      bool          // user can't send message when the value is true (only returned in presence responses)
	Users          []int         // array of userID that are in the channel (not included for PUBLIC channels)
}

type ChatMessage struct {
	MessageID int         // unique Identifier for message
	SenderID  int         // userID of the sender
	ChannelID int         // channelID of where the message was sent
	Timestamp Timestamp   // when the message was sent, ISO-8601
	Content   string      // message content
	IsAction  bool        // was this an action i.e. /me dances
	Sender    UserCompact // embeded UserCompact object to save additional api lookups
}
type Comment struct {
	CommentableID   int    // ID of the object the comment is attached to
	CommentableType string // type of object the comment is attached to
	CreatedAt       string // ISO 8601 date
	DeletedAt       string // ISO 8601 date if the comment was deleted; null, otherwise
	EditedAt        string // ISO 8601 date if the comment was edited; null, otherwise
	EditedByID      int    // user ID of the user that edited the post; null, otherwise
	ID              int    // the ID of the comment
	LegacyName      string // username displayed on legacy comments
	Message         string // markdown of the comment's content
	MessageHTML     string // html version of the comment's content
	ParentID        int    // ID of the comment's parent
	Pinned          bool   // Pin status of the comment
	RepliesCount    int    // int of replies to the comment
	UpdatedAt       string // ISO 8601 date
	UserID          int    // user ID of the poster
	VotesCount      int    // int of votes
}
type CommentBundle struct {
	CommentableMeta  CommentableMeta // ID of the object the comment is attached to
	Comments         []Comment       // array of comments ordered according to sort;
	Cursor           Cursor          // no desc
	HasMore          bool            // If there are more comments or replies available
	HasMoreID        int             // no desc
	IncludedComments []Comment       // related comments; e.g. parent comments and nested replies
	PinnedComments   []Comment       // pinned comments
	Sort             string          // one of the CommentSort types
	TopLevelCount    int             // num of comments at the top level. Not returned for replies.
	Total            int             // total num of comments. Not retuned for replies.
	UserFollow       bool            // is the current user watching the comment thread
	UserVotes        []int           // IDs of the comments in the bundle the current user has upvoted
	Users            []UserCompact   // array of users related to the comments
}

// TODO: IMPLEMENT TYPE "CommentSort" once I know what it actually is

type CommentableMeta struct {
	ID         int    // the ID of the object
	Title      string // display title
	ObjectType string // the type of the object
	URL        string // URL of the object
}

type CurrentUserAttributes struct {
	CanDestroy        bool // can delete the discussion.
	CanReopen         bool // can reopen the discussion.
	CanModerateKudosu bool // can allow or deny kudosu.
	CanResolve        bool // can resolve the discussion.
	VoteScore         int  // current vote given to the discussion.
}

type ForumPost struct {
	CreatedAt  Timestamp
	DeletedAt  Timestamp
	EditedAt   Timestamp
	EditedByID int
	ForumID    int
	ID         int
	TopicID    int
	UserID     int

	// Optional fields

	BodyHTML string // post content in HTML format.
	BodyRaw  string // post content in BBCode format.
}

type ForumTopic struct {
	CreatedAt   Timestamp
	DeletedAt   Timestamp
	FirstPostID int
	ForumID     int
	ID          int
	IsLocked    bool
	LastPostID  int
	PostCount   int
	Title       string
	TopicType   string    // normal, sticky, or announcement
	EditedAt    Timestamp // normally called "updatedAt"
	UserID      int
}
