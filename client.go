// +build ignore

package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"os"
	"io/ioutil"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func main() {
	flag.Parse()
	var hostName string

	if flag.NArg() < 1 {
		log.Fatal("Filename not specified")
	}
	if (flag.NArg() == 2) {
		hostName = flag.Args()[1]
		addr := flag.String("addr", hostName+":8080", "http service address")
		_ = addr
	} else {
		addr := flag.String("addr", "localhost:8080", "http service address")
		_ = addr
	}

	// Get filename
	filename := flag.Args()[0]

	// Upload file to server
	upload(filename, hostName)
}

func upload(filename string, hostName string) {
	file, err := os.Open("./send-files/"+filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post("http://"+hostName+":8080/upload?filename="+filename, "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}