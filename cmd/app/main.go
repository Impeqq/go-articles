package main

import (
	"gin/internal/handlers"
	"gin/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=password dbname=articlesdb port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if err := db.AutoMigrate(&models.Article{}); err != nil {
		panic("failed to migrate: " + err.Error())
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/articles", func(c *gin.Context) {
		articles, err := handlers.GetAllArticles(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, articles)
	})

	r.GET("/articles/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		article, err := handlers.GetArticle(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
			return
		}
		c.JSON(http.StatusOK, article)
	})

	r.POST("/articles", func(c *gin.Context) {
		var article models.Article
		if err := c.ShouldBindJSON(&article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := handlers.CreateArticle(db, &article); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, article)
	})

	r.PUT("/articles/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		article, err := handlers.GetArticle(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
			return
		}
		if err := c.ShouldBindJSON(article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := handlers.UpdateArticle(db, article); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, article)
	})

	r.DELETE("/articles/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		if err := handlers.DeleteArticle(db, uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	r.Run()
}
