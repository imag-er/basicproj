package handlers

import (
	"backend/dal/models"
	"context"

	"backend/dal"
	myutils "backend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c context.Context, ctx *app.RequestContext) {
	user := models.User{}

	if err := ctx.BindAndValidate(&user); err != nil {
		ctx.JSON(consts.StatusBadRequest, utils.H{
			"error": err.Error(),
		})
		return
	}
	hlog.Infof("register info: %+v", user)
	hashedPassword, err := myutils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"error": "Failed to hash password",
		})
		return
	}

	user.Password = hashedPassword

	if err := dal.DB.Create(&user).Error; err != nil {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"error": "Failed to register user",
		})
		return
	}

	ctx.JSON(consts.StatusCreated, utils.H{
		"message": "successfully registered",
	})
}

func Login(c context.Context, ctx *app.RequestContext) (interface{}, error) {
	var input struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}
	if err := ctx.BindAndValidate(&input); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	// Check user credentials
	var user models.User
	if err := dal.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	
	// **校验密码的地方不能直接判断hash值相等
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return &user, nil
}
