package main

import (
    "fmt"
    "flag"
)

func main() {
    flag.Parse()
    for _, arg := range flag.Args() {
        CheckGetText(arg)
    }
    fmt.Println("ok")
}