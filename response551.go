package response551

import "net/http"

type RedirectType struct {
	uri  string
	code int
	text string
}

func Redirect(uri string, code int) RedirectType {
	if code != 301 && code != 302 {
		code = 302
	}

	return RedirectType{
		uri:  uri,
		code: code,
		text: http.StatusText(code),
	}
}

func (r RedirectType) Code() int {
	return r.code
}

func (r RedirectType) Text() string {
	return r.text
}

func (r RedirectType) Uri() string {
	return r.uri
}

type ErrorType struct {
	code    int
	text    string
	message string
}

func Error(code int, message string) ErrorType {
	return ErrorType{
		code:    code,
		text:    http.StatusText(code),
		message: message,
	}
}

func (e ErrorType) Code() int {
	return e.code
}

func (e ErrorType) Text() string {
	return e.text
}

func (e ErrorType) Message() string {
	return e.message
}

func (e ErrorType) String() string {
	return e.message
}
