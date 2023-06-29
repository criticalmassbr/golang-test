package handlers

import (
	"DialogBookStore/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func BookCreate(w http.ResponseWriter, r *http.Request) {
	var books models.Books

	err := json.NewDecoder(r.Body).Decode(&books)
	if err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := models.BookInsert(books)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Message": fmt.Sprintf("Failed trying to insert book: %v", err),
		}
	} else {
		resp = map[string]any{
			"Message": fmt.Sprintf("Book successfully inserted! Id: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}