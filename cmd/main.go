package main

import (
	"fmt"
	"io"
    "encoding/json"
	"net/http"

    "github.com/ansakharov/few_options_of_code_composition/1_simple_handler_that_fetch_data_from_repository/handler"
	"github.com/ansakharov/few_options_of_code_composition/server_dto"
)

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

	_ = body
	_ = ctx
	var in server_dto.QueryDtoIn

	err = json.Unmarshal(body, &in)
	if err != nil {
		fmt.Println(err)
	
		return
	}

    

	out := server_dto.QueryDtoOut{
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
