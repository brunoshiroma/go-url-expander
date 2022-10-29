package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle200RemoveUTMWithRedirect(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	t.Run("Test handle http without redirect", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("GET", fmt.Sprintf("/expander?url=%s&no-utm=1&redirect=1", svr.URL), nil)
		if err != nil {
			t.Fatal(err)
		}

		handle(recorder, req)

		assert.Equal(t, http.StatusFound, recorder.Code)
		assert.Equal(t, svr.URL, recorder.Header().Get("Location"))
	})
}

func TestHandle200(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	t.Run("Test handle http without redirect", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("GET", fmt.Sprintf("/expander?url=%s", svr.URL), nil)
		if err != nil {
			t.Fatal(err)
		}

		handle(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, svr.URL, recorder.Body.String())
	})
}

func TestHandle302(t *testing.T) {
	const targetUrl = "https://www.google.com"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", targetUrl)
		w.WriteHeader(http.StatusMovedPermanently)
	}))

	t.Run("Test handle http without redirect", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("GET", fmt.Sprintf("/expander?url=%s", svr.URL), nil)
		if err != nil {
			t.Fatal(err)
		}

		handle(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, targetUrl, recorder.Body.String())
	})
}
