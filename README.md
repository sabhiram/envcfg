# envcfg

load tagged structs from the OS environment.

## Install

```
go get github.com/sabhiram/envcfg
```

## Usage

The main (and only) library API is the `envcfg.Load()`. It accepts an  
`interface{}` backed by a tagged struct as shown below.

For example:
```
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
	C string `envcfg:"C,,FOOBAR"`
}

func main() {
	config := Config{}

	if err := envcfg.Load(&config); err != nil {
		fmt.Printf("config error :: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("A: %v\n", config.A)
	fmt.Printf("B: %v\n", config.B)
	fmt.Printf("C: %v\n", config.C)
}

```

To run the example (from the root directory of the library), and have it fail 
since `B` is not set (and neither is `A` but it is also not required).

```
$ go run ./example/main.go
config error :: missing required field in env (B)
exit status 1
```

Which can be fixed by setting the required `B` param in the env.
```
$ A=HELLO B=42 go run ./example/main.go
A: HELLO
B: 42
```

## All else

Leave it here: https://github.com/sabhiram/envcfg/issues
