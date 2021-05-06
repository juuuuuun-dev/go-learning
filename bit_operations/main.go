package main

import "fmt"

func main() {
	i := 8
	j := 9
	fmt.Printf("%b ", i)    //1000
	fmt.Printf("%b ", j)    //1001
	fmt.Printf("%b ", i&j)  // 1000
	fmt.Printf("%b ", i|j)  // 1001
	fmt.Printf("%b ", i^j)  // 1
	fmt.Printf("%b ", i&^j) // 0
	fmt.Printf("%b ", i<<1) // 10000
	fmt.Printf("%b ", j<<1) // 10010
	fmt.Printf("%b ", i>>1) // 100
	fmt.Printf("%b ", j>>1) // 100
}
