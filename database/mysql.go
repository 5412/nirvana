// Copyright (c) 2018. Weixiong-Piao, All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var Orm *xorm.Engine

func init() {
	var err error

	Orm, err = xorm.NewEngine("mysql", "root:root@tcp(mysql:3306)/hrms?charset=utf8")

	if err != nil {
		log.Fatal(err.Error())
	}

	Orm.ShowSQL(true)
}
