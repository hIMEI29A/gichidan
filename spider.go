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
	"bufio"
	//"errors"
	"fmt"
	"net"
	"strings"
	//"time"

	"github.com/antchfx/htmlquery"
	"github.com/hIMEI29A/gotorsocks"
	"golang.org/x/net/html"
)

// Consts for connecting to search engine
const (
	ICHIDAN     string = "ichidanv34wrx7m7.onion"
	ICHIDANPORT        = ":80"
	SEARCH             = "/search?query="
)

// Spider is an async urls handler
type Spider struct {
	// Urls to crawl
	Urls []string
}

func NewSpider(urls []string) *Spider {
	spider := &Spider{urls}

	return spider
}

func requestProvider(request string) string {
	var fullRequest string
	switch {
	case string(request[0]) == "/":
		fullRequest = "GET " + request + "\n"

	case string(request[0]) != "/" &&
		string(request[0]) != NONE &&
		strings.Contains(request, NONE) == true:

		splitted := strings.Split(request, NONE)
		fullRequest = "GET " + SEARCH + splitted[0] + "\n"

	default:
		fullRequest = "GET " + SEARCH + request + "\n"
	}

	return fullRequest
}

func connectProvider(url string) net.Conn {
	tor, err := gotorsocks.NewTorGate()
	ErrFatal(err)

	connect, err := tor.DialTor(url)
	ErrFatal(err)

	return connect
}

// getContents makes request to search engine and gets response body
func getContents(request string) *html.Node {
	url := ICHIDAN + ICHIDANPORT
	connect := connectProvider(url)
	defer connect.Close()

	fmt.Fprintf(connect, request)
	resp := bufio.NewReader(connect)

	node, err := htmlquery.Parse(resp)
	ErrFatal(err)

	return node
}

func (s *Spider) Crawl() <-chan *html.Node {
	fmt.Printf("Need to handle %d urls \n", len(s.Urls))
	fmt.Println(len(s.Urls))

	var count int
	recieve := make(chan *html.Node, 12)

	fmt.Println(s.Urls[0])
	for i := 0; i < len(s.Urls); i++ {
		fmt.Println("BEBE")
		go func() {
			fmt.Println("MEME")
			SLEEPER()
			//time.Sleep(500 * time.Millisecond)
			fmt.Println(s.Urls[i])
			node := getContents(s.Urls[i])
			if node == nil {
				fmt.Println("node not ready")
			}
			recieve <- node
		}()
	}

	fmt.Printf("Started %d workers\n", count)

	return recieve
}

func (s *Spider) Collect(recieve <-chan *html.Node) []*html.Node {
	var nodes []*html.Node
	for i := 0; i < len(s.Urls); i++ {
		select {
		case node := <-recieve:
			fmt.Println("node recieved")
			nodes = append(nodes, node)
		default:
			//			SLEEPER()
			//			fmt.Println(i, "Waiting...")
		}
	}

	return nodes
}
