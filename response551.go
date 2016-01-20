package response551

import (
	"fmt"
	"html/template"
	"net/http"
)

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

func Response(w http.ResponseWriter, r *http.Request, data interface{}, packageName, routeName string) {
	if redirectType, ok := interface{}(data).(RedirectType); ok {
		// Redirect Type
		http.Redirect(w, r, redirectType.uri, redirectType.code)
		return
	} else if errorType, ok := interface{}(data).(ErrorType); ok {
		// Error Type
		http.Error(w, errorType.message, errorType.code)
		return
	} else if _, ok := interface{}(data).(map[string]interface{}); ok {
		// View template rendering
		html(w, data, packageName, routeName)
		return
	}

	fmt.Fprintf(w, "%v", data)

}

func html(w http.ResponseWriter, data interface{}, packageName, routeName string) {

	templates := []string{
		"view/template/base.html",
		"view/" + packageName + "/" + routeName + ".html",
	}

	tmpl, err := template.New(packageName + routeName).Funcs(funcMap()).ParseFiles(templates...)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func funcMap() template.FuncMap {
	funcMap := template.FuncMap{}

	funcMap["raw"] = func(text string) template.HTML { return template.HTML(text) }

	return funcMap

}
