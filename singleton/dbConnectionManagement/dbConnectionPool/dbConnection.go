package dbconnectionpool

type DBConnection struct {
	id int
}

func NewDBConnection(id int) *DBConnection {
	return &DBConnection{
		id: id,
	}
}
