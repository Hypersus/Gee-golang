

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		main.go
//	Author:			Zeke
//	Date:			2022.05.15
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package main

import (
	"gee"
	"net/http"
)

func main() {
	e := gee.New()
	e.GET("/",func(c *gee.Context){
		c.HTML(http.StatusOK,"<h1>Hello Gee</h1>\n")
	})
	e.GET("/hello",func(c *gee.Context){
		c.String(http.StatusOK,"hello %s, you're at %s\n",c.Query("name"),c.Path)
	})

	e.POST("/login",func(c *gee.Context){
		c.JSON(http.StatusOK,gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	e.Run(":9999")
}
