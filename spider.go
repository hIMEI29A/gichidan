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

/** file spider.go contains data types and its methods for web-crawling */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"

	"github.com/antchfx/htmlquery"
	"github.com/hIMEI29A/gotorsocks"
	"golang.org/x/net/html"
)

// Spider is an async urls handler
type Spider struct {
	// Urls already being handled
	HandledUrls map[string]bool
}

// NewSpider is a constructor for Spider
func NewSpider() *Spider {
	handled := make(map[string]bool)
	spider := &Spider{}
	spider.HandledUrls = handled

	return spider
}

// ConnectProvider provides connect to Ichidan with gotorsocks package
func connectProvider() net.Conn {
	tor, err := gotorsocks.NewTorGate()
	ErrFatal(err)

	connect, err := tor.DialTor(ICHIDAN)
	ErrFatal(err)

	return connect
}

// GetContents makes request to Ichidan search engine and gets response body
func getContents(request string) chan *html.Node {
	chanNode := make(chan *html.Node)
	go func() {
		connect := connectProvider()
		defer connect.Close()

		fmt.Fprintf(connect, request)
		resp := bufio.NewReader(connect)

		node, err := htmlquery.Parse(resp)
		ErrFatal(err)
		chanNode <- node
	}()

	return chanNode
}

// CheckResult controls empty search results
func (s *Spider) checkResult(node *html.Node) bool {
	ch := true

	resultNoresult := findEntry(node, NORESULT)
	if resultNoresult != nil {
		ch = false
	}

	return ch
}

func (s *Spider) checkAuth(node *html.Node) bool {
	ch := true

	resultNoauth := findEntry(node, NOAUTH)
	if resultNoauth != nil {
		ch = false
	}

	return ch
}

// CheckRoot checks if given page is first or single page
func (s *Spider) checkRoot(node *html.Node) bool {
	ch := false

	if s.checkSingle(node) == false || getTag(findEntry(node, PAGINATION), CURRENT) == "1" {
		ch = true
	}

	return ch
}

// CheckDone checks last pagination's page
func (s *Spider) checkDone(node *html.Node) bool {
	ch := false

	pagination := findEntry(node, PAGINATION)

	if findEntry(pagination, DISABLED) != nil {
		ch = true
	}

	return ch
}

// CheckSingle checks if given page is single (have not pagination)
func (s *Spider) checkSingle(node *html.Node) bool {
	ch := true

	if findEntry(node, PAGINATION) == nil {
		ch = false
	}

	return ch
}

// Crawl is a async crawler that takes request as first argument, gets it content
// and sends it to channel given as second argument
func (s *Spider) Crawl(url string, channelBody chan map[string]*html.Node) {
	bodyMap := make(map[string]*html.Node)
	chanNode := getContents(url)
	body := <-chanNode

	if s.checkResult(body) == false {
		errString := makeErrString(NOTHING)
		err := errors.New(errString)
		ErrFatal(err)
	}

	if s.checkAuth(body) == false {
		errString := makeErrString(ERRAUTH)
		err := errors.New(errString)
		ErrFatal(err)
	}

	bodyMap[trimString(trimUrl(url))] = body

	channelBody <- bodyMap

	return
}

// GetPagination finds pagination <div> and gets all links from it.
// Also it checks for single-paged result
func (s *Spider) getPagination(node *html.Node, chanUrls chan string) {
	pagination := findEntry(node, PAGINATION)

	if pagination != nil {
		current := toInt(getTag(pagination, CURRENT))
		hrefs := findEntrys(pagination, LINK)

		for _, newtag := range hrefs {
			if htmlquery.InnerText(newtag) != PREVIOUS &&
				htmlquery.InnerText(newtag) != NEXT &&
				toInt(htmlquery.InnerText(newtag)) > current {
				req := NewRequest(getHref(newtag))

				chanUrls <- req.RequestStrings[0]
			}
		}
	} else {
		fmt.Println(makeMessage(ONLYONE))
	}

	return
}
