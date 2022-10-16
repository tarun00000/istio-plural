package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var mapforEP = map[string]int{
	"exp": 0,
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{ep:[0-9,a-z,/]+}", counter).Methods("GET", "POST", "PUT", "DELETE")
	http.ListenAndServe(":9080", r)
}

func counter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varEP := vars["ep"]
	w.Write([]byte("Hello" + varEP))

	if mapforEP == nil {
		mapforEP[varEP] = 1
	} else {
		mapforEP[varEP] = mapforEP[varEP] + 1
	}

	for url, count := range mapforEP {
		fmt.Println(url, count)
	}
	fmt.Println("======================")
}
