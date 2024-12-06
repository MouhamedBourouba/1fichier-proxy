package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sync"
)

const PROXIES_FILE_PATH string = "proxies.txt"
const REGEX_PATTERN = `<button id="dlw" style=".*" />.*</button>`
const URL_TO_PARSE = "https://1fichier.com/?n5tf6ye85hq8i1a8as8a&af=4186346"

func main() {

	buffer, err := os.ReadFile(PROXIES_FILE_PATH)
	re := regexp.MustCompile(URL_TO_PARSE)

	if err != nil {
		log.Fatalf("Error openning proxies file error: %s", err.Error())
	}

	lines := bytes.Split(buffer, []byte{'\n'})
	output := []byte{}
	var wg sync.WaitGroup

	for _, line := range lines {
		wg.Add(1)

		go func() {
			defer wg.Done()
			if len(line) > 0 {
				url, err := url.Parse("http://" + string(line[:len(line)-1]))

				if err != nil {
					log.Printf("Error parssing proxy: %s", err.Error())
				} else {
					clinet := &http.Client{
						Transport: &http.Transport{
							Proxy: http.ProxyURL(url),
						},
					}

					res, err := clinet.Get(URL_TO_PARSE)

					if err != nil {
						log.Printf("Error making request: %s", err.Error())
					} else {
						defer res.Body.Close()
						body, err := io.ReadAll(res.Body)
						if err != nil {
							log.Printf("Error reading response body: %s", err.Error())
						} else {
							output = append(output, []byte(url.String())...)
							output = append(output, []byte("        ")...)
							output = append(output, re.Find(body)...)
							output = append(output, []byte("\n\n------------------------------------------------------------------------------------------------------------\n\n")...)
						}
					}
				}
			}
		}()
	}

	wg.Wait()
	os.WriteFile("output", output, 0644)
}
