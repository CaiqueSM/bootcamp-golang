package main

import(
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request){
	fmt.Println(w, "Olá\n")
}

func main(){
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}