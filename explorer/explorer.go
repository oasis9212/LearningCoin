package explorer

import (
	"fmt"
	"log"
	"net/http"
	"ralo/blockchain"
	"text/template"
)

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

type HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

//template.ParseFiles("templates/home.html")
func home(rw http.ResponseWriter, r *http.Request) {
	data := HomeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)

}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockChain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

var templates *template.Template

func Start() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home) // url 함수
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
