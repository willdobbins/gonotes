package main

import (
	"github.com/gin-gonic/gin"
	"upper.io/db.v3/mysql"
	"log"
	"upper.io/db.v3/lib/sqlbuilder"
)

type Server struct {
	session sqlbuilder.Database
}

func (s Server) pingAction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}

func (s Server) listNotesAction(c *gin.Context) {
	var notes = []Note{}
	err := s.session.Collection("notes").Find().All(&notes)
	if err != nil {
		log.Print(err)
	}
	c.JSON(200, notes)
}

func (s Server) getNoteAction(c *gin.Context) {
	id := c.Param("id")
	note := Note{}
	err := s.session.Collection("notes").Find("id", id).One(&note)
	if err != nil {
		log.Print(err)
	}

	c.JSON(200, note)
}

func (s Server) createNoteAction(c *gin.Context) {

}

func (s Server) deleteNoteAction(c *gin.Context) {

}

func main() {
	var settings = mysql.ConnectionURL{
		User: "root",
		Password: "trololol",
		Host: "db",
		Database: "notes",
	}

	sess, err := mysql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	server := &Server{session: sess}

	router := gin.Default()
	router.GET("/ping", server.pingAction)
	router.GET("/notes", server.listNotesAction)
	router.POST("/notes", server.createNoteAction)
	router.GET("/notes/:id", server.getNoteAction)
	router.DELETE("/notes/:id", server.deleteNoteAction)
	router.Run()
}


type Note struct {
	ID   uint  `db:"id" json:"id"`
	Body string `db:"body" json:"body"`
}
