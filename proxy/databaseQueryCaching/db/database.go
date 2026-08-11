package db

import (
	"fmt"
	"time"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (this Database) Query(sql string) string {
	fmt.Println("[Database] Executing query: ", sql)
	time.Sleep(1000 * time.Millisecond)
	return fmt.Sprintf("Result for [%s]\n", sql)
}
