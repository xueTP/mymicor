package server

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	pd "github.com/xueTP/gen-proto/mymicor-user"
	"time"
)

type TokenServer struct{}

type TokenServerInterface interface {
	Encode(user *pd.User) (string, error)
	Decode(token string) (*pd.User, error)
}

const salt = "ADSf34434Awqef23@#@$"

// NewTokenServer New TokenServer Object
func NewTokenServer() TokenServer {
	return TokenServer{}
}

type tokenContent struct {
	user *pd.User
	jwt.StandardClaims
}

func (TokenServer) Encode(user *pd.User) (string, error) {
	expiresTime := time.Now().Add(72 * time.Hour).Unix()
	tokenContent := tokenContent{
		user,
		jwt.StandardClaims{
			ExpiresAt: expiresTime,
			Issuer: "go.micro.srv.user",
		},
	}
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, tokenContent)
	return token.SignedString(salt)
}

func (TokenServer) Decode(token string) (*pd.User, error) {
	// parse Token
	tokenDn, err := jwt.ParseWithClaims(token, &tokenContent{}, func(token *jwt.Token) (i interface{}, e error) {
		return salt, nil
	})
	if err != nil {
		return &pd.User{}, err
	}
	if tokenContent, ok := tokenDn.Claims.(*tokenContent); ok {
		return tokenContent.user, nil
	}
	return &pd.User{}, errors.New("token content is false")
}
