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
	"flag"
	"fmt"
	"os"
	"os/exec"
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

// Gichidan represents main app type
type Gichidan struct {
	Data

	Parser
}

func main() {
	var (
		// "Search" subcommand
		//searchCmd = flag.NewFlagSet("search", flag.ExitOnError)
		//requestFlag = searchCmd.String("r", "", "your search request to Ichidan")

		// Version flag gets current app's version
		version    = getVersion()
		versionCmd = flag.Bool("v", false, "print current version")
	)

	flag.Parse()

	if *versionCmd {
		fmt.Println(version())
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		fmt.Println("Usage: gichidan <command> [<args>]")
		os.Exit(1)
	}

	p := NewParser("ichidan")
	p.Parse()
}
