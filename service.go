package notes

//Note is a simple text post consisting of a title and body, referenced by ID.
type Note struct {
	ID    uint   `db:"id,omitempty" json:"id"`
	Title string `db:"title" json:"title" form:"title" binding:"required"`
	Body  string `db:"body" json:"body" form:"body" binding:"required"`
}

//NoteManager defines basic interface for CRUD operations on Notes.
type NoteManager interface {
	Note(id uint64) (*Note, error)
	All() (*[]Note, error)
	CreateNote(n *Note) (*Note, error)
	DeleteNote(id uint64) error
}
