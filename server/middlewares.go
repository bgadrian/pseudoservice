package server

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
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

func CustomHeaders(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Server", "pseudoservice")
		w.Header().Add("Content-Type", "application/json")

		//TODO make cache policies for each endpoint
		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Add("Pragma", "no-cache")
		w.Header().Add("Expires", "0")

		inner.ServeHTTP(w, r)
	})
}

//write in gzip and Header() from http
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Gzip(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		algSupp := r.Header.Get("Accept-Encoding")
		supportGzip := strings.Contains(algSupp, "gzip")

		if supportGzip {
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)

			defer func() {
				err := gz.Close()
				if err != nil {
					log.Printf("error closing gzip: %+v\n", err)
				}
			}()

			gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
			inner.ServeHTTP(gzr, r)
			return
		}

		//fallback
		inner.ServeHTTP(w, r)

	})
}
