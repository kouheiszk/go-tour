package main

import (
    "fmt"
    "math"
)

func main() {
    // math.pi は小文字から始まるので、公開されていない
    // math.Pi は外部に公開されているので実行できる
    // fmt.Println(math.pi)
    fmt.Println(math.Pi)
}