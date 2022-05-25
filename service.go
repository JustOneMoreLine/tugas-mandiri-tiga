package main

import (
	"fmt"
	"net/http"
)

func EchoService(w http.ResponseWriter, headers map[string]interface{}) map[string]interface{} {
	for _, value := range headers {
		switch c := value.(type) {
		case string:
			fmt.Fprintf(w, "%q\n", c)
		case int:
			fmt.Fprintf(w, "%d\n", c)
		default:
			fmt.Fprintf(w, "%T\n", c)
		}
	}
	return headers
}