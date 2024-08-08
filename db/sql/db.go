package sql

import (
	"database/sql"
	"fmt"
	"gorestAPI/configs"
	"log"

	_ "github.com/lib/pq"
)

type DbObj struct {
	DB   *sql.DB
	Conf configs.Configs
}

func (sqlDB *DbObj) DbConnect() {
	var err error
	DBconf := sqlDB.Conf.Database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBconf.Host, DBconf.Port, DBconf.Username, DBconf.Password, sqlDB.Conf.Database.DBName)

	// connStr := os.Getenv("DATABASE_URL")

	sqlDB.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to connecting DB", err)
		return
	}
	err = sqlDB.DB.Ping()

	log.Println("connected DB", err)
}
