package todos

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/thailsonbezerra/go-todo/src/config/database"
	"github.com/thailsonbezerra/go-todo/src/models/todos"
)

func CriarTodo(w http.ResponseWriter, r *http.Request) {

	var t todos.Todo
	json.NewDecoder(r.Body).Decode(&t)

	db := database.GetDB()

	err := db.QueryRow("INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING id", t.Title, t.Description).Scan(&t.ID)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(t)

}

func BuscarTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando todos os TODOs!"))
}

func BuscarTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um TODO!"))
}

func AtualizarTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando um TODO!"))
}

func DeletarTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um TODO!"))
}
