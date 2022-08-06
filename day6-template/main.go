

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		main.go
//	Author:			Zeke
//	Date:			2022.05.29
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package main

import (
	"gee"
	"net/http"
	"time"
	//"log"
	"fmt"
	"html/template"
)

type student struct {
	Name	string
	Age		int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate":FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")
	stu1 := &student{Name: "Zeke", Age: 22}
	stu2 := &student{Name: "Hypersus", Age: 100}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":	"gee",
			"stuArr":	[2]*student{stu1,stu2},
		})
	})
	r.GET("/date", func(c *gee.Context){
		c.HTML(http.StatusOK,"custom_func.tmpl", gee.H{
			"title":	"gee",
			"now":		time.Now(),
		})
	})
	r.Run(":9999")
}
