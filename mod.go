package main

import(
	"log"
	// "fmt"
	"encoding/json"
	"net/http"
 	"github.com/gorilla/mux"
)

type Block struct{
	Pos 		int
	Data		BookCheckout
	Hash		string
	Timestamp	string
	PrevHash	string

}

type BookCheckout struct{
	BookId 		 string   'json:"book_id"'
	User         string   'json:"user"'
	CheckoutDate string	  'json:"checkout_date"'
	IsGenesis    bool     'json:"is_genesis"'
}

type Book struct{
	ID 			 string   'json:"id"'
	Title		 string	  'json:"title"'	
	Author		 string   'json:"author"'
	PublishDate	 string   'json:"publish_date"'
	ISBN		 string   'json:"isbn"'
}

type Blockchain struct{
	blocks []*Block //Slice of multiple blocks

}

var Blockchain *Blockchain

func newBook(w http.ResponseWriter , r *http.Request){
	var book Book
	if err := json.NewDecoder(r.body).Decode(&book); err!=nil{
		w.writeHeader(http.StatusInternalServerError);
		log.Printf("could not create")
	}


}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/",getBlockchain).Methods("GET")
	r.HandleFunc("/",writeBlock).Methods("POST")
	r.HandleFunc("/new",newBook).Methods("POST")

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000",r))

}