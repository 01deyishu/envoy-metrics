package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var contour map[string]domainCounter

type domainCounter struct {
	domain    string
	qps       int
	bondwidth int
	codes     ResCodes
}

type ResCodes struct {
	two   int
	three int
	four  int
	five  int
	other int
}

func main() {
	fmt.Println("start a daemon")
	res, err := http.Get("http://183.136.237.67:8002/stats/prometheus")
	if err != nil {
		fmt.Println("error of get prometheus metrics")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error of read body")
	}
	metricstext := string(body)
	lines := strings.Split(metricstext, "\n")
	analysisMetrics(lines)
}

func analysisMetrics(mt []string) (counter map[string]domainCounter) {
	for _, line := range mt {
		if code := regexp.MustCompile(`^envoy_cluster_upstream_rq{`); code.MatchString(line) {
			fmt.Println(line)
		}
		if bd := regexp.MustCompile(`^envoy_cluster_upstream_cx_rx_bytes_total{`); bd.MatchString(line) {
			fmt.Println(line)
			counter["domain1"].bondwidth += bd
		}
	}
	return nil
}
