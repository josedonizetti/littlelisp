package littlelisp

type Value interface {
  Eval(env *Env) (Value, error)
  String() string
}
