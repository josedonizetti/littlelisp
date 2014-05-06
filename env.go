package littlelisp

import "strconv"
import "fmt"

type function func(env *Env, args evaluable) evaluable

type Env struct {
  parent *Env
  defaults map[string]function
}

func DefaultEnv() *Env {
  defaults := make(map[string]function)

  defaults["+"] = func(env *Env, args evaluable) evaluable {
    pair, _ := args.(*Pair)
    a, _ := strconv.Atoi(pair.car.(*Atom).value)
    return atom(string(a + 1))
  }

  env := &Env{nil,defaults}
  return env
}

func (e *Env) lookup(symbol string) function {
  return e.defaults[symbol]
}
