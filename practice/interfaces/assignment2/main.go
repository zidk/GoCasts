package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	var mf string

	if len(args) > 1 {
		mf = args[1]
	} else {
		fmt.Printf("Please add filename arg")
		os.Exit(1)
	}

	f, err := os.Open(mf)
	if err != nil {
		fmt.Printf("Getting err trying to open %v : %v", f, err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		fmt.Printf("Getting err trying to output the data: %v", err)
		os.Exit(1)
	}
}
