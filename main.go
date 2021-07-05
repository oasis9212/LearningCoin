package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

// MarshalText는  json string 으로써 어떻게 보여줄지 결정하는 interface이다.
func (u URL) MarshalText() ([]byte, error) { // 묵시적으로 어떤 타입을 쓸것인지만 선정해준다.
	url := fmt.Sprintf("http://locahost%s%s", port, u) // url 명시
	return []byte(url), nil                            // URL 타입이 변환.
}

// 마샬하던 언마샬 을 하던 결과물을 수정하는 interface가 존재해야한다.

type URLDescrption struct {
	Url         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` //omitempty 필수가 아니다.  없으면 냅둔다.
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescrption{
		{
			Url:         URL("/"), // URL 타입 결정.
			Method:      "GET",
			Description: "See documention",
		},
		{
			Url:         URL("/blocks"),
			Method:      "POST",
			Description: "See documention",
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

func main() {

	http.HandleFunc("/", documentation)
	//explorer.Start()
	fmt.Printf("Listening on http://localhost%s", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
