package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.String("port", "8888", "HTTP port")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello!")
	})

	httpSrv := http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}

	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err)
	}
}
