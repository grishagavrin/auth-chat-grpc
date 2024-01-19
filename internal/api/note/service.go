package note

import (
	"github.com/grishagavrin/auth-chat-grpc/internal/service"
	desc "github.com/grishagavrin/auth-chat-grpc/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
