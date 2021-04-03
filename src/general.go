package gosu

// Note that sort option should be specified for cursor to work
type Cursor struct {
	Next string // request string to get next set of results
	Prev string // request string to get previous set of results
}

type Timestamp struct {
	Time string
}
