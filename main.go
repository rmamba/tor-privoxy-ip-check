package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No URL provided!!!")
		os.Exit(3)
	}

	requestURL := os.Args[1]
	res, err := http.Get(requestURL)
	if err != nil {
		log.Printf("Error making http request: %s\n", err)
		os.Exit(2)
	}

	defer res.Body.Close()
	originalIP, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error parsing body: %s\n", err)
		os.Exit(2)
	}
	fmt.Printf("OriginalIP: %s\n", originalIP)

	proxyURL, _ := url.Parse("http://localhost:8118")
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := http.Client{
		Transport: tr,
	}

	res2, err2 := client.Get(requestURL)
	if err2 != nil {
		log.Printf("Error making http proxy request: %s\n", err2)
		os.Exit(2)
	}
	defer res2.Body.Close()
	maskedIP, err2 := io.ReadAll(res2.Body)
	if err2 != nil {
		log.Printf("Error parsing body: %s\n", err2)
		os.Exit(2)
	}
	fmt.Printf("MaskedIP: %s\n", maskedIP)

	c := 0
	if bytes.Equal(originalIP, maskedIP) {
		// Mark fail
		c = 1
	}
	os.Exit(c)
}
