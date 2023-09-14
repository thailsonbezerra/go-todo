package todos

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thailsonbezerra/go-todo/src/config/database"
	modelTodo "github.com/thailsonbezerra/go-todo/src/models/todos"
)

func CriarTodo(w http.ResponseWriter, r *http.Request) {

	var todo modelTodo.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	db := database.GetDB()

	err := db.QueryRow("INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING id, createdAt", todo.Title, todo.Description).Scan(&todo.ID, &todo.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(todo)

}

func BuscarTodos(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}

	var todos []modelTodo.Todo
	for rows.Next() {

		var todo modelTodo.Todo
		if err := rows.Scan(&todo.ID, &todo.CreatedAt, &todo.Title, &todo.Description, &todo.Completed); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(todos)

}

func BuscarTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var todo modelTodo.Todo
	db := database.GetDB()
	err := db.QueryRow("SELECT * FROM todos WHERE id = $1", id, ).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r) // Retorna status 404 se o TODO não for encontrado
			return
		}
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(todo)
}

func AtualizarTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var todo modelTodo.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	db := database.GetDB()
	statement, err := db.Prepare("UPDATE todos SET title = $2, description = $3, completed = $4 WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	if _, err = statement.Exec(id, todo.Title, todo.Description, todo.Completed); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Todo atualizado")
}

func DeletarTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := database.GetDB()
	_, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Todo deletado")

}
