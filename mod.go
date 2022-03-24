package main

import(

)

func main(){
	r := mux.newRouter()
	r.Handle("/").Methods("GET")
}