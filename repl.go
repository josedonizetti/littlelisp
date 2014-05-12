package main

import (
  "fmt"
  . "github.com/littlelisp/context"
  . "github.com/littlelisp/parser"
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
