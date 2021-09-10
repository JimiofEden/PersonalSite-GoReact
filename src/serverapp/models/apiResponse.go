package models

type ApiResponse struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

func NewApiResponse(status string, data interface{}) ApiResponse {
	return ApiResponse {
		Status: status,
		Data: data,
	}
}