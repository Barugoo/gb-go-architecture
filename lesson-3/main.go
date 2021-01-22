package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var count float64
var goodResp float64

func httpGet(url string) (int, float64) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	elapsed := time.Since(start).Seconds()
	return resp.StatusCode, elapsed

}

func httpPost(url string, body []byte) (int, float64) {
	start := time.Now()

	r := bytes.NewReader(body)
	resp, err := http.Post(url, "application/json", r)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	elapsed := time.Since(start).Seconds()
	return resp.StatusCode, elapsed
}

func worker(w int, url string, jobs <-chan int, results chan<- int, method string, body []byte) {
	if method == "GET" {
		for j := range jobs {
			code, t := httpGet(url)
			if code >= 200 && code <= 299 {
				fmt.Println("Время отклика: ", t)
				count = +t
				goodResp++
			}

			results <- j * 2
		}
	} else {
		for j := range jobs {
			code, t := httpPost(url, body)
			if code >= 200 && code <= 299 {
				fmt.Println("Время отклика: ", t)
				count = +t
			}

			results <- j * 2
		}
	}

}

func main() {

	var url string
	var workerCount int
	var jobCount int
	var httpMethod string
	var bodyReques []byte

	flag.StringVar(&url, "url", "https://yandex.ru", "URL сайта")
	flag.IntVar(&workerCount, "w", 5, "number of workers")
	flag.IntVar(&jobCount, "j", 5, "number of jobs")
	flag.StringVar(&httpMethod, "m", "GET", "method GET or POST")

	flag.Parse()

	jobs := make(chan int, jobCount)
	results := make(chan int, jobCount)

	for w := 1; w <= workerCount; w++ {
		go worker(w, url, jobs, results, httpMethod, bodyReques)
	}

	for j := 1; j <= jobCount; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= jobCount; a++ {
		<-results
	}

	rpc := goodResp / count
	sred := count / goodResp
	fmt.Println("RPC: ", rpc)
	fmt.Println("Среднее значение: ", sred)

}
