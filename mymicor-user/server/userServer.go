package server

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	pd "github.com/xueTP/gen-proto/mymicor-user"
	"mymicor/mymicor-user/data"
)

type userServer struct{
	userModel data.UserModeler
}

func NewUserServer(conn *gorm.DB) userServer {
	return userServer{
		userModel: data.NewUserModel(conn),
	}
}

func (this userServer) Create(ctx context.Context, user *pd.User, res *pd.Response) error {
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

func (this userServer) Auth(ctx context.Context, user *pd.User, token *pd.Token) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("user param is empty")
	}
	userRes, err := this.userModel.Get(&pd.User{Email: user.Email, Password: user.Password})
	if err != nil {
		logrus.Errorf("userServer Auth error：%s, email: %s, password: %v", err, user.Email, user.Password)
	}
	if userRes.Id != "" {
		token.Token = "testing abc"
	}
	return nil
}

func (this userServer) ValidateToken(ctx context.Context, token *pd.Token, tokenRes *pd.Token) error {
	// TODO : 待完善user服务
	return nil
}

