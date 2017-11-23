package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strconv"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for index, url := range os.Args[1:] {
		go fetch(url, ch, index)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, index int) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	f, err := os.OpenFile("out" + strconv.Itoa(index) + ".txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
	  fmt.Printf("file open: %v", err)
	  return
	}

	f.Write(b)
}