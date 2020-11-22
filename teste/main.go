package main

import (
	"fmt"

	"./mymap"
)

func main() {
	m := make(map[string]string)
	m["nome4"] = "s"
	pointer := mymap.SharedMap{m}
	fmt.Println(pointer.Register("nome1"))
	fmt.Println(pointer.Register("nome2"))
	fmt.Println(pointer.Register("nome3"))
	fmt.Println(pointer.Lookup("nome4"))

	fmt.Println(pointer.Lookup("nome1"))
}
