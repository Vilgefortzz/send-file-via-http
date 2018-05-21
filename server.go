// +build ignore

package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"os"
	"io"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func main() {
	flag.Parse()
	log.SetFlags(0)

	http.HandleFunc("/upload", uploadHandler)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	file, err := os.Create("./get-files/"+filename)
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bytes:", n)
}