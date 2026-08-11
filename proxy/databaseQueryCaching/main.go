package main

import (
	"dbquerycaching/db"
	"fmt"
)

func main() {
	db := db.NewDBCache()
	fmt.Println(db.Query("SELECT * FROM users"))

	fmt.Println(db.Query("SELECT * FROM users"))
}
