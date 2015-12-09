package main

import (
	"net/http"
	"os"
	"path/filepath"
)

const defaultPort = "3000"

func symlinkRedirectWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Isolate relative path; strip trailing slash to prevent redirect
		// loops.
		p := filepath.Clean(r.URL.Path[1:])
		// Resolve symlinks.
		realPath, err := filepath.EvalSymlinks(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Redirect to resolved path.
		if p != realPath {
			http.Redirect(w, r, "/"+realPath, http.StatusFound)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", symlinkRedirectWrapper(fs))
	http.ListenAndServe(":"+port, nil)
}
