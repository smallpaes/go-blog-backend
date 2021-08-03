package errcode

var (
	Success                   = NewError(0, "Success")
	ServerError               = NewError(10000000, "Internal Server Error")
	InvalidParams             = NewError(10000001, "Invalid Parameter Provided")
	NotFound                  = NewError(10000002, "Not Found")
	UnauthorizedAuthNotExist  = NewError(10000003, "Auth Failed, cannot find AppKey and AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "Auth Failed, token error")
	UnauthorizedTokenTimeout  = NewError(10000005, "Auth Failed, token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "Auth Failed, token generation failed")
	TooManyRequest            = NewError(10000007, "Excessive Request")
)
