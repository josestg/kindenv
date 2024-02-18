# kindenv

Set of constants for the environment mode.

## Installation

```bash
go get github.com/josestg/kindenv
```


## Usage

```go
package main

import (
    "fmt"
    "os"

    "github.com/josestg/kindenv"
)

func main() {
    env, err := kindenv.Parse(os.Getenv("APP_ENV"))
    if err != nil {
        panic(err)
    }
    
    if env < kindenv.Release {
        fmt.Println("verbose mode")
    } else {
        fmt.Println("quiet mode")
    }
}
```

```shell
APP_ENV=develop go run main.go
# verbose mode
```