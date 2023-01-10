package service

import (
	"context"
	"errors"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/canergulay/bilgipedia/internal/pkg/utils"
	"github.com/google/uuid"
)

const USER_ALREADY_EXIST = "user already exist"

func (s *Service) SignUP(ctx context.Context, user *model.User) (*string, error) {
	u, _ := s.db.GetUser(ctx, user.Email)
	if u != nil {
		return nil, errors.New(USER_ALREADY_EXIST)
	}

	hashedPW, err := utils.HashMyPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPW
	user.ID = uuid.NewString()
	err = s.db.AddUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	credentials, err := s.jwt.JwtSignUpCredentialsCreator(user)
	if err != nil {
		return nil, err
	}

	return &credentials.AccessToken, nil
}
