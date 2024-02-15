package main

import (
	"github.com/silvergama/transations/config"
	"github.com/silvergama/transations/internal/api"
)

func main() {
	api.Start(config.ReadProperties())
}
