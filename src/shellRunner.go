package main

import (
	"bytes"
	"log"
	"os/exec"
)

func execCmd(cmdString string) (string, error) {
	cmd := exec.Command("sh", "-c", cmdString)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Command execution '%s' failed %s, %s", cmdString, stdout.String(), stderr.String())
		return "", err
	} else {
		return string(stdout.Bytes()), nil
	}
}
