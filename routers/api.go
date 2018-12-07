// Copyright (c) 2018. Weixiong-Piao, All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routers

import (
	"nirvana/app/controllers"
	"nirvana/middlewares"
)

func initApi()  {
	apiAuthGroup := Router.Group("/api").Use(middlewares.ApiTokenAuthMiddleware())
	apiAuthGroup.GET("/user/:id", controllers.GetUser)
}
