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
	//"bufio"
	"fmt"
	//"log"
	//"strings"
)

// Gichidan represents main app type
type Gichidan struct {
	Data

	Parser
}

func main() {
	p := NewParser("bitcoin")

	tot := p.getTotalr()

	fmt.Println(tot)

	services := p.getServices()

	servMap := p.getServMap(services)

	hosts := p.getHosts()

	hStructs := p.getHostsStructs(hosts)

	page := NewPage(tot, servMap, hStructs)

	fmt.Println(page)
}
