package dtos

type BaseResponse struct {
	Errors []interface{} `json:"errors,omitempty"`
}
