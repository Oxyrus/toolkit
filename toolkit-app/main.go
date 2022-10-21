package main

import (
	"fmt"
	"github.com/Oxyrus/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)

	fmt.Println("Random string: ", s)
}
