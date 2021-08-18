package monads

func IntToIntFn(fn func(x int) int) func(x interface{}) interface{} {
	return func(x interface{}) interface{} {
		return fn(x.(int))
	}
}

func IntToIntErrorFn(fn func(x int) (int, error)) func(x interface{}) (interface{}, error) {
	return func(x interface{}) (interface{}, error) {
		return fn(x.(int))
	}
}

func IntToValidatorFn(fn func(x int) Validator) func(x interface{}) Validator {
	return func(x interface{}) Validator {
		return fn(x.(int))
	}
}



