package explorer

import (
	"fmt"
	"log"
	"net/http"
	"ralo/blockchain"
	"text/template"
)

const (
	templateDir string = "explorer/templates/"
)

type HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

//template.ParseFiles("templates/home.html")
func home(rw http.ResponseWriter, r *http.Request) {
	data := HomeData{"Home", nil}
	templates.ExecuteTemplate(rw, "home", data)

}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.BlockChain().Addblcok(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

var templates *template.Template

func Start(port int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home) // url 함수
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
