package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	sc := make(chan string)

	for _, link := range links {
		go checkLink(link, sc)
	}

	// for idx := 0; idx < len(links); idx++ {
	// 	fmt.Println(<- sc)
	// }

	// Infinite loop go routine
	for {
		link :=<- sc
		go func() {
			time.Sleep(5 * time.Second)
			go checkLink(link, sc)
		}()
	}
}

func checkLink(link string, sc chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println("Error:", error)
		sc <- link
		return
	}
	fmt.Println(link, "is up!")
	sc <- link
}
