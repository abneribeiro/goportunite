package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	// show opening 
	v1.GET("/opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "GET oppening",})
	})

	v1.POST("/opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "POST oppening",})
	})

	v1.DELETE("/opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "GET oppening",})
	})

	v1.PUT("/opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "GET oppening",})
	})

	v1.GET("/openings", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "GET oppenings",})
	})
}
