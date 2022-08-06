

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		router.go
//	Author:			Zeke
//	Date:			2022.05.25
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package gee

import (
	"net/http"
	"strings"
)

type router struct {
	// root per method
	roots		map[string]*node
	handlers	map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:		make(map[string]*node),
		handlers:	make(map[string]HandlerFunc),
	}
}

// Universal inner method to parse parameters
// e.g. pattern=="/static/css/icon.jpg"
// then return []string{"static","css","icon.jpg"}
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern,"/")
	parts := make([]string,0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// Bind route and handler to provide service (i.e. enroll routing)
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_,ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}	
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}


func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	// see if the root of the given method exists
	root, ok := r.roots[method]

	if !ok {
		return nil,nil
	}

	n:=root.search(searchParts,0)

	if n != nil {
		parts:=parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:],"/")
				break
			}
		}
		return n, params
	}

	return nil,nil
}

// Dynamically match the routing and call binded function
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n",c.Path)
	}
	c.Next()
}
