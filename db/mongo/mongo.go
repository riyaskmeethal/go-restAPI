package mongo

import (
	"gorestAPI/configs"
	"gorestAPI/models"
	"log"
)

type MongoTest struct {
	DB   string
	Conf configs.Configs
}

func DbConnect() (sqlDB *MongoTest) {

	sqlDB = &MongoTest{}

	log.Println("connected DB")
	return
}
func NewDBObj(Conf configs.Configs) *MongoTest {
	Mdb := MongoTest{
		Conf: Conf,
	}
	return &Mdb
}
func (mdb *MongoTest) DbConnect() {
}
func (mdb *MongoTest) ConnectionCheck() bool {
	return true
}
func (mdb *MongoTest) GetTests() (Tests []models.Test, err error) {
	return
}
