package handler

import (
	"fmt"
	"latihan/config"
)

func Sales() {
	query := "SELECT book_type, SUM(price) FROM orders  JOIN books_detail ON orders.book_detail_id = books_detail.book_detail_id GROUP BY book_type"
	rows, err := config.DB.Query(query)
	if err != nil {
		fmt.Println("Failed to select data : ", err.Error())
		return
	}
	fmt.Println("Total sales for each book type:")
	for rows.Next() {
		var bookType string
		var totalPrice float64
		err = rows.Scan(&bookType, &totalPrice)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(bookType, ":", totalPrice)
	}
}
