package errcode

var (
	ErrorCreateTagFail = NewError(20010002, "Create Tag Failed")
	ErrorUpdateTagFail = NewError(20010003, "Update Tag Failed")
)
