# gichidan

**gichidan** - command line wrapper with enhansed pentest features for 
(_onion link_) [**Ichidan**](http://ichidanv34wrx7m7.onion) - deep-web search engine.

               ███           █████       ███      █████                     
              ░░░           ░░███       ░░░      ░░███                    
      ███████ ████   ██████  ░███████   ████   ███████   ██████   ████████  
     ███░░███░░███  ███░░███ ░███░░███ ░░███  ███░░███  ░░░░░███ ░░███░░███ 
    ░███ ░███ ░███ ░███ ░░░  ░███ ░███  ░███ ░███ ░███   ███████  ░███ ░███ 
    ░███ ░███ ░███ ░███  ███ ░███ ░███  ░███ ░███ ░███  ███░░███  ░███ ░███ 
    ░░███████ █████░░██████  ████ █████ █████░░████████░░████████ ████ █████
     ░░░░░███░░░░░  ░░░░░░  ░░░░ ░░░░░ ░░░░░  ░░░░░░░░  ░░░░░░░░ ░░░░ ░░░░░ 
     ███ ░███           ___onion secrets for console cowboys___
    ░░██████
    ░░░░░░

[![Go Report Card](https://goreportcard.com/badge/github.com/hIMEI29A/gichidan)](https://goreportcard.com/report/github.com/hIMEI29A/gichidan) [![GoDoc](https://godoc.org/github.com/hIMEI29A/gichidan?status.svg)](http://godoc.org/github.com/hIMEI29A/gichidan) [![Apache-2.0 License](https://img.shields.io/badge/license-Apache--2.0-red.svg)](LICENSE)

Copyright 2017 hIMEI


## TOC
- [About](#about)
- [Features](#features)
- [Version](#version)
- [Install](#install)
- [Usage](#usage)
- [TODO](#todo)
- [Contributing](#contributing)

## About

Forget about Tor Browser. Parse onion hosts from your console with **Gichidan** now.

##### License

Apache-2.0 License

##### About Gichidan

**Gichidan** is a CLI utility designed to collect information about deep-web hosts.

###### Dependencies

    github.com/antchfx/htmlquery
    github.com/antchfx/xpath
    github.com/hIMEI29A/gotorsocks
    golang.org/x/net/html

###### Short Ichidan's info 

Short info about Ichidan search engine from [here](https://www.cylance.com/en_us/blog/ichidan-a-search-engine-for-the-dark-web.html)

> Ichidan is a type of Japanese verb which implies the first (“ichi”) time something is done. Now, Ichidan is also a search engine for looking up websites that are hosted through the Tor network, which may be the first time that's been done at this scale.

> The search engine is less like Google and more like Shodan, in that it allows users to see technical information about .onion websites, including their connected network interfaces, such as TCP/IP ports.

> Ichidan is a valuable resource for security researchers and law enforcement agencies who want to learn about what's happening on the Dark Web.

## Features

**NEW!** Since version 1.0.0 (current) search with logical expressions supported.
See **Usage** section of this paper for details.

## Version

**v1.0.0**

## Install

Check the [release page](https://github.com/hIMEI29A/gichidan/releases)!

###### Install from source

Progect uses `glide` to manage dependencies, so install it first

```sh
curl https://glide.sh/get | sh
```
Clone repo, install deps, then install **Gichidan**

```sh
mkdir -p $GOPATH/src/github.com/hIMEI29A/gichidan
cd $GOPATH/src/github.com/hIMEI29A/gichidan
git clone https://github.com/hIMEI29A/gichidan.git .
glide install
go install
```

## Usage

Gichidan's CLI options are:

    -b    show ASCII banner
    -f string
          save results to file
    -h    help message
    -m    Don't print GET request's messages (non-verbose output)
    -r string
          your search request to Ichidan
    -s    print hosts urls only
    -v    print current version

Typical request to Ichidan looks like

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
        Set-Cookie: _ichidan_session=NXQ5NWc4ZmJiSHRnVVM2TDFmblVzcmo4NnY1aUdtUFZFY0VmcVpCTzJHUUx2T25XOUhKa0hMT2F4QS9LanVEMGNYeXlKaEwyNGFITjA1bjdsSE1PRnR3TTIrNEJuc3dtMS9JczM1c3haL0xsa0U5K3E4RytSbHNWakxYVTdhYmZ3dFdhRGhzTWR4SXdlT2VhMlhFRzNRPT0tLWpiOU9SMFJnbTFXeTJFamN6Q3FmU3c9PQ%3D%3D--6281f0c900799f334e5f8eb76589c89c38212d37; path=/; HttpOnly
        X-Request-Id: 1e002391-0137-41e1-83cd-acc6b69b5019
        X-Runtime: 0.005388
    
        (Request type: HEAD)
    
    http-server-header:
        nginx/1.10.3 (Ubuntu)
    http-title:
        Ichidan

You may search by keywords (only to know what bad guys do):

    gichidan -r hacking

    gichidan -r paypal

As well as by protocol, application name or service detail:

    gichidan -r ssh

    gichidan -r irc

    gichidan -r apache

    gichidan -r tcpwrapped

    gichidan -r prosody

    gichidan -r raspbian

To save results in file use flag `-f` with full file path followed:

    gichidan -r telnet -f ~/my_folder/telnet_search.txt

If you don't want to see all details info about collected servers, use `-s` ("short") option. In case of short info and output to file mode, your file will contains all details anymore: 

    gichidan -r apache -s -f ~/my_folder/paypal_search.txt

To run program in non-verbose ("mute") mode, use `-m` flag. GET requests messages will not be printed in this case:

    gichidan -r accounts -m

To print oldschool ASCII banner before crawling start, use `-b` flag:

    gichidan -r ejabberd -b

If Ichidan can not find anything by your request, application  will display error:

    gichidan -r jdfhchgbverugbvcevcegrfvcew

Output:
    
    2013/01/20 16:12:12 Nothing found there!

###### Logical expressions

**NEW!** Since version 
v1.0.0 (current) search with logical expressions supported. Here is a simple rules for its usage:

Expression MUST contain no more than two words (_yet_) with an operator between them and MUST NOT contain spaces between words and operator. Operators are:

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

**Request MUST NOT contain spaces**. In case of request such as `gichidan -r prosody client`, only first word will be processed. Also search by host url is not supported (in most case) by app (and Ichidan too). 

**NOTE:** Tor Network it is not your vanilla Internet. It may be unstable or slow and there may be unexpected delays and errors. In this case you may try to simply restart tor service on your mashine:

    sudo service tor restart

**NOTE:** Ichidan it is not your vanilla Google, Yandex or Baidu. On its [page](http://ichidanv34wrx7m7.onion) you wont even find contact info or credits. In first january days of new 2018 it was absolutely unavailable! So there is no guarantee to recieve any response! 

## TODO

* Tests!!!
* Ichidan's authorisation support
* Third party tools for possible discovery of found hosts

## Contributing

Feel free to contribute!
