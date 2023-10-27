package errors

import "errors"

var (
	Internal      = errors.New("internal error")
	DescrNotFound = errors.New("decription not found")
	ArtNotFound   = errors.New("article not found")
	NoArt         = errors.New("not new article for post")
	NewDoc        = errors.New("create new doc")
	LinkNoFound   = errors.New("link not found")
	TitleNotFound = errors.New("not found tittle")
)
