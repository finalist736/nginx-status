package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"net/http"
	"time"
	"net"
	"bufio"
	"io"
)

/*
nginx status data
Active connections: 22
server accepts handled requests
 2607076 2607076 2588670
Reading: 0 Writing: 1 Waiting: 13
 */
var params = []string{"connections", "accepts", "handled", "requests", "reading", "writing", "waiting"}
var nginxUrl = flag.String("nginx", "http://localhost/server-status", "where is nginx status located")

func main() {

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("no parameter specified")
		os.Exit(1)
	}
	var param string
	for _, iterator := range os.Args[1:] {
		if strings.Contains(iterator, "-") {
			continue
		}
		param = iterator
		break
	}
	if !IsValueInList(param, params) {
		fmt.Println("invalid parameter specified")
		os.Exit(1)
	}
	fmt.Println(getStatus(param))
}

func getStatus(param string) string {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var netClient = &http.Client{
		Timeout: time.Second * 10,
		Transport: netTransport,
	}

	response, err := netClient.Get(*nginxUrl)
	if err != nil {
		os.Exit(1)
	}

	defer response.Body.Close()

	return GetParameter(response.Body, param)
}

func GetParameter(reader io.Reader, param string) string {

	results := make(map[string]string)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	//fmt.Printf("connections: %s\n", scanner.Text())
	results["connections"] = scanner.Text()
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	//fmt.Printf("accepts: %s\n", scanner.Text())
	results["accepts"] = scanner.Text()
	scanner.Scan()
	//fmt.Printf("handled: %s\n", scanner.Text())
	results["handled"] = scanner.Text()
	scanner.Scan()
	//fmt.Printf("requests: %s\n", scanner.Text())
	results["requests"] = scanner.Text()
	scanner.Scan()
	scanner.Scan()
	//fmt.Printf("reading: %s\n", scanner.Text())
	results["reading"] = scanner.Text()
	scanner.Scan()
	scanner.Scan()
	//fmt.Printf("writing: %s\n", scanner.Text())
	results["writing"] = scanner.Text()
	scanner.Scan()
	scanner.Scan()
	//fmt.Printf("waiting: %s\n", scanner.Text())
	results["waiting"] = scanner.Text()

	return results[param]
}
