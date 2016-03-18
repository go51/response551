package response551

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go51/string551"
	"github.com/go51/time551"
	"html/template"
	"net/http"
)

type RedirectType struct {
	uri  string
	code int
	text string
}

func Redirect(uri string, code int) RedirectType {
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

func Response(w http.ResponseWriter, r *http.Request, data interface{}, packageName, routeName string, user interface{}, appConfig interface{}) {
	if redirectType, ok := interface{}(data).(RedirectType); ok {
		// Redirect Type
		http.Redirect(w, r, redirectType.uri, redirectType.code)
		return
	} else if errorType, ok := interface{}(data).(ErrorType); ok {
		// Error Type
		http.Error(w, errorType.message, errorType.code)
		return
	} else if param, ok := interface{}(data).(map[string]interface{}); ok {
		if isJSON(r) {
			jsonOutput(w, param)
		} else {
			// View template rendering
			param["user"] = user
			param["config"] = appConfig
			htmlOutput(w, param, packageName, routeName)
		}
		return
	}

	fmt.Fprintf(w, "%v", data)

}

func isJSON(r *http.Request) bool {
	format := string551.Lower(r.FormValue("format"))

	if format == "json" {
		return true
	} else {
		return false
	}
}

func htmlOutput(w http.ResponseWriter, data interface{}, packageName, routeName string) {

	templates := []string{
		"view/template/base.html",
		"view/" + packageName + "/" + routeName + ".html",
	}

	tmpl, err := template.New(packageName + routeName).Funcs(funcMap()).ParseFiles(templates...)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	buf := &bytes.Buffer{}

	err = tmpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, buf)

}

func jsonOutput(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func funcMap() template.FuncMap {
	funcMap := template.FuncMap{}

	funcMap["raw"] = func(text string) template.HTML {
		return template.HTML(text)
	}
	funcMap["url"] = UrlFunction
	funcMap["rightRune"] = string551.RightRune
	funcMap["right"] = string551.Right
	funcMap["elapsed"] = time551.Elapsed

	return funcMap

}

type urlFunc func(name string, parameter ...string) string

var UrlFunction urlFunc
