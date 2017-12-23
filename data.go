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
	//"strconv"
	//"time"
	//"golang.org/x/net/html"
)

const (
	ADDED        string = "Added on "
	LONGFORM            = "2017-09-09 01:30:35 UTC"
	PRE                 = "//pre"
	SPAN                = "//span"
	LINK                = "//a"
	HREF                = "href"
	DETAILS             = "//a[@class='details']"
	PAGINATION          = "//div[@class='pagination']"
	TOTALR              = "//div[@class='bignumber']"
	SERVICE             = "//div[@class='service']"
	SERVICES            = "//div[@class='services']"
	SERVICENAME         = "//div[@class='span8 name']"
	SERVICECOUNT        = "//div[@class='span4 count']"
	HOST                = "//div[@class='search-result row-fluid']"
	TOPSERVICES         = "Top Services"
	TOPSOFT             = "Top Software"
	TOPSYS              = "Top Operating Systems"
)

var TOPS = []string{
	//TOPRES,
	TOPSERVICES,
	TOPSOFT,
	TOPSYS,
}

var parser *Parser

type Data struct {
	Query string
	Date  time.Time
	Page  *Page
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

// Constructors //////////////////////////////////

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

// Stringer implementations for data types ///////

func (h *Host) String() string {
	return fmt.Sprintf("Host Url = %s\n Added on = %s\n Details Link = %s\n Pre = %s\n",
		h.HostUrl, h.AddDate, h.DetailsLink, h.Pre)
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

	return fmt.Sprintf("Total found = %d\n Services:\n %s Hosts found:\n %s\n",
		p.Totalr, servString, hostString)
}

func (s *Service) String() string {
	return fmt.Sprintf("ServiceName = %s\n, ServiceCount = %s\n", s.ServiceName, s.ServiceCount)
}
