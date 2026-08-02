package main

import (
	dbconnectionpool "dbconnectionmanagement/dbConnectionPool"
	"fmt"
)

func aquireConnections(dbPool *dbconnectionpool.DBConnectionPool, n int) []*dbconnectionpool.DBConnection {
	var dbConnections []*dbconnectionpool.DBConnection

	for i := 0; i < n; i++ {
		db, err := dbPool.GetConnection()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("connection aquired!")
			dbConnections = append(dbConnections, db)
		}
	}

	return dbConnections
}

func releaseConnections(dbPool *dbconnectionpool.DBConnectionPool, connArray []*dbconnectionpool.DBConnection) {
	for i := 0; i < len(connArray); i++ {
		dbPool.ReleaseConnection(connArray[i])
		fmt.Println("connection released!")
	}
}

func main() {
	dbPool := dbconnectionpool.GetDBConnection(10)

	dbcon := aquireConnections(dbPool, 10)
	releaseConnections(dbPool, dbcon)
}
