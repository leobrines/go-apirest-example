package main

import "github.com/gin-gonic/gin"

type Webserver struct {
	server *gin.Engine
	controller *Controller
}
func (this *Webserver) start () {
	this.server = gin.Default()
	this.setRoutes()
	this.server.Run()
}
func (this *Webserver) setRoutes () {
	v1Group := this.server.Group("/v1")
	this.setItemRoutes(v1Group)
}
func (this *Webserver) setItemRoutes (group *gin.RouterGroup) {
	group.GET("/items", this.controller.getItems)
	group.POST("/item", this.controller.createItems)
}
