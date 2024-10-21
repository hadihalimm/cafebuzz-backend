package main

import (
	"fmt"

	"github.com/hadihalimm/cafebuzz-backend/internal/api"
)

func main() {
	server := api.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Cannot start server: %s", err))
	}
}
