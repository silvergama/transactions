// @title Transaction APIs
// @description API for managing transactions
// @version 1.0
// @host localhost:8080
// @BasePath /
package main

import (
	"github.com/silvergama/transations/cmd/api"
	_ "github.com/silvergama/transations/docs"
)

func main() {
	api.Run()
}
