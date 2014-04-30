package littlelisp


type Null struct {
}

func null() *Null {
  return new(Null)
}

func (p *Null) eval() string {
  return ""
}

func (p *Null) evalList() string {
  return p.eval()
}

func isNull(e evaluable) bool {
  _, ok := e.(*Null)
  return ok
}
