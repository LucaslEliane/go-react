package render

import (
	"html/template"
	"net/http"
	"github.com/robertkrimen/otto"
)

var templates = template.Must(template.ParseFiles("../../src/public/index.html"))

type IndexPage struct {
	HTML			template.HTML
}

type handler func (http.ResponseWriter, *http.Request) error

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Init() {
	http.Handle("/", handler(handleIndex))
}

func handleIndex(w http.ResponseWriter, r *http.Request) error {
	var renderHTML string

	renderHTML, err := render()

	if err != nil {
		return err
	}

	return templates.ExecuteTemplate(w, "../../src/public/index.html", IndexPage {
		HTML: template.HTML(renderHTML),
	})
}

func render() (string, error) {
	var renderResult otto.Value

	renderResult, err := callRenderJS()

	if err != nil {
		return "", err
	}

	var renderHTML otto.Value

	renderHTML, err = renderResult.Object().Get("html")

	if err != nil {
		return "", err
	}

	return renderHTML.String(), err
}

func callRenderJS()(otto.Value, error) {
	vm := otto.New()

	var v, renderJS otto.Value

	script, err := vm.Compile("../../src/App.js", nil)
	if err != nil {
		return v, err
	}

	v, err = vm.Run(script)

	if err != nil {
		return v, err
	}

	v, err = vm.Get("server")

	if err != nil {
		return v, err
	}

	v, err = v.Object().Get("render")

	if err != nil {
		return v, err
	}

	return renderJS.Call(otto.NullValue())
}