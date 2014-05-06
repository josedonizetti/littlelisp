package littlelisp

type Atom struct {
  value string
}

func (p *Atom) eval(env *Env) string {
  return p.value
}

func (p *Atom) evalList(env *Env) string {
  return p.eval(env)
}

func atom(value string) *Atom {
  return &Atom{value}
}

func isAtom(e evaluable) bool {
  _, ok := e.(*Atom)
  return ok
}
