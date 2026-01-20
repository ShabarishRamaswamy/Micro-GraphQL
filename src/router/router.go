package router

import (
	"fmt"
	"io"
	"net/http"
)

func GetNewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.HandleFunc("/sendQuery", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			type byteArr []byte

			body, err := io.ReadAll(r.Body)
			fmt.Printf("%+v, %+v", string(body), err)
		}
	})

	return router
}
