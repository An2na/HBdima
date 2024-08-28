package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	//Success string `json:"success"`
}

const (
	StatusOK      = "OK"
	StatusError   = "Error"
	StatusSuccess = "Success"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}
func Success(msg string) Response {
	return Response{
		Status: StatusSuccess,
		//Success: msg,
	}
}
