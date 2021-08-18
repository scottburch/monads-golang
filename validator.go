package monads

type Validator struct {
	v      interface{}
	isSome bool
	err    error
}

func Some(v interface{}) Validator {
	return Validator{
		v:      v,
		isSome: true,
	}
}

func Error(err error) Validator {
	return Validator{
		err: err,
	}
}

func None() Validator {
	return Validator{
		isSome: false,
	}
}

func (m Validator) Map(fn func(interface{}) interface{}) Validator {
	if m.isSome {
		return Some(fn(m.v))
	}
	return m

}

func (m Validator) Bind(fn func(interface{}) Validator) Validator {
	if m.isSome {
		return fn(m.v)
	}
	return m
}

func (m Validator) MapError(fn func(interface{}) (interface{}, error)) Validator {
	if m.isSome {
		v, err := fn(m.v)
		if err != nil {
			return Error(err)
		}
		return Some(v)
	}
	return m
}

func (m Validator) JoinError() (interface{}, error) {
	return m.v, m.err
}

func (m Validator) Join() interface{} {
	return m.v
}

func (m Validator) CatchMap(fn func (err error)) Validator {
	if m.err != nil {
		fn(m.err)
	}
	return m
}
