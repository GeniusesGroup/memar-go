//Copyright 2017 SabzCity
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Package ehttputil is Extended HTTP package utility
// https://en.wikipedia.org/wiki/List_of_HTTP_header_fields
package ehttputil

import (
	"strings"
)

// New return a new tree with given root element at the top of the tree
func New() *URLNode {
	return &URLNode{}
}

// URLNode struct, represent a name tree
type URLNode struct {
	component     string
	entry         interface{}
	childrens     []*URLNode
	paramChildren *URLNode
}

// FindExact finds the exact match for a given prefix
func (node *URLNode) FindExact(prefix string) interface{} {

	components := strings.Split(strings.Trim(prefix, "/"), "/")

	lastNode, remaining := node.walkTree(components)

	if remaining > 0 {
		return nil
	}

	return lastNode.entry
}

// Insert adds a prefix to the tree with the given entry.
func (node *URLNode) Insert(prefix string, entry interface{}) {

	components := strings.Split(strings.Trim(prefix, "/"), "/")

	someNode := node

L1:
	for i, component := range components {
		if strings.HasPrefix(component, "{") && strings.HasSuffix(component, "}") {
			if someNode.paramChildren == nil {
				someNode.paramChildren = &URLNode{}
			}

			someNode = someNode.paramChildren

			if i == len(components)-1 {
				someNode.entry = entry
			}
		} else {
			for _, child := range someNode.childrens {
				if component == child.component {
					someNode = child

					if i == len(components)-1 {
						someNode.entry = entry
					}

					continue L1
				}
			}

			newNode := &URLNode{component: component}

			if i == len(components)-1 {
				newNode.entry = entry
			}

			if len(someNode.childrens) == 0 {
				someNode.childrens = []*URLNode{newNode}
			} else {
				someNode.childrens = append(someNode.childrens, newNode)
			}

			someNode = newNode
		}
	}
}

func (node *URLNode) walkTree(components []string) (*URLNode, int) {

	if len(components) == 0 {
		return node, 0
	}

	component := components[0]

	if len(node.childrens) > 0 {
		for _, child := range node.childrens {
			if child.component == component {
				return child.walkTree(components[1:])
			}
		}
	}

	if node.paramChildren != nil {
		innerNode, remaining := node.paramChildren.walkTree(components[1:])

		if remaining == 0 {
			return innerNode, 0
		}

		return node.paramChildren, len(components)
	}

	return node, len(components)
}
