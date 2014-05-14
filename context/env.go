package context

import "strings"

type procedure func(value Value) Value

type Env struct {
  parent *Env
  procedures map[string]procedure
}

func NewEnv(parent *Env) *Env {
  procedures := make(map[string]procedure)

  procedures["+"] = Add
  procedures["-"] = Sub
  procedures["/"] = Div
  procedures["*"] = Mul

  return &Env{parent,procedures}
}

func (env *Env) Lookup(symbol string) procedure {
  procedure := env.procedures[strings.TrimSpace(symbol)]

  if procedure == nil {
    return func(value Value) Value { return nil }
  }

  return procedure
}

func (env *Env) Define(symbol string, value Value) {
  env.procedures[strings.TrimSpace(symbol)] = func(v Value) Value {
    return value
  }
}
