// Copyright (c) 2018. Weixiong-Piao, All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"nirvana/database"
	"nirvana/routers"
)

func main() {

	defer func() {
		err := database.Orm.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	router := routers.InitRouter()

	err := router.Run("80")

	if err != nil {
		log.Println(err)
	}

}
