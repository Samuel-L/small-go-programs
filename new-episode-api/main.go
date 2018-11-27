package main

import (
	"fmt"
	"net/http"

	"github.com/Samuel-L/new-episode-api/internal/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/new-episode/", api.NewEpisode).Methods("POST")

	fmt.Println(http.ListenAndServe(":8000", router))
}
