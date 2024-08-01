package server

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tmplPath = "./html/"

func NewHtmlRouter(w http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/api/") {
		err := errors.New("not found api handler")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// https://github.com/gorilla/mux?tab=readme-ov-file#serving-single-page-applications
	// can it possible ../../etc/hosts ?
	path := filepath.Join(tmplPath, req.URL.Path)
	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if (os.IsNotExist(err) || fi.IsDir()) && strings.HasPrefix(req.URL.Path, "/assets/") {
		err := errors.New("file in assests not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if os.IsNotExist(err) || fi.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, req, filepath.Join(tmplPath, "index.html"))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if strings.HasSuffix(req.URL.Path, "/api") {
		err := errors.New("not found api handler")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(tmplPath)).ServeHTTP(w, req)
}
