package main

import (
	"fmt"
	"golang"
)

func main() {
	fmt.Println(golang.BasicAtoi("12345"))
	fmt.Println(golang.BasicAtoi("0000000012345"))
	fmt.Println(golang.BasicAtoi("000000"))
}