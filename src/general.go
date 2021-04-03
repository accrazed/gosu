package gosu

// Note that sort option should be specified for cursor to work
type Cursor struct {
	Next string // request string to get next set of results
	Prev string // request string to get previous set of results
}

type Timestamp struct {
	Time string
}

type WikiPage struct {
	Layout   string   // The layout type for the page.
	Locale   string   // All lowercase BCP 47 language tag.
	Markdown string   // Markdown content.
	Path     string   // Path of the article.
	Subtitle string   // 	The article's subtitle.
	Tags     []string // Associated tags for the article.
	Title    string   // The article's title.
}
