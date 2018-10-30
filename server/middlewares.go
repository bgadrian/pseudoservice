package server

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

type ApiKeyHandler struct {
	keys map[string]struct{}
}

func NewApiKeyHandler(keys []string) *ApiKeyHandler {
	result := &ApiKeyHandler{}
	result.keys = make(map[string]struct{})
	for _, key := range keys {
		result.keys[key] = struct{}{}
	}
	return result
}

func (h ApiKeyHandler) Handle(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		token := values.Get("token")
		if _, ok := h.keys[token]; ok {
			inner.ServeHTTP(w, r)
			return
		}

		log.Printf(
			"access denied: wrong api key for %s %s",
			r.Method,
			r.RequestURI,
		)

		w.WriteHeader(http.StatusBadRequest)
		written, err := w.Write(NewErrorJson(ERROR_APIKEY, "wrong api key"))
		if written == 0 || err != nil {
			log.Printf("error while writing the error %s", err)
		}

	})
}
