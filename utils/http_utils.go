package utils

import "github.com/gin-gonic/gin"

func SendSuccess(response *DataExchange, message string, status int, c *gin.Context) {
	response.SetState("Success", message)
	c.JSON(status, response)
	c.Set("msg", message)
	c.Set("Status", status)
}

func SendFailure(response *DataExchange, message string, status int, c *gin.Context) {
	response.SetState("Failure", message)
	c.JSON(status, response)
	c.Set("msg", message)
	c.Set("Status", status)
	c.Abort()
}
