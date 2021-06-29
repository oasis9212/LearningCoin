package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ralo/blockchain"
)

const port string = ":4000"

type HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

//template.ParseFiles("templates/home.html")
func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml")) // template.Must 자동으로 에러처리를 해주는 방식.
	data := HomeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	tmpl.Execute(rw, data)

}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
