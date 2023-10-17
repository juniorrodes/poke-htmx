package router

import (
	"net/http"
)

type router struct {
	mux http.ServeMux
}

type HandleFunc func(r *http.Request) Response

func NewRouter() *router {
	return &router{
		mux: *http.NewServeMux(),
	}
}

func (r *router) GET(path string, handle HandleFunc) {
	r.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
		}

		response := handle(r)

		w.WriteHeader(response.StatusCode)
		w.Write(response.Body)
	})
}

func (r *router) POST(path string, handle HandleFunc) {
	r.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
		}

		response := handle(r)

		w.WriteHeader(response.StatusCode)
		w.Write(response.Body)
	})
}

func (r *router) ListenAndServe(addr string) {
	fs := http.FileServer(http.Dir("static"))
	r.mux.Handle("/", http.StripPrefix("/", fs))

	server := http.Server{
		Handler: &r.mux,
		Addr:    addr,
	}

	server.ListenAndServe()
}
