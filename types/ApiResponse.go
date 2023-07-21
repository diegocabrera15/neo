package types

type ApiResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
	//Errors          ErrorResponse     `json:"errors"`
}

type ErrorResponse struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Timestamp    string `json:"timestamp"`
}
