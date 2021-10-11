package middleware

import (
	"github.com/Jhonnycpp/lds-0221-common/cmd/constant"
	"github.com/Jhonnycpp/lds-0221-common/cmd/errors"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	service iAuthService
}

func NewAuthMiddleware(service iAuthService) *AuthMiddleware {
	return &AuthMiddleware{
		service: service,
	}
}

func (a *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			errors.SendErrorRespose(c, errors.UnauthorizedError)
			return
		}

		token := header[len(Bearer_schema):]

		if email, err := a.service.ValidateToken(token); err == nil {
			c.Set(constant.EmailParam, email)
			return
		}

		errors.SendErrorRespose(c, errors.UnauthorizedError)
	}
}

func (a *AuthMiddleware) GetEmail(context *gin.Context) (email string, err error) {
	if value, exists := context.Get(constant.EmailParam); exists {
		return value.(string), nil
	}
	return "", errors.InvalidAuth
}
