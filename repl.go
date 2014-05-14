package main

import (
  "fmt"
  "bufio"
  "os"
  . "github.com/littlelisp/context"
  . "github.com/littlelisp/parser"
)

func main() {
  env := NewEnv(nil)
  for {
    fmt.Print("> ")
    in := bufio.NewReader(os.Stdin)
    line, _ := in.ReadString('\n')
    value, _ := Parse(line).Eval(env, GetForms())
    fmt.Println(value)
  }
}
