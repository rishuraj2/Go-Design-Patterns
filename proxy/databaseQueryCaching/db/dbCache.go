package db

import "fmt"

type DBCache struct {
	database *Database
	cache    map[string]string
}

func NewDBCache() *DBCache {
	return &DBCache{
		database: NewDatabase(),
		cache:    make(map[string]string),
	}
}

func (this DBCache) Query(sql string) string {
	if result, exists := this.cache[sql]; exists {
		fmt.Println("[Cache] HIT for: ", sql)
		return result
	}

	fmt.Println("[Cache] MISS for: ", sql)
	res := this.database.Query(sql)
	this.cache[sql] = res
	return res
}

func (this *DBCache) ClearCache() {
	this.cache = make(map[string]string)
	fmt.Println("[Cache] Cleared")
}
