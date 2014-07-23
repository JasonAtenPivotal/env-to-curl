package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Result struct {
	Pipeline  string `json:"pipeline"`
	Pipecount int    `json:"pipecount"`
}

func toInt(a string) int {
	val, err := strconv.Atoi(a)
	if err != nil {
		panic(fmt.Sprintf("could not convert '%s' to int: %s", a, err))
	}
	return val
}

func main() {

	os.Setenv("GO_PIPELINE_NAME", "test-pipe")
	os.Setenv("GO_PIPELINE_COUNT", "42")

	r := Result{
		Pipeline:  os.Getenv("GO_PIPELINE_NAME"),
		Pipecount: toInt(os.Getenv("GO_PIPELINE_COUNT")),
	}

	json, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("json = %s\n", json)

	resp, err := http.Post("http://localhost:3000/results", "application/json", bytes.NewBuffer(json))

	body := bytes.NewBuffer(nil)
	io.Copy(body, resp.Body)

	fmt.Printf("posted: %s\nresp.Body is '%s'\n", json, string(body.Bytes()))

}
