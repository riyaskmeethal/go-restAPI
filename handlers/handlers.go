package handlers

import (
	"gorestAPI/configs"
	"gorestAPI/db"
	"gorestAPI/db/sql"
	"gorestAPI/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	DB   db.DBinterface
	Conf configs.Configs
}

func NewController(conf configs.Configs) (C *Controller) {
	return &Controller{
		DB:   sql.NewDBObj(conf),
		Conf: conf,
	}
}
func (HC *Controller) GetTests(ctx *gin.Context) {
	response := new(utils.DataExchange)
	response.Header = new(utils.Header)
	response.Body = new(utils.Body)
	Test, err := HC.DB.GetTest()
	if err != nil {
		log.Println(err)
		message := "Failed to fetch data."
		utils.SendFailure(response, message, http.StatusBadRequest, ctx)
		return
	}
	response.Body.Test = &Test
	message := "Test data loaded"
	utils.SendSuccess(response, message, http.StatusOK, ctx)
}
