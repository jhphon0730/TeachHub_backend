package context

import (
	"html/template"
	"path/filepath"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Data    map[string]interface{}
	Flash   *Flash
}

type Flash struct {
	SuccessMsg string
	ErrorMsg   string
}

type AppHandler struct {
	HandleFunc func(*Context)
}

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Writer:  w,
		Request: r,
		Data:    make(map[string]interface{}),
		Flash:   &Flash{},
	}
	h.HandleFunc(ctx)
}

func (c *Context) Success(templateName string) {
	templatePath := filepath.Join("templates", templateName)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(c.Writer, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(c.Writer, c.Data)
	if err != nil {
		http.Error(c.Writer, "Error executing template", http.StatusInternalServerError)
		return
	}
}
