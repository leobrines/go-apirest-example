package main

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Webserver struct {
	server *gin.Engine
}
func (this *Webserver) Configure () {
	this.server = gin.Default()
	this.setRoutes()
}
func (this *Webserver) Start () {
	this.server.Run()
}
func (this *Webserver) setRoutes () {
	v1Group := this.server.Group("/v1")
	this.setItemRoutes(v1Group)
}
func (this *Webserver) setItemRoutes (group *gin.RouterGroup) {
	group.GET("/items", getItemsCtrl)
	group.POST("/item", createItemsCtrl)
}

func getItemsCtrl (c *gin.Context) {
	fmt.Println("Obteniendo items...")
	c.Status(http.StatusOK)
}
func createItemsCtrl (c *gin.Context) {
	fmt.Println("Creando items...")
	c.Status(http.StatusOK)
}

