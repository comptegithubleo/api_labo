package utils

import (
	"io"
	"log"
	"os/exec"
)

// executes an os commands, returns stdout & stderr
// test the function with stdin stdout : "sh", "-c", "echo stdout; echo 1>&2 stderr"
func Exec(name string, args ...string) (string, string) {

	cmd := exec.Command(name, args...)
	stdoutBytes, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderrBytes, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Println("[X] Error starting cmd : ")
	}

	stdout_, _ := io.ReadAll(stdoutBytes)
	stderr_, _ := io.ReadAll(stderrBytes)

	// we do not error check, we already read stderr and let callee deal with it (throw error or not)
	// still needed because we wait for command to stop and all stdout stderr pipe to be read
	cmd.Wait()

	return string(stdout_), string(stderr_)
}
