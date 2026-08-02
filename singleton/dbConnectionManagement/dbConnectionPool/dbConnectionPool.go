package dbconnectionpool

import (
	"errors"
	"sync"
)

type DBConnectionPool struct {
	mu                   sync.Mutex
	dbConnections        map[*DBConnection]bool
}

var (
	instance *DBConnectionPool
	once     sync.Once
)

func GetDBConnection(maxConnections int) *DBConnectionPool {
	once.Do(func() {
		connectionMap := make(map[*DBConnection]bool, maxConnections)

		for i := 0; i < maxConnections; i++ {
			conn := NewDBConnection(i)
			connectionMap[conn] = true
		}

		instance = &DBConnectionPool{
			dbConnections:        connectionMap,
		}
	})

	return instance
}

func (db *DBConnectionPool) GetConnection() (*DBConnection, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for conn, isAvailable := range db.dbConnections {
		if isAvailable {
			db.dbConnections[conn] = false
			return conn, nil
		}
	}

	return nil, errors.New("DB connection unavailable")
}

func (db *DBConnectionPool) ReleaseConnection(conn *DBConnection) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.dbConnections[conn]; !exists {
		return
	}

	db.dbConnections[conn] = true
}

func (db *DBConnectionPool) GetAvailableCount() int {
	db.mu.Lock()
	defer db.mu.Unlock()

	counter := 0

	for _, isAvailable := range db.dbConnections {
		if isAvailable {
			counter++
		}
	}

	return counter
}
