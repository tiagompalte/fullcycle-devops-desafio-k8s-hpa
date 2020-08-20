package main

import (
    "math"
    "net/http"
)

func greeting(message string) string {
  return "<b>" + message + "</b>"
}

func index(w http.ResponseWriter, req *http.Request) {
  message := greeting("Code.education Rocks!")

  sum := 0.0001
	for i := 0; i <= 1000000; i++ {
		sum += math.Sqrt(sum)
	}

  w.WriteHeader(200)
  w.Write([]byte(message))
}

func main() {

    http.HandleFunc("/", index)

    http.ListenAndServe(":8000", nil)
}

