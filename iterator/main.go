package main

import (
	"fmt"
	"./iterator"
)

func main() {
	a := iterator.Book{}
	a.Name = "aaa"
	fmt.Println(a.Name)
}
