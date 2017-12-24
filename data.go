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
	"fmt"
	"time"
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

const (
	ADDED        string = "Added on "
	DELIM               = "]===================================================[\n"
	LONGFORM            = "2017-09-09 01:30:35 UTC"
	PRE                 = "//pre"
	SPAN                = "//span"
	LINK                = "//a"
	HREF                = "href"
	PAGINATION          = "//div[@class='pagination']"
	TOTALR              = "//div[@class='bignumber']"
	SERVICE             = "//div[@class='service']"
	SERVICES            = "//div[@class='services']"
	SERVICENAME         = "//div[@class='span8 name']"
	SERVICECOUNT        = "//div[@class='span4 count']"
	HOST                = "//div[@class='search-result row-fluid']"
	NORESULT            = "//div[@class='msg alert alert-info']"
	TOPSERVICES         = "Top Services"
	TOPSOFT             = "Top Software"
	TOPSYS              = "Top Operating Systems"
)

var TOPS = []string{
	TOPSERVICES,
	TOPSOFT,
	TOPSYS,
}

type Data struct {
	Query string
	Date  time.Time
	Page  []*Page
}

type Page struct {
	Totalr   int
	Services map[string][]*Service
	Hosts    []*Host
}

type Service struct {
	ServiceName  string
	ServiceCount string
}

type Host struct {
	HostUrl     string
	AddDate     string
	DetailsLink string
	Pre         string
}

//////////////////////////////////////////////////
// Constructors //////////////////////////////////
//////////////////////////////////////////////////

func NewService(name string, count string) *Service {
	service := &Service{name, count}

	return service
}

func NewPage(totalr int, servMap map[string][]*Service, hosts []*Host) *Page {
	page := &Page{totalr, servMap, hosts}

	return page
}

func NewHost(fields []string) *Host {
	host := &Host{fields[0], fields[1], fields[2], fields[3]}

	return host
}

//////////////////////////////////////////////////
// Stringer implementations for data types ///////
//////////////////////////////////////////////////

func (h *Host) String() string {
	return fmt.Sprintf("Host Url %s\n Added on %s\n Details:  %s\n\n %s\n %s\n",
		h.HostUrl, h.AddDate, h.DetailsLink, h.Pre, DELIM)
}

func (p *Page) String() string {
	var hostString string
	var servString string

	for key, value := range p.Services {
		for _, val := range value {
			servString += key + ": " + val.String() + "\n"
		}
	}

	for _, host := range p.Hosts {
		hostString += host.String() + "\n"
	}

	return fmt.Sprintf("Total found = %d\n\n Services:\n\n %s Hosts found:\n %s\n",
		p.Totalr, servString, hostString)
}

func (s *Service) String() string {
	return fmt.Sprintf("Name = %s\n Number = %s\n", s.ServiceName, s.ServiceCount)
}
