package config

import "database/sql"

func ConnectDB() *sql.DB {
	dbDriver, dbUser, dbPass, dbName := "mysql", "eep", "1903", "CRUD_API_MEDIUM_COM"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	return db
}
