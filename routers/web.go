// Copyright (c) 2018. Weixiong-Piao, All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nirvana/app/controllers"
)

func initWeb()  {
	// 获取单条客户信息
	Router.GET("/user/:id",controllers.GetUser)
	Router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "layout.html", gin.H{
			"title": "Nirvana CRM",
			"smallTitle": "NCRM",
		})
	})

}
