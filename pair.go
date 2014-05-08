package littlelisp

type Pair struct {
  car *Atom
  cdr *Pair
}

func (p *Pair) Eval(env *Env) *Atom {
  return nil
}

func (p *Pair) String() string {
  return ""
}

func NewPair(car *Atom, cdr *Pair) *Pair {
  return &Pair{car,cdr}
}
