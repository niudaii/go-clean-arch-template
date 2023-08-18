package entity

type User struct {
	BasicModel
	Username    string    `json:"username"`
	Password    string    `json:"-"`
	AuthorityID uint      `json:"authorityId"`
	Authority   Authority `json:"authority" gorm:"foreignKey:AuthorityID"`
}

type Authority struct {
	AuthorityID   uint   `json:"-" gorm:"not null;unique;primary_key;"`
	AuthorityName string `json:"authorityName"`
	Menus         []Menu `json:"-" gorm:"many2many:authority_menus"`
}

type Menu struct {
	ID        int    `json:"id" gorm:"not null;unique;primary_key"`
	ParentID  int    `json:"parentId"`
	Path      string `json:"path,omitempty"`
	Name      string `json:"name"`
	Redirect  string `json:"redirect,omitempty"`
	Component string `json:"component"`
	Meta      `json:"meta"`
}

type Meta struct {
	Title        string `json:"title"`
	Show         bool   `json:"show"`
	HideChildren bool   `json:"hideChildren,omitempty"`
}

func (User) TableName() string {
	return "system_user"
}

func (Authority) TableName() string {
	return "system_user_authority"
}

func (Menu) TableName() string {
	return "system_user_authority_menu"
}

type UserRepository interface {
	Create(user User) (err error)
	Update(user User) (err error)
	Delete(f *UserFilter) (err error)
	Find(f *UserFilter) (user User, err error)
	FindList(f *UserFilter) (list []User, total int64, err error)
	CheckExist(f *UserFilter) bool
}

type UserFilter struct {
	User
	PageInfo
}

type UserUsecase interface {
	FindByID(id uint) (user User, err error)
}

type Role struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}

type UserInfo struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	HomePath string `json:"homePath"`
	Roles    []Role `json:"roles"`
}
