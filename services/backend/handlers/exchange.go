package handlers

import (
	"context"
	"time"

	"backend/dal"
	"backend/dal/models"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func CreateExchangeRate(c context.Context, ctx *app.RequestContext) {
	var exchangeRate models.ExchangeRate

	if err := ctx.BindAndValidate(&exchangeRate); err != nil {
		ctx.JSON(consts.StatusBadRequest, utils.H{
			"message": "Invalid exchange rate data",
		})
		return
	}

	exchangeRate.Date = time.Now()
	hlog.Warnf("Creating exchange rate: %+v", exchangeRate)
	if err := dal.DB.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(consts.StatusInternalServerError, err)
		return
	}

	ctx.JSON(consts.StatusCreated, exchangeRate)
}

func GetExchangeRates(c context.Context, ctx *app.RequestContext) {
	var exchangeRates []models.ExchangeRate
	if err := dal.DB.Find(&exchangeRates).Error; err != nil {
		ctx.JSON(consts.StatusInternalServerError, err)
		return
	}

	ctx.JSON(consts.StatusOK, exchangeRates)
}
