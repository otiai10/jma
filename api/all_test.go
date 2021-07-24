package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

var s *httptest.Server

func TestMain(m *testing.M) {
	s = createTestServer()
	code := m.Run()
	os.Exit(code)
}

func TestForecast(t *testing.T) {
	c := &Client{BaseURL: s.URL}
	res, err := c.Forecast("130000")
	Expect(t, err).ToBe(nil)
	Expect(t, len(res)).ToBe(2)
}

func createTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("." + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if _, err := io.Copy(w, f); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	return httptest.NewServer(mux)
}
