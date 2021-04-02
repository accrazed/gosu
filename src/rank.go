package gosu

type RankStatus int

const (
	Graveyard = iota - 2
	WIP
	Pending
	Ranked
	Approved
	Qualified
	Loved
)
