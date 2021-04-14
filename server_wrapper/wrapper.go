package server_wrapper

import "net/http"

type Middleware interface {
	Before(writer http.ResponseWriter, request *http.Request)
	After(writer http.ResponseWriter, request *http.Request)
}

type ServerWrapper struct {
	server      http.Handler
	middlewares []Middleware
}

func New(s http.Handler) ServerWrapper {
	return ServerWrapper{s, nil}
}

func (s *ServerWrapper) AddMiddleware(m Middleware) {
	s.middlewares = append(s.middlewares, m)
}

func (s ServerWrapper) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, m := range s.middlewares {
		m.Before(writer, request)
	}
	s.server.ServeHTTP(writer, request)
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		s.middlewares[i].After(writer, request)
	}
}
