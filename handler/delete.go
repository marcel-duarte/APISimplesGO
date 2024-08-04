package handler

import (
	"apisimples/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id_indice, err := strconv.Atoi(chi.URLParam(r, "id_indice"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := model.Delete(int64(id_indice))
	if err != nil {
		log.Printf("Erro ao remover o registro: %v ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: Foram removidos %d registros ", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Dados removidos com sucesso.",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
