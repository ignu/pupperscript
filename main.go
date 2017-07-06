package main

import (
	"fmt"
	"os"
	"os/user"
	"pupperscript/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("H*ckin' Bork, %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
