package handlers

import (
	"compress/gzip"
	"github.com/go-openapi/runtime"
	"io"
	"log"
	"net/http"
	"strings"
)

func CustomHeaders(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Server", "pseudoservice")
		w.Header().Add(runtime.HeaderContentType, "application/json")

		if strings.Contains(r.URL.RawQuery, "seed") {
			//deterministic calls can be cached
			w.Header().Add("Expires", "36000")
		} else {
			w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Add("Pragma", "no-cache")
			w.Header().Add("Expires", "0")
		}

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
