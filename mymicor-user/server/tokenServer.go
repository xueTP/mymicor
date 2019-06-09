package server

import (
	pd "github.com/xueTP/gen-proto/mymicor-user"
)

type TokenServer struct{}

type TokenServerInterface interface {
	Encode(user *pd.User) (string, error)
	Decode(token string) (*pd.User, error)
}

// NewTokenServer New TokenServer Object
func NewTokenServer() TokenServer {
	return TokenServer{}
}

func (TokenServer) Encode(user *pd.User) (string, error) {
	// TODO: tokenServer encode
	panic("implement me")
}

func (TokenServer) Decode(token string) (*pd.User, error) {
	// TODO: tokenServer decode
	panic("implement me")
}