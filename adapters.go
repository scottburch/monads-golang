package monads

type MapFn = func(x interface{}) interface{}
type MapErrorFn = func(x interface{}) (interface{}, error)
type BindFn = func(x interface{}) Validator

func IntToIntFn(fn func(x int) int) MapFn {
	return func(x interface{}) interface{} {
		return fn(x.(int))
	}
}

func IntToIntErrorFn(fn func(x int) (int, error)) MapErrorFn {
	return func(x interface{}) (interface{}, error) {
		return fn(x.(int))
	}
}

func IntToValidatorFn(fn func(x int) Validator) BindFn {
	return func(x interface{}) Validator {
		return fn(x.(int))
	}
}




