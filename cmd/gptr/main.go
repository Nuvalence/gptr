package main

import (
	"fmt"
	"os"

	"github.com/nuvalence/gptr"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must specify a permission to collect the associated roles")
		os.Exit(1)
	}

	gptr.Run(os.Args[1])
}
