package http

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/willdobbins/notes"
)

type Server struct {
	Service notes.NoteService
}

func (s Server) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func (s Server) ListNotes(c *gin.Context) {
	set, err := s.Service.Notes()
	if err != nil {
		log.Print(err)
	}
	c.JSON(200, set)
}

func (s Server) GetNote(c *gin.Context) {
	id := c.Param("id")
	idNumber, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not parse that ID"})
	}

	note, err := s.Service.Note(idNumber)

	if err != nil {
		log.Print(err)
	}

	c.JSON(200, note)
}

func (s Server) CreateNote(c *gin.Context) {
	var add = new(notes.Note)
	add.Body = c.PostForm("body")

	note, err := s.Service.CreateNote(add)
	if err != nil {
		c.JSON(500, gin.H{"message": "CreateNote fail"})
	}
	c.JSON(200, note)
}

func (s Server) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	idNumber, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"message":"delete: could not parse ID"})
	}
	err = s.Service.DeleteNote(idNumber)
	if err != nil {
		c.JSON(500, gin.H{"message":"delete service fail"})
	}

	c.JSON(200, gin.H{"message": "OK"})
}
