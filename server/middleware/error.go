package middleware

import (
	"encoding/json"
	"net/http"
)

type xerr struct {
	Message  string `json:"message"`
	Response string `json:"response"`
}

type ErrHandler func(http.ResponseWriter, *http.Request) error

func (fn ErrHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			xerr := xerr{
				Message:  "Panic",
				Response: "Error",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
	}()
	if err := fn(w, r); err != nil {
		xerr := xerr{
			Message:  err.Error(),
			Response: "Error",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(xerr)
	}

}
