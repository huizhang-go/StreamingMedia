package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSc int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSc: 400,
		Error:Err{
			Error:"Request body is not correct",
			ErrorCode: "001",
		},
	}

	ErrorNotAuthUser = ErrorResponse{
		HttpSc: 401,
		Error: Err{
			Error:"User authentication failed.",
			ErrorCode: "002",
		},
	}
)