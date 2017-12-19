//
// @Author: hIMEI
// @Date:   2017-12-17 21:29:46
// @Last Modified by:   hIMEI
// @Last Modified time: 2017-12-17 21:29:46

package main

import (
	//"net/url"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	H6           string = "<h6>"
	CH6                 = "</h6>"
	DIV                 = "<div>"
	CDIV                = "</div>"
	LESSER              = "<"
	GREATER             = ">"
	LINK                = "<a "
	CLINK               = "</a>"
	HREFLINK            = "<a href=\""
	PRE                 = "<pre>"
	CPRE                = "</pre>"
	SPACE               = " "
	BACKSEP             = "</"
	ADDED               = "<span>Added on "
	CADDED              = "</span><br"
	LONGFORM            = "2017-09-09 01:30:35 UTC"
	TOTALR              = "Total Results"
	COLXS12             = "<div class=\"col-xs-12\""
	SERVICE             = "<div class=\"service\""
	SERVICENAME         = "<div class=\"span8 name\""
	SERVICECOUNT        = "<div class=\"span4 count\""
	SUMMARY             = "<div class=\"search-result-summary col-xs-4\""
	ONION               = "<div class=\"onion\""
	COLXS8              = "<div class=\"col-xs-8\""
	DETAILSPREF         = "<a class=\"details\" href=\"/onions/"
	DETAILSSUFF         = "\">Details</a>"
	PAGINATION          = "<div class=\"pagination\">"
)

type Parser struct {
}

func ParseData(data []string) *Data {
	query := "paypal"
	pages := NewPage(data)

	parsedData := &Data{query, pages}

	return parsedData
}

func NewPage(data []string) *Page {
	var headers []*Header6
	var summarys []*Summary
	pagination := getPagination(data)
	h6bodies := getH6bodies(data)

	for key, value := range h6bodies {
		services := newServices(value)
		headername := key
		header6 := newHeader6(headername, services)
		headers = append(headers, header6)

		summs := getSummarysBodies(data)

		for _, sum := range summs {
			summary := newSummary(sum)
			summarys = append(summarys, summary)
		}
	}

	page := &Page{headers, summarys, pagination}

	return page
}

func findIndex(data []string, line string) int {
	index := -1
	for i, str := range data {
		if str == line {
			index = i
			return index
			break
		}
	}

	return index
}

func getH6name(data string) string {
	splitted := strings.Split(data, BACKSEP)
	h6Name := strings.TrimPrefix(splitted[0], H6)
	return h6Name
}

func splitText(line string) []string {
	text := strings.Split(line, "\n")
	return text
}

func getH6names(data []string) []string {
	var h6names []string
	for _, str := range data {
		if strings.Contains(str, H6) == true {
			h6Name := getH6name(str)
			h6names = append(h6names, h6Name)
		}
	}

	return h6names
}

func getH6bodies(data []string) map[string][]string {
	h6bodies := make(map[string][]string)
	dataSplitted := strings.Split(strings.Join(data, SPACE), H6)

	for i := 1; i < len(dataSplitted); i++ {
		name := getH6name(dataSplitted[i])
		nameString := H6 + name + CH6
		bodyString := strings.TrimPrefix(dataSplitted[i], nameString)
		bodySplitted := splitText(bodyString)
		h6bodies[name] = bodySplitted
	}

	return h6bodies
}

func newServices(h6body []string) []*Service {
	var services []*Service
	bodys := strings.Split(strings.Join(h6body, SPACE), SERVICE)
	for i, str := range bodys {
		bodySplitted := splitText(str)
		if i != 0 && findIndex(bodySplitted, SERVICE) != -1 {
			namestr := bodySplitted[findIndex(bodySplitted, SERVICENAME)+1]
			nameSplitted := strings.Split(namestr, GREATER)
			name := strings.TrimSuffix(nameSplitted[1], CLINK)

			countstr := bodySplitted[findIndex(bodySplitted, SERVICECOUNT)+1]
			count, err := strconv.Atoi(countstr)
			ErrFatal(err)

			service := &Service{name, count}
			services = append(services, service)

		}

		return services
	}

	return services
}

func newHeader6(headername string, services []*Service) *Header6 {
	header6 := &Header6{headername, services}

	return header6
}

func getSummarysBodies(data []string) []string {
	dataSplitted := strings.Split(strings.Join(data, SPACE), SUMMARY)

	return dataSplitted
}

func getHostUrl(body string) string {
	bodySplitted := strings.Split(body, GREATER)
	namestr := bodySplitted[findIndex(bodySplitted, ONION)+1]
	urlTrimmed := strings.TrimSuffix(namestr, HREFLINK)
	urlSplitted := strings.Split(urlTrimmed, "\"")
	hostUrl := urlSplitted[0]

	return hostUrl
}

func getDate(body string) time.Time {
	bodySplitted := strings.Split(body, GREATER)
	daterStr := bodySplitted[findIndex(bodySplitted, ONION)+3]
	dateTrimmed := strings.TrimPrefix(daterStr, ADDED)
	dateInTrimmed := strings.TrimSuffix(dateTrimmed, CADDED)
	date, _ := time.Parse(LONGFORM, dateInTrimmed)

	return date
}

func getDetailsLink(body string) string {
	bodySplitted := strings.Split(body, GREATER)

	//	for _, w := range bodySplitted {
	//		fmt.Println(w)
	//	}

	index := findIndex(bodySplitted, COLXS8)
	fmt.Println(index)

	detailsLink := bodySplitted[index-2]
	detailsTrimmed := strings.TrimPrefix(detailsLink, DETAILSPREF)
	details := strings.TrimSuffix(detailsTrimmed, DETAILSSUFF)

	return details
}

func getPre(body string) string {
	preSplitted := strings.Split(body, PRE)
	preStr := preSplitted[1]
	prePostSplitted := strings.Split(preStr, CPRE)
	pre := prePostSplitted[0]

	return pre
}

func newSummary(body string) *Summary {
	hostUrl := getHostUrl(body)
	date := getDate(body)
	details := getDetailsLink(body)
	pre := getPre(body)

	summary := &Summary{hostUrl, date, details, pre}

	return summary
}

func getPagination(data []string) string {
	pagination := ""

	if findIndex(data, PAGINATION) != -1 {
		dataSplitted := strings.Split(strings.Join(data, SPACE), PAGINATION)
		pagination = strings.TrimSuffix(dataSplitted[1], CDIV)
	}

	return pagination

}
