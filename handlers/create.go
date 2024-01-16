package handlers

import (
	"encoding/json"
	"fmt"
	"go-crud/models"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"status":     "error",
			"statusCode": http.StatusInternalServerError,
			"message":    fmt.Sprintf("Erro ao inserir o registro: %s", err),
		}
	} else {
		resp = map[string]any{
			"status":     "success",
			"statusCode": http.StatusCreated,
			"message":    "Registro inserido com sucesso",
			"id":         id,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
