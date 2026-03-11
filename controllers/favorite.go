package controllers

import (
	"net/http"
	"zhix-backend/config"
	"zhix-backend/models"

	"github.com/gin-gonic/gin"
)

func AddFavorite(c *gin.Context) {
	articleID := c.Param("id")
	userID, _ := c.Get("userId")

	// 检查是否存在软删除的记录
	var existing models.Favorite
	if err := config.DB.Unscoped().Where("user_id = ? AND article_id = ?", userID, parseUint(articleID)).First(&existing).Error; err == nil {
		// 存在记录，恢复它
		config.DB.Unscoped().Model(&existing).Update("deleted_at", nil)
		c.JSON(http.StatusOK, gin.H{"message": "Added to favorites"})
		return
	}

	// 不存在记录，创建新的
	favorite := models.Favorite{
		UserID:    userID.(uint),
		ArticleID: parseUint(articleID),
	}

	if err := config.DB.Create(&favorite).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already favorited"})
		return
	}

	// 增加用户收藏统计
	config.DB.Model(&models.User{}).Where("id = ?", userID).Update("favorite_count", config.DB.Raw("favorite_count + 1"))

	// 增加作者被收藏统计
	var article models.Article
	if err := config.DB.First(&article, parseUint(articleID)).Error; err == nil {
		config.DB.Model(&models.User{}).Where("id = ?", article.AuthorID).Update("total_favorited", config.DB.Raw("total_favorited + 1"))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Added to favorites"})
}

func RemoveFavorite(c *gin.Context) {
	articleID := c.Param("id")
	userID, _ := c.Get("userId")

	if err := config.DB.Where("user_id = ? AND article_id = ?", userID, parseUint(articleID)).Delete(&models.Favorite{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 减少用户收藏统计
	config.DB.Model(&models.User{}).Where("id = ? AND favorite_count > 0", userID).Update("favorite_count", config.DB.Raw("favorite_count - 1"))

	// 减少作者被收藏统计
	var article models.Article
	if err := config.DB.First(&article, parseUint(articleID)).Error; err == nil {
		config.DB.Model(&models.User{}).Where("id = ? AND total_favorited > 0", article.AuthorID).Update("total_favorited", config.DB.Raw("total_favorited - 1"))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Removed from favorites"})
}

func GetFavorites(c *gin.Context) {
	userID, _ := c.Get("userId")

	var favorites []models.Favorite
	if err := config.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&favorites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var articleIDs []uint
	for _, fav := range favorites {
		articleIDs = append(articleIDs, fav.ArticleID)
	}

	var articles []models.Article
	if len(articleIDs) > 0 {
		config.DB.Where("id IN ?", articleIDs).Find(&articles)
	}

	c.JSON(http.StatusOK, articles)
}

func CheckFavorite(c *gin.Context) {
	articleID := c.Param("id")
	userID, _ := c.Get("userId")

	var count int64
	config.DB.Model(&models.Favorite{}).Where("user_id = ? AND article_id = ?", userID, parseUint(articleID)).Count(&count)

	c.JSON(http.StatusOK, gin.H{"isFavorited": count > 0})
}

func parseUint(s string) uint {
	var id uint
	for _, c := range s {
		id = id*10 + uint(c-'0')
	}
	return id
}
