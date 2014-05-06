package littlelisp

type evaluable interface {
  eval(env *Env) string
  evalList(env *Env) string
}
