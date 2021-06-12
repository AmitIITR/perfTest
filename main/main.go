package main

import (
	"fmt"
	"net/http"
	"sync"
)


func main(){

	var wg sync.WaitGroup
	wg.Add(1)

	defer wg.Wait()
   // this is for 5* 10000 go routines
	for i:=1; i <10000; i++ {
		go test()
	}
}

func test() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//fmt.Println(<-c)

        fmt.Println("starting curl!")
	// Doing a for loop over a channel will keep waiting for the new data to show up in the channel
	// which it will retrieve in l
	//these are 5 go-routines
	for l := range c {
		// its a function literal which is similar to lamda in java
		go func(link string) {
			//time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
