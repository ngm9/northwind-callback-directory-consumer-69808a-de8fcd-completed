package response

type BodyStruct struct {
	StatusCode int     `json:"statucode"`
	Message    *string `json:"message,omitempty"`
	Err        *string `json:"err,omitempty"`
}

func BuildResponseBody(statusCode int, message *string, err error) *BodyStruct {
	var errStr *string
	if err != nil {
		s := err.Error()
		errStr = &s
	}
	return &BodyStruct{
		StatusCode: statusCode,
		Message:    message,
		Err:        errStr,
	}
}
