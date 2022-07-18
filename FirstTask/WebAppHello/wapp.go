package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

/*
curl localhost:8080
OR
curl -X GET http://localhost:8080/
curl: (7) Failed to connect to localhost port 8080 after 0 ms: Connection refused
*/
//BUT:
/* go run wapp.go
And navigate to http://localhost:8080 in the browser
it gives: Hello, world!
*/
