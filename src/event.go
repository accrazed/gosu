package gosu

type EventType int

const (
	TypeAchievement EventType = iota
	TypeBeatmapPlaycount
	TypeBeatmapsetApprove
	TypeBeatmapsetDelete
	TypeBeatmapsetRevive
	TypeBeatmapsetUpdate
	TypeBeatmapsetUpload
	TypeRank
	TypeRankLost
	TypeUserSupportAgain
	TypeUserSupportFirst
	TypeUserSupportGift
	TypeUsernameChange
)

type Event struct {
	createdAt Timestamp
	id        int
	eventType EventType
}

type EventBeatmap struct {
	Event
	title string
	url   string
}

type EventBeatmapset struct {
	Event
	title string
	url   string
}

type EventUser struct {
	Event
	username         string
	url              string
	previousUsername string // only for EventUsernameChange event
}

// When user obtained an achievement.
type EventAchievement struct {
	achievement Achievement
	user        EventUser
}

// When a beatmap has been played for certain number of times.
type EventBeatmapPlaycount struct {
	beatmap EventBeatmap
	count   int
}

// When a beatmapset changes state.
type EventBeatmapsetApprove struct {
	approval   RankStatus
	beatmapset EventBeatmapset
	user       EventUser // Beatmapset owner
}

// When a beatmapset is deleted.
type EventBeatmapsetDelete struct {
	beatmapset EventBeatmapset
}

// When a beatmapset in graveyard state is updated.
type EventBeatmapsetRevive struct {
	beatmapset EventBeatmapset
	user       EventUser // Beatmapset owner
}

// When a beatmapset is updated.
type EventBeatmapsetUpdate struct {
	beatmapset EventBeatmapset
	user       EventUser // Beatmapset owner
}

// When a new beatmapset is uploaded.
type EventBeatmapsetUpload struct {
	beatmapset EventBeatmap
	user       EventUser // Beatmapset owner
}

// When a user achieves a certain rank on a beatmap.
type EventRank struct {
	scoreRank string
	rank      int
	mode      GameMode
	beatmap   EventBeatmap
	user      EventUser
}

// When a user loses first place to another user.
type EventRankLost struct {
	mode    GameMode
	beatmap EventBeatmap
	user    EventUser
}

// When a user supports osu! for the second and onwards.
type EventUserSupportAgain struct {
	user EventUser
}

// When a user becomes a supporter for the first time.
type EventUserSupportFirst struct {
	user EventUser
}

// When a user is gifted a supporter tag by another user.
type EventUserSupportGift struct {
	user EventUser
}

// When a user changes their username.
type EventUsernameChange struct {
	user EventUser // Includes previousUsername
}

type Achievement struct {
	// TODO: implement once specified in docs
}
