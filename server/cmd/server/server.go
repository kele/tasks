package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kele/tasks/server/tasks"
	"github.com/pacedotdev/oto/otohttp"
)

func main() {
	// Setting up reverse proxy for the frontend.
	localhost, _ := url.Parse("http://localhost:4200")
	proxy := httputil.NewSingleHostReverseProxy(localhost)
	http.Handle("/", proxy)

	t := &tasks.ServiceImpl{}
	server := otohttp.NewServer()
	tasks.RegisterService(server, t)
	http.Handle("/oto/", server)
	fmt.Printf("Listening on :8080...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
