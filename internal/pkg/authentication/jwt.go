package authentication

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/golang-jwt/jwt"
)

const tokenInfo string = "sign_in"

type JwtManager struct {
	SecretKey string
}

type Claimer struct {
	*jwt.StandardClaims
	TokenType string
}
type JwtSignUpCredentials struct {
	AccessToken  string
	RefreshToken string
}

func (jm JwtManager) JwtSignUpCredentialsCreator(user *model.User) (*JwtSignUpCredentials, error) {
	accesTokenExpire, refreshTokenExpire := getExpires()
	t := jwt.New(jwt.SigningMethodHS256)
	accessToken, err := tokenCreator(t, accesTokenExpire, user, jm.SecretKey)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("an unexpected error has occured, so sorry")
	}
	refreshToken, err := tokenCreator(t, refreshTokenExpire, user, jm.SecretKey)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("an unexpected error has occured, so sorry")
	}
	credentials := &JwtSignUpCredentials{AccessToken: accessToken, RefreshToken: refreshToken}
	return credentials, nil
}

// Will return the user object if token is verified, otherwise will return an indicating error
func (helper JwtManager) JwtCredentialsVerifier(token string) (*string, error) {
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(strings.ReplaceAll(token, "Bearer ", ""), &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(helper.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.Subject, nil
}

func tokenCreator(t *jwt.Token, expiresAt int64, user *model.User, SecretKey string) (string, error) {

	t.Claims = &Claimer{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Subject:   user.ID,
		},
		tokenInfo,
	}
	return t.SignedString([]byte(SecretKey))
}

func getExpires() (int64, int64) {
	accesTE := time.Now().Add(time.Hour * 72).Unix()
	refreshTE := time.Now().Add(time.Hour * 1400).Unix()
	return accesTE, refreshTE
}
