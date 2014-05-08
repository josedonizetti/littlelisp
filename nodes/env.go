package littlelisp

type Env struct {
  parent *Env
  defaults map[string]*Atom
}

func NewEnv(parent *Env) *Env {
  return &Env{parent,make(map[string]*Atom)}
}

func (env *Env) Define(symbol string, val *Atom) {
  env.defaults[symbol] = val
}

func (env *Env) Lookup(symbol string) *Atom {
  return env.defaults[symbol]
}
