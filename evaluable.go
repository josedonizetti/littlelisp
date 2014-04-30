package littlelisp

type evaluable interface {
  eval() string
  evalList() string
}
