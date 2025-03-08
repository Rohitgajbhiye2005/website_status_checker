package main

import (
	"fmt"
	"net/http"
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
	for _, link := range links {
		go checklink(link, c)
	}
	for i := 0; i < len(links); i++ {
		fmt.Print(<-c)
	}

}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Print(link, "might be down")
		c <- "Might be down i think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "yup its up\n"
}
