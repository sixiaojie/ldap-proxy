package controllers

import (
	"github.com/gin-gonic/gin"
)
//Ok 测试页面
func Ok(c *gin.Context) {
	c.JSON(200,gin.H{
		"status":"ok",
	})
}


//Test 测试页面
func Test(c *gin.Context) {
	c.JSON(200,gin.H{
		"status":"ok",
	})
}