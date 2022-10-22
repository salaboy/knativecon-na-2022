package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/tjarratt/babble"
)



func main() {



	r := mux.NewRouter()
	r.HandleFunc("/", InputHandler).Methods("GET")

	log.Printf("Strange app Started in port 8080!")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}


func InputHandler(writer http.ResponseWriter, request *http.Request) {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = 3

	fmt.Fprintf(writer, babbler.Babble())

}

