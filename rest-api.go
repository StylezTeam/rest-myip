package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintf(w, "Hello, %q, %s\n", html.EscapeString(r.URL.Path), now.Format("2006-01-02 15:04:05"))

		fmt.Fprintf(w, "\n# Host Name\n")
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%v\n", hostname)

		fmt.Fprintf(w, "\n# Local IP\n")
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			panic(err)
		}
		for i, addr := range addrs {
			fmt.Fprintf(w, "%d %v\n", i, addr)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
