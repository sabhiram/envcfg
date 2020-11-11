package main

import (
	"fmt"
	"os"

	"github.com/sabhiram/envcfg"
)

// Config implements a structure that we want to use to inject variables in from
// the OS environment.
type Config struct {
	A string `envcfg:"A"`
	B int    `envcfg:"B,required"`
}

func main() {
	config := Config{}

	if err := envcfg.Load(&config); err != nil {
		fmt.Printf("config error :: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("A: %s\n", config.A)
	fmt.Printf("B: %d\n", config.B)
}
