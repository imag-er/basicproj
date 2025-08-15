package router

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"backend/config"
	"backend/handlers"
	"backend/middleware"
	"fmt"
)

func InitRouter() (h *server.Hertz) {
	h = server.Default(
		server.WithHostPorts(fmt.Sprintf("%s:%d", config.Config.App.Host, config.Config.App.Port)),
	)

	h.Use(middleware.CORS())

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	authMW := middleware.InitJWT() // Initialize JWT middleware
	auth := h.Group("/api/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", authMW.LoginHandler)
		auth.POST("/logout", authMW.LogoutHandler)
	}

	api := h.Group("/api")
	api.GET("/exchange-rates", handlers.GetExchangeRates)

	api.GET("/articles/:id", handlers.GetArticlesById)
	api.GET("/articles", handlers.GetArticles)
	api.POST("/articles", handlers.CreateArticle)

	api.GET("/likes/:id", handlers.GetArticleLikes)
	api.POST("/likes/:id", handlers.LikesArticle)

	api.Use(authMW.MiddlewareFunc()) // Apply JWT middleware to the API group
	{
		api.POST("/exchange-rates", handlers.CreateExchangeRate)
	}

	return
}
