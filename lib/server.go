package lib

import (
	"fmt"
	"log"
	"net/http"
)

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprintf(w, getStaticFileForHost(req.Host))
}

func Serve(port int) {
	var r router
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), &r)
	if err != nil {
		log.Printf("Failed to listen and serve :(: %s\n", err)
		return
	}
}