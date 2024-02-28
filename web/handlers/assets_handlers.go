package handlers

import "net/http"

func AssetsHandler(dir string, prefix string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))
}
