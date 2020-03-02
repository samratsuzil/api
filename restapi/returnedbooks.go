package restapi

import (

	"net/http"
	"strconv"
	"github.com/samratsuzil/api/models"
	"github.com/gin-gonic/gin"
)

func getAllReturnedBooks(c *gin.Context) {
	//restapi for returnedbookdbooks
	var returnedbook []models.ReturnedBook
	err := models.GetAllReturnedBooks(&returnedbook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, returnedbook)
}

func getReturnedBook(c *gin.Context) {

	var returnedbook models.ReturnedBook
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetReturnedBook(&returnedbook, uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, returnedbook)

}

func addNewReturnedBook(c *gin.Context) {
	var returnedbook models.ReturnedBook
	if err := c.ShouldBindJSON(&returnedbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err := models.AddNewReturnedBook(&returnedbook)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.GetReturnedBook(&returnedbook, returnedbook.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, returnedbook)

}

func updateReturnedBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var returnedbook models.ReturnedBook
	if err := c.ShouldBindJSON(&returnedbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err := models.UpdateReturnedBook(&returnedbook, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := models.GetReturnedBook(&returnedbook, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &returnedbook)

}

func deleteReturnedBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteReturnedBook(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "returnedbook has been deleted",
	})
}
