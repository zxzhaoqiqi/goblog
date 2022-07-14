package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf8")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>测试</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 ):</h1>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf8")

	fmt.Fprint(w, "about 页面")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":3000", nil)
}
