package models

// Entry is a blog post entry
type Entry struct {
	Body             string
	Meta             map[string]string
	Plugins          map[string]interface{}
	URL              string
	HomeURL          string
	OriginalFilename string
}
