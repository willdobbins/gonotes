package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/willdobbins/notes/http"
	"github.com/willdobbins/notes/mysql"
)

func main() {
	connectionString := "notes:cr0wdst4r@tcp(db:3306)/notes" // TODO - pull this from a config.

	service, err := mysql.New(connectionString) // Makes a new mysql.Service (after trying to parse connectionString
	if err != nil {
		log.Fatal("fatal: " + err.Error())
	}

	//Sets up our controller with a copy of the mysql.Service
	server := http.Server{Service: service}

	//Defines routes and handlers for our app & then starts the router.
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("./templates/*")
	router.GET("/health", server.Health)
	router.GET("/notes", server.ListNotes)
	router.POST("/notes", server.CreateNote)
	router.GET("/notes/:id", server.GetNote)
	router.POST("/notes/:id", server.UpdateNote)
	router.DELETE("/notes/:id", server.DeleteNote)
	router.Run()
}
