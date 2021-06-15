package main

import (
	"net/http"
	"webservice/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run(":8080")
}

//https://github.com/gin-gonic/examples/blob/master/basic/main.go
//Setup router using gin engine with basic configuration.
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	articleRepo := controllers.New()
	r.POST("/article", articleRepo.CreateArticle)
	r.GET("/articles", articleRepo.GetArticles)
	r.GET("/article/:id", articleRepo.GetArticle)
	r.PUT("/article/:id", articleRepo.UpdateArticle)
	r.DELETE("/article/:id", articleRepo.DeleteArticle)

	return r
}
