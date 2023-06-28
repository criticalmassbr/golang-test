package models

import "DialogBookStore/db"

func GetAll() (bookslist []Books, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM books`)
	if err != nil {
		return
	}

	for rows.Next() {
		var books Books

		err = rows.Scan(&books.ID, &books.BookName, &books.Edition, &books.PublicationYear, &books.Authors)
		if err != nil {
			continue
		}

		bookslist = append(bookslist, books)
	}

	return
}
