package handlers

import (
	"gin/internal/models"

	"gorm.io/gorm"
)

// GetAllArticles возвращает список статей
// @Summary      Получить все статьи
// @Description  Возвращает массив статей
// @Tags         articles
// @Produce      json
// @Success      200  {array}   models.Article
// @Failure      500  {object}  map[string]string
// @Router       /articles [get]
func GetAllArticles(db *gorm.DB) ([]models.Article, error) {
	var articles []models.Article
	if err := db.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticle возвращает одну статью по ID
// @Summary      Получить статью по ID
// @Description  Возвращает одну статью
// @Tags         articles
// @Produce      json
// @Param        id   path      int  true  "Article ID"
// @Success      200  {object}  models.Article
// @Failure      404  {object}  map[string]string
// @Router       /articles/{id} [get]
func GetArticle(db *gorm.DB, id uint) (*models.Article, error) {
	var article models.Article
	if err := db.First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// CreateArticle создает новую статью
// @Summary      Создать статью
// @Description  Создает новую статью
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        article  body      models.Article  true  "Article object"
// @Success      201      {object}  models.Article
// @Failure      400      {object}  map[string]string
// @Router       /articles [post]
func CreateArticle(db *gorm.DB, input *models.Article) error {
	return db.Create(input).Error
}

// UpdateArticle обновляет существующую статью
// @Summary      Обновить статью
// @Description  Обновляет статью по её ID
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Article ID"
// @Param        article  body      models.Article  true  "Article object"
// @Success      200      {object}  models.Article
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /articles/{id} [put]
func UpdateArticle(db *gorm.DB, input *models.Article) error {
	return db.Save(input).Error
}

// DeleteArticle удаляет статью по ID
// @Summary      Удалить статью
// @Description  Удаляет статью по её ID (soft delete)
// @Tags         articles
// @Produce      json
// @Param        id   path  int  true  "Article ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /articles/{id} [delete]
func DeleteArticle(db *gorm.DB, id uint) error {
	return db.Delete(&models.Article{}, id).Error
}
