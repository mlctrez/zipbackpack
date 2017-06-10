package main

import (
	"github.com/mlctrez/zipbackpack/httpfs"
	"net/http"
)

func main() {

	sf, err := httpfs.NewStaticFileSystem("")

	if err != nil {
		panic(err)
	}

	handler := http.FileServer(sf)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// http.FileServer does a permanent redirect of /index.html to /
		// so for this path we serve /_index.html to avoid a redirect loop
		if r.URL.Path == "/" {
			r.URL.Path = "/_index.html"
		}
		handler.ServeHTTP(rw, r)
	})
	http.HandleFunc("/api/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("api handler for " + r.URL.Path))
	})

	http.ListenAndServe(":8888", nil)

}
