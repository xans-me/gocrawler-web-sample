package response

// ResponseStruct model response
type ResponseStruct struct {
	Data   interface{} `json:"result"`
	TimeIn string      `json:"timeIn"`
}

// ErrorStruct model response
type ErrorStruct struct {
	Code    int
	Message string
}

// we implement the built-in package 'error' interface by creating this function
func (e ErrorStruct) Error() string {
	return e.Message
}
