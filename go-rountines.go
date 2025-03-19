package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	websites := []string{
		"https://google.com",
		"https://vithsutra.com",
		"https://github.com",
	}
	for _, web := range websites {
		go getStatus(web)
		wg.Add(1)

	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatus(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("ooops error in site %s\n", endpoint)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d for the site %s\n", res.StatusCode, endpoint)
	}
}
