package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var copyCmd *exec.Cmd

func InputBox(title, message, defaultAnswer string) (string, bool) {
	out, err := exec.Command(
		"zenity", "--entry",
		"--title", title,
		"--text", message,
		"--entry-text", defaultAnswer).Output()
	if err != nil {
		return "", false
	}
	return strings.TrimSpace(string(out)), true
}

func main() {
	str, _ := InputBox("><>", "><>", "")

	copyCmd = exec.Command("xclip", "-selection", "c")

	in, err := copyCmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := copyCmd.Start(); err != nil {
		log.Fatal(err)
	}

	if _, err := in.Write([]byte(str)); err != nil {
		log.Fatal(err)
	}

	if err := in.Close(); err != nil {
		log.Fatal(err)
	}

	copyCmd.Wait()

	fmt.Print(str)
}
