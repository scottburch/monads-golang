package monads

type MapFn = func(x interface{}) interface{}
type MapErrorFn = func(x interface{}) (interface{}, error)
type BindFn = func(x interface{}) Validation

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

func IntToValidationFn(fn func(x int) Validation) BindFn {
	return func(x interface{}) Validation {
		return fn(x.(int))
	}
}




