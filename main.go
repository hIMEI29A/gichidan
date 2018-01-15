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
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"golang.org/x/net/html"
)

// getVersion parses git to get current app's version
func getVersion() func() string {
	return func() string {
		getVerCmd := exec.Command("git", "describe", "--tags")
		getVerOut, err := getVerCmd.Output()
		ErrFatal(err)

		return string(getVerOut)
	}
}

/*
func toFile(filename string, output []*Host) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0666)
	ErrFatal(err)
	defer file.Close()

	for _, s := range output {
		file.WriteString(s.String() + "\n\n\n")
	}
}
*/

func main() {
	var (
		// "Search" subcommand
		searchCmd   = flag.NewFlagSet("search", flag.ExitOnError)
		requestFlag = searchCmd.String("r", "", "your search request to Ichidan")

		// Version flag gets current app's version
		version    = getVersion()
		versionCmd = flag.Bool("v", false, "print current version")

		//toFileCmd = flag.String("f", "", "save output to file")

		// usage prints short help message
		usage = func() {
			fmt.Println(BOLD, RED, "\t", "Usage:", RESET)
			fmt.Println(WHT, "\t", "gichidan <command> [<args>] [options]")
			fmt.Println("\t\t", BLU, "Commands:", GRN, "\t", "search")
			fmt.Println("\t\t", BLU, "Args:", GRN, "\t", "-r", "\t", "your search request to Ichidan")
			fmt.Println("\t\t", BLU, "Options:", GRN, "\t", "-v", "\t", "app's current version")
			fmt.Println("\t\t\t\t", "-h", "\t", "prints this message", RESET)
			fmt.Println("\n")
		}
	)

	flag.Parse()

	if *versionCmd {
		fmt.Println(version())
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		searchCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if searchCmd.Parsed() {
		if *requestFlag == "" {
			usage()
			os.Exit(1)
		}
	}

	request := requestProvider(*requestFlag)

	channelBody := make(chan *html.Node, 120)

	var parsedHosts []*Host
	wg := &sync.WaitGroup{}

	s := NewSpider(request)
	p := NewParser(request)

	go s.Crawl(request, channelBody, wg)

	for {
		recievedNode := <-channelBody
		newHosts := p.parseOne(recievedNode)

		urlStr := BOLD + YEL + "Collected hosts: " + RESET
		fmt.Println(urlStr)

		for _, h := range newHosts {
			parsedHosts = append(parsedHosts, h)
			fmt.Println(h.HostUrl)
		}

		wg.Wait()

		if s.checkSingle(recievedNode) != false {
			if s.checkDone(recievedNode) == true {
				fmt.Println(BOLD, GRN, "finished", RESET)
				break
			}
		} else {
			fmt.Println(BOLD, GRN, "finished", RESET)
			break
		}

	}

	fmt.Println(BOLD, RED, "Full info:\n", RESET)
	for _, m := range parsedHosts {
		fmt.Println(m.String())
	}
}
