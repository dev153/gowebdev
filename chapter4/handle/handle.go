package main

import (
	"fmt"
	"net/http"
)

func messageHandler(replyStr string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, replyStr)
	})
}

func main() {
	mux := http.NewServeMux()
	msgPairs := map[string]string{
		"/welcome": "Welcome to Go Web Development",
		"/about":   "About to Go Web Development",
	}

	for httpURL, httpReplyStr := range msgPairs {
		mux.Handle(httpURL, messageHandler(httpReplyStr))
	}

	http.ListenAndServe(":9099", mux)
}
