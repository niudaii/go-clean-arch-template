package v1

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/jwt"
	"go-clean-template/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	UserUsecase entity.UserUsecase
}

const (
	LogoutSuccess = "注销成功"
)

func (a UserApi) GetInfo(c *gin.Context) {
	userID := jwt.GetUserID(c)
	user, err := a.UserUsecase.FindByID(userID)
	if err != nil {
		response.ErrorWithMessage(FindFail, err, c)
	} else {
		role := entity.Role{
			RoleName: user.Authority.AuthorityName,
			Value:    user.Authority.AuthorityName,
		}
		data := entity.UserInfo{
			UserID:   user.ID,
			Username: user.Username,
			Avatar:   "https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640",
			HomePath: "/task/index",
			Roles: []entity.Role{
				role,
			},
		}
		response.Ok(data, FindSuccess, c)
	}
}

func (a UserApi) Logout(c *gin.Context) {
	response.OkWithMessage(LogoutSuccess, c)
}
