// Copyright (c) 2018. Weixiong-Piao, All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"nirvana/app/models"
	"strconv"
)

func GetUser(c *gin.Context) {
	tmpId := c.Param("id")
	id, err := strconv.Atoi(tmpId)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "格式错误",
		})
	} else {
		user := models.GetUserById(id)
		if user == nil {
			c.JSON(200, gin.H{
				"success": true,
				"msg":     "",
			})
		} else {
			c.JSON(200, gin.H{
				"success": true,
				"msg":     user,
			})
		}
	}
}
