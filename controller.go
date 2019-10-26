package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	itemService *ItemService
}

func (this *Controller) getItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": this.itemService.getItems(),
	})
}
func (this *Controller) createItems(c *gin.Context) {
	fmt.Println("Creando items...")
	c.Status(http.StatusOK)
}
