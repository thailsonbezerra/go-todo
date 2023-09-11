package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/thailsonbezerra/go-todo/src/config/database"
	"github.com/thailsonbezerra/go-todo/src/router"
)

func main() {

	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	router.Router(r)

	defer database.GetDB().Close()

	fmt.Printf("Servidor rodando na porta 3000\n")
	http.ListenAndServe(":3000", r)

}
