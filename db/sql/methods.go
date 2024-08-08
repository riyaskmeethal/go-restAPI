package sql

import (
	"gorestAPI/models"
	"log"
)

func (db *DbObj) ConnectionCheck() bool {
	return db.DB != nil && db.DB.Ping() == nil
}

func (DbObj *DbObj) GetTest() (Testdata []models.Test, err error) {
	err = DbObj.DB.Ping() // ping database to verify/establish connection
	if err != nil {
		return nil, err
	}
	_, err = DbObj.DB.Exec(`CREATE TABLE IF NOT EXISTS test(id int primary key , title text)`)
	if err != nil {
		return nil, err
	}
	rows, err := DbObj.DB.Query("SELECT * FROM test")
	if err != nil {
		return nil, err
	} else {

		b := models.Test{}

		for rows.Next() {

			if err := rows.Scan(&b.ID, &b.Title); err != nil {
				log.Println("reading err", err)
			}
			Testdata = append(Testdata, b)
		}
		if err != nil {
			return nil, err
		} else {
			log.Println("test data fetched")
		}
	}
	return Testdata, err
}
