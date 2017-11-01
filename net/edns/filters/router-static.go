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

package filters

import (
	"strings"

	"github.com/SabzCity/go-library/net/edns"
	"github.com/SabzCity/go-library/net/edns/ednsutil"
)

// StaticZoneCache : Set DNS zones to this variable in static process.
var StaticZoneCache DNSRouteTree

// StaticRouter : will ready response of question from selected zone.
// Needed filter: before(), after()
func StaticRouter(ctx *edns.Context) {
	// Multiple queries in a question section have not been supported in DNS
	// due the applicability of some DNS Message Header flags (such as AA) and
	// of the RCODE field only to a single QNAME, QTYPE, and QCLASS.
	if len(ctx.Request.Question) > 1 {
		return
	}
	question := ctx.Request.Question[0]

	zone := StaticZoneCache.FindExact(question.Name)
	if zone == nil {
		return
	}

	responseHandler, ok := responseHandlers[question.Qtype]
	if !ok {
		return
	}

	responseHandler(ctx, zone, question.Name)

	// https://tools.ietf.org/html/rfc2308#section-2.2.1
	if len(ctx.Response.Answer) == 0 {
		zone = StaticZoneCache.FindExact(zone.Origin)
		if zone == nil {
			return
		}
		SOAHandler(ctx, zone, question.Name)
	}
}

// DNSRouteTree struct, represent a name tree
type DNSRouteTree struct {
	domainLabel   string
	zone          *ednsutil.MiniDNS
	children      []*DNSRouteTree
	wildCardChild *DNSRouteTree
}

// FindExact finds the exact match for a given domain
func (node *DNSRouteTree) FindExact(domainName string) *ednsutil.MiniDNS {
	domainLabels := strings.Split(strings.Trim(strings.ToLower(domainName), "."), ".")
	for i, j := 0, len(domainLabels)-1; i < j; i, j = i+1, j-1 {
		domainLabels[i], domainLabels[j] = domainLabels[j], domainLabels[i]
	}

	lastNode := node.walkTree(domainLabels)
	if lastNode == nil {
		return nil
	}

	return lastNode.zone
}

// Set adds or replace a domainNames to the tree with the given zone.
func (node *DNSRouteTree) Set(zone *ednsutil.DNS) {
	for domainName := range zone.IN {
		domainLabels := strings.Split(strings.Trim(strings.ToLower(domainName), "."), ".")
		someNode := node

	Loop:
		for i := len(domainLabels) - 1; i >= 0; i-- {
			label := domainLabels[i]
			if label == "*" {
				if someNode.wildCardChild == nil {
					someNode.wildCardChild = &DNSRouteTree{}
				}

				if i == 0 {
					someNode.wildCardChild = &DNSRouteTree{zone: &ednsutil.MiniDNS{Origin: zone.Origin, TTL: zone.TTL, IN: zone.IN[domainName]}}
				}
				someNode = someNode.wildCardChild
				continue Loop
			} else {
				for _, child := range someNode.children {
					if label == child.domainLabel {
						if i == 0 {
							child.zone = &ednsutil.MiniDNS{Origin: zone.Origin, TTL: zone.TTL, IN: zone.IN[domainName]}
						}
						someNode = child
						continue Loop
					}
				}

				newNode := &DNSRouteTree{domainLabel: label}
				if i == 0 {
					newNode.zone = &ednsutil.MiniDNS{Origin: zone.Origin, TTL: zone.TTL, IN: zone.IN[domainName]}
				}
				someNode.children = append(someNode.children, newNode)
				someNode = newNode
			}
		}
	}
}

func (node *DNSRouteTree) walkTree(domainLabels []string) *DNSRouteTree {
	if len(domainLabels) == 0 {
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			if child.domainLabel == domainLabels[0] {
				return child.walkTree(domainLabels[1:])
			}
		}
	}

	if node.wildCardChild != nil {
		return node.wildCardChild.walkTree(domainLabels[1:])
	}

	return nil
}
