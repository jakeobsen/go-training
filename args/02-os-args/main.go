package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for i, v := range arg {
		fmt.Println("Argument:", i, "- Value:", v)
	}
}
