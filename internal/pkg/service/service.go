package service

import (
	"context"

	"github.com/canergulay/bilgipedia/internal/pkg/authentication"
	"github.com/canergulay/bilgipedia/internal/pkg/model"
)

type IDBManager interface {
	AddUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, email string) (*model.User, error)
	GetNotes(ctx context.Context, course string) (*[]model.Note, error)
	AddNote(ctx context.Context, note model.Note) error
	GetCourses(ctx context.Context) (*[]model.Course, error)
	GetCourse(ctx context.Context, courseName string) (*model.Course, error)
	AddCourse(ctx context.Context, course model.Course) error
}

type Service struct {
	db  IDBManager
	jwt *authentication.JwtManager
}

func NewService(db IDBManager, auth authentication.JwtManager) *Service {
	return &Service{
		db:  db,
		jwt: &auth,
	}
}
