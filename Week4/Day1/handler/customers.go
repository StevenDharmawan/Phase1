package handler

import (
	"fmt"
	"latihan/config"
)

func Customers() {
	query := "SELECT customers.customer_name, COUNT(*) AS `order` FROM orders  JOIN customers ON orders.customer_id = customers.customer_id GROUP BY customers.customer_name HAVING COUNT(*) > 1"
	rows, err := config.DB.Query(query)
	if err != nil {
		fmt.Println("Failed to select data : ", err.Error())
		return
	}
	fmt.Println("Customers who ordered more than one book:")
	for rows.Next() {
		var customerName string
		var totalOrder int
		err = rows.Scan(&customerName, &totalOrder)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(customerName, ":", totalOrder, "orders")

	}
}
