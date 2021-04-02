package gosu

// See https://osu.ppy.sh/docs/index.html?javascript#chatchannel
type Channel string

const (
	PUBLIC      = "PUBLIC"
	PRIVATE     = "PRIVATE"
	MULTIPLAYER = "MULTIPLAYER"
	SPECTATOR   = "SPECTATOR"
	PM          = "PM"
	GROUP       = "GROUP"
)

const ()

type BeatmapsetDiscussion struct {
	beatmap               BeatmapCompact
	beatmapID             int
	beatmapset            BeatmapsetCompact
	beatmapsetID          int
	canBeResolved         bool
	canGrandKudosu        bool
	createdAt             string
	currentUserAttributes CurrentUserAttributes
	deletedAt             string // TODO: may become a Timestamp object
	deletedByID           int
	id                    int
	kudosuDenied          bool
	lastPostAt            string // TODO: may become a Timestamp object
	messageType           string // FIXME: originally a "MessageType" object, not in docs
	parentID              int
	posts                 []BeatmapsetDiscussionPost
	resolved              bool
	startingPost          BeatmapsetDiscussionPost
	timestamp             int    // TODO: may become a Timestamp object
	updatedAt             string // TODO: may become a Timestamp object
	userID                int
	votes                 []BeatmapsetDiscussionVote // FIXME: structure will change
}

type BeatmapsetDiscussionPost struct {
	beatmapsetDiscussionID int
	createdAt              string // TODO: may become a Timestamp object
	deletedAt              string // TODO: may become a Timestamp object
	deletedByID            int
	id                     int
	lastEditorID           int
	message                string
	system                 bool
	updatedAt              string // TODO: may become a Timestamp object
	userID                 int
}
type BeatmapsetDiscussionVote struct {
	beatmapsetDiscussionID int
	createdAt              string // TODO: may become a Timestamp object
	id                     int
	score                  int
	updatedAt              string // TODO: may become a Timestamp object
	userID                 int
}

type ChatChannel struct {
	channelID      int
	name           string
	description    string
	icon           string        // display icon for the channel
	channelType    Channel       // see channel types below
	firstMessageID int           // messageID of first message (only returned in presence responses)
	lastReadID     int           // messageID of last message read (only returned in presence responses)
	lastMessageID  int           // messageID of last known message (only returned in presence responses)
	recentMessages []ChatMessage // up to 50 most recent messages
	moderated      bool          // user can't send message when the value is true (only returned in presence responses)
	users          []int         // array of userID that are in the channel (not included for PUBLIC channels)
}

type ChatMessage struct {
	messageID int         // unique identifier for message
	senderID  int         // userID of the sender
	channelID int         // channelID of where the message was sent
	timestamp string      // when the message was sent, ISO-8601
	content   string      // message content
	isAction  bool        // was this an action i.e. /me dances
	sender    UserCompact // embeded UserCompact object to save additional api lookups
}
type Comment struct {
	commentableID   int    // ID of the object the comment is attached to
	commentableType string // type of object the comment is attached to
	createdAt       string // ISO 8601 date
	deletedAt       string // ISO 8601 date if the comment was deleted; null, otherwise
	editedAt        string // ISO 8601 date if the comment was edited; null, otherwise
	editedByID      int    // user id of the user that edited the post; null, otherwise
	id              int    // the ID of the comment
	legacyName      string // username displayed on legacy comments
	message         string // markdown of the comment's content
	messageHTML     string // html version of the comment's content
	parentID        int    // ID of the comment's parent
	pinned          bool   // Pin status of the comment
	repliesCount    int    // int of replies to the comment
	updatedAt       string // ISO 8601 date
	userID          int    // user ID of the poster
	votesCount      int    // int of votes
}
type CommentBundle struct {
	commentableMeta  CommentableMeta // ID of the object the comment is attached to
	comments         []Comment       // array of comments ordered according to sort;
	cursor           Cursor          // no desc
	hasMore          bool            // If there are more comments or replies available
	hasMoreID        int             // no desc
	includedComments []Comment       // related comments; e.g. parent comments and nested replies
	pinnedComments   []Comment       // pinned comments
	sort             string          // one of the CommentSort types
	topLevelCount    int             // num of comments at the top level. Not returned for replies.
	total            int             // total num of comments. Not retuned for replies.
	userFollow       bool            // is the current user watching the comment thread
	userVotes        []int           // IDs of the comments in the bundle the current user has upvoted
	users            []UserCompact   // array of users related to the comments
}

// TODO: IMPLEMENT TYPE "CommentSort" once I know what it actually is

type CommentableMeta struct {
	id         int    // the ID of the object
	title      string // display title
	objectType string // the type of the object
	url        string // url of the object
}

type CurrentUserAttributes struct {
	canDestroy        bool // can delete the discussion.
	canReopen         bool // can reopen the discussion.
	canModerateKudosu bool // can allow or deny kudosu.
	canResolve        bool // can resolve the discussion.
	voteScore         int  // current vote given to the discussion.
}
