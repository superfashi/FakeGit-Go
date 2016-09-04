package main

import (
	"github.com/hanbang-wang/FakeGit-Go"
	"fmt"
)

func main() {
	addition, name, email, forever := fakegit.ProcArgs()
	cfg := fakegit.NewGitConf(".git/config")
	if forever && name == "" {
		fakegit.Fatal(fakegit.ARGUMENT_ERROR_USERNAME)
	}
	if name != "" {
		cfg.Change(name, email)
	}
	if !forever {
		fmt.Println("git " + addition)
		fakegit.RunCommand("git " + addition)
		if name != "" {
			cfg.Recover()
		}
	}
}