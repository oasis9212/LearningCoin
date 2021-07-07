package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ralo/blockchain"
	"ralo/utils"
)

var port string

type url string

// MarshalText는  json string 으로써 어떻게 보여줄지 결정하는 interface이다.
func (u url) MarshalText() ([]byte, error) { // 묵시적으로 어떤 타입을 쓸것인지만 선정해준다.
	url := fmt.Sprintf("http://locahost%s%s", port, u) // url 명시
	return []byte(url), nil                            // URL 타입이 변환.
}

// 마샬하던 언마샬 을 하던 결과물을 수정하는 interface가 존재해야한다.

type urlDescrption struct {
	Url         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` //omitempty 필수가 아니다.  없으면 냅둔다.
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("documentation 진입")
	data := []urlDescrption{
		{
			Url:         url("/"), // URL 타입 결정.
			Method:      "GET",
			Description: "See documention",
		},
		{
			Url:         url("/blocks"),
			Method:      "GET",
			Description: "See All blocks",
			Payload:     "data:string",
		},
		{
			Url:         url("/blocks"),
			Method:      "POST",
			Description: "Add A blocks",
			Payload:     "data:string",
		},
		{
			Url:         url("/blocks/{id [0-9}+}"),
			Method:      "GET",
			Description: "See a Block",
			Payload:     "data:string",
		},
	}
	fmt.Println(data)
	rw.Header().Add("Content-Type", "application/json")
	//result,err:=json.Marshal(data)
	//
	//utils.HandleErr(err)
	//똑같다.
	//fmt.Fprintf(rw,"%s",result)
	json.NewEncoder(rw).Encode(data)

}

func blocks(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("blocks 진입")
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockChain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody

		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))

		blockchain.GetBlockChain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)

	}

}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)

}

func Start(aport int) {

	router := mux.NewRouter() // url(/block) 와 url 함수(blocks) 를 연결해주는 역할.
	port = fmt.Sprintf(":%d", aport)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{id: [0-9}+}", block).Methods("GET")
	//explorer.Start()
	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, router))
}
