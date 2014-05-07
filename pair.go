package littlelisp

type Pair struct {
  car *Atom
  cdr *Pair
}

func (p *Pair) Eval(env *Env) *Atom {

  if p.car.IsSymbol() {
    atom, _ := p.car.Eval(env)
    procedure := atom.Procedure()
    return procedure(p.cdr)
  }

  return nil
}

func (p *Pair) String() string {
  return "(1 2)"
}

func NewPair(car *Atom, cdr *Pair) *Pair {
  return &Pair{car,cdr}
}
