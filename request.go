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
				primr = append(primr, s)
			}
		}
	}

	return fullr, primr
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
	case string(req[0]) == "/":
		reqString := "GET " + req + "\n"
		fullRequest = append(fullRequest, reqString)

	case string(req[0]) != "/" &&
		string(req[0]) != NONE &&
		strings.Contains(req, NONE) == true:

		splitted := strings.Split(req, NONE)
		reqString := "GET " + SEARCH + splitted[0] + "\n"
		fullRequest = append(fullRequest, reqString)

	case strings.Contains(req, AND) == true:
		fullr, primr := makeLogicRequest(req)
		fullRequest = fullr
		primStrings = primr
		op = AND

	case strings.Contains(req, OR) == true:
		fullr, primr := makeLogicRequest(req)
		fullRequest = fullr
		primStrings = primr

		op = OR

	case strings.Contains(req, NOT) == true:
		fullr, primr := makeLogicRequest(req)
		fullRequest = fullr
		primStrings = primr

		op = NOT

	default:
		reqString := "GET " + SEARCH + req + "\n"
		fullRequest = append(fullRequest, reqString)
	}

	request.RequestStrings = fullRequest
	request.PrimaryStrings = primStrings
	request.Operator = op

	return request
}

// sortResult sorts received hosts by its primary request's strings
func (r *Request) sortResult(hosts []*Host) map[string][]*Host {
	hostMap := make(map[string][]*Host)

	if r.PrimaryStrings != nil {
		for _, h := range hosts {
			if h.PrimaryRequest == r.PrimaryStrings[0] {
				hostMap[r.PrimaryStrings[0]] = append(hostMap[r.PrimaryStrings[0]], h)
			}

			if h.PrimaryRequest == r.PrimaryStrings[1] {
				hostMap[r.PrimaryStrings[1]] = append(hostMap[r.PrimaryStrings[1]], h)
			}
		}
	} else {
		hostMap["total"] = hosts
	}

	return hostMap
}

// InRange checks if given slice of *Host contains given Host
func (r *Request) inRange(host *Host, hosts []*Host) bool {
	check := false

	for _, h := range hosts {
		if host.HostUrl == h.HostUrl {
			check = true
			break
		}
	}

	return check
}

// ResultProvider makes logical operations NOT, OR and AND against found hosts.
func (r *Request) resultProvider(hosts []*Host) []*Host {
	var finalHosts []*Host

	hostMap := r.sortResult(hosts)

	if hostMap[r.PrimaryStrings[0]] != nil && hostMap[r.PrimaryStrings[1]] != nil {
		switch {
		case r.Operator == AND:
			for _, h := range hostMap[r.PrimaryStrings[0]] {
				if r.inRange(h, hostMap[r.PrimaryStrings[1]]) == true {
					finalHosts = append(finalHosts, h)
				}
			}

		case r.Operator == NOT:
			for _, h := range hostMap[r.PrimaryStrings[0]] {
				if r.inRange(h, hostMap[r.PrimaryStrings[1]]) == false {
					finalHosts = append(finalHosts, h)
				}
			}

		case r.Operator == OR:
			for _, h := range hostMap[r.PrimaryStrings[0]] {
				finalHosts = append(finalHosts, h)
			}

			for _, hh := range hostMap[r.PrimaryStrings[1]] {
				finalHosts = append(finalHosts, hh)
			}
		}
	}

	return finalHosts
}
