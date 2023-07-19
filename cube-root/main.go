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

type CbRoot struct {
	InputNumber float64
	CubeRoot    float64
}

func cubeRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println(r.URL.Path)
	inputNumber, err := strconv.ParseFloat(r.URL.Path[len("/cube-root/"):], 64)
	if err != nil {
		panic(err)
	}
	cr := CbRoot{inputNumber, math.Cbrt(inputNumber)}
	json.NewEncoder(w).Encode(cr)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlMap := map[string]string{
		"/square-root/x": "Square root of x",
		"/cube-root/x":   "Cube root of x",
	}
	json.NewEncoder(w).Encode(urlMap)
}

func main() {
	http.HandleFunc("/cube-root/", cubeRootHandler)
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
