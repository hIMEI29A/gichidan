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
	//"strconv"
	"time"

	"golang.org/x/net/html"
)

const (
	CLINK        string = "</a>"
	PRE                 = "pre"
	CPRE                = "</pre>"
	SPACE               = " "
	BACKSEP             = "</"
	ADDED               = "Added on "
	LONGFORM            = "2017-09-09 01:30:35 UTC"
	COLXS12             = "\"col-xs-12\""
	ONION               = "\"onion\""
	ONIONHREF           = "\"/onions/"
	COLXS8              = "\"col-xs-8\""
	SPAN                = "//span"
	INDEXSUMMARY        = "//div[@class='indexsummary']"
	SUMMARY             = "//div[@class='search-result-summary col-xs-4']"
	LINK                = "//a"
	HREF                = "//a[@href]"
	DETAILS             = "//a[@class='details']"
	PAGINATION          = "//div[@class='pagination']"
	TOTALR              = "//div[@class='bignumber']"
	SERVICE             = "//div[@class='service']"
	SERVICES            = "//div[@class='services']"
	SERVICENAME         = "//div[@class='span8 name']"
	SERVICECOUNT        = "//div[@class='span4 count']"
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
	Pages *Page
}

type Page struct {
	IndSummary *IndexSummary
	Summarys   []*Summary
	//Pagination string
}

type IndexSummary struct {
	Totalr string
	Servs  map[string]*Services
}

type Services struct {
	Serv []*Service
}

type Service struct {
	ServiceName  string
	ServiceCount string
}

type Summary struct {
	HostUrl     string
	Date        time.Time
	DetailsLink string
	//Pre         string
}

// Constructors //////////////////////////////

func NewData(query string, page *Page) *Data {
	data := &Data{query, page}

	return data
}

func NewPage(ind *IndexSummary, fields map[int][]string) *Page {
	var summarys []*Summary

	for _, value := range fields {
		summary := NewSummary(value)
		summarys = append(summarys, summary)
	}

	page := &Page{ind, summarys}

	return page
}

func NewService(name string, count string) *Service {
	service := &Service{name, count}

	return service
}

func NewServices(serv []*Service) *Services {
	services := &Services{serv}

	return services
}

func NewSummary(fields []string) *Summary {
	summary := &Summary{fields[0], getTime(fields[1]), fields[2]}

	return summary
}

func NewIndSummary(totalr string, svs map[string]*html.Node) *IndexSummary {
	var servs map[string]*Services
	var serv []*Service

	for key, svsnode := range svs {
		servnodes := getServiceNodes(svsnode)

		for _, svnode := range servnodes {
			name := getServiceName(svnode)
			count := getServiceCount(svnode)
			service := NewService(name, count)
			serv = append(serv, service)
			s := NewServices(serv)

			servs[key] = s
		}
	}

	ind := &IndexSummary{totalr, servs}

	return ind
}

// Stringer Implementations of Stringer interface for all data types ///////////////////

func (d *Data) String() string {
	return fmt.Sprintf("Query = %s\n, Page = %+v\n", d.Query, d.Pages)
}

func (s *Summary) String() string {
	return fmt.Sprintf("HostUrl = %s\n, Date = %s\n, DetailsLink = %s\n",
		s.HostUrl, s.Date.String(), s.DetailsLink)
}

func (s *Service) String() string {
	return fmt.Sprintf("ServiceName = %s\n, ServiceCount = %s\n", s.ServiceName, s.ServiceCount)
}

func (ss *Services) String() string {
	var servstring string

	for _, s := range ss.Serv {
		servstring += s.String() + ", "
	}

	return fmt.Sprintf("%s\n", servstring)
}

func (i *IndexSummary) String() string {
	var servsstring string

	for key, value := range i.Servs {
		servsstring += key + ": " + value.String() + "\n"
	}

	return fmt.Sprintf("Totalr = %s\n, Services = %s\n", i.Totalr, servsstring)
}

func (p *Page) String() string {
	var summstring string

	for _, s := range p.Summarys {
		summstring += s.String() + "\n"
	}

	return fmt.Sprintf("IndSumm = %+v\n, Summ = %s\n", p.IndSummary, p.Summarys)
}
