package tests

import (
	"os/exec"
	"fmt"
	"os"
)

func compileFileAndGetResult(filename string) string{
	var (
		cmdOut []byte
		err error
	)
	cmdName := "go"
	cmdArgs := []string{"run", "../get-files/"+filename}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error compiling go source file: ", err)
		return "compiling error"
	}
	output := string(cmdOut)
	return output
}