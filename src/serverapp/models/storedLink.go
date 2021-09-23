package models

type StoredLink struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
}

func NewStoredLink(name string, url string) StoredLink {
	return StoredLink {
		Name: name,
		Url: url,
	}
}