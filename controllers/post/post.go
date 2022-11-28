package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordnr/login-register/database"
	"github.com/lordnr/login-register/models"
)

func NewPost(c *gin.Context) {
	var post models.Post
	if c.Bind(&post) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Failed to read body",
		})

		return
	}

	result := database.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error,
		})

		return
	}

	if result.RowsAffected > 0 {
		c.JSON(http.StatusCreated, gin.H{
			"status":  "ok",
			"message": "Post Create Success",
			"post":    post,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Post Create Failed",
		})
	}

}

func Posts(c *gin.Context) {
	var posts []models.Post
	result := database.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"post":   posts,
	})
}

func PostById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := database.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"post":   post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	c.Bind(&post)

	var updatePost models.Post
	postId := database.DB.First(&updatePost, id)
	if postId.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID does not exist",
		})

		return
	}

	result := database.DB.Model(&updatePost).Updates(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Update Success",
		"post":    updatePost,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error,
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
