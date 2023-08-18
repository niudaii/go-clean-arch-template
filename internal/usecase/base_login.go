package usecase

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/jwt"
)

type loginUsecase struct {
	userRepository entity.UserRepository
}

func NewLoginUsecase(userRepository entity.UserRepository) entity.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
	}
}

func (l loginUsecase) GetUserByUsername(username string) (user entity.User, err error) {
	f := &entity.UserFilter{}
	f.Username = username
	return l.userRepository.Find(f)
}

func (l loginUsecase) CreateAccessToken(user *entity.User) (token string, err error) {
	claims := jwt.CreateClaims(jwt.BaseClaims{
		UserID:      user.ID,
		AuthorityID: user.AuthorityID,
		Username:    user.Username,
	})
	token, err = jwt.CreateToken(claims)
	return
}
