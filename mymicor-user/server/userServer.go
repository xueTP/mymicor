package server

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	pd "github.com/xueTP/gen-proto/mymicor-user"
	"golang.org/x/crypto/bcrypt"
	"mymicor/mymicor-user/data"
	"strings"
)

type userServer struct{
	userModel data.UserModeler
	TokenServer TokenServerInterface
}

func NewUserServer(conn *gorm.DB) userServer {
	return userServer{
		userModel: data.NewUserModel(conn),
		TokenServer: NewTokenServer(),
	}
}

func (this userServer) Create(ctx context.Context, user *pd.User, res *pd.Response) error {
	user.Password = strings.Trim(user.Password, " ")
	if user.Password == "" {
		return errors.New("user Password is nil")
	}
	// password 加密
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("userServer Create bcrypt.GenerateFromPassword error: %v, password: %s", err, user.Password)
		return err
	}
	user.Password = string(hashPwd)
	// user create
	if err := this.userModel.Add(user); err != nil {
		logrus.Errorf("userServer Create error: %v, param: %+v", err, user)
		return err
	}
	res.User = user
	return nil
}

func (this userServer) Get(ctx context.Context, user *pd.User, res *pd.Response) error {
	user, err := this.userModel.Get(user)
	if err != nil {
		logrus.Errorf("userServer Get error:%v", err)
		return err
	}
	res.User = user
	return nil
}

func (this userServer) GetAll(ctx context.Context, req *pd.Request, res *pd.Response) error {
	users := []*pd.User{}
	var err error
	if users, err = this.userModel.GetAll(&pd.User{}); err != nil {
		logrus.Errorf("userServer GetAll error:%v", err)
		return err
	}
	res.Users = users
	return nil
}

// Auth Auth user password and get token
func (this userServer) Auth(ctx context.Context, user *pd.User, token *pd.Token) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("user param is empty")
	}
	// get user info by user.Email
	userRes, err := this.userModel.Get(&pd.User{Email: user.Email})
	if err != nil {
		logrus.Errorf("userServer Auth error：%s, email: %s, password: %v", err, user.Email, user.Password)
		return err
	}
	// check password
	if err = bcrypt.CompareHashAndPassword([]byte(userRes.Password), []byte(user.Password)); err != nil {
		logrus.Errorf("userServer Auth bcrypt.CompareHashAndPassword error：%s, email: %s, password: %v", err, user.Email, user.Password)
		return err
	}
	// get token
	tk, err := this.TokenServer.Encode(userRes)
	if err != nil {
		logrus.Errorf("userServer Auth TokenServer.Encode error：%s, user: %+v", err, userRes)
		return err
	}
	token.Token = tk
	return nil
}

func (this userServer) ValidateToken(ctx context.Context, token *pd.Token, tokenRes *pd.Token) error {
	// TODO : 待完善user服务
	return nil
}

