package service

import (
	"context"
	"errors"
	"time"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/google/uuid"
)

const COURSE_ALREADY_EXIST = "course already exist"

func (s *Service) AddCourse(ctx context.Context, course *model.Course) error {
	course.CorrectCourseName()
	c, _ := s.db.GetCourse(ctx, course.Name)
	if c != nil {
		return errors.New(COURSE_ALREADY_EXIST)
	}
	course.ID = uuid.NewString()
	course.CreatedAt = time.Now()
	return s.db.AddCourse(ctx, *course)
}
