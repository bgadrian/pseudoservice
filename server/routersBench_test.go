package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func BenchmarkGenerateUsersHTTPHandler100RandSeed(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		request := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
		request = mux.SetURLVars(request, map[string]string{"count": "100"})
		w := httptest.NewRecorder()

		b.StartTimer()
		UsersCountGet(w, request)
		b.StopTimer()

		if w.Code != http.StatusOK {
			b.Error("wrong response status code")
		}
	}
}

func BenchmarkGenerateUsersHTTPHandler100FixedSeed(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		request := httptest.NewRequest(http.MethodGet, "http://dummy?seed=42", nil)
		request = mux.SetURLVars(request, map[string]string{"count": "100"})
		w := httptest.NewRecorder()

		b.StartTimer()
		UsersCountGet(w, request)
		b.StopTimer()

		if w.Code != http.StatusOK {
			b.Error("wrong response status code")
		}
	}
}

func BenchmarkGenerateUsersHTTPHandler300RandSeed(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		request := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
		request = mux.SetURLVars(request, map[string]string{"count": "300"})
		w := httptest.NewRecorder()

		b.StartTimer()
		UsersCountGet(w, request)
		b.StopTimer()

		if w.Code != http.StatusOK {
			b.Error("wrong response status code")
		}
	}
}

func BenchmarkGenerateUsersHTTPHandler300FixedSeed(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		request := httptest.NewRequest(http.MethodGet, "http://dummy?seed=42", nil)
		request = mux.SetURLVars(request, map[string]string{"count": "300"})
		w := httptest.NewRecorder()

		b.StartTimer()
		UsersCountGet(w, request)
		b.StopTimer()

		if w.Code != http.StatusOK {
			b.Error("wrong response status code")
		}
	}
}
