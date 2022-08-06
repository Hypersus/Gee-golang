

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		recovery.go
//	Author:			Zeke
//	Date:			2022.08.06
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package gee

import(
	"strings"
	"log"
	"runtime"
	"fmt"
	"net/http"
)

func trace(message string) string {
	pcs := make([]uintptr,32)
	n := runtime.Callers(3, pcs)

	var str strings.Builder
	str.WriteString(message+"\nTraceBack:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s",err)
				log.Printf("%s\n\n",trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
