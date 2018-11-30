// Copyright (c) 2018. Weixiong-Piao, All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routers

import (
	"github.com/gin-gonic/gin"
	"nirvana/app/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//accouts := map[string]string{"solar": "solar"}
	//router.Use(gin.BasicAuth(accouts))
	// 初始化静态资源与网站小图标
	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	// 获取单条客户信息
	router.GET("/user/:id", controllers.GetUser)
	return router
}
