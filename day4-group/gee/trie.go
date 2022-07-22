

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		trie.go
//	Author:			Zeke
//	Date:			2022.05.16
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package gee

import (
	"strings"
)

type node struct {
	pattern		string
	part		string
	children	[]*node
	isWild		bool
}

// return first matched node
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// return all matched nodes
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node,0)
	for _,child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes,child)
		}
	}
	return nodes
}


func (n *node) insert(pattern string, parts []string, height int) {
	// leave node in this situation
	if len(parts)==height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part : part, isWild: part[0]==':' || part[0]=='*'}
		n.children=append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search from node n with given parts
func (n *node) search(parts []string, height int) *node {
	// if current searching has reached the last item of parts or node n has wildcard matching
	if len(parts)==height || strings.HasPrefix(n.part,"*") {
		// node n is not a valid address, routing failed
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _,child := range children {
		result := child.search(parts,height+1)
		if result != nil {
			return result
		}
	}
	return nil
}


