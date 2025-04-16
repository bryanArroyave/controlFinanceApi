package utils

type SuccessResponse struct {
	Code     string `json:"code"`
	Data     any    `json:"data"`
	Metadata any    `json:"metadata,omitempty"`
}

func BuildSuccessResponse(data, metadata any) *SuccessResponse {
	return &SuccessResponse{
		Code:     "OK",
		Data:     data,
		Metadata: metadata,
	}
}
