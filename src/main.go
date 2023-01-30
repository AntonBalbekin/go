package main

import (
	db "usr/api/pkg/renplus/sqlconect"

	. "usr/api/pkg/renplus/logic"

	"github.com/gin-gonic/gin"
	
)

func main() {
	defer db.SqlDB.Close()

	router := initRouter()
	router.Run(":4000")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/go/chekclient", ChekClient)
	return router
}
