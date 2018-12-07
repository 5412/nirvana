// Copyright (c) 2018. Weixiong-Piao, All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"nirvana/app/models"
)

func responseError(code int, err string, c *gin.Context) {
	resp := gin.H{"code":code, "error":err}
	c.JSON(code, resp)
	c.Abort()
}

func ApiTokenAuthMiddleware() gin.HandlerFunc {
	log.Println("binding ApiTokenAuthMiddleware")
	return func(context *gin.Context) {
		token := context.Request.FormValue("api_token")

		if token == "" {
			responseError(401, "Unauthorized:api_token required", context)
			return
		}

		user := models.GetUserByApiToken(token)

		if user.Id == 0 {
			responseError(401, "Unauthorized:error api_token", context)
			return
		}
		context.Next()
	}
}