/* Here I'm using two time API services - worldtimeapi.org and timeapi.io for
comparsion, but you can use this snippet for whatever services you'd like to check.
Visiting these links below should give the current time in Asia/Kolkata timezone:

https://worldtimeapi.org/api/timezone/Asia/Kolkata
https://timeapi.io/api/Time/current/zone\?timeZone\=Asia/Kolkata

However, here we will be discarding the responses as we're just concerned
about which service returns faster.
*/

package main

import (
	"fmt"
	"net/http"
)

const (
	worldTimeAPIURL = "https://worldtimeapi.org/api/timezone/Asia/Kolkata"
	timeAPIURL      = "https://timeapi.io/api/Time/current/zone?timeZone=Asia/Kolkata"
)

func makeHTTPReq(url string, signalChan chan<- struct{}) {
	http.Get(url)
	signalChan <- struct{}{}
}

func main() {
	timeAPIChan := make(chan struct{})
	worldTimeAPIChan := make(chan struct{})

	go makeHTTPReq(timeAPIURL, timeAPIChan)
	go makeHTTPReq(worldTimeAPIURL, worldTimeAPIChan)

	select {
	case <-timeAPIChan:
		fmt.Println("timeapi.io returned first")
	case <-worldTimeAPIChan:
		fmt.Println("worldtimeapi.org returned first")
	}
}
