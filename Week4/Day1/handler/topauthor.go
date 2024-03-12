package handler

import (
	"fmt"
	"latihan/config"
)

func Topauthor() {
	query := "SELECT authors.author_name, SUM(price) as total FROM orders JOIN books_detail ON orders.book_detail_id = books_detail.book_detail_id JOIN books ON books_detail.book_id = books.book_id JOIN authors ON books.author_id = authors.author_id GROUP BY authors.author_name ORDER BY total DESC LIMIT 1"
	rows, err := config.DB.Query(query)
	if err != nil {
		fmt.Println("Failed to select data : ", err.Error())
		return
	}
	for rows.Next() {
		var authorName string
		var totalEarn float64
		err = rows.Scan(&authorName, &totalEarn)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Top earning author", authorName, "with earnings:", totalEarn)

	}
}
