package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var numberOfWorker int = 5
var url string = "http://localhost:8080"
var method string = "GET"
var listOfData = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
}

func main() {
	client := &http.Client{}
	wg := sync.WaitGroup{}

	i := 0
	for i < len(listOfData) {
		j := 0
		for i < len(listOfData) && j < numberOfWorker {
			payload := fmt.Sprintf("{\n	\"order_uuid\": %s \n}", listOfData[i])
			wg.Add(1)
			go call(client, payload, &wg)
			j += 1
			i += 1
		}
		wg.Wait()
	}

}

func call(client *http.Client, payload string, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
