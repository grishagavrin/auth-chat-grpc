package note

import (
	"github.com/grishagavrin/auth-chat-grpc/internal/client/db"
	"github.com/grishagavrin/auth-chat-grpc/internal/repository"
	"github.com/grishagavrin/auth-chat-grpc/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
	txManager      db.TxManager
}

func NewService(
	noteRepository repository.NoteRepository,
	txManager db.TxManager,
) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}
