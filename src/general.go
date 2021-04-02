package gosu

// Note that sort option should be specified for cursor to work
type Cursor struct {
	next string // request string to get next set of results
	prev string // request string to get previous set of results
}

type Timestamp struct {
	time string
}
