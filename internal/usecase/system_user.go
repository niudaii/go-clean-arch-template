package usecase

import (
	"fmt"
	"go-clean-template/internal/entity"

	"github.com/niudaii/util/crypto"
)

type systemUserUsecase struct {
	userRepository entity.UserRepository
}

func NewSystemUserUsecase(userRepository entity.UserRepository) entity.SystemUserUsecase {
	return &systemUserUsecase{
		userRepository: userRepository,
	}
}

func (s systemUserUsecase) Create(user entity.CreateUser) (err error) {
	// 判断用户名是否已经存在
	f := &entity.UserFilter{}
	f.Username = user.Username
	if s.userRepository.CheckExist(f) {
		err = fmt.Errorf("用户名已经注册")
		return
	}
	// md5 加密
	newUser := entity.User{
		Username:    user.Username,
		Password:    crypto.Md5([]byte(user.Password)),
		AuthorityID: user.AuthorityID,
	}
	return s.userRepository.Create(newUser)
}

func (s systemUserUsecase) Delete(id uint) (err error) {
	f := &entity.UserFilter{}
	f.ID = id
	return s.userRepository.Delete(f)
}

func (s systemUserUsecase) FindList(pageInfo entity.PageInfo) (list []entity.User, total int64, err error) {
	f := &entity.UserFilter{}
	f.PageInfo = pageInfo
	return s.userRepository.FindList(f)
}
