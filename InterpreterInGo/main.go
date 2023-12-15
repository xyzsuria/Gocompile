package main

import (
	"InterpreterInGo/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s,succeed\n", user.Username)
	fmt.Printf("输入：\n")
	repl.Start(os.Stdin, os.Stdout)
}
