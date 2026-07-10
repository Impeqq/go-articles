package handlers

import (
	"gin/internal/models"

	"gorm.io/gorm"
)

func GetAllArticles(db *gorm.DB) ([]models.Article, error) {
	var articles []models.Article
	if err := db.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func GetArticle(db *gorm.DB, id uint) (*models.Article, error) {
	var article models.Article
	if err := db.First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func CreateArticle(db *gorm.DB, input *models.Article) error {
	return db.Create(input).Error
}

func UpdateArticle(db *gorm.DB, input *models.Article) error {
	return db.Save(input).Error
}

func DeleteArticle(db *gorm.DB, id uint) error {
	return db.Delete(&models.Article{}, id).Error
}
