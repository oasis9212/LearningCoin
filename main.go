package main

import (
	"fmt"
	"ralo/goElement/modules"
)

func main() {
	fmt.Println("it works")
	nako := modules.Person{}

	fmt.Println(nako)
	nako.SeDetails("나코", 21)
	fmt.Println(nako)

}
