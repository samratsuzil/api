package restapi

import (

	"net/http"
	"strconv"
	"github.com/samratsuzil/api/models"
	"github.com/gin-gonic/gin"
)

func getAllPublications(c *gin.Context) {
	//restapi for publications
	var publication []models.Publication
	err := models.GetAllPublications(&publication)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, publication)
}

func getPublication(c *gin.Context) {

	var publication models.Publication
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetPublication(&publication, uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, publication)

}

func addNewPublication(c *gin.Context) {
	var publication models.Publication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err := models.AddNewPublication(&publication)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.GetPublication(&publication, publication.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, publication)

}

func updatePublication(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var publication models.Publication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err := models.UpdatePublication(&publication, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := models.GetPublication(&publication, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &publication)

}

func deletePublication(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeletePublication(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "publication has been deleted",
	})
}
