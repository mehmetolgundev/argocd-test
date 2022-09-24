package main

import "net/http"

func main() {
	var page = "<html><body><h1>HELLO WORLD</h1></body></html>"
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(page))
	})
	http.ListenAndServe(":8080", nil)
}
