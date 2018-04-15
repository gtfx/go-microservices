package main

import (
	"io"
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"os"
    "github.com/gtfx/go-microservices/registry"
)

const frontendSrvName = "srv-frontend"

func sendRequest() (string, error) {
	address := "backend"
	port := 80

	responseChan := make(chan []byte)
	errorChan := make(chan error)

	var netClient = &http.Client{
	  Timeout: time.Second * 10,
	}
	go func() {
		target := fmt.Sprintf("http://%v:%v", address, port)


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
	consuladdr := "svc-consul.kube-system:8500"
	consul, err := registry.NewClient(consuladdr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	id, err := consul.Register(frontendSrvName, 8000)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to register service: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Registered service [%s] id [%s]", frontendSrvName, id)
	defer consul.Deregister(id)

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

