package models

type Books struct {
	ID              int64  `json:"id"`
	BookName        string `json:"book_name"`
	Edition         string `json:"edition"`
	PublicationYear string `json:"publication_year"`
}

type Authors struct {
	ID         int64  `json:"id"`
	AuthorName string `json:"Author_name"`
}
