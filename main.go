package main

import (
	"fmt"
	"net/http"
	"strings"
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

	fmt.Fprint(w, "ceshi")
	fmt.Fprint(w, "about 页面")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)

	router.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.SplitN(r.URL.Path, "/", 3)[2]
		fmt.Fprint(w, "文章ID："+id)
	})

	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "访问文章列表")
		case "POST":
			fmt.Fprint(w, "创建新的文章")
		}
	})

	http.ListenAndServe(":3000", router)
}
