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

/** file data.go contains types and its methods for representing of collected data*/

package main

import (
	"encoding/json"
	"fmt"
)

// Host struct is a basic data type
type Host struct {
	// HostUrl is an url of host
	HostUrl string `json:"hosturl"`
	// AddDate is a date in which host was added to Ichidan index
	AddDate string `json:"adddate"`
	// PrimaryRequest is a request starter, e.g. search word
	PrimaryRequest string `json:"request"`
	// Services on host
	Services []*Service `json:"services"`
}

// Service contains all info about found Host
type Service struct {
	// Name is a service name: "OpenSSH" or "Apache httpd" for example
	Name string `json:"name"`
	// Port is a service listening port
	Port string `json:"port"`
	// Protocol is a service protocol
	Protocol string `json:"protocol"`
	// State is a service state: "http" or "ssh" for example
	State string `json:"state"`
	// Version is a service version if parsed
	Version string `json:"version"`
	// ServDetails is a <pre> tag's content of parsed page
	ServDetails string `json:"servdetails"`
}

// NewService is a constructor for Service struct
func NewService(fields []string) *Service {
	service := &Service{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5]}

	return service
}

// NewHost is a constructor for Host struct
func NewHost(fields []string, services []*Service) *Host {
	host := &Host{fields[0], fields[1], fields[2], services}

	return host
}

// String is a Stringer implementation for Service to output
func (s *Service) String() string {
	return fmt.Sprintf("%s\n %s\n %s\n %s\n %s\n %s\n",
		s.Name, s.Port, s.Protocol, s.State, s.Version, s.ServDetails)
}

// String is a Stringer implementation for Host to output
func (h *Host) String() string {
	var servs string

	for _, s := range h.Services {
		servs += s.String() + "\n"
	}

	return fmt.Sprintf("%s\n %s\n %s\n", h.HostUrl, h.AddDate, servs)
}

// HostToJson converts output to JSON
func (host *Host) hostToJson() []byte {
	nosj, err := json.Marshal(host)
	ErrFatal(err)

	return nosj
}

/*
func (host *Host) jsonToHost(jsoned []byte) *Host {
	h := &Host{}
	err := json.Unmarshal(jsoned, h)
	ErrFatal(err)

	return h
}
*/
