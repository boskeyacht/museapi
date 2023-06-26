package main

import (
	cmd "github.com/boskeyacht/museapi/cmd/museapi"
	"github.com/boskeyacht/museapi/internal/types"
)

func main() {
	c := types.InitConfig()

	cmd.Execute(c)
}
