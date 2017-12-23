// Copyright 2017 hIMEI
//
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
/////////////////////////////////////////////////////////////////////////
// @Author: hIMEI
// @Date:   2017-12-17 21:29:46
/////////////////////////////////////////////////////////////////////////

package main

import (
	//"net/url"
	"fmt"
	"strings"
	//"strconv"
	"bufio"
	//"errors"
	//"time"

	"github.com/antchfx/htmlquery"
	"github.com/hIMEI29A/gotorsocks"
	"golang.org/x/net/html"
)

// Consts for connecting to search engine end create requests
const (
	ICHIDAN   string = "ichidanv34wrx7m7.onion:80"
	GETPARAMS        = "GET /search?query="
)

// Parser is a html and xpath parser
type Parser struct {
	Request string
	Root    *html.Node
}

// NewParser creates instance of Parser
func NewParser(request string) *Parser {
	rootNode := getBody(request)
	parser := &Parser{request, rootNode}

	return parser
}

// Get body makes request to search engine and gets response body
func getBody(request string) *html.Node {
	requestString := GETPARAMS + request + "\n"
	tor, err := gotorsocks.NewTorGate()
	ErrFatal(err)

	connect, err := tor.DialTor(ICHIDAN)
	ErrFatal(err)

	fmt.Fprintf(connect, requestString)
	resp := bufio.NewReader(connect)

	node, err := htmlquery.Parse(resp)
	ErrFatal(err)

	return node
}

// FindEntry finds html element on the page
func findEntry(node *html.Node, entryexp string) *html.Node {
	return htmlquery.FindOne(node, entryexp)
}

// FindEntrys finds set of html elements on the page
func findEntrys(node *html.Node, entryexp string) []*html.Node {
	return htmlquery.Find(node, entryexp)
}

func getTag(node *html.Node, tagexp string) string {
	return htmlquery.InnerText(findEntry(node, tagexp))
}

func (p *Parser) getServices() []*html.Node {
	return findEntrys(p.Root, SERVICES)
}

func (p *Parser) getService(node *html.Node) []*html.Node {
	return findEntrys(node, SERVICE)
}

func (p *Parser) getServMap(services []*html.Node) map[string][]*Service {
	servMap := make(map[string][]*Service)

	for i, servs := range services {
		sNodes := p.getService(servs)
		var servSlice []*Service

		for _, srvnodes := range sNodes {
			name := p.getServiceName(srvnodes)
			count := p.getServiceCount(srvnodes)
			service := NewService(name, count)
			servSlice = append(servSlice, service)
		}

		servMap[TOPS[i]] = servSlice
	}

	return servMap
}

func (p *Parser) getServiceName(node *html.Node) string {
	nodeSpan := findEntry(node, SERVICENAME)
	return getTag(nodeSpan, LINK)
}

func (p *Parser) getServiceCount(node *html.Node) string {
	return (getTag(node, SERVICECOUNT))
}

func (p *Parser) getTotalr() int {
	return toInt(trimString((getTag(p.Root, TOTALR))))
}

func (p *Parser) getHref(node *html.Node) string {
	return htmlquery.SelectAttr(node, HREF)
}

func (p *Parser) getHosts() []*html.Node {
	return findEntrys(p.Root, HOST)
}

func (p *Parser) getHostFields(node *html.Node) []string {
	var fields []string

	links := findEntrys(node, LINK)
	hostUrl := p.getHref(links[0])

	addDate := getTag(node, SPAN)
	dateTrimmed := strings.TrimPrefix(addDate, ADDED)

	detailsLink := p.getHref(links[1])
	pre := getTag(node, PRE)

	fields = append(fields, hostUrl)
	fields = append(fields, dateTrimmed)
	fields = append(fields, detailsLink)
	fields = append(fields, pre)

	return fields
}

func (p *Parser) getHostsStructs(hosts []*html.Node) []*Host {
	var newhosts []*Host

	for _, h := range hosts {
		fields := p.getHostFields(h)
		hStruct := NewHost(fields)
		newhosts = append(newhosts, hStruct)
	}

	return newhosts
}

func (p *Parser) getPagination() []string {
	var links []string

	pagination := findEntry(p.Root, PAGINATION)
	hrefs := findEntrys(pagination, LINK)

	for _, newtag := range hrefs {
		if htmlquery.InnerText(newtag) != "Previous" && htmlquery.InnerText(newtag) != "Next" {
			link := p.getHref(newtag)
			links = append(links, link)
		}
	}

	return links
}
