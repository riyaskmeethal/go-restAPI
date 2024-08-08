package server

import (
	"gorestAPI/configs"
	"gorestAPI/handlers"
	"gorestAPI/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRoutes() (r *gin.Engine) {
	conf := configs.ReadConf()
	NC := handlers.NewController(conf)
	r = gin.Default()
	RG := r.Group("/test")
	RG.Use(middlewares.DBConnectionCheck(NC.DB))
	RG.Use(middlewares.Logger())
	{
		RG.GET("", NC.GetTests)

	}
	return
}
