package handlers

import (
	"DialogBookStore/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func BookList(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Default values if page or limit are not provided or invalid
	defaultPage := 1
	defaultLimit := 10

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	books, totalCount, err := models.BookGetAll(page, limit)
	if err != nil {
		log.Printf("Error obtaining book registers: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"books":      books,
		"page":       page,
		"pageSize":   limit,
		"totalCount": totalCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
