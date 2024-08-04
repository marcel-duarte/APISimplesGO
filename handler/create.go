package handler

import (
	"apisimples/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var indice model.Indice

	err := json.NewDecoder(r.Body).Decode(&indice)

	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return		
	}

	id, err := model.Insert(indice)

	var resp map[string] any

	if err != nil {
		map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("Indice inserido com sucesso! Id_Indice: %v", id)
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}