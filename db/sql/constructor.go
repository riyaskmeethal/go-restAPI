package sql

import "gorestAPI/configs"

func NewDBObj(Conf configs.Configs) *DbObj {
	db := DbObj{
		Conf: Conf,
	}
	db.DbConnect()
	return &db
}
