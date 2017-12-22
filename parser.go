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
	//"strconv"
	//"strings"
	"bufio"
	"errors"
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

func (p *Parser) getTotalr() int {
	return toInt(getTag(p.Root, TOTALR))
}

func getServiceName(node *html.Node) string {
	nodeSpan := findEntry(node, SERVICENAME)
	return getTag(nodeSpan, LINK)
}

func getServiceCount(node *html.Node) int {
	return toInt(getTag(node, SERVICECOUNT))
}

func (p *Parser) getServicesNodes() map[string]*html.Node {
	servsnodes := make(map[string]*html.Node)
	newtags := findEntrys(p.Root, SERVICES)

	if len(newtags) != len(TOPS) {
		err := errors.New("Number of services not equals number of headers6")
		ErrFatal(err)
	}

	for i, s := range newtags {
		for j, h6 := range TOPS {
			if i == j {
				servsnodes[h6] = s
			}
		}
	}

	return servsnodes
}

func getServiceNodes(services *html.Node) []*html.Node {
	var servnodes []*html.Node
	newtags := findEntrys(services, SERVICE)

	for _, node := range newtags {
		servnodes = append(servnodes, node)
	}

	return servnodes
}

func (p *Parser) getSummaryFields() map[int][]string {
	nodes := findEntrys(p.Root, SUMMARY)
	fields := make(map[int][]string)
	var body []string

	for i, newtag := range nodes {
		body = append(body, htmlquery.InnerText(findEntry(newtag, HREF)))
		body = append(body, htmlquery.InnerText(findEntry(newtag, SPAN)))
		body = append(body, htmlquery.InnerText(findEntry(newtag, DETAILS)))
		//details := findEntry(newtag, DETAILS)
		/*
			for _, attr := range details.Attr {
				if details.Key == "href" {
					body = append(body, details.Val)
					break
				}
			}*/

		body = append(body, htmlquery.InnerText(findEntry(newtag, PRE)))
		fields[i] = body
	}

	return fields
}

/*
func (p *Parser) getPagination() string {
	node := htmlquery.InnerText(findEntry(p.Root, PAGINATION))

	return node
}
*/
