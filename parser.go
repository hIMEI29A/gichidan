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

/*
func createExp(entryname, attrname, attrvalue string) string {
	switch entryname {
	case entryname == "div":

	}
}
*/

func (p *Parser) getH6names() []string {
	var tags []string
	nodes := findEntrys(p.Root, H6)

	for _, newtag := range nodes {
		text := htmlquery.InnerText(newtag)
		tags = append(tags, text)
	}

	return tags
}

func (p *Parser) getServiceFields(h6names []string) map[string][]string {
	nodes := findEntrys(p.Root, H6)
	fields := make(map[string][]string)
	var body []string

	for i, newtag := range nodes {
		fmt.Println(newtag)
		fname := h6names[i]
		//fmt.Println(fname)
		if htmlquery.InnerText(newtag) == fname {
			body = append(body, htmlquery.InnerText(findEntry(newtag, LINK)))
			body = append(body, htmlquery.InnerText(findEntry(newtag, SERVICECOUNT)))
			fields[fname] = body
		}
	}

	return fields
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
