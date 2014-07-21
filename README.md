Goexile
==========
Golang based library for the Path of Exile API.  It implements all endpoints for both fun and profits.

## Usage
`go get github.com/taion809/goexile`

## Example
```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/taion809/goexile"
    "log"
)

func main() {
    api := goexile.NewApi()
    results, err := api.GetAllLeagues()

    if err != nil {
        log.Fatal(err)
    }

    for _, v := range results {
        fmt.Printf("Id: %s \n", v.Id)
    }

    l, err := json.Marshal(results)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Marshalled: ", string(l))
}
```

## License
Goexile is licensed under the new bsd license, see license.md

