package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/hIMEI29A/gotorsocks"
	"golang.org/x/net/html"
)

const (
	ICHIDAN   string = "ichidanv34wrx7m7.onion:80"
	GETPARAMS        = "GET /search?query="
)

func ErrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Gichidan struct {
}

//status, err := bufio.NewReader(connect).ReadString('\n')

func Search(request string) []string {
	requestString := GETPARAMS + request + "\n"

	tor, err := gotorsocks.NewTorGate()

	ErrFatal(err)

	connect, err := tor.DialTor(ICHIDAN)

	ErrFatal(err)

	fmt.Fprintf(connect, requestString)

	var result []string

	scanner := bufio.NewScanner(connect)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func main() {
	res := Search("paypal")

	fmt.Println(res[69])

	h6names := getH6names(res)

	for _, str := range h6names {
		fmt.Println(str)
	}

	data := ParseData(res)

	fmt.Println(data.String())
}
