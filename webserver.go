package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Webserver struct {
	srv        *http.Server
	router     *gin.Engine
	controller *Controller
	isRunning  bool
}

func (this *Webserver) start() {
	this.router = gin.Default()
	this.setRoutes()
	this.srv = &http.Server{
		Addr:    ":8080",
		Handler: this.router,
	}
	go this.listen()
	this.quit()
}
func (this *Webserver) listen() {
	if err := this.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	} else {
		this.isRunning = true
	}
}
func (this *Webserver) quit() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}
func (this *Webserver) stop() {
	if err := this.srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
	log.Println("Server exiting")
}
func (this *Webserver) setRoutes() {
	v1Group := this.router.Group("/v1")
	this.setItemRoutes(v1Group)
}
func (this *Webserver) setItemRoutes(group *gin.RouterGroup) {
	group.GET("/items", this.controller.getItems)
	group.POST("/item", this.controller.createItems)
}
