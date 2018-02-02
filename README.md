# gichidan

**gichidan** - command line wrapper to 
(warning: _onion link_) [**Ichidan**](http://ichidanv34wrx7m7.onion) deep-web search engine with enhansed pentest features (TODO).

               ███           █████       ███      █████                     
      v1.0.0  ░░░           ░░███       ░░░      ░░███            © hIMEI
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

1. About
* License
* About **Gichidan**
* Features
* Dependencies
* Short Ichidan's info 
2. Version
3. Install
4. Usage
5. TODO
6. Contributing

## About


##### License

Apache-2.0 License

##### About Gichidan

**Gichidan** is a CLI utility designed to collect information about deep-web hosts.

###### Features

See **Usage** section of this paper.

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

## Version

**0.1.1**

## Install

Check the [release page](https://github.com/hIMEI29A/gichidan/releases)!

Progect uses `glide` to manage dependencies, so install it first

```sh
curl https://glide.sh/get | sh
```

```sh
mkdir -p $GOPATH/src/github.com/hIMEI29A/gichidan
cd $GOPATH/src/github.com/hIMEI29A/gichidan
git clone https://github.com/hIMEI29A/gichidan.git .
glide install
go install
```

## Usage

Gichidan's CLI options are:

    gichidan
        -r <request>    search request (required)
        -s              short info (only hosts urls will be printed)
        -f <filepath>   save results to given file
        -v              version
        -h              help

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

    gichidan -r prosody

    gichidan -r raspbian

If Ichidan can not find anything by your request, application  will display error:

    gichidan -r jdfhchgbverugbvcevcegrfvcew

Output:
    
    2013/01/20 16:12:12 Nothing found there, Neo!

In current version (0.1.1) **request must not contains spaces**. In case of request such as `gichidan search -r prosody client`, only first word will be processed. Also search by host url is not supported (in most case) by app (and Ichidan too). 
Options with **compound requests**, **search by url** and **search with logical operators** will be implemented in future.

**NOTE:** Tor Network it is not your vanilla Internet. It may be unstable or slow and there may be unexpected delays and errors. In this case you may try to simply restart tor service on your mashine:

    sudo service tor restart

**NOTE:** Ichidan it is not your vanilla Google, Yandex or Baidu. On its [page](http://ichidanv34wrx7m7.onion) you wont even find contact info or credits. In first january days of new 2018 it was absolutely unavailable! So there is no guarantee to receive any response! 

## TODO

* Tests!!!
* Ichidan's authorisation support
* Search by url
* Logical operators in requests
* Third party tools for possible discovery of found hosts

## Contributing

Feel free to contribute!
