package note

import (
	"context"
	"github.com/grishagavrin/auth-chat-grpc/internal/converter"
	desc "github.com/grishagavrin/auth-chat-grpc/pkg/note_v1"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	id, err := i.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
