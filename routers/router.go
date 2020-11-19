package routers

import "github.com/gin-gonic/gin"

func initRouter(){
	router := gin.Default()
	err := router.Run(":8888")
	if err != nil{

	}
}

