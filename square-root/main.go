package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"
	// "fmt"
	"encoding/json"
	"strconv"
)

const (
	defaultHTTPPort = "8080"
)

var hostname string

func SquareRoot(x float64) float64 {
	return math.Sqrt(x)
}

func GetTimestamp() string {
	t := time.Now()
	return t.Format("20060102150405")
}

type SQRoot struct {
	InputNumber float64
	SquareRoot  float64
}

func squareRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	inputNumber, err := strconv.ParseFloat(r.URL.Path[len("/square-root/"):], 64)
	if err != nil {
		panic(err)
	}
	sq := SQRoot{inputNumber, SquareRoot(inputNumber)}
	strSq, err := json.Marshal(sq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%s:%s:square-root v3] Request: /square-root/%v, Response: %s\n", hostname, GetTimestamp(), inputNumber, strSq)
	json.NewEncoder(w).Encode(sq)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlMap := map[string]string{
		"/square-root/x": "Square root of x",
	}
	fmt.Printf("[%s:%s:square-root v3] Request: %s\n", hostname, GetTimestamp(), r.URL.Path)
	json.NewEncoder(w).Encode(urlMap)
}

func main() {
	h, error := os.Hostname()
	if error != nil {
		hostname = "none"
	} else {
		hostname = h
	}

	http.HandleFunc("/square-root/", squareRootHandler)
	http.HandleFunc("/", rootHandler)

	fmt.Printf("Starting server on port :8080\n")
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Shutting down server\n")
	} else if err != nil {
		fmt.Printf("Cannot start server: %s\n", err)
		os.Exit(1)
	}
}
