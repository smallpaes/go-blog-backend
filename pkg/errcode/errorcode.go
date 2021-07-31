package errcode

import "fmt"

// Define error response template
type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

// Store codes globally
var codes = make(map[int]string)

func NewError(code int, msg string) *Error {
	// check if code inputted already existed in global
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("Error: Code %d already existed", code))
	}
	// add code to global
	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error Code: %d, Error Message: %s", e.Code(), e.Msg())
}

func (e *Error) Msgf(args ...[]interface{}) string {
	return fmt.Sprintf(e.msg, args)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

// func (e *Error) StatusCode() int {
// 	switch e.Code() {
//   case
// 	}
// }
