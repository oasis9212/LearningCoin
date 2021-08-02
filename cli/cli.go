package cli

import (
	"flag"
	"fmt"
	"os"
	"ralo/explorer"
	"ralo/rest"
	"runtime"
)

func usage() {
	fmt.Printf("useless coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port: Set the Port of the server \n")
	fmt.Printf("-mode:  chosse 'html' and 'rest' \n")
	runtime.Goexit()

}

func Start() {
	fmt.Println(os.Args)
	// Array[2:]  3번째 배열부터 끝까지 를 뜻한다.
	// os.Args 는 길이가 1인 전재로 시작한다 인자값 go run main.go 자체가 길이가 1일다.
	//go explorer.Start(3000)
	if len(os.Args) < 2 {
		usage()
	}

	port := flag.Int("port", 4000, "Set Port Server")

	mode := flag.String("mode", "rest", "chosse 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
