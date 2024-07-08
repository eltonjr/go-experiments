package server

import (
	"net/http"

	"github.com/eltonjr/go-experiments/decorator-pattern/internal/item"
)

const pathPrefix = "/items/"

type server struct {
	mux        *http.ServeMux
	itemModule item.Module
}

func NewServer(addr string, m item.Module) error {
	mux := http.NewServeMux()

	s := &server{
		mux:        mux,
		itemModule: m,
	}
	mux.HandleFunc("/items/", s.GetItem)
	mux.HandleFunc("/items", s.SetItem)

	return http.ListenAndServe(addr, mux)
}
