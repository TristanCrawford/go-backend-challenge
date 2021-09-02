package controllers

import (
	"github.com/gin-gonic/gin"
	"backend.com/m/v2/models"
)

func (a *API) GetPerson(c *gin.Context) {
	var person models.Person

	id := c.Params.ByName("id")
	result := a.DB.Where("id = ?", id).First(&person)

	if result.Error != nil {
		c.JSON(200, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"person": person,
	})
}
