package tests

import (
	"os"
	"net/http"
	"io/ioutil"
	"fmt"
)

func upload(filename string, hostName string) string {
	file, err := os.Open("../send-files/"+filename)
	if err != nil {
		return "file not found"
	}
	defer file.Close()

	res, err := http.Post("http://"+hostName+":8080/upload?filename="+filename, "binary/octet-stream", file)
	if err != nil {
		return "file not send"
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))

	return string(message)
}