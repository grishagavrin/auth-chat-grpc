package repository

import (
	"context"
	"github.com/grishagavrin/auth-chat-grpc/internal/model"
)

type NoteRepository interface {
	Create(ctx context.Context, info *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
}
