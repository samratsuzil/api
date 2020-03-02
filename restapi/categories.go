package restapi

import (

	"net/http"
	"strconv"
	"github.com/samratsuzil/api/models"
	"github.com/gin-gonic/gin"
)

func getAllCategories(c *gin.Context) {
	var category []models.Category
	err := models.GetAllCategories(&category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, category)
}

func getCategory(c *gin.Context) {

	var category models.Category
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetCategory(&category, uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, category)

}

func addNewCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err := models.AddNewCategory(&category)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.GetCategory(&category, category.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, category)

}

func updateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err := models.UpdateCategory(&category, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := models.GetCategory(&category, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &category)

}

func deleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "category has been deleted",
	})
}
