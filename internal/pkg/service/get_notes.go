package service

import (
	"context"
	"fmt"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
)

// TODO : ADD VERIFICATION
func (s *Service) GetNotes(ctx context.Context, course string) (*[]model.Note, error) {
	c := model.Course{Name: course}
	c.CorrectCourseName()
	fmt.Println(c.Name)
	return s.db.GetNotes(ctx, c.Name)
}
