package main

import (
	"fmt"
	"latihan/config"
	"latihan/handler"
	"os"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	if len(os.Args) < 2 {
		fmt.Println("Please specify a command: books, sales, customers, topauthor")
	}

	command := os.Args[1]
	switch command {
	case "books":
		handler.Books()
	case "sales":
		handler.Sales()
	case "customers":
		handler.Customers()
	case "topauthor":
		handler.Topauthor()
	default:
		fmt.Println("Unknown command:", command)
	}
}
