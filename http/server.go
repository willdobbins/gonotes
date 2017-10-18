package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/willdobbins/notes"
)

type Server struct {
	Service notes.NoteService
}

func (s Server) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (s Server) ListNotes(c *gin.Context) {
	set, err := s.Service.Notes()
	if err != nil {
		log.Print(err)
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"results": set})
}

func (s Server) GetNote(c *gin.Context) {
	id := c.Param("id")
	idNumber, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not parse that ID"})
	}

	note, err := s.Service.Note(idNumber)

	if err != nil {
		log.Print(err)
	}
	c.HTML(http.StatusOK, "single.tmpl", note)
}

func (s Server) CreateNote(c *gin.Context) {
	var note notes.Note
	if c.Bind(&note) == nil {
		_, err := s.Service.CreateNote(&note)
		if err != nil {
			log.Print(err)
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/notes/")
}

func (s Server) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	idNumber, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Print(err)
	}
	err = s.Service.DeleteNote(idNumber)
	if err != nil {
		log.Print(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/notes/")
}
