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
	"flag"
	"fmt"
	"os"
	"path"
	"sync"

	"golang.org/x/net/html"
)

var (
	requestFlag   = flag.String("r", "", "your search request to Ichidan")
	shortInfoFlag = flag.Bool("s", false, "print hosts urls only")

	// Save output to file
	outputFlag = flag.String("f", "", "save results to file")
	Parsed     []*Host
	Filepath   string

	// Version flag gets current app's version
	version     = "v1.0.0"
	versionFlag = flag.Bool("v", false, "print current version")

	// Print ASCII banner for oldschool guys
	bannerFlag = flag.Bool("b", false, "show ASCII banner")

	// Don't print GET request's messages
	muteFlag = flag.Bool("m", false, "Don't print GET request's messages (non-verbose output)")

	helpCmd = flag.Bool("h", false, "help message")
)

// ToFile saves results to given file
func toFile(filepath string, parsed []*Host) {
	dir := path.Dir(filepath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errString := makeErrString(NOTEXIST)
		newerr := errors.New(errString)
		ErrFatal(newerr)
	}

	if _, err := os.Stat(filepath); os.IsExist(err) {
		errString := makeErrString(EXIST)
		newerr := errors.New(errString)
		ErrFatal(newerr)
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0666)
	ErrFatal(err)
	defer file.Close()

	for _, s := range parsed {
		file.WriteString(s.String() + "\n\n\n")
		ErrFatal(err)
	}
}

func main() {
	// Cli options parsing
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(1)
	}

	if len(os.Args) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *requestFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *outputFlag != "" {
		Filepath = *outputFlag
	}

	var (
		parsedHosts []*Host
		rootHosts   = make(map[string]string)
		mutex       = &sync.Mutex{}
		// len(parsedHosts) must be less than totalHosts at the start of crawling
		totalHosts = 1
	)

	// Channels
	var (
		channelBody = make(chan map[string]*html.Node, BUFFSIZE)
		chanUrls    = make(chan string, BUFFSIZE)
		chanHost    = make(chan []*Host, BUFFSIZE)
	)

	// Actors
	var (
		s = NewSpider()
		p = NewParser()
	)

	request := NewRequest(*requestFlag)

	// Start crawling
	for _, req := range request.RequestStrings {
		go s.Crawl(req, channelBody)
	}

	if *bannerFlag {
		banner()
	}

	fmt.Println(makeMessage(WAIT))
	if *muteFlag {
		SLEEPER()
		fmt.Println(makeMessage(CONN))
	}

	for len(parsedHosts) < totalHosts {
		select {
		case recievedNode := <-channelBody:
			primUrl, hostNode := unMap(recievedNode)
			if s.checkRoot(hostNode) == true {
				total := p.getTotal(hostNode)
				rootHosts[primUrl] = total
				// Get total number of all hosts. If here is first found root page,
				// totalHosts value must be decremented for happy loop exiting.
				if len(rootHosts) == 1 {
					totalHosts += (toInt(total) - 1)
				}

				if len(rootHosts) > 1 {
					totalHosts += toInt(total)
				}
			}

			go s.getPagination(hostNode, chanUrls)
			go p.parseOne(recievedNode, chanHost)

		case newUrl := <-chanUrls:
			// Check if link was visited
			mutex.Lock()
			if s.HandledUrls[newUrl] == false {
				go s.Crawl(newUrl, channelBody)
				s.HandledUrls[newUrl] = true
				SLEEPER()

				// verbose output
				if !*muteFlag {
					fmt.Println(makeValMessage(newUrl), makeMessage(PROCESSING))
				}

			} else {
			}
			mutex.Unlock()

		case newhosts := <-chanHost:
			for _, h := range newhosts {
				parsedHosts = append(parsedHosts, h)

			}
		}
	}

	finalHosts := request.resultProvider(parsedHosts)

	fmt.Println(getTotalStats(rootHosts, finalHosts, totalHosts))

	pressAny()

	// Results output. If shortInfoFlag was parsed, only collected urls will be printed.
	if !*shortInfoFlag {
		fmt.Println(makeMessage(FULL))
		for _, m := range finalHosts {
			fmt.Println(makeUrlMessage(m.String()))
		}
	} else {
		fmt.Println(makeMessage(SHORT))
		for _, m := range finalHosts {
			fmt.Println(makeUrlMessage(m.HostUrl))
		}
	}

	// Save results to file if flag parsed
	if Filepath != "" {
		fmt.Println(makeMessage(SAVED))
		Parsed = finalHosts
		toFile(Filepath, Parsed)
	}

}
