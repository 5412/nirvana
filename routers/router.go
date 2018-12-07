// Copyright (c) 2018. Weixiong-Piao, All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routers

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	// 初始化静态资源与网站小图标
	Router.Static("/assets", "./assets")
	Router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// 初始化Api接口路由
	initApi()

	// 初始化web路由
	initWeb()
}
