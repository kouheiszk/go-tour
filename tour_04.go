package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main () {
    // rand.Intn はSeedが変わらないと常に同じ値を返す
    rand.Seed(time.Now().UnixNano())
    fmt.Println("My favorite number is", rand.Intn(10))
}
