package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/willdobbins/notes/mysql"
	"github.com/willdobbins/notes/http"
)

func main() {
	service, err := mysql.NewService()

	if err != nil {
		log.Fatal(err)
	}

	server := new(http.Server)
	server.Service = service

	router := gin.Default()
	router.GET("/health", server.Health)
	router.GET("/notes", server.ListNotes)
	router.POST("/notes", server.CreateNote)
	router.GET("/notes/:id", server.GetNote)
	router.DELETE("/notes/:id", server.DeleteNote)
	router.Run()
}
