

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		main.go
//	Author:			Zeke
//	Date:			2022.05.14
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
		case "/":
			fmt.Fprintf(w,"URL.Path=%q\n",req.URL.Path)
		case "/hello":
			for k, v := range req.Header {
				fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
			}
		default:
			fmt.Fprintf(w,"404 NOT FOUND: %s\n",req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
