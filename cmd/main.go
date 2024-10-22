package main

import (
	"fmt"

	"github.com/hadihalimm/cafebuzz-backend/internal/api"
)

func main() {
	httpServer, _ := api.NewServer()
	err := httpServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Cannot start server: %s", err))
	}
}
