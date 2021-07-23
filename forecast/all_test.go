package forecast

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/otiai10/jma"
	"github.com/otiai10/marmoset"
	. "github.com/otiai10/mint"
)

func TestForecastByCity(t *testing.T) {
	r := marmoset.NewRouter()
	r.GET("/data/forecast/130000.json", func(w http.ResponseWriter, r *http.Request) {
		if f, err := os.Open("examples/130000.json"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			io.Copy(w, f)
		}
	})
	s := httptest.NewServer(r)
	defer s.Close()

	c := NewClient()
	c.BaseURL = s.URL

	res, err := c.Forecast(jma.Tokyo)
	Expect(t, err).ToBe(nil)
	Expect(t, len(res)).ToBe(2)
}
