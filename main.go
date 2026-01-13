package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pronto <command>")
		return
	}

	fmt.Println("Command:", os.Args[1])
}
