package todos

import (
	"fmt"
	"net/http"
)

// GetTodos List all todos from database
func GetTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Todos")
}
