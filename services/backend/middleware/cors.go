package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
)

func CORS() app.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"localhost"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

}
