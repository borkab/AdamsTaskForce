package wapp

import (
	"fmt"
	"io"
	"net/http"
)

type Wapp struct {
}

func (wa *Wapp) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	bs, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	fmt.Println(string(bs))
}
