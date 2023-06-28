package handlers

import (
	"DialogBookStore/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAll()
	if err != nil {
		log.Printf("Error obtaining book registers: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
