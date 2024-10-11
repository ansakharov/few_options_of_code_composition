package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type QueryDtoIn struct {
	TariffID int64 `json:"tariff_id"`
}

type QueryDtoOut struct {
	AmountCommission float64 `json:"amount_commission"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	defer func() {
		err := r.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	ctx := r.Context()

	_ = ctx
	var in QueryDtoIn

	err = json.Unmarshal(body, &in)
	if err != nil {
		fmt.Println(err)

		return
	}

	out := QueryDtoOut{
		AmountCommission: float64(in.TariffID),
	}

	output, err := json.Marshal(&out)
	if err != nil {
		fmt.Println(err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, _ = w.Write(output)
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
