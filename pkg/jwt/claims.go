package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type BaseClaims struct {
	UserID      uint
	AuthorityID uint
	Username    string
}

func getClaims(c *gin.Context) (claims *CustomClaims, err error) {
	token := c.Request.Header.Get(config.HeaderName)
	claims, err = ParseToken(token)
	return claims, err
}

func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := getClaims(c); err != nil {
			return 0
		} else {
			return cl.UserID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.UserID
	}
}

func GetAuthorityID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := getClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.AuthorityID
	}
}

func GetUsername(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := getClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.Username
	}
}
