package entity

type SystemUserUsecase interface {
	Create(user CreateUser) (err error)
	Delete(id uint) (err error)
	FindList(pageInfo PageInfo) (list []User, total int64, err error)
}

type CreateUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	AuthorityID uint   `json:"authorityId"`
}

type DeleteUser struct {
	ID uint `json:"id"`
}
