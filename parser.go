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
	"errors"
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// Usefull string constants for html parsing
const (
	ADDED          string = "Added on "
	DELIM                 = "]===================================================[\n"
	LONGFORM              = "2017-09-09 01:30:35 UTC"
	PRE                   = "//pre"
	SPAN                  = "//span"
	LINK                  = "//a"
	HREF                  = "href"
	H2                    = "//h2"
	H3                    = "//h3"
	VERSION               = "//small"
	NONE                  = " "
	SEARCHRESULT          = "//div[@class='search-results']"
	PAGINATION            = "//div[@class='pagination']"
	DETAILS               = "//a[@class='details']"
	SUMMARY               = "//div[@class='search-result-summary col-xs-4']"
	ONION                 = "//div[@class='onion']"
	TOTALR                = "//div[@class='bignumber']"
	SERVICE               = "//div[@class='service']"
	SERVICES              = "//div[@class='services']"
	SERVICELONG           = "//li[@class='service service-long']"
	SERVICEDETAILS        = "//div[@class='service-details col-sm-2']"
	HOST                  = "//div[@class='search-result row-fluid']"
	NORESULT              = "//div[@class='msg alert alert-info']"
	RESULT                = "//div[@class='col-sm-9']"
	PORT                  = "//div[@class='port']"
	PROTO                 = "//div[@class='protocol']"
	STATE                 = "//div[@class='state']"
)

// Parser is a html and xpath parser
type Parser struct {
	// Spider is an asynch urls handler
	Spider
	// Request string
	Request string
	//	Hosts   []*Host
	// Urls already being handled
	HandledUrls []string
}

// NewParser creates instance of Parser
func NewParser(request string) *Parser {
	parser := &Parser{}
	parser.Request = requestProvider(request)

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

// CheckResult controls empty search results
func checkResult(node *html.Node) bool {
	ch := true

	result := findEntry(node, NORESULT)
	if result != nil {
		ch = false
	}

	return ch
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

func (p *Parser) parseAll(request string) []*Host {
	var hosts []*Host

	fmt.Println(request)
	root := getContents(request)

	if checkResult(root) == false {
		err := errors.New("Nothing found there, Neo!")
		ErrFatal(err)
	}

	urls := p.getLinks(root)
	fmt.Printf("Must have %d urls \n", len(urls))

	spider := NewSpider(urls)
	// paginated root-nodes
	recieved := spider.Crawl()
	nodes := spider.Collect(recieved)

	for _, n := range nodes {
		newhosts := p.parseOne(n)

		for _, m := range newhosts {
			hosts = append(hosts, m)
		}

		newlinks := p.getLinks(n)

		for _, l := range newlinks {
			morehosts := p.parseAll(l)

			for _, j := range morehosts {
				hosts = append(hosts, j)
			}
		}
	}

	return hosts
}

func (p *Parser) parseOne(node *html.Node) []*Host {
	var hosts []*Host

	hostsNodes := p.getHosts(node)

	for _, h := range hostsNodes {
		fields := p.getHostFields(h)
		var services []*Service

		detailslink := p.getHref(findEntry(h, DETAILS))
		req := requestProvider(detailslink)
		dnode := getContents(req) // TODO: gorutines
		srvNodes := findEntrys(dnode, SERVICELONG)

		for _, srv := range srvNodes {
			srvFields := p.getServiceFields(srv)
			service := NewService(srvFields)
			services = append(services, service)
		}

		host := NewHost(fields, services)
		hosts = append(hosts, host)
	}

	return hosts
}

/*
// parseNode is a main Parser's method
func (p *Parser) parseNode(request string) []*Host {
	var hosts []*Host

	req := requestProvider(request)
	//fmt.Println(req)
	root := getContents(req)

	if checkResult(root) == false {
		err := errors.New("Nothing found there, Neo!")
		ErrFatal(err)
	}

	urls := p.getLinks(root)
	fmt.Println(urls[0])
	fmt.Printf("Must have %d urls \n", len(urls))

	spider := NewSpider(urls)
	// paginated root-nodes
	recieved := spider.Crawl()
	nodes := spider.Collect(recieved)

	for _, node := range nodes {
		// hosts <div>'s
		hostsNodes := p.getHosts(node)

		for _, h := range hostsNodes {
			fields := p.getHostFields(h)
			var services []*Service

			detailslink := p.getHref(findEntry(h, DETAILS))
			req := requestProvider(detailslink)
			dnode := getContents(req) // TODO: gorutines
			srvNodes := findEntrys(dnode, SERVICELONG)

			for _, srv := range srvNodes {
				srvFields := p.getServiceFields(srv)
				service := NewService(srvFields)
				services = append(services, service)
			}

			host := NewHost(fields, services)
			hosts = append(hosts, host)
		}

		paginations := p.getPagination(node)

		for _, page := range paginations {

			// recursion is here
			newhosts := p.parseNode(page)

			for _, newbody := range newhosts {
				hosts = append(hosts, newbody)
			}

		}
	}

	return hosts
}
*/

func (p *Parser) getLinks(node *html.Node) []string {
	var urls []string
	//fmt.Println(p.Request)
	urls = append(urls, p.Request)
	p.HandledUrls = append(p.HandledUrls, p.Request)
	//fmt.Println(urls[0])

	pagination := p.getPagination(node)

	for _, link := range pagination {
		urls = append(urls, link)
		p.HandledUrls = append(p.HandledUrls, link)
	}

	return urls
}

func (p *Parser) getHostFields(node *html.Node) []string {
	var fields []string

	hostUrl := p.getHref(findEntry(findEntry(findEntry(node, SUMMARY), ONION), LINK))
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
		fields = append(fields, getTag(node, NONE))
	}

	fields = append(fields, getTag(node, PRE))

	return fields
}

// getServices gets <div>'s of class "service"
func (p *Parser) getService(node *html.Node) []*html.Node {
	return findEntrys(node, SERVICE)
}

// getHref gets content of href attribute of <a> tag
func (p *Parser) getHref(node *html.Node) string {
	return htmlquery.SelectAttr(node, HREF)
}

// getHosts gets data to instantiate Host structs
func (p *Parser) getHosts(node *html.Node) []*html.Node {
	return findEntrys(node, HOST)
}

func (p *Parser) checkVisited(url string) bool {
	check := false

	for _, handled := range p.HandledUrls {
		if handled == url {
			check = true
			break
		}
	}

	return check
}

func (p *Parser) getPagination(node *html.Node) []string {
	var links []string

	pagination := findEntry(node, PAGINATION)

	if pagination != nil {
		hrefs := findEntrys(pagination, LINK)

		for _, newtag := range hrefs {
			if htmlquery.InnerText(newtag) != "Previous" &&
				htmlquery.InnerText(newtag) != "Next" &&
				htmlquery.InnerText(newtag) != "1" {
				link := requestProvider(p.getHref(newtag))

				if p.checkVisited(link) == false {
					links = append(links, link)
				}
			}
		}
	} else {
		fmt.Println("No more pages")
	}

	return links
}
