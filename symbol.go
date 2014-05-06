package littlelisp

type Symbol struct {
  value string
}

func (p *Symbol) eval(env *Env) string {
  return p.Call(env, null())
}

func (p *Symbol) evalList(env *Env) string {
  return p.eval(env)
}

func (p *Symbol) Call(env *Env, args evaluable) string {
  method := env.lookup(p.value)
  return method(env,args).(*Atom).value
}

func symbol(value string) *Symbol {
  return &Symbol{value}
}

func isSymbol(e evaluable) bool {
  _, ok := e.(*Symbol)
  return ok
}
