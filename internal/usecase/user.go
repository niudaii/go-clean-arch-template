package usecase

import "go-clean-template/internal/entity"

type userUsecase struct {
	userRepository entity.UserRepository
}

func NewUserUsecase(userRepository entity.UserRepository) entity.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u userUsecase) FindByID(id uint) (user entity.User, err error) {
	f := &entity.UserFilter{}
	f.ID = id
	return u.userRepository.Find(f)
}
