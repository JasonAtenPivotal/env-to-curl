package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("request r = %#v\n", r)
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, r.Body)
		fmt.Printf("request r.Body = %s\n", buf)
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "server got request body:, %s", string(buf.Bytes()))
	})

	addr := "localhost:3000"
	fmt.Printf("listening on %s and responding to /results\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
