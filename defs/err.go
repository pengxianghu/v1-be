package defs

type Err struct {
	ErrorCode int
	Error     Result
}

var (
	// 解析错误.
	ErrorRequestBodyParseFailed = Err{
		ErrorCode: 400,
		Error: Result{
			Code: 1,
			Msg:  "Request body is not correct.",
			Data: nil,
		},
	}
	// 用户不存在.
	ErrorNotAuthUser = Err{
		ErrorCode: 401,
		Error: Result{
			Code: 2,
			Msg:  "user authentication failed.",
			Data: nil,
		},
	}

	// dberror status. code: 500
	ErroeDBError = Err{
		ErrorCode: 500,
		Error: Result{
			Code: 3,
			Msg:  "db ops failed.",
			Data: nil,
		},
	}

	// internal faults. status code: 500
	ErrorInternalFaults = Err{
		ErrorCode: 500,
		Error: Result{
			Code: 4,
			Msg:  "Internal service error",
			Data: nil,
		},
	}

	// status code: 404
	ErrorNotFound = Err{
		ErrorCode: 404,
		Error: Result{
			Code: 5,
			Msg:  "Not Found",
			Data: nil,
		},
	}
)
