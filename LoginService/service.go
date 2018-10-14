package main

import (
	"fmt"
	"net/http"

	ep "dcg_demo/LoginService/presentationlayer/endpoints"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	ep.InitEndpoints(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
