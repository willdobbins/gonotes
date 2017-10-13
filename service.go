package notes


type Note struct {
	ID   uint  `db:"id" json:"id"`
	Body string `db:"body" json:"body"`
}

type NoteService interface {
	Note(id uint64) (*Note, error)

	Notes() (*[]Note, error)

	CreateNote(n *Note) (*Note, error)
	DeleteNote(id uint64) error
}
