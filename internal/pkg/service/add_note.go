package service

import (
	"context"
	"fmt"
	"time"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/canergulay/bilgipedia/internal/pkg/utils"
	"github.com/google/uuid"
)

// TODO : ADD VERIFICATION
func (s *Service) AddNote(ctx context.Context, note *model.Note) error {
	userID := utils.GetUserID(ctx)
	fmt.Println(userID)
	note.Creator = userID
	note.CreatedAt = time.Now()
	note.ID = uuid.NewString()
	return s.db.AddNote(ctx, *note)
}
