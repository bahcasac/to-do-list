package main

import (
	"fmt"
	"github.com/bahcasac/stock/config"
)

func main() {
	envs := config.LoadEnvVars()
	fmt.Println(envs)
}
