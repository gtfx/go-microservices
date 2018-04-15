package main

import (
	"io"
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
)

func sendRequest() (string, error) {
	address := "backend"
	port := 80

	responseChan := make(chan []byte)
	errorChan := make(chan error)

	var netClient = &http.Client{
	  Timeout: time.Second * 10,
	}
	go func() {
		target := fmt.Sprintf("%v:%v", address, port)


		r, err := netClient.Get(target)
		
		if err != nil {
			errorChan <- err
			return
		}

		defer r.Body.Close()
		bs, _ := ioutil.ReadAll(r.Body)

		responseChan <- bs
		return
	}()

	select {
	case r := <-responseChan:
		return string(r[:]), nil
		
	case err := <-errorChan:
		return "", err
	}
}


func hello(w http.ResponseWriter, r *http.Request) {
	response, err := sendRequest()
	if err != nil {
		io.WriteString(w, err.Error())	
	}
	io.WriteString(w, response)	
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

