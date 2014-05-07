package littlelisp

type Env struct {
  parent *Env
  defaults map[string]*Atom
}

func NewEnv() *Env {
  return &Env{nil,make(map[string]*Atom)}
}

func (env *Env) Define(symbol string, val *Atom) {
  env.defaults[symbol] = val
}

func (env *Env) Lookup(symbol string) *Atom {
  return env.defaults[symbol]
}
