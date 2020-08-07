package helloworld

import (
	"fmt"
	"net/http"
)

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello, World!7")
}

func Verify(r *http.Request) error {
	return nil
}
