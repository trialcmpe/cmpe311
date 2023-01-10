package service

import (
	"context"
	"errors"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/canergulay/bilgipedia/internal/pkg/utils"
)

func (s *Service) SignIN(ctx context.Context, u *model.User) (*string, error) {
	user, err := s.db.GetUser(ctx, u.Email)
	if err != nil {
		return nil, err
	}

	if !utils.ComparePasswords(user.Password, u.Password) {
		return nil, errors.New("PASSWORD IS WRONG")
	}

	credentials, err := s.jwt.JwtSignUpCredentialsCreator(user)
	if err != nil {
		return nil, err
	}

	return &credentials.AccessToken, err
}
