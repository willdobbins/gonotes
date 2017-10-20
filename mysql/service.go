package mysql

import (
	"errors"
	"fmt"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"

	"github.com/willdobbins/notes"
)

// Implements notes.NoteManager, holds a ConnectionURL describing the database it will connect to.
type Service struct {
	ConnectionUrl db.ConnectionURL
}

// Attempts to find a Note with the given ID.
func (service Service) One(id uint64) (*notes.Note, error) {

	session, collection, err := service.connect() //Get db connection and a collection
	if err != nil {
		return nil, err
	}
	defer session.Close() // At the end of this function, close DB session if it is not already closed.

	result := collection.Find("id", id) // Query the Notes table by id.
	exists, err := result.Exists() // Does it have non-zero results?
	if !exists {
		return nil, errors.New(fmt.Sprintf("could note find note %d", id))
	}
	if err != nil {
		return nil, err
	}

	note := notes.Note{} //Allocate space for a note to store the results of the query.

	err = result.One(&note) // Store the result of the query in the allocated note.
	return &note, err
}

// Finds and returns a slice of all Notes.
func (service Service) All() (*[]notes.Note, error) {
	session, collection, err := service.connect()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var result = new([]notes.Note)
	err = collection.Find().All(result)
	return result, err
}

// Persists a new Note to the database.
func (service Service) CreateNote(n *notes.Note) (*notes.Note, error) {
	session, collection, err := service.connect()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	id, err := collection.Insert(*n)
	if err != nil {
		parsedId, ok := id.(uint64)
		if ok {
			return service.One(parsedId)
		}
		err = errors.New(fmt.Sprintf("could not parse %+v as id", id))
	}
	return nil, err
}

//Given a note ID, will try to find it and then delete it.
func (service Service) DeleteNote(id uint64) error {
	session, collection, err := service.connect()
	if err != nil {
		return err
	}
	defer session.Close()

	res := collection.Find("id", id)
	exists, err := res.Exists()
	if err == nil {
		return err
	}
	if !exists {
		return errors.New(fmt.Sprintf("could note find note %d", id))
	}

	err = res.Delete()
	return err
}

//Update an existing Note record.
func (service Service) UpdateNote(id uint64, n *notes.Note) (*notes.Note, error) {
	session, collection, err := service.connect()
	if err != nil {
		return n, err
	}
	defer session.Close()

	oldNote, err  := service.One(id)
	if err != nil {
		return n, err
	}
	oldNote.Title = n.Title
	oldNote.Body = n.Body

	//Try to update this record.
	err = collection.UpdateReturning(oldNote)
	if err != nil {
		return n, err
	}

	return n, nil
}

//Establishes a connection to the specified MySQL database, gets a collection, and returns both.
//Please remember to Close() your db connections.
func (service Service) connect() (db.Database, db.Collection, error) {
	//Lowercase indicates that this method is not exported - it is a helper for other parts of the Service.

	session, err := mysql.Open(service.ConnectionUrl) // Connect to MySQL db
	if err != nil {
		return nil, nil, err
	}

	collection := session.Collection("notes") // Then we try to see if there is a table of Notes
	if !collection.Exists() {
		return nil, nil, errors.New("notes collection does not exist")
	}

	return session, collection, nil
}

//Factory method to make a mysql.Service
func New(connectionString string) (*Service, error) {
	service := new(Service)

	connectionSettings, err := mysql.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}
	service.ConnectionUrl = connectionSettings

	return service, nil
}

