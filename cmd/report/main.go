package main

import (
	"fmt"
	"os"
)

const usage string = `Usage: report file.json`

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}

}
