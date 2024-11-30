package main

import (
	"fmt"

	"github.com/hadihalimm/cafebuzz-backend/internal/api"
)

// @title CafeBuzz API
// @version 1.0
// @description CafeBuzz API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 9Mf1o@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth
func main() {
	httpServer, _ := api.NewServer()
	err := httpServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Cannot start server: %s", err))
	}
}
