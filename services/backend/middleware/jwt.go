package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"

	"backend/dal/models"
	"backend/handlers"
)

type login struct {
	Username string `form:"username,required" json:"username"`
	Password string `form:"password,required" json:"password"`
}

var identityKey = "id"

// ** JWT原理, 实际运用, 八股
func InitJWT() *jwt.HertzJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &models.User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: handlers.Login,
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			return true
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		hlog.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
