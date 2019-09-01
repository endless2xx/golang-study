package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
	)

	cmd = exec.Command("D:\\software\\Git\\git-bash.exe", "-c", "echo1")

	err = cmd.Run()

	fmt.Println(err)
}
