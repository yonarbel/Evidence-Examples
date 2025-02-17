package main

import (
	"fmt"
	"net/http"

	"github.com/hbollon/go-edlib"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	token:= "0d6a9bb5af126f73350a2afc058492765446aaad"
	key := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2V4YW1wbGUuYXV0aDAuY29tLyIsImF1ZCI6Imh0dHBzOi8vYXBpLmV4YW1wbGUuY29tL2NhbGFuZGFyL3YxLyIsInN1YiI6InVzcl8xMjMiLCJpYXQiOjE0NTg3ODU3OTYsImV4cCI6MTQ1ODg3MjE5Nn0.CA7eaHjIHz5NxeIJoFK9krqaeZrPLwmMmgI_XiQiIkQ"
	url := "https://myuser:" + token + "@api.github.com"
	res, err := edlib.StringsSimilarity("string1", "string2", edlib.Levenshtein)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Similarity: %f", res)
	}

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":9001", nil)
}
