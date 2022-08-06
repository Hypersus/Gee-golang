

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
//	"html/template"
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
	r.Use(gee.Recovery())
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Zeke\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"zeke"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
