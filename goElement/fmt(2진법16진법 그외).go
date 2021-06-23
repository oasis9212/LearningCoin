package main

import "fmt"

func main() {
	x := 19
	fmt.Printf("%b\n", x)               // 2진법.
	fmt.Printf("%o\n", x)               // 8진법.
	fmt.Printf("%x\n", x)               // 16진법.
	fmt.Printf("%U\n", x)               // 유니코드  영어이외의 언어를 byte로 나타낸거.
	xAsBinary := fmt.Sprintf("%b\n", x) // 2진법.

	fmt.Println(x, xAsBinary)

}
