package context

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
  return env.procedures[symbol]
}
