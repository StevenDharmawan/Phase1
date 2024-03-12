package handler

import (
	"fmt"
	"latihan/config"
)

func Books() {
	query := "SELECT books.book_title FROM books JOIN authors ON books.author_id = authors.author_id WHERE authors.author_id = 1"
	rows, err := config.DB.Query(query)
	if err != nil {
		fmt.Println("Failed to select data : ", err.Error())
		return
	}
	fmt.Println("Books by Jane Smith:")
	for rows.Next() {
		var bookTitle string
		err = rows.Scan(&bookTitle)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(bookTitle)
	}

}
