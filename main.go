package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go Testing")
}

// GetNum return a number
func GetNum() int {
	return 5
}
