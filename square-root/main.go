package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"

	// "fmt"
	"encoding/json"
	"strconv"
)

const (
	defaultHTTPPort = "8080"
)

func SquareRoot(x float64) float64 {
	return math.Sqrt(x)
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
	json.NewEncoder(w).Encode(sq)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlMap := map[string]string{
		"/square-root/x": "Square root of x",
	}
	json.NewEncoder(w).Encode(urlMap)
}

func main() {
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
