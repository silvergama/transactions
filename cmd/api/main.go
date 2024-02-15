// @title Transaction APIs
// @description API for managing transactions
// @version 1.0
// @host localhost:8080
// @BasePath /
package main

import (
	_ "github.com/silvergama/transations/cmd/api/docs"
	"github.com/silvergama/transations/config"
	"github.com/silvergama/transations/internal/api"
)

func main() {
	api.Start(config.ReadProperties())
}
