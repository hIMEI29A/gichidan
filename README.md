# gichidan

**gichidan** - command line wrapper to 
(warning: _onion link_) [**Ichidan**](http://ichidanv34wrx7m7.onion) deep-web search engine with enhansed pentest features (TODO).

     ____, __, ____,__, _,__, ____,____,_,  _,
    (-/ _,(-| (-/  (-|__|(-| (-|  (-/_|(-|\ | 
     _\__| _|_,_\__,_|  |,_|_,_|__//  |,_| \|,
    (     (   (    (     (   (   (     (      

[![Go Report Card](https://goreportcard.com/badge/github.com/hIMEI29A/gichidan)](https://goreportcard.com/report/github.com/hIMEI29A/gichidan) [![GoDoc](https://godoc.org/github.com/hIMEI29A/gichidan?status.svg)](http://godoc.org/github.com/hIMEI29A/gichidan)

Copyright 2017 hIMEI


## TOC

1. About
* License
* About **Gichidan**
* Features
* Dependencies
* Short Ichidan's info 
2. Version
3. Install
4. Usage

## About


##### License

[![Apache-2.0 License](http://img.shields.io/badge/License-Apache-2.0-yellow.svg)](LICENSE)

##### About Gichidan

**Gichidan** is a CLI utility designed to collect information about deep-web hosts.

###### Features


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

## Install

Check the [release page](https://github.com/hIMEI29A/gichidan/releases)!

Progect uses `glide` to manage dependencies, so install it first

```sh
mkdir -p $GOPATH/src/github.com/hIMEI29A/gichidan
cd $GOPATH/src/github.com/hIMEI29A/gichidan
git clone https://github.com/hIMEI29A/gichidan.git .
glide install
go install
```

You may **install** Gichidan also by `go install`:

```sh
go get github.com/hIMEI29A/gichidan
```

## Usage

Type in your console to get help:

```sh
gichidan -h
```

You will see next message:

             Usage: 
            gichidan <command> [<args>] [options]
    Commands:       search
    Args:       -r  your search request to Ichidan
    Options:
        
    -h        help message
    -v        prints current version

Typical request to Ichidan may looks such as

    gichidan search -r ichidan

    Hosts found:   1 
    Only one page 
    parsed   http://ichidanv34wrx7m7.onion
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


You may search by keywords (to only know what bad guys do):

    gichidan search -r drugs

    gichidan search -r guns

As well as by protocol, application name or service detail:

    gichidan search -r ssh

    gichidan search -r irc

    gichidan search -r apache

    gichidan search -r prosody

    gichidan search -r raspbian

If Ichidan can not find anything by your request, application  will display error:

    gichidan search -r jdfhchgbverugbvcevcegrfvcew

Output:
    
    2013/01/20 16:12:12 Nothing found there, Neo!

In current version (0.1.0) request must not contains space. In case of request such as `gichidan search -r prosody client`, only first word will be processed.
Option with compound requests and logical operators will be implemented in future.

**NOTE:** Tor Network it is not your vanilla Internet. It may be unstable or slow and there may be unexpected delays and errors. In this case you may simply restart tor service on your mashine:

    sudo service tor restart

**NOTE:** Ichidan it is not your vanilla Google, Yandex or Baidu. On its [page](http://ichidanv34wrx7m7.onion) you wont even find contact info or credits. In first january days of new 2018 it even was absolutely unavailable! So there is possible to recieve not any response! 
