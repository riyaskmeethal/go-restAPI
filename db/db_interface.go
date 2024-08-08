package db

import "gorestAPI/models"

type DBinterface interface {
	DbConnect()
	ConnectionCheck() bool
	GetTest() ([]models.Test, error)
}
