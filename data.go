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
)

const (
	H6           string = "//h6"
	CH6                 = "</h6>"
	DIV                 = "<div>"
	CDIV                = "</div>"
	SPAN                = "span"
	LESSER              = "<"
	GREATER             = ">"
	LINK                = "//a"
	CLINK               = "</a>"
	HREF                = "//a[@href]"
	PRE                 = "pre"
	CPRE                = "</pre>"
	SPACE               = " "
	BACKSEP             = "</"
	ADDED               = "Added on "
	LONGFORM            = "2017-09-09 01:30:35 UTC"
	TOTALR              = "Total Results"
	COLXS12             = "\"col-xs-12\""
	SERVICE             = "\"service\""
	SERVICENAME         = "//div[@class='span8 name']" ////div[@id='merchant-info']/a
	SERVICECOUNT        = "//div[@class='span4 count']"
	SUMMARY             = "//div[@class='search-result-summary col-xs-4']"
	ONION               = "\"onion\""
	ONIONHREF           = "\"/onions/"
	COLXS8              = "\"col-xs-8\""
	DETAILS             = "//a[@class='details']"
	PAGINATION          = "//div[@class='pagination']"
)

type Data struct {
	Query string
	Pages *Page
}

type Page struct {
	Headers6 []*Header6
	Summarys []*Summary
	//Pagination string
}

type Header6 struct {
	Name     string
	Services []*Service
}

type Service struct {
	ServiceName  string
	ServiceCount int
}

type Summary struct {
	HostUrl     string
	Date        time.Time
	DetailsLink string
	Pre         string
}

// Constructors //////////////////////////////

func NewData(query string, page *Page) *Data {
	data := &Data{query, page}

	return data
}

func NewHeader6(name string, services []*Service) *Header6 {
	header := &Header6{name, services}

	return header
}

func NewPage(h6names []string, serviceFields map[string][]string, summaryFields map[int][]string) *Page {
	var services []*Service
	var summarys []*Summary
	var headers []*Header6

	for _, value := range serviceFields {
		serv := NewService(value)
		services = append(services, serv)
	}

	if len(h6names) == len(services) {
		for _, name := range h6names {
			header := NewHeader6(name, services)
			headers = append(headers, header)
		}
	}

	for _, value := range summaryFields {
		summ := NewSummary(value)
		summarys = append(summarys, summ)
	}

	page := &Page{headers, summarys /*pagination*/}

	return page
}

func NewService(fields []string) *Service {
	service := &Service{fields[0], toInt(fields[1])}

	return service
}

func NewSummary(fields []string) *Summary {
	summary := &Summary{fields[0], getTime(fields[1]), fields[2], fields[3]}

	return summary
}

// Stringer implementations for all types ///////////////////

func (d *Data) String() string {
	return fmt.Sprintf("Query = %s\n, Page = %v\n", d.Query, d.Pages)
}

func (s *Summary) String() string {
	return fmt.Sprintf("HostUrl = %s\n, Date = %s\n, DetailsLink = %s\n, Pre = %s\n",
		s.HostUrl, s.Date.String(), s.DetailsLink, s.Pre)
}

func (s *Service) String() string {
	return fmt.Sprintf("ServiceName = %s\n, ServiceCount = %i\n", s.ServiceName, s.ServiceCount)
}

func (h *Header6) String() string {
	var strserv string
	serv := func(s []*Service) string {
		for _, service := range s {
			strserv += service.String() + ", "
		}

		return fmt.Sprintf("Header %s = \n %s", h.Name, strserv)
	}

	return fmt.Sprintf("Name = %s\n, Services = %v\n", h.Name, serv(h.Services))
}

func (p *Page) String() string {
	var strhead, strsumm string
	head := func(h []*Header6) string {
		for _, header := range h {
			strhead += header.String() + ", "
		}

		return fmt.Sprintf("Headers = %s", strhead)
	}

	sum := func(s []*Summary) string {
		for _, summ := range s {
			strsumm += summ.String() + ", "
		}

		return fmt.Sprintf("Summarys = %s", strsumm)
	}

	return fmt.Sprintf("Headers6 = %v\n, Summarys = %v\n", head(p.Headers6), sum(p.Summarys))
}
