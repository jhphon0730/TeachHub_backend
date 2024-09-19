package context

import (
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
