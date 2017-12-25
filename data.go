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

// Usefull string constants for html parsing
const (
	ADDED          string = "Added on "
	DELIM                 = "]===================================================[\n"
	LONGFORM              = "2017-09-09 01:30:35 UTC"
	PRE                   = "//pre"
	SPAN                  = "//span"
	LINK                  = "//a"
	HREF                  = "href"
	H3                    = "//h3"
	SMALL                 = "//small"
	PAGINATION            = "//div[@class='pagination']"
	TOTALR                = "//div[@class='bignumber']"
	SERVICE               = "//div[@class='service']"
	SERVICES              = "//div[@class='services']"
	SERVICENAME           = "//div[@class='span8 name']"
	SERVICECOUNT          = "//div[@class='span4 count']"
	SERVICELONG           = "//li[@class='service service-long']"
	SERVICEDETAILS        = "//div[@class='service-details col-sm-2']"
	HOST                  = "//div[@class='search-result row-fluid']"
	NORESULT              = "//div[@class='msg alert alert-info']"
	PORT                  = "//div[@class='port']"
	PROTO                 = "//div[@class='protocol']"
	STATE                 = "//div[@class='state']"
	TOPSERVICES           = "Top Services"
	TOPSOFT               = "Top Software"
	TOPSYS                = "Top Operating Systems"
)

// TOPS is a slice of string constants represents <h6> tags from Ichidan search result
var TOPS = []string{
	TOPSERVICES,
	TOPSOFT,
	TOPSYS,
}

// Data is the common data type represents whole search result
type Data struct {
	// Request string
	Query string
	// Date of our request
	Date time.Time
	// Page is a recieved web page(s)
	Page []*Page
}

// Page is the parsed data from search results one page
type Page struct {
	// Totalr is a number of found hosts
	Totalr int
	// Services is a set of found Service
	Services map[string][]*Service
	// Hosts is a set of found Hosts
	Hosts []*Host
}

// Service is the data type describes service that is active on one host
type Service struct {
	// ServiceName represents name of service
	ServiceName string
	// ServiceCount represents number of service
	ServiceCount string
}

// Host store data about one onion host
type Host struct {
	// HostUrl is an url of host
	HostUrl string
	// AddDate is an date in which host was added to Ichidan index
	AddDate string
	// DetailsLink is a link to host's detail description
	DetailsLink string
	// DetailsList is a parsed host's detail description
	DetailsList []map[string]string
	// Pre
	Pre string
}

//////////////////////////////////////////////////
// Constructors //////////////////////////////////
//////////////////////////////////////////////////

// NewService creates instance of Service
func NewService(name string, count string) *Service {
	service := &Service{name, count}

	return service
}

// NewPage creates instance of Page
func NewPage(totalr int, servMap map[string][]*Service, hosts []*Host) *Page {
	page := &Page{totalr, servMap, hosts}

	return page
}

// NewHost creates instance of Host
func NewHost(fields []string, detailsList []map[string]string) *Host {
	host := &Host{fields[0], fields[1], fields[2], detailsList, fields[3]}

	return host
}

//////////////////////////////////////////////////
// Stringer implementations for data types ///////
//////////////////////////////////////////////////

// Stringer interface implementation for Host
func (h *Host) String() string {
	return fmt.Sprintf("Host Url %s\n Added on %s\n Details:  %s\n\n %s\n %s\n",
		h.HostUrl, h.AddDate, h.DetailsLink, h.Pre, DELIM)
}

// Stringer interface implementation for Page
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

// Stringer interface implementation for Service
func (s *Service) String() string {
	return fmt.Sprintf("Name = %s\n Number = %s\n", s.ServiceName, s.ServiceCount)
}
