package data

import (
	"github.com/jinzhu/gorm"
	pd "github.com/xueTP/gen-proto/mymicor-user"
	"strings"
)

type UserModeler interface {
	Add(user *pd.User) error
	Get(user *pd.User) (*pd.User, error)
	GetAll(user *pd.User) ([]*pd.User, error)
}

type userModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) userModel {
	return userModel{
		db: db.Model(pd.User{}),
	}
}

func (this userModel) Add(user *pd.User) error {
	user.Id = GetUUID()
	err := this.db.Create(user).Error
	return err
}

// getUserCondition 根据proto规定的user模型获取条件
func (this userModel) getUserCondition(user *pd.User) (string, []interface{}) {
	var where []string
	var value []interface{}
	if user.Id != "" {
		where = append(where, "Id = ?")
		value = append(value, user.Id)
	}
	if user.Email != "" {
		where = append(where, "Email = ?")
		value = append(value, user.Email)
	}
	return strings.Join(where, " and "), value
}

func (this userModel) Get(user *pd.User) (*pd.User, error) {
	where, value := this.getUserCondition(user)
	err := this.db.Where(where, value...).First(user).Error
	return user, err
}

func (this userModel) GetAll(user *pd.User) ([]*pd.User, error) {
	res := []*pd.User{}
	where, value := this.getUserCondition(user)
	err := this.db.Where(where, value...).Find(&res).Error
	return res, err
}