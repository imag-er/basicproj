package handlers

import (
	"backend/dal"
	"backend/dal/models"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// ** redis旁路由模式
var articleKey = "article"

func CreateArticle(c context.Context, ctx *app.RequestContext) {
	var article models.Article

	if err := ctx.BindAndValidate(&article); err != nil {
		ctx.JSON(consts.StatusBadRequest, utils.H{
			"message": "Invalid Article data",
		})
		return
	}

	if err := dal.DB.Create(&article).Error; err != nil {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "Failed to create article",
		})
		return
	}

	// 清除缓存, 以便下次请求时重新加载数据
	dal.RedisClient.Del(articleKey)

	ctx.JSON(consts.StatusCreated, article)
}

func GetArticles(c context.Context, ctx *app.RequestContext) {
	var articles []models.Article
	// 如果从redis里能找到数据, 则返回
	// 否则从db查询, 序列化成json后存入redis
	cachedData, err := dal.RedisClient.Get(articleKey).Result()
	if err == redis.Nil {
		// 无缓存数据, 从DB读取, 并存入redis
		hlog.Info("DB")
		if err := dal.DB.Find(&articles).Error; err != nil {
			ctx.JSON(consts.StatusInternalServerError, utils.H{
				"message": "Failed to retrieve articles",
			})
			return
		}

		// 从DB读取成功后, 更新缓存
		data, err := json.Marshal(articles)
		if err != nil {
			ctx.JSON(consts.StatusInternalServerError, utils.H{
				"message": "Failed to serialize articles",
			})
			return
		}
		dal.RedisClient.Set(articleKey, data, 1*time.Minute)
		return
	} else {
		hlog.Info("REDIS")
		if err := json.Unmarshal([]byte(cachedData), &articles); err != nil {
			ctx.JSON(consts.StatusInternalServerError, utils.H{
				"message": "Failed to parse cached articles",
			})
			return
		}

		ctx.JSON(consts.StatusOK, articles)
	}

}

func GetArticlesById(c context.Context, ctx *app.RequestContext) {
	var article models.Article
	id := ctx.Param("id")

	// **细化错误处理
	if err := dal.DB.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(consts.StatusNotFound, utils.H{
				"message": "Article not found",
			})
			return
		}
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "Failed to retrieve article: " + err.Error(),
		})
		return
	}

	ctx.JSON(consts.StatusOK, article)
}
