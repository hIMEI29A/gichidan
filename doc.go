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

/*
Description

Package gichidan is a console utility that is a wrapper for the Ichidan deep-web search engine.
The purpose of the application is to collect information about hosts in Tor Network, e.g. in .onion
zone.

A little information about Ichidan. The search engine is less like Google and more like Shodan,
in that it allows users to see technical information about .onion websites, including softwares names,
services details, used protocols, connected network interfaces, such as TCP/IP ports.

Details

As Ichidan is located in .onion zone too, Gichidan uses package github.com/hIMEI29A/gotorsocks
for making requests through Tor proxy.

When app receives response from search engine, it asynchronously parses all results with Golang
concurrency model, even if result's pagination contains a lot of web pages.

NEW!!! Since v1.0.0 (current) search with logical expressions is implemented. See details below.

Dependencies

    github.com/antchfx/htmlquery
    github.com/antchfx/xpath
    github.com/hIMEI29A/gotorsocks
    golang.org/x/net/html

Usage

Gichidan's CLI options are:

    gichidan
        -r <request>    search request (required)
        -s              short info (only hosts urls will be printed)
        -f <filepath>   save results to given file
        -v              version
        -h              help

You may search with app by keyword, by software name, by network protocol and by many others things.
In most cases, Gichidan cannot search by url as main search engine cannot too. But you may try it.

Examples

To get usage help, type in console:

      gichidan -h

To get current app's version number (1.0.0), try

      gichidan -v

To get info about same Ichidan server, type

      gichidan -r ichidan

Output:

    Hosts found:   1
    Only one page

    Full info:

    http://ichidanv34wrx7m7.onion
      2017-09-18 13:08:58 UTC
      tcpwrapped
      80
      tcp
      tcpwrapped
      unknown VERSION

    http-headers:

        Server: nginx/1.10.3 (Ubuntu)
        Date: Mon, 18 Sep 2017 13:08:55 GMT
        Content-Type: text/html; charset=utf-8
        Connection: close
        X-Frame-Options: SAMEORIGIN
        X-XSS-Protection: 1; mode=block
        X-Content-Type-Options: nosniff
        ETag: W/"7e087af022204d46cb9b655936aa2915"
        Cache-Control: max-age=0, private, must-revalidate
        Set-Cookie: _ichidan_session=NXQ5NWc4ZmJiSHRnVVM2TDFmblVzcmo4NnY1aUdtUFZFY0VmcVpCTz
        JHUUx2T25XOUhKa0hMT2F4QS9LanVEMGNYeXlKaEwyNGFITjA1bjdsSE1PRnR3TTIrNEJuc3dtMS9JczM1c3haL0
        xsa0U5K3E4RytSbHNWakxYVTdhYmZ3dFdhRGhzTWR4SXdlT2VhMlhFRzNRPT0tLWpiOU9SMFJnbTFXeTJFamN6Q3
        FmU3c9PQ%3D%3D--6281f0c900799f334e5f8eb76589c89c38212d37; path=/; HttpOnly
        X-Request-Id: 1e002391-0137-41e1-83cd-acc6b69b5019
        X-Runtime: 0.005388
        (Request type: HEAD)

    http-server-header:
        nginx/1.10.3 (Ubuntu)
    http-title:
        Ichidan

To collect info about .onion sites which have "paypal" keyword in metatags, and save it to file, try:

    gichidan -r paypal -f ~/my_folder/paypal_search.txt

You may want to know about .onion Raspberry Pi hosts with Raspbian OS?

    gichidan -r raspbian

There is many private XMPP(Jabber) servers in Tor network. To know about it, type in console:

    gichidan -r xmpp

Or to collect info about Prosody XMPP servers only:

    gichidan -r prosody

To run program in non-verbose ("mute") mode, use `-m` flag. GET requests messages
will not be printed in this case:

    gichidan -r accounts -m

To print oldschool ASCII banner before crawling start, use `-b` flag:

    gichidan -r ejabberd -b

If you don't want to see all details info about collected servers, use -s ("short") option:

    gichidan -r ssh -s

In case of short info and output to file mode, your file will contains all details anymore

    gichidan -r apache -s -f ~/my_folder/paypal_search.txt

Try to search by URL:

    gichidan -r facebookcorewwwi.onion

If Ichidan can not find anything by your request, application  will display error:

    gichidan -r jdfhchgbverugbvcevcegrfvcew

Output:

    2013/01/20 16:12:12 Nothing found there!

Logical operators (NEW)

Here is a simple rules for its usage:

Expression MUST contain no more than two words (_yet_) with an operator between them and
MUST NOT contain spaces between words and operator. Operators are:

    AND "+"
    NOT "-"
    OR  "="

Examples:

It will show only results which satisfy "prosody" and "ejabberd" requests both:

    gichidan -r prosody+ejabberd

It will show only results of "paypal" request wich not satisfy "crime" request:

    gichidan -r paypal-crime

It will show results of "bbs" and "telnet" requests separately:

    gichidan -r bbs=telnet

If search engine cannot find anything by one of words, application  will display error:

    gichidan -r ssh+jdfhchgbverugbvcevcegrfvcew

Notes

Tor network may be slow. In case of long delay, restart Tor:

    sudo service tor restart
*/
package main
