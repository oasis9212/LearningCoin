package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ralo/blockchain"
	"ralo/utils"
	"strconv"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string `json:"message"`
}

type errorResponse struct {
	ErrorMessage string `json:"errormessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See all block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "see A block",
		},
	}

	json.NewEncoder(rw).Encode(data)
	//result,err:= json.Marshal(data)
	//utils.HandleErr(err)
	//fmt.Fprintf(rw,"%s",result)

}

// adapter
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockChain().AllBlocks())
	case "POST":
		var addblockbody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addblockbody))
		blockchain.GetBlockChain().AddBlock(addblockbody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)

	block, error := blockchain.GetBlockChain().GetBlock(id)
	encoder := json.NewEncoder(rw)
	if error == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(error)})
	} else {
		json.NewEncoder(rw).Encode(block)

	}
}

func Start(aport int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aport)
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")

	fmt.Printf("go go http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
