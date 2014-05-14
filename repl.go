package main

import (
  "fmt"
  "bufio"
  "os"
  . "github.com/littlelisp/context"
  . "github.com/littlelisp/parser"
)

func main() {
  for {
    fmt.Print("> ")
    in := bufio.NewReader(os.Stdin)
    line, _ := in.ReadString('\n')
    env := NewEnv(nil)
    value, _ := Parse(line).Eval(env)
    fmt.Println(value)
  }

}
