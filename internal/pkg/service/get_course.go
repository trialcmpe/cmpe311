package service

import (
	"context"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
)

func (s *Service) GetCourse(ctx context.Context, courseName string) (*model.Course, error) {
	return s.db.GetCourse(ctx, courseName)
}

func (s *Service) GetCourses(ctx context.Context) (*[]model.Course, error) {
	return s.db.GetCourses(ctx)
}
