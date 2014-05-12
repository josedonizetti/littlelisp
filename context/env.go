package context

type Env struct {
  parent *Env
  defaults map[string]Value
}

func NewEnv(parent *Env) *Env {
  return &Env{parent,make(map[string]Value)}
}

func (env *Env) Define(symbol string, val Value) {
  env.defaults[symbol] = val
}

func (env *Env) Lookup(symbol string) Value {
  return env.defaults[symbol]
}
