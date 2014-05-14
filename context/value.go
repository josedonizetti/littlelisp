package context

type Value interface {
  Eval(env *Env, forms *Forms) (Value, error)
  String() string
}
