package context

type form func(value Value, env *Env, forms *Forms) Value
type Forms struct {
  forms map[string]form
}

func GetForms() *Forms {
  forms := map[string]form{
    "define": Define,
  }
  return &Forms{forms}
}

func (f *Forms) Lookup(name string) form {
  return f.forms[name]
}

func Define(value Value, env *Env, forms *Forms) Value {
  pair := pair(value)
  name := symbol(car(pair))
  value, _ = cdr(pair).Eval(env, forms)

  env.Define(name, value)

  return nil
}
