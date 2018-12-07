// Copyright (c) 2018. Weixiong-Piao, All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import (
	"log"
	"nirvana/database"
	"time"
)

type User struct {
	Id        int       `json:"id" form:"id"`
	Nickname  string    `json:"nickname" form:"nickname"`
	Email     string    `json:"email" form:"email"`
	Tel       string    `json:"tel" form:"tel"`
	Account   string    `json:"account" form:"account"`
	Password  string    `json:"password" form:"password"`
	ApiToken string  	`json:"api_token" form:"api_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUserById(id int) (u *User) {
	user := &User{Id: id}
	is, _ := database.Orm.Get(user)
	if !is {
		log.Println("can't find user by id:", id)
	}
	return user
}

func GetUserByApiToken(token string) (u *User)  {
	user := &User{ApiToken: token}
	is, err := database.Orm.Where("api_token = ?", token).Get(user)
	if !is {
		log.Println("can't find user by api_token:", token)
	}
	if err != nil {
		log.Println(err)
	}
	return user
}
