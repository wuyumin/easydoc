package easydoc

// PostSource defines post source struct.
type PostSource struct {
	Id      int
	Title   string
	AbsPath string
	UrlPath string
}

var postSourceById map[int]*PostSource

func store(postSource PostSource) {
	postSourceById[postSource.Id] = &postSource
}
