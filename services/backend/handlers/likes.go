package handlers

import (
	"backend/dal"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func LikesArticle(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")

	likesArticle := "article:" + id + ":like"

	// **redis使用
	if err := dal.RedisClient.Incr(likesArticle).Err(); err != nil {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "failed to like article: " + err.Error(),
		})
		return
	}

	ctx.JSON(consts.StatusOK, utils.H{
		"message": "success liked",
	})
}

func GetArticleLikes(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")

	likesArticle := "article:" + id + ":like"

	likes, err := dal.RedisClient.Get(likesArticle).Result()
	if err != nil {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "Failed to get article likes: " + err.Error(),
		})
		return
	}

	ctx.JSON(consts.StatusOK, utils.H{
		"likes": likes,
	})
}
