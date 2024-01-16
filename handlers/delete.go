package handlers

import (
	"encoding/json"
	"go-crud/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: Foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"status":     "success",
		"statusCode": http.StatusOK,
		"message":    "Registro removido com sucesso",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
