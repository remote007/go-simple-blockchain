package main

import(
	"log"
	// "fmt"
	"net/http"
 	"github.com/gorilla/mux"
)

type Block struct{

}

type BookCheckout struct{

}

type Book struct{
	ID string
	Title string
	Author string
	PublishDate string
	ISBN string
}

type Blockchain struct{
	blocks []*Block //Slice of multiple blocks

}

var Blockchain *Blockchain

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/",getBlockchain).Methods("GET")
	r.HandleFunc("/",writeBlock).Methods("POST")
	r.HandleFunc("/new",newBook).Methods("POST")

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000",r))

}