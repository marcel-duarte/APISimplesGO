package handler

import (
	"apisimples/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id_indice, err := strconv.Atoi(chi.URLParam(r, "id_indice"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	indice, err := model.Get(int64(id_indice))
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(indice)
}
