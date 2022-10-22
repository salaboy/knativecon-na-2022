package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/tjarratt/babble"
	"os"
	"strconv"
)


var wordsCount = os.Getenv("WORDS_COUNT")

func main() {



	r := mux.NewRouter()
	r.HandleFunc("/", InputHandler).Methods("GET")

	log.Printf("Strange app Started in port 8080!")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}


func InputHandler(writer http.ResponseWriter, request *http.Request) {
	var wordsCountInt int
	if wordsCount == ""{
		wordsCountInt = 3
	}else{
		wordsCountInt, _ = strconv.Atoi(wordsCount)
	}
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = wordsCountInt

	fmt.Fprintf(writer, babbler.Babble())

}

