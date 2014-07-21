Goexile
==========
Golang based library for the Path of Exile API.  It implements all endpoints for both fun and profits.

## Usage
`go get github.com/taion809/goexile`

## Example
```go
package main

import (
    "github.com/taion809/goexile"
    "fmt"
)

func main() {
    api := goexile.NewApi()
    results, err := api.GetAllLeagues()

    for _, v := range results {
        fmt.Printf("Name: %s \n", v.Name)
    }
}
```

## License
Goexile is licensed under the new bsd license, see license.md

