// +build ignore

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Filename not specified")
	}

	// Get filename
	filename := flag.Args()[0]

	// Upload file to server
	upload(filename)
}

func upload(filename string) {
	file, err := os.Open("./send-files/"+filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post("http://127.0.0.1:8080/upload?filename="+filename, "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
}