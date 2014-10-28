package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	addr := flag.String("a", ":5000", "address to serve(host:port)")
	prefix := flag.String("p", "/", "prefix path under")
	root := flag.String("r", ".", "root path to serve")
	flag.Parse()

	var err error
	*root, err = filepath.Abs(*root)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("serving %s as %s on %s", *root, *prefix, *addr)
	http.Handle(*prefix, http.StripPrefix(*prefix, http.FileServer(http.Dir(*root))))

	logger := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		http.DefaultServeMux.ServeHTTP(w, r)
	})
	err = http.ListenAndServe(*addr, logger)
	if err != nil {
		log.Fatalln(err)
	}
}
