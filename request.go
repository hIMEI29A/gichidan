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

/** file request.go contains data types and methods for http requests creating */

package main

import (
	"strings"
)

// MakeLogicRequest parses given string for logical operators and create request with its if found
func makeLogicRequest(req string) ([]string, []string) {
	var fullr, primr []string

	for _, l := range LOGIC {
		if strings.Contains(req, l) == true {
			splitted := strings.Split(req, l)

			for _, s := range splitted {
				reqString := "GET " + SEARCH + s + "\n"
				fullr = append(fullr, reqString)
				primr = append(primr, trimString(s))
			}
		}
	}

	return fullr, primr
}

// TrimUrl takes string as argument and cuts everything but primary request
func trimUrl(url string) string {
	splitted := strings.Split(url, "query")
	primary := trimString(strings.TrimPrefix(splitted[1], "="))

	return primary
}

// Request is a data type for representing requests to search engine
type Request struct {
	// RequestStrings contains prepared GET request(s)
	RequestStrings []string
	// PrimaryStrings contains searched words
	PrimaryStrings []string
	// Operator is a logical operator (NOT, OR and AND ) in case of logic request
	Operator string
}

// NewRequest creates instance of Request type
func NewRequest(req string) *Request {
	request := &Request{}
	var fullRequest []string
	var primStrings []string
	var op string

	switch {
	// Case for program's inner logic
	case string(req[0]) == "/":
		reqString := "GET " + req + "\n"
		fullRequest = append(fullRequest, reqString)

	// Case for program's inner logic
	case string(req[0]) != "/" &&
		string(req[0]) != NONE &&
		strings.Contains(req, NONE) == true:

		splitted := strings.Split(req, NONE)
		reqString := "GET " + SEARCH + splitted[0] + "\n"
		fullRequest = append(fullRequest, reqString)

	// Search with operators
	case strings.Contains(req, AND) == true:
		fullr, primr := makeLogicRequest(req)
		fullRequest = fullr
		primStrings = primr

		op = AND

	// Search with operators
	case strings.Contains(req, OR) == true:
		fullr, primr := makeLogicRequest(req)
		fullRequest = fullr
		primStrings = primr

		op = OR

	// Search with operators
	case strings.Contains(req, NOT) == true:
		fullr, primr := makeLogicRequest(req)
		fullRequest = fullr
		primStrings = primr

		op = NOT

	// Default is a case without search operators, e.g. "gichidan -r ichidan"
	default:
		reqString := "GET " + SEARCH + req + "\n"
		fullRequest = append(fullRequest, reqString)
		primStrings = append(primStrings, req)
	}

	request.RequestStrings = fullRequest
	request.PrimaryStrings = primStrings
	request.Operator = op

	return request
}

// InRange checks if given slice of *Host contains given Host
func (r *Request) inRange(host *Host, hosts []*Host) bool {
	check := false

	for i := range hosts {
		if hosts[i].HostUrl == host.HostUrl {
			check = true
			break
		}
	}

	return check
}

// SortResult sorts received hosts by its primary request's strings
func (r *Request) splitResult(hosts []*Host) chan []*Host {
	// Channel for output
	chHosts := make(chan []*Host, 2)

	go func() {
		var (
			hostsFirst []*Host
			hostsSec   []*Host
		)

		if len(r.PrimaryStrings) > 1 {
			for i := range hosts {
				if hosts[i].PrimaryRequest == r.PrimaryStrings[0] {
					hostsFirst = append(hostsFirst, hosts[i])
				}

				if hosts[i].PrimaryRequest == r.PrimaryStrings[1] {
					hostsSec = append(hostsSec, hosts[i])
				}
			}
		}
		if len(r.PrimaryStrings) == 1 {
			hostsFirst = hosts
		}

		chHosts <- hostsFirst
		chHosts <- hostsSec
	}()

	return chHosts
}

// ResultProvider makes logical operations NOT, OR and AND against found hosts.
func (r *Request) resultProvider(hosts []*Host) []*Host {
	var finalHosts []*Host

	chHosts := r.splitResult(hosts)
	//hostsFirst, hostsSec := r.sortResult(hosts)
	hostsFirst := <-chHosts
	hostsSec := <-chHosts

	if len(hostsSec) != 0 {
		switch {
		case r.Operator == AND:
			for i := range hostsFirst {
				if r.inRange(hostsFirst[i], hostsSec) == true {
					finalHosts = append(finalHosts, hostsFirst[i])
				}
			}

		case r.Operator == NOT:
			for i := range hostsFirst {
				if r.inRange(hostsFirst[i], hostsSec) == false {
					finalHosts = append(finalHosts, hostsFirst[i])
				}
			}

		case r.Operator == OR:
			for i := range hostsFirst {
				finalHosts = append(finalHosts, hostsFirst[i])
			}

			for i := range hostsSec {
				finalHosts = append(finalHosts, hostsSec[i])
			}
		}
	} else {
		finalHosts = hosts
	}

	return finalHosts
}
