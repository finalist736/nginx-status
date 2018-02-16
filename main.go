package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"net/http"
	"time"
	"net"
	"regexp"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		os.Exit(1)
	}

	return GetParameterReg(body, param)
}

func GetParameterReg(data []byte, param string) string {
	re := regexp.MustCompile(`\d+`)
	arr := re.FindAllString( string(data), 7)
	if len(arr) != 7 {
		return ""
	}
	results := make(map[string]string)
	results["connections"] = arr[0]
	results["accepts"] = arr[1]
	results["handled"] = arr[2]
	results["requests"] = arr[3]
	results["reading"] = arr[4]
	results["writing"] = arr[5]
	results["waiting"] = arr[6]
	return results[param]
}
