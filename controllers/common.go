package controllers

import (
	"github.com/Test214-dev/mooncascade-task/models"
	"github.com/gin-gonic/gin"
)

func setError(c *gin.Context, appError *models.AppError) {
	c.JSON(appError.Code, appError)
}
