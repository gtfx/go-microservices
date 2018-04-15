package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"sync"
)

var (
	counter = 0
)

type ServerHandler struct {
	Name string
	Counter int
	counterMutex *sync.Mutex
}

func (s *ServerHandler) increment()  {
	s.counterMutex.Lock()
	s.Counter += 1
	defer s.counterMutex.Unlock()
}

func NewServerHandler(name string) *ServerHandler {
	serverHandler := &ServerHandler{
		Name: name,
		Counter: 0,
		counterMutex: &sync.Mutex{},
	}	
	return serverHandler
} 

func getName() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return name, nil
}

func (s *ServerHandler) Serve(w http.ResponseWriter, r *http.Request) {
	s.increment()
	response := fmt.Sprintf("Server [%s]: served %d requests", s.Name, s.Counter)
	io.WriteString(w, response)
}

func main() {
	name, err := getName()
	if err != nil {
		fmt.Println("Could not get name", err)
		return
	}

	handler := NewServerHandler(name)

	fmt.Printf("Server [%s] reporting for duty\n", name)
	http.HandleFunc("/", handler.Serve)
	http.ListenAndServe(":8000", nil)
}

