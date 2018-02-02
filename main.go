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
	versionFlag = flag.Bool("v", false, "\tprint current version")

	bannerFlag = flag.Bool("b", false, "\tshow ASCII banner")

	// usage prints short help message
	usage = func() {
		fmt.Println(BOLD, RED, "\t", "Usage:", RESET)
		fmt.Println(WHT, "\t", "gichidan [<args>] [options]")
		fmt.Println(BLU, "Args:", GRN, "\t", "-r", "\t", CYN, "your search request to Ichidan")
		fmt.Println(BLU, "Options:\n", GRN, "\t\t")
	}

	// helpCmd prints usage()
	helpCmd = flag.Bool("h", false, "\thelp message")
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

	if *helpCmd {
		usage()
	}

	if len(os.Args) < 1 {
		usage()
		os.Exit(1)
	}

	if *requestFlag == "" {
		usage()
		os.Exit(1)
	}

	if *outputFlag != "" {
		Filepath = *outputFlag
	}

	var (
		parsedHosts []*Host
		rootHosts   = make(map[string]string)
		mutex       = &sync.Mutex{}
		totalHosts  = 1
	)

	// Channels
	var (
		channelBody = make(chan *html.Node, BUFFSIZE)
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

	for len(parsedHosts) < totalHosts {
		select {
		case recievedNode := <-channelBody:
			if s.checkRoot(recievedNode) == true {
				prim := p.getPrimary(recievedNode)
				total := p.getTotal(recievedNode)
				rootHosts[prim] = total
				// Get total number of all requests
				totalHosts += toInt(total)
			}

			go s.getPagination(recievedNode, chanUrls)
			go p.parseOne(recievedNode, chanHost)

		case newUrl := <-chanUrls:
			// Check if link was visited
			mutex.Lock()
			if s.HandledUrls[newUrl] == false {
				go s.Crawl(newUrl, channelBody)
				s.HandledUrls[newUrl] = true
				SLEEPER()
				fmt.Println(makeValMessage(newUrl), makeMessage(PROCESSING))
			} else {
			}
			mutex.Unlock()

		case newhosts := <-chanHost:
			for _, h := range newhosts {
				parsedHosts = append(parsedHosts, h)
			}
		}
	}

	fmt.Println(getTotalStats(rootHosts, totalHosts))

	pressAny()

	finalHosts := request.resultProvider(parsedHosts)

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
