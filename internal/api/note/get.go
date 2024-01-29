package note

import (
	"context"
	"github.com/grishagavrin/auth-chat-grpc/internal/converter"
	desc "github.com/grishagavrin/auth-chat-grpc/pkg/note_v1"
	"github.com/opentracing/opentracing-go"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	getId := req.GetId()

	span, ctx := opentracing.StartSpanFromContext(ctx, "get note")
	defer span.Finish()

	span.SetTag("id", getId)

	noteObj, err := i.noteService.Get(ctx, getId)
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, title: %s, body: %s, created_at: %v, updated_at: %v\n", noteObj.ID, noteObj.Info.Title,
		noteObj.Info.Content, noteObj.CreatedAt, noteObj.UpdatedAt)

	return &desc.GetResponse{
		Note: converter.ToNoteFromService(noteObj),
	}, nil
}
