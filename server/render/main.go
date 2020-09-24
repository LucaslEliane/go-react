package render

import (
	"html/template"
	"log"
	"net/http"
	"github.com/robertkrimen/otto"
	"google.golang.org/appengine"
	"sync"
)

var templates = template.Must(template.ParseFiles("../../src/public/index.html"))

const jsFile = "../../src/src/server.js"

type IndexPage struct {
	HTML			template.HTML
}

type handler func (http.ResponseWriter, *http.Request) error

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var callRenderJS = func() (otto.Value, error) {
	vm := otto.New()
	renderJS, err := compileRenderJS(vm)

	if err != nil {
		return otto.NullValue(), err
	}

	return renderJS.Call(otto.NullValue())
}

func Init() {
	if !appengine.IsDevAppServer() {
		var renderJS otto.Value
		var once sync.Once
		callRenderJS = func() (otto.Value, error) {
			once.Do(func() {
				vm := otto.New()
				var err error
				renderJS, err = compileRenderJS(vm)
				if err != nil {
					log.Fatal(err)
				}
			})

			return renderJS.Call(otto.NullValue())
		}
	}
	http.Handle("/", handler(handleIndex))
}

func compileRenderJS(vm *otto.Otto) (otto.Value, error) {
	var v otto.Value
	script, err := vm.Compile(jsFile, nil)
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

	return v, nil
}

func handleIndex(w http.ResponseWriter, r *http.Request) error {
	var renderHTML string

	if len(r.Header["X-Devserver"]) > 0 {
		return templates.ExecuteTemplate(w, "../../src/public/index.html", IndexPage {
			HTML: "",
		})
	}

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