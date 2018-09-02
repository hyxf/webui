package server

import (
	"fmt"
	_ "github.com/hyxf/webui/statik"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

func RunServer(host string, port string) error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	log.Printf("http://%s:%s/", host, port)
	runError := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	return runError
}
