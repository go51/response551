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
