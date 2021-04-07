package gosu

import "time"

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
	CreatedAt time.Time
	ID        int
	EventType EventType
}

type EventBeatmap struct {
	Event
	Title string
	URL   string
}

type EventBeatmapset struct {
	Event
	Title string
	URL   string
}

type EventUser struct {
	Event
	Username         string
	URL              string
	PreviousUsername string // only for EventUsernameChange event
}

// When user obtained an achievement.
type EventAchievement struct {
	Achievement Achievement
	User        EventUser
}

// When a beatmap has been played for certain number of times.
type EventBeatmapPlaycount struct {
	Beatmap EventBeatmap
	Count   int
}

// When a beatmapset changes state.
type EventBeatmapsetApprove struct {
	Approval   RankStatus
	Beatmapset EventBeatmapset
	User       EventUser // Beatmapset owner
}

// When a beatmapset is deleted.
type EventBeatmapsetDelete struct {
	Beatmapset EventBeatmapset
}

// When a beatmapset in graveyard state is updated.
type EventBeatmapsetRevive struct {
	Beatmapset EventBeatmapset
	User       EventUser // Beatmapset owner
}

// When a beatmapset is updated.
type EventBeatmapsetUpdate struct {
	Beatmapset EventBeatmapset
	User       EventUser // Beatmapset owner
}

// When a new beatmapset is uploaded.
type EventBeatmapsetUpload struct {
	Beatmapset EventBeatmap
	User       EventUser // Beatmapset owner
}

// When a user achieves a certain rank on a beatmap.
type EventRank struct {
	ScoreRank string
	Rank      int
	Mode      GameMode
	Beatmap   EventBeatmap
	User      EventUser
}

// When a user loses first place to another user.
type EventRankLost struct {
	Mode    GameMode
	Beatmap EventBeatmap
	User    EventUser
}

// When a user supports osu! for the second and onwards.
type EventUserSupportAgain struct {
	User EventUser
}

// When a user becomes a supporter for the first time.
type EventUserSupportFirst struct {
	User EventUser
}

// When a user is gifted a supporter tag by another user.
type EventUserSupportGift struct {
	User EventUser
}

// When a user changes their username.
type EventUsernameChange struct {
	User EventUser // Includes previousUsername
}

type Achievement struct {
	// TODO: implement once specified in docs
}

type Notification struct {
	ID           int
	Name         string //Name of the event
	CreatedAt    string //ISO 8601 date
	ObjectType   string
	ObjectID     int
	SourceUserID int
	IsRead       bool
	Details      string // message_id of last known message (only returned in presence responses)
}
