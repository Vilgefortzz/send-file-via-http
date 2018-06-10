// +build ignore

package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/websocket"
	"os/exec"
	"os"
	"io"
)

var upgrader = websocket.Upgrader{} // use default options
const GO_EXTENSION string = ".go"

func main() {
	flag.Parse()
	log.SetFlags(0)
	var addr *string

	if (flag.NArg() >= 1) {
		hostName := flag.Args()[0]
		addr = flag.String("addr", hostName+":8080", "http service address")
	} else {
		addr = flag.String("addr", "localhost:8080", "http service address")
	}

	http.HandleFunc("/upload", uploadHandler)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	ext := getExtensionFromFile(filename)

	if filename == "" {
		log.Fatal("Filename is empty")
	}

	file, err := os.Create("./get-files/"+filename)
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, r.Body)
	_ = n
	if err != nil {
		panic(err)
	}

	if ext != GO_EXTENSION {
		w.Write([]byte(fmt.Sprintf("Source file extension is not valid\n")))
	} else {
		result := compileFileAndGetResult(filename)
		w.Write([]byte(fmt.Sprintf("__Output__\n------------\n%s------------\n", result)))
	}
}

func getExtensionFromFile(filename string) string {
	return path.Ext(filename)
}

func compileFileAndGetResult(filename string) string{
	var (
		cmdOut []byte
		err error
	)
	cmdName := "go"
	cmdArgs := []string{"run", "./get-files/"+filename}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error compiling go source file: ", err)
		os.Exit(1)
	}
	output := string(cmdOut)
	return output
}