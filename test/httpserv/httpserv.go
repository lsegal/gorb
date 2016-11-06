package httpserv

import (
	"fmt"
	"net/http"
)

func Serve(addr string, handler func(string) string) {
	http.HandleFunc("/gorb", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", handler(r.URL.String()))
	})
	http.ListenAndServe(addr, nil)
}
