package main

import "fmt"

func main() {
	a := 2
	b := &a //  주소를 저장만 하고싶다면 & 만 써라.
	fmt.Println(b)
	a = 33
	fmt.Println(*b)

}
