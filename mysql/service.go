package mysql

import (
	"errors"

	"upper.io/db.v3/mysql"
	"upper.io/db.v3"

	"github.com/willdobbins/notes"
)

type NoteService struct {
	collection db.Collection
}

func (ns NoteService) Note(id uint64) (*notes.Note, error) {
	var result = new(notes.Note)
	err := ns.collection.Find("id", id).One(result)
	return result, err
}

func (ns NoteService) Notes() (*[]notes.Note, error) {
	var result = new([]notes.Note)
	err := ns.collection.Find().All(result)
	return result, err
}

func (ns NoteService) CreateNote(n *notes.Note) (*notes.Note, error) {

	id, err:= ns.collection.Insert(*n)
	if err != nil {
		id, ok := id.(uint64)
		if ok  {
			return ns.Note(id)
		}
		err = errors.New("couldn't parse id")
	}
	return nil, err
}

func (ns NoteService) DeleteNote(id uint64) error {
	res := ns.collection.Find("id", id)
	err := res.Delete()
	return err
}

func NewService() (*NoteService, error) {// TODO pull settings out of here, to become arg.
	var service = new(NoteService)

	var settings = mysql.ConnectionURL{
		User: "root",
		Password: "trololol",
		Host: "db",
		Database: "notes",
	}

	sess, err := mysql.Open(settings)
	if err != nil {
		return service, err
	}

	coll := sess.Collection("notes")

	service.collection = coll
	return service, nil
}
