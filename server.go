package main

import (
	"flag"
	"fmt"
	"net/http"
)

var expectedid = flag.String("id", "", "An unique id")

func handler(w http.ResponseWriter, r *http.Request) {
	gotid := r.URL.Path[1:]
	if *expectedid != gotid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "BAD -- Expected '%s', Got '%s'\n", *expectedid, gotid)
	} else {
		fmt.Fprintf(w, "OK\n")
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
