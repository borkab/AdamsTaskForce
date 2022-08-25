package wapp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	h := &Wapp{}

	serv := httptest.NewServer(h)
	defer serv.Close()

	t.Run("foo n point test", func(t *testing.T) {

		requestURL := serv.URL + "/foo"
		requestBody := strings.NewReader("borkab a legmenobb")
		req, err := http.NewRequest(http.MethodPost, requestURL, requestBody)
		if err != nil {
			t.Fatal(err)
		}

		http.DefaultClient.Do(req)
	})
}
