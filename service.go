package notes

//Note is a simple text post consisting of a title and body, referenced by ID.
type Note struct {
	ID    uint   `db:"id" json:"id"`
	Title string `db:"title" json:"title" form:"title" binding:"required"`
	Body  string `db:"body" json:"body" form:"body" binding:"required"`
}

//NoteService defines basic interface for CRUD operations on Notes.
type NoteService interface {
	Note(id uint64) (*Note, error)
	Notes() (*[]Note, error)
	CreateNote(n *Note) (*Note, error)
	DeleteNote(id uint64) error
}
