package littlelisp


type Null struct {
}

func null() *Null {
  return new(Null)
}

func (p *Null) eval(env *Env) string {
  return ""
}

func (p *Null) evalList(env *Env) string {
  return p.eval(env)
}

func isNull(e evaluable) bool {
  _, ok := e.(*Null)
  return ok
}
