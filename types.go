package lambda_common

type ApiError struct {
	Id         int
	Body       ApiErrorBody
	Err        error
	StatusCode int
}

type ApiErrorBody struct {
	Id         int
	Message    string
	StatusCode int
}

type ApiSuccess struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}
