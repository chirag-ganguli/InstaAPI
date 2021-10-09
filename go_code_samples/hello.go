package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var a int8
	var b int32
	fmt.Println(a)
	fmt.Println(b)

	var s string = "Hello World"
	fmt.Println(s)

	c := []byte(s) // convert string to []byte type
	c[0] = 'c'
	s2 := string(c) // convert back to string type
	fmt.Printf("%s\n", s2)
}
