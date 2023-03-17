package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/haha", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hahah")
		writer.Write([]byte("haha"))
		return
	})
	err := http.ListenAndServe(":80", server)
	if err != nil {
		fmt.Println(err)
		return
	}
}
