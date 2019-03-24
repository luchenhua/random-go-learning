package entity

// Result represents the response of the HTTP request.
type Result struct {
	Data    string `json:"data"`
	Message string `json:"msg"`
}
