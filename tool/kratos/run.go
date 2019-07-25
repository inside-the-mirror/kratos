package main

import (
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func runAction(c *cli.Context) error {
	base, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := buildDir(base, "cmd", 5)
	args := append([]string{"run", "main.go"}, c.Args()...)
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	return nil
}
