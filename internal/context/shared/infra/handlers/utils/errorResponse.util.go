package utils

type ErrorResponse struct {
	Code     string `json:"code"`
	Data     any    `json:"data"`
	Metadata any    `json:"metadata,omitempty"`
}

func BuildErrorResponse(code string, err error, metadata any) *ErrorResponse {
	return &ErrorResponse{
		Code:     code,
		Data:     err.Error(),
		Metadata: metadata,
	}
}
