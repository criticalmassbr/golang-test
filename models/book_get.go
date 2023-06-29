package models

import "DialogBookStore/db"

func BookGet(id int64) (books Books, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM books WHERE id=$1`, id)

	err = row.Scan(&books.ID, &books.BookName, &books.Edition, &books.PublicationYear, &books.Authors)

	return
}
