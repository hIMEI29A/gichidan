# gichidan

**gichidan** - command line wrapper to 
(_onion link_) [**Ichidan**](http://ichidanv34wrx7m7.onion) deep-web search with enhansed pentest features.

**NOTE** Since 01.01.2018 Ichidan seems to be dead so **Gichidan** also does not work. Reborn, Ichidan, we hope!

    ▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆
     ____, __, ____,__, _,__, ____,____,_,  _,
    (-/ _,(-| (-/  (-|__|(-| (-|  (-/_|(-|\ | 
     _\__| _|_,_\__,_|  |,_|_,_|__//  |,_| \|,
    (     (   (    (     (   (   (     (      
    ▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆

[![Go Report Card](https://goreportcard.com/badge/github.com/hIMEI29A/gichidan)](https://goreportcard.com/report/github.com/hIMEI29A/gichidan) [![GoDoc](https://godoc.org/github.com/hIMEI29A/gichidan?status.svg)](http://godoc.org/github.com/hIMEI29A/gichidan)

Copyright 2017 hIMEI


## TOC
- [About](#about)
- [Install](#install)

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

Then run make

```sh
make
```

You may **install** Gichidan also by `go get`:

```sh
go get github.com/hIMEI29A/gichidan
```
