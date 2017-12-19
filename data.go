//
// @Author: hIMEI
// @Date:   2017-12-17 20:34:41
// @Last Modified by:   hIMEI
// @Last Modified time: 2017-12-17 20:34:41

package main

import (
	"fmt"
	"time"
)

type Data struct {
	Query string
	Pages *Page
}

func (d *Data) String() string {
	return fmt.Sprintf(" Query = %v\n, Page = %v\n", d.Query, d.Pages)
}

type Page struct {
	Headers6   []*Header6
	Summarys   []*Summary
	Pagination string
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
