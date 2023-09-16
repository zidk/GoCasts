package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for l := range c {
		go func(li string, ch chan string) {
			time.Sleep(3 * time.Second)
			checkLink(li, ch)
		}(l, c)
	}

	/*
		for l := range c {
			go checkLink(l, c)
		}

	*/

	/* similar than while(true)
	for {
		go checkLink(<-c, c)
	}
	*/

	/*
		for i := 0; i < len(links); i++ {
			fmt.Printf(<-c)
		}
	*/

}

func checkLink(l string, c chan string) {

	_, err := http.Get(l)
	if err != nil {
		fmt.Println("Error getting ", l, err)
		//c <- l + " is down!"
		c <- l
		return
	}
	fmt.Println(l, "is up.")
	//c <- l + " is up!"
	c <- l
}
