package service

import (
	"awesomeProject/utils"
	"errors"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/repository"
	uuid "github.com/satori/go.uuid"
)

type UserSrv interface {
	Get(user model.User) (*model.User, error)
	Exist(user model.User) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(id string) (bool, error)
}

type UserService struct {
	Repo repository.UserRepository
}

func (srv *UserService) Get(user model.User) (*model.User, error) {
	return srv.Repo.Get(user)
}

func (srv *UserService) Exist(user model.User) *model.User {
	return srv.Repo.Exist(user)
}

func (srv *UserService) Add(user model.User) (*model.User, error) {
	result := srv.Repo.ExistByMobile(user.Mobile)
	if result != nil {
		return nil, errors.New("用户已经存在")
	}
	user.UserId = uuid.NewV4().String()
	if user.Password == "" {
		user.Password = utils.Md5("123456")
	}
	user.IsDeleted = false
	user.IsLocked = false
	return srv.Repo.Add(user)
}

func (srv *UserService) Edit(user model.User) (bool, error) {
	if user.UserId == "" {
		return false, errors.New("参数错误")
	}
	exist := srv.Repo.ExistByUserID(user.UserId)
	if exist == nil {
		return false, errors.New("参数错误")
	}
	exist.NickName = user.NickName
	exist.Mobile = user.Mobile
	exist.Address = user.Address
	return srv.Repo.Edit(user)
}

func (srv *UserService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("参数错误")
	}
	user := srv.Repo.ExistByUserID(id)
	if user == nil {
		return false, errors.New("参数错误")
	}
	user.IsDeleted = !user.IsDeleted

	return srv.Repo.Delete(*user)
}
