// Copyright 2017 hIMEI

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	//"errors"
	//"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// Parser is a html and xpath parser
type Parser struct {
	// Spider is an asynch urls handler
	Spider
	//	Hosts   []*Host
}

// NewParser creates instance of Parser
func NewParser(request string) *Parser {
	parser := &Parser{}

	return parser
}

// FindEntry finds html element on the page
func findEntry(node *html.Node, entryexp string) *html.Node {
	return htmlquery.FindOne(node, entryexp)
}

// FindEntrys finds set of html elements on the page
func findEntrys(node *html.Node, entryexp string) []*html.Node {
	return htmlquery.Find(node, entryexp)
}

// GetTag gets inner value of html tag
func getTag(node *html.Node, tagexp string) string {
	return htmlquery.InnerText(findEntry(node, tagexp))
}

// getHref gets content of href attribute of <a> tag
func getHref(node *html.Node) string {
	return htmlquery.SelectAttr(node, HREF)
}

// checkPage returns true if page is a root page and false if it is a host details page
func (s *Parser) checkPage(node *html.Node) bool {
	ch := false

	result := findEntry(node, SEARCHRESULT)
	if result != nil {
		ch = true
	}

	return ch
}

func (p *Parser) parseOne(node *html.Node) []*Host {
	var hosts []*Host
	//go func(node *html.Node) {
	hostsNodes := p.getHosts(node)

	for _, h := range hostsNodes {
		fields := p.getHostFields(h)
		var services []*Service

		detailslink := getHref(findEntry(h, DETAILS))
		req := requestProvider(detailslink)

		chanNode := getContents(req)
		dnode := <-chanNode

		srvNodes := findEntrys(dnode, SERVICELONG)

		for _, srv := range srvNodes {
			srvFields := p.getServiceFields(srv)
			service := NewService(srvFields)
			services = append(services, service)
		}

		host := NewHost(fields, services)
		hosts = append(hosts, host)
	}
	//}(node)
	return hosts
}

func (p *Parser) getHostFields(node *html.Node) []string {
	var fields []string

	hostUrl := getHref(findEntry(findEntry(findEntry(node, SUMMARY), ONION), LINK))
	fields = append(fields, hostUrl)

	addDate := strings.TrimPrefix(getTag(findEntry(node, SUMMARY), SPAN), ADDED)
	fields = append(fields, addDate)

	return fields
}

func (p *Parser) getServiceFields(node *html.Node) []string {
	var fields []string

	if findEntry(node, H3) != nil {
		fields = append(fields, getTag(node, H3))
	} else {
		fields = append(fields, getTag(node, STATE))
	}

	fields = append(fields, getTag(node, PORT))
	fields = append(fields, getTag(node, PROTO))
	fields = append(fields, getTag(node, STATE))

	if findEntry(node, VERSION) != nil {
		fields = append(fields, getTag(node, VERSION))
	} else {
		fields = append(fields, "unknown VERSION")
	}

	fields = append(fields, getTag(node, PRE))

	return fields
}

// getServices gets <div>'s of class "service"
func (p *Parser) getService(node *html.Node) []*html.Node {
	return findEntrys(node, SERVICE)
}

// getHosts gets data to instantiate Host structs
func (p *Parser) getHosts(node *html.Node) []*html.Node {
	return findEntrys(node, HOST)
}
