package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Args := os.Args[1:]

	file, err := os.Open(Args[0])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
