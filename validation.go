package monads

type Validation struct {
	v      interface{}
	err    error
}

func Success(v interface{}) Validation {
	return Validation{
		v:      v,
	}
}

func Fail(err error) Validation {
	return Validation{
		err: err,
	}
}


func (m Validation) Map(fn func(interface{}) interface{}) Validation {
	if m.err == nil {
		return Success(fn(m.v))
	}
	return m

}

func (m Validation) Bind(fn func(interface{}) Validation) Validation {
	if m.err == nil {
		return fn(m.v)
	}
	return m
}

func (m Validation) MapError(fn func(interface{}) (interface{}, error)) Validation {
	if m.err == nil {
		v, err := fn(m.v)
		if err != nil {
			return Fail(err)
		}
		return Success(v)
	}
	return m
}

func (m Validation) JoinError() (interface{}, error) {
	return m.v, m.err
}

func (m Validation) Join() interface{} {
	return m.v
}

func (m Validation) CatchMap(fn func (err error)) Validation {
	if m.err != nil {
		fn(m.err)
	}
	return m
}
