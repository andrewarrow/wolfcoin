package network

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func DoPost(ipAndPort, route string, payload []byte) string {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("http://%s%s", ipAndPort, route)
	request, _ := http.NewRequest("POST", urlString, body)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 5}
	return DoHttpRead("POST", route, client, request)
}

func DoHttpRead(verb, route string, client *http.Client, request *http.Request) string {
	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			os.Exit(1)
			return ""
		}
		if resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 204 {
			return string(body)
		} else {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, string(body))
			os.Exit(1)
			return ""
		}
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	os.Exit(1)
	return ""
}
