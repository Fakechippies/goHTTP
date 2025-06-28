package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func isPortAvailable(port string) bool {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return false
	}
	_ = ln.Close()
	return true
}

func main() {
	directoryPath := "."
	port := ":8000"

	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	if len(os.Args) > 2 {
		directoryPath = os.Args[2]
	}

	directory := http.Dir(directoryPath)
	fileServer := http.FileServer(directory)

	if !isPortAvailable(port) {
		log.Fatalf("Port %s is already in use or unavailable", port)
	}

	fmt.Printf("HTTP Server started on port %s \n", port)
	err := http.ListenAndServe(port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var color string
		switch r.Method {
		case "GET":
			color = "\033[32m"
		case "POST":
			color = "\033[34m"
		default:
			color = "\033[33m"
		}

		reset := "\033[0m"
		fmt.Printf("[*] %s\t%s\t[%s%s%s]\t%s\n", r.RemoteAddr, time.Now().Format("2006-01-02 15:04:05"), color, r.Method, reset, r.URL.Path)
		fileServer.ServeHTTP(w, r)
	}))
	if err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}

}
