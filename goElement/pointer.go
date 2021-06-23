package main

import (
	"fmt"
)

func main() {
	a := 2
	b := &a //  주소를 저장만 하고싶다면 & 만 써라.
	fmt.Println(b)
	a = 33
	fmt.Println(*b)

	//nako:=person.Person{"나코",21}
	//fmt.Println(nako.name) 소문자로 되어서 임포트 불가.
}
