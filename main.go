package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("useless coin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:   this starts the HTML explorer\n")
	fmt.Printf("rest:   this starts the REST API (recommended)\n")
	os.Exit(0)

}

func main() {

	fmt.Println(os.Args)
	// Array[2:]  3번째 배열부터 끝까지 를 뜻한다.
	// os.Args 는 길이가 1인 전재로 시작한다 인자값 go run main.go 자체가 길이가 1일다.
	//go explorer.Start(3000)
	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	portFlag := rest.Int("port", 4000, "Sets the port of the server ")
	switch os.Args[1] {
	case "explorer":
		fmt.Println("start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Println("Start server")
	}
	fmt.Println(*portFlag)
}
