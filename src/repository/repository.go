package repository

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var mu = &sync.Mutex{}
var databaseClient *ParkingLotRepository

type ParkingLotRepository struct {
	DbClient *sql.DB
}

func GetDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/pk_db")
	if err != nil || db == nil {
		fmt.Println("Error while opening db driver for sql-lite. Error: ", err)
		panic(err)
	}
	fmt.Println("Successfully initialized driver for my-sql")
	return db
}

func InitialiseDatabase() *ParkingLotRepository {
	// singleton design pattern
	isFreshDB := false
	if databaseClient == nil || databaseClient.DbClient == nil {
		mu.Lock()
		defer mu.Unlock()
		if databaseClient == nil || databaseClient.DbClient == nil {
			isFreshDB = true
			databaseClient = &ParkingLotRepository{
				DbClient: GetDatabase(),
			}
		}
	}

	if isFreshDB {
		databaseClient.DeleteTables()
		databaseClient.CreateTables()
		databaseClient.InsertValuesIntoEnums()
	}
	return databaseClient
}
