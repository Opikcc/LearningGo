package main

import (
	"fmt"

	"gitlab.com/opikcc/strcon"
)

func main() {
	s := strcon.SwapCase("Gopher")
	fmt.Println("Converted string is :", s)
}
