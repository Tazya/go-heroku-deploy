package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello!")
	})

	httpSrv := http.Server{
		Addr:    ":8888",
		Handler: mux,
	}

	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err)
	}
}
