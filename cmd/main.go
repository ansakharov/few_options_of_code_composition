package main

import (
	"context"
	"encoding/json"
	"fmt"
	repository "github.com/ansakharov/few_options_of_code_composition/1_simple_handler_that_fetch_data_from_repository/domain/repository/tariff"
	"github.com/ansakharov/few_options_of_code_composition/1_simple_handler_that_fetch_data_from_repository/rpc/simple_query_commission_for_tariff_handler"
	"io"
	"net/http"

	"github.com/ansakharov/few_options_of_code_composition/server_dto"
)

func handleRequestTplFunc(handlerFunc func(ctx context.Context,
	in *server_dto.QueryDtoIn, out *server_dto.QueryDtoOut) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		var in server_dto.QueryDtoIn
		var out server_dto.QueryDtoOut

		err = json.Unmarshal(body, &in)
		if err != nil {
			fmt.Println(err)

			return
		}

		err = handlerFunc(ctx, &in, &out)
		if err != nil {
			fmt.Println(err)

			return
		}

		output, err := json.Marshal(&out)
		if err != nil {
			fmt.Println(err)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		_, _ = w.Write(output)
	}
}

func main() {
	simpleQueryHandler := simple_query_commission_for_tariff_handler.New(repository.New())

	http.HandleFunc("/simple_query", handleRequestTplFunc(simpleQueryHandler.Handle))

	fmt.Println("Server is starting on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
