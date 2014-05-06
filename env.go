package littlelisp

import "strconv"

type function func(env *Env, args evaluable) evaluable

type Env struct {
  parent *Env
  defaults map[string]function
}

func DefaultEnv() *Env {
  defaults := make(map[string]function)

  defaults["+"] = func(env *Env, args evaluable) evaluable {
    pair, _ := args.(*Pair)

    if isNull(pair.cdr) {
      return pair.car
    }

    if isAtom(pair.cdr) {
      v1 := pair.car.(*Atom).value
      v2 := pair.cdr.(*Atom).value

      v3, _ := strconv.Atoi(v1)
      v4, _ := strconv.Atoi(v2)
      sum := strconv.Itoa(v3 + v4)

      return atom(sum)
    }

    return null()
  }

  env := &Env{nil,defaults}
  return env
}

func (e *Env) lookup(symbol string) function {
  return e.defaults[symbol]
}
