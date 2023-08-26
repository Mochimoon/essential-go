package main 

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Hello, World")

	// convert int to str
	var i1 int = -38
	fmt.Printf("i1: %s\n", strconv.Itoa(i1))

	s1 := fmt.Sprintf("%d", i1)
	fmt.Printf("i1: %s\n", s1)
}