package note

import (
	"github.com/grishagavrin/auth-chat-grpc/internal/repository"
	"github.com/grishagavrin/auth-chat-grpc/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
}

func NewService(
	noteRepository repository.NoteRepository,
) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
	}
}

func NewMockService(deps ...interface{}) service.NoteService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.NoteRepository:
			srv.noteRepository = s
		}
	}

	return &srv
}
