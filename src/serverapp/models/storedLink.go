package models

type StoredLink struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

func NewStoredLink(name string, url string) StoredLink {
	return StoredLink {
		Name: name,
		Url: url,
	}
}