package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willdobbins/notes/http"
	"github.com/willdobbins/notes/mysql"
	"log"
)

func main() {
	//Creates a new instance of a mysql implementation of the NoteService.
	service, err := mysql.NewService()

	//if it fails to build (error connecting, etc, dies)
	if err != nil {
		log.Fatal(err)
	}

	//Sets up a Server - not sure about name, functionally it's a controller
	//with a dependency on NoteService.
	server := new(http.Server)
	server.Service = service

	//Defines routes for
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("./templates/*")
	router.GET("/health", server.Health)
	router.GET("/notes", server.ListNotes)
	router.POST("/notes", server.CreateNote)
	router.GET("/notes/:id", server.GetNote)
	router.DELETE("/notes/:id", server.DeleteNote)
	router.Run()
}
