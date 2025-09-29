package utils

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

// test the function with stdin stdout : "sh", "-c", "echo stdout; echo 1>&2 stderr"
func Exec(name string, args ...string) (stdin, stderr string) {

	cmd := exec.Command(name, args)
	stder_, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
