package models

type ApiData struct {
	Message string `json:"message"`
}

func NewApiData(message string) ApiData {
	return ApiData{Message: message}
}