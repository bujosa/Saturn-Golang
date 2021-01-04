package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",Index)
	router.HandleFunc("/planets",Planets)
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola mundo desde el servidor web con GO")
}

func Planets(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Pagina de los planetas")
}

