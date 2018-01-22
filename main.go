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
	version     = "0.1.2"
	versionFlag = flag.Bool("v", false, "\tprint current version")

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
		errString := BOLD + RED + "Given path does not exist" + RESET
		newerr := errors.New(errString)
		ErrFatal(newerr)
	}

	if _, err := os.Stat(filepath); os.IsExist(err) {
		errString := BOLD + RED + "File already exist, we'll not rewrite it " + RESET
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

	request := requestProvider(*requestFlag)

	var (
		parsedHosts []*Host
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

	// Start crawling
	go s.Crawl(request, channelBody)

	for len(parsedHosts) < totalHosts {
		select {
		case recievedNode := <-channelBody:
			if s.checkRoot(recievedNode) == true {
				// Get results total number
				total := s.getTotal(recievedNode)
				totalHosts = toInt(total)
				fmt.Println(BOLD, YEL, "Hosts found: ", CYN, total, RESET)
			}

			go s.getPagination(recievedNode, chanUrls)
			go p.parseOne(recievedNode, chanHost)

		case newUrl := <-chanUrls:
			// Firstly check if link was visited
			mutex.Lock()
			if s.HandledUrls[newUrl] == false {
				go s.Crawl(newUrl, channelBody)
				s.HandledUrls[newUrl] = true
				SLEEPER()
				fmt.Println(BOLD, CYN, newUrl, YEL, " in processing", RESET)
			} else {
			}
			mutex.Unlock()

		case newhosts := <-chanHost:
			for _, h := range newhosts {
				parsedHosts = append(parsedHosts, h)
			}
		}
	}

	// Results output. If shortInfoFlag was parsed, only hosts urls will be printed.
	if !*shortInfoFlag {
		fmt.Println(BOLD, YEL, "Full info:\n", RESET)
		for _, m := range parsedHosts {
			fmt.Println(BOLD, GRN, m.String(), RESET)
		}
	} else {
		fmt.Println(BOLD, YEL, "\nShort info:\n", RESET)
		for _, m := range parsedHosts {
			fmt.Println(BOLD, GRN, m.HostUrl, RESET)
		}
	}

	// Save results to file if flag parsed
	if Filepath != "" {
		fmt.Println(BOLD, YEL, "Saved to ", CYN, Filepath, RESET)
		Parsed = parsedHosts
		toFile(Filepath, Parsed)
	}

}
