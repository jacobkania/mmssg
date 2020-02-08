package models

// Entry is a blog post entry
type Entry struct {
	Body             string
	Meta             map[string]string
	URL              string
	HomeURL          string
	OriginalFilename string
}
