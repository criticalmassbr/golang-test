package models

import "DialogBookStore/db"

func Insert(books Books) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO books (book_name, edition, publication_year, authors) VALUES ($1, $2, $3, $4) RETURNING id`

	err = conn.QueryRow(sql, books.BookName, books.Edition, books.PublicationYear, books.Authors).Scan(&id)

	return
}
