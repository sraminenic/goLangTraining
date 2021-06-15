package controllers

import (
	"errors"
	"net/http"
	"webservice/database"
	"webservice/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleRepo struct {
	Db *gorm.DB
}

func New() *ArticleRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Article{})
	return &ArticleRepo{Db: db}
}

//create Article
func (repository *ArticleRepo) CreateArticle(context *gin.Context) {
	var article models.Article
	context.BindJSON(&article)
	fmt.Println(article)
	err := models.CreateArticle(repository.Db, &article)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, article)
}

//get Article
func (repository *ArticleRepo) GetArticles(context *gin.Context) {
	var article []models.Article
	err := models.GetArticles(repository.Db, &article)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, article)
}

//get article by id
func (repository *ArticleRepo) GetArticle(context *gin.Context) {
	id, _ := context.Params.Get("id")
	var article models.Article
	err := models.GetArticle(repository.Db, &article, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, article)
}

// update article
func (repository *ArticleRepo) UpdateArticle(context *gin.Context) {
	var article models.Article
	id, _ := context.Params.Get("id")
	err := models.GetArticle(repository.Db, &article, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.BindJSON(&article)
	err = models.UpdateArticle(repository.Db, &article)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, article)
}

// delete Article
func (repository *ArticleRepo) DeleteArticle(context *gin.Context) {
	var article models.Article
	id, _ := context.Params.Get("id")
	err := models.DeleteArticle(repository.Db, &article, id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "article deleted successfully"})
}
