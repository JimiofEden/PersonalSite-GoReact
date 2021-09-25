package models

type StoredLink struct {
	Id int `json:"id"`
	LinkName string `json:"linkName"`
	Url string `json:"url"`
}

func NewStoredLink(linkName string, url string) StoredLink {
	return StoredLink {
		LinkName: linkName,
		Url: url,
	}
}