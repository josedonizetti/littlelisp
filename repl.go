package main

import (
  "fmt"
  . "github.com/littlelisp/parse"
  . "github.com/littlelisp/nodes"
)

func main() {
  for {
    fmt.Print("> ")
    var input string
    fmt.Scan(&input)
    env := NewEnv(nil)
    value, _ := Parse(input).Eval(env)
    fmt.Println(value)
  }

}
