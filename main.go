package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

const name = "serve"

const version = "0.0.4"

var revision = "HEAD"

func main() {
	addr := flag.String("a", ":5000", "address to serve(host:port)")
	prefix := flag.String("p", "/", "prefix path under")
	root := flag.String("r", ".", "root path to serve")
	certFile := flag.String("cf", "", "tls cert file")
	keyFile := flag.String("kf", "", "tls key file")
	dumpPost := flag.Bool("dumpPost", false, "dump post data")
	showVersion := flag.Bool("v", false, "show version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return
	}

	var err error
	*root, err = filepath.Abs(*root)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("serving %s as %s on %s", *root, *prefix, *addr)

	http.Handle(*prefix, http.StripPrefix(*prefix, http.FileServer(http.Dir(*root))))

	mux := http.DefaultServeMux.ServeHTTP

	var logger http.HandlerFunc
	if *dumpPost {
		logger = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Print(r.RemoteAddr + " " + r.Method + " " + r.URL.String())
			io.Copy(os.Stderr, r.Body)
			os.Stderr.Write([]byte{'\n'})
			mux(w, r)
		})
	} else {
		logger = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Print(r.RemoteAddr + " " + r.Method + " " + r.URL.String())
			mux(w, r)
		})
	}

	if *certFile != "" && *keyFile != "" {
		err = http.ListenAndServeTLS(*addr, *certFile, *keyFile, logger)
	} else {
		err = http.ListenAndServe(*addr, logger)
	}
	if err != nil {
		log.Fatalln(err)
	}
}
