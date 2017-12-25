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
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/hIMEI29A/gotorsocks"
	"golang.org/x/net/html"
)

// Consts for connecting to search engine
const (
	ICHIDAN     string = "ichidanv34wrx7m7.onion"
	ICHIDANPORT        = ":80"
	GETPARAMS          = "GET /search?query="
)

// Parser is a html and xpath parser
type Parser struct {
	// Spider is an asynch urls handler
	Spider
	// Request string
	Request string
	// Root is a result recieved from Ichidan
	Root *html.Node
}

// Spider is an asynch urls handler
type Spider struct {

	// Urls to crawl
	Urls []string
}

// NewParser creates instance of Parser
func NewParser(request string) *Parser {
	rootNode := getBody(request)

	if checkResult(rootNode) == false {
		err := errors.New("Nothing found there, Neo!")
		ErrFatal(err)
	}

	parser := &Parser{request, rootNode}

	return parser
}

// GetBody makes request to search engine and gets response body
func getBody(request string) *html.Node {
	requestString := GETPARAMS + request + "\n"
	tor, err := gotorsocks.NewTorGate()
	ErrFatal(err)

	url := ICHIDAN + ICHIDANPORT
	connect, err := tor.DialTor(url)
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

// GetTag gets inner value of html tag
func getTag(node *html.Node, tagexp string) string {
	return htmlquery.InnerText(findEntry(node, tagexp))
}

// CheckResult controls empty search results
func checkResult(node *html.Node) bool {
	check := true

	result := findEntry(node, NORESULT)
	if result != nil {
		check = false
	}

	return check
}

// Parse is a main Parser's method
func (p *Parser) Parse() {
	tot := p.getTotalr()
	services := p.getServices()
	servMap := p.getServMap(services)
	hosts := p.getHosts()
	hStructs := p.getHostsStructs(hosts)
	page := NewPage(tot, servMap, hStructs)
	fmt.Println(page)
}

// getServices gets <div>'s of class "services"
func (p *Parser) getServices() []*html.Node {
	return findEntrys(p.Root, SERVICES)
}

// getServices gets <div>'s of class "service"
func (p *Parser) getService(node *html.Node) []*html.Node {
	return findEntrys(node, SERVICE)
}

// getServMap prepares data to instantiate Service struct
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

// getServiceName gets name for Service struct
func (p *Parser) getServiceName(node *html.Node) string {
	nodeSpan := findEntry(node, SERVICENAME)
	return getTag(nodeSpan, LINK)
}

// getServiceName gets count for Service struct
func (p *Parser) getServiceCount(node *html.Node) string {
	return trimString((getTag(node, SERVICECOUNT)))
}

// getTotalr gets number of found hosts per request
func (p *Parser) getTotalr() int {
	totalr := toInt(trimString((getTag(p.Root, TOTALR))))

	return totalr
}

// getHref gets content of href attribute of <a> tag
func (p *Parser) getHref(node *html.Node) string {
	return htmlquery.SelectAttr(node, HREF)
}

// getHosts gets data to instantiate Host structs
func (p *Parser) getHosts() []*html.Node {
	return findEntrys(p.Root, HOST)
}

func (p *Parser) getDetails(url string) string {
	detailsLink := ICHIDAN + url + ICHIDANPORT
}

// getHosts gets data to instantiate Host structs
func (p *Parser) getHostFields(node *html.Node) []string {
	var fields []string

	links := findEntrys(node, LINK)
	hostUrl := p.getHref(links[0])

	addDate := getTag(node, SPAN)
	dateTrimmed := strings.TrimPrefix(addDate, ADDED)

	details := p.getHref(links[1])
	detailsLink := p.getDetails(details)
	pre := getTag(node, PRE)

	fields = append(fields, hostUrl)
	fields = append(fields, dateTrimmed)
	fields = append(fields, detailsLink)
	fields = append(fields, pre)

	return fields
}

func (p *Parser) getDetailsList(body *html.Node) []map[string]string {
	var list []map[string]string

	listNodes := findEntrys(body, SERVICELONG)

	for _, node := range listNodes {
		nodeMap := make(map[string]string)
		nodeMap["port"] = getTag(node, PORT)
		nodeMap["proto"] = getTag(node, PROTO)
		nodeMap["state"] = getTag(node, STATE)
		nodeMap["service name"] = getTag(node, H3)
		nodeMap["version"] = getTag(node, SMALL)
		nodeMap["pre"] = getTag(node, PRE)

		list = append(list, nodeMap)
	}

	return list
}

func (p *Parser) getHostsStructs(hosts []*html.Node) []*Host {
	var newhosts []*Host

	for _, h := range hosts {
		fields := p.getHostFields(h)
		body := p.getDetailsBody(fields[2])
		detailsList := getDetailsList(body)
		hStruct := NewHost(fields, detailsList)
		newhosts = append(newhosts, hStruct)
	}

	return newhosts
}

func (p *Parser) getPagination() []string {
	var links []string

	pagination := findEntry(p.Root, PAGINATION)

	if pagination != nil {
		hrefs := findEntrys(pagination, LINK)

		for _, newtag := range hrefs {
			if htmlquery.InnerText(newtag) != "Previous" && htmlquery.InnerText(newtag) != "Next" {
				link := p.getHref(newtag)
				links = append(links, link)
			}
		}
	}

	return links
}

///////////////////////////////////////////////////////
// Spider's constructor and methods ////////////////////
///////////////////////////////////////////////////////

func NewSpider(urls []string) *Spider {
	spider := &Spider{urls}

	return spider
}

func (s *Spider) getDetailsBody(url string) *html.Node {
	tor, err := gotorsocks.NewTorGate()
	ErrFatal(err)

	connect, err := tor.DialTor(url)
	ErrFatal(err)

	requestString := "GET " + url + " HTTP/1.0\r\n"
	fmt.Fprintf(connect, requestString)
	resp := bufio.NewReader(connect)

	node, err := htmlquery.Parse(resp)
	ErrFatal(err)

	return node
}

func (s *Spider) Collect() []*html.Node {
	var urls []*html.Node

	for _, url := range s.Urls {
		body, err := htmlquery.LoadURL(url)
		ErrFatal(err)
		urls = append(urls, body)
	}

	return urls
}
