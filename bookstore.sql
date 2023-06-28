CREATE TABLE authors (
                         id SERIAL PRIMARY KEY,
                         author TEXT
);

CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       book_name TEXT,
                       edition TEXT,
                       publication_year TEXT,
                       authors TEXT
);

COPY authors(authors) FROM './var/lib/postgresql/data/authors.csv' DELIMITER ';' CSV;