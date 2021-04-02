package gosu

type RankStatus int

const (
	Graveyard RankStatus =  iota - 2
	WIP
	Pending
	Ranked
	Approved
	Qualified
	Loved
)
