package handlers

type ErrResponse struct {
	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func ErrInvalidRequest(err error) *ErrResponse {
	return &ErrResponse{
		StatusText: "invalid request",
		ErrorText:  err.Error(),
	}
}

func ErrNotFound(err error) *ErrResponse {
	return &ErrResponse{
		StatusText: "not found",
		ErrorText:  err.Error(),
	}
}

func ErrRender(err error) *ErrResponse {
	return &ErrResponse{
		StatusText: "error rendering response",
		ErrorText:  err.Error(),
	}
}

func ErrInternalServer(err error) *ErrResponse {
	return &ErrResponse{
		StatusText: "internal server error",
		ErrorText:  err.Error(),
	}
}
