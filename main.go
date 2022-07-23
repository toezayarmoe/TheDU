package main

import (
	"fmt"
	"github.com/toezayarmoe/TheDU/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! The is The Damn Useless programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands!")
	repl.Start(os.Stdin, os.Stdout)

}
