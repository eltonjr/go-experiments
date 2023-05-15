package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/eltonjr/go-experiments/interface-adapter/internal/item"
)

func (s *server) GetItem(w http.ResponseWriter, req *http.Request) {
	i, err := s.itemModule.Get(req.Context(), parseID(req.URL.Path))
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, item.ErrItemNotFound) {
			status = http.StatusNotFound
		}

		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(i.JSON()))
}

func (s *server) SetItem(w http.ResponseWriter, req *http.Request) {
	var i item.Item
	err := json.NewDecoder(req.Body).Decode(&i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = s.itemModule.Set(req.Context(), &i)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, item.ErrMissingID) {
			status = http.StatusNotFound
		}

		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(i.JSON()))
}

func parseID(url string) int {
	idraw := strings.TrimPrefix(url, pathPrefix)
	id, err := strconv.Atoi(idraw)
	if err != nil {
		return 0
	}

	return id
}
