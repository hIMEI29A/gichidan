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

	"golang.org/x/net/html"
)

const (
	RED   string = "\x1B[31m"
	GRN          = "\x1B[32m"
	YEL          = "\x1B[33m"
	BLU          = "\x1B[34m"
	MAG          = "\x1B[35m"
	CYN          = "\x1B[36m"
	WHT          = "\x1B[97m"
	RESET        = "\x1B[0m"
	BOLD         = "\x1B[1m"
	LINE         = "\x1B[4m"
	INV          = "\x1B[7m"
	ITAL         = "\x1B[3m"
)

var gichidan *Gichidan

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

// Gichidan represents main app type
type Gichidan struct {
	Host
	Spider
	Parser
}

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
	//fmt.Println(request)

	channelBody := make(chan *html.Node, 120)
	channelDone := make(chan bool)

	var parsedHosts []*Host

	s := NewSpider(request)
	p := NewParser(request)
	//fmt.Println(request)

	go s.Crawl(request, channelDone, channelBody)

	for {
		select {
		case recievedNode := <-channelBody:
			newHosts := p.parseOne(recievedNode)
			for _, h := range newHosts {
				parsedHosts = append(parsedHosts, h)
				fmt.Println("added", h.HostUrl)
			}

		case <-channelDone:
			fmt.Println("Jobs finished")
			break
		}

		break
	}

	for _, m := range parsedHosts {
		fmt.Println(m.String())
	}
}
