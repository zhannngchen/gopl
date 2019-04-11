package main

import (
    "fmt"
    "flag"
    "strings"
)

var n = flag.Bool("n", false, "omit tailing newline")
var sep = flag.String("s", "", "seperator")

func main() {
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()
    }
}
