package middleware

import "net/http"

type MiddleWares func(http.Handler) http.Handler

/* MiddleWares which acts as a function which takes a Handler as Input and Return a Handler*/

type Manager struct {
	globalMiddleWares []MiddleWares
}

func NewManager() *Manager { // NewManager return pointer of an Object of Middlewares
	return &Manager{
		globalMiddleWares: make([]MiddleWares, 0), // make globalMiddleWares Empty
	}
}

func (mngr *Manager) Use(middleWares ...MiddleWares) {
	mngr.globalMiddleWares = append(mngr.globalMiddleWares, middleWares...) // Every MiddlerWare Store in Global MiddlerWares
}

func (mngr *Manager) With(next http.Handler, middlewares ...MiddleWares) http.Handler {

	n := next
	for _, middleware := range middlewares {
		n = middleware(n)
	}
	return n
}
func (mngr *Manager) WrapMux(next http.Handler) http.Handler {

	n := next

	for _, middleWare := range mngr.globalMiddleWares {
		n = middleWare(n)
	}
	return n

}
