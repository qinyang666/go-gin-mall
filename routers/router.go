package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-mall/api/admin"
)

func InitRouter() *gin.Engine{
	r := gin.Default()

	r.POST("/category", admin.AddProductCategory)
	//authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//	// nested group
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}

	// Listen and serve on 0.0.0.0:8080
	return r
}

