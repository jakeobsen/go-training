package main

import (
	"flag"
	"fmt"
)

func main() {
	age := flag.Int("age", 30, "Persons age")
	name := flag.String("name", "John doe", "Persons name")
	flag.Parse()
	fmt.Println(*name, "is", *age, "years old")
}
