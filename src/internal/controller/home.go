package controller

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"

	"github.com/nukahaha/car_store/src/internal/configuration"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (hc *HomeController) GetHome(ctx *gin.Context) {
	ginview.HTML(ctx, http.StatusOK, "home", gin.H{
		"isAuthenticated": configuration.IsAuthenticated(ctx),
	})
}
