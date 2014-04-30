package littlelisp

type Atom struct {
  value string
}

func (p *Atom) eval() string {
  return p.value
}

func (p *Atom) evalList() string {
  return p.eval()
}

func atom(value string) *Atom {
  return &Atom{value}
}
