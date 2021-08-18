## A monad library for GoLang

#### Features

* Works with types by using adapters to convert typed functions to untyped functions
* Works with an error return value through special mapping functions like MapError() and BindError()
* Cleaner more easily readable code

#### Validator Monad

The Validator monad has three types:
* Some = contains a regular value runs all .Map() .Bind()...
* None = no value does not run .Map() .Bind()...
* Error = contains an error value does not run .Map() .Bind()

The Error type eliminates the need for 
```go
if x != nil {
	return errors.New(....)
}
```

#### Example
```go
Some(10).
	Map(IntToIntFn(double)).
	Bind(IntToValidatorFn(doubleBind)).
	Map(IntToIntFn(
		func(x int) int {
			return x + 10
		},
	)).
	Bind(IntToValidatorFn(
		func(x int) Validator {
			return Some(x * 10)
		},
	))

// ***********************************************
func double(x int) int {
    return x * 2
}


func doubleBind(x int) Validator {
    return Some(x * 2)
}

```

#### Type Adapters

Type adapters look like this
```
func IntToIntFn(fn func(x int) int) MapFn {
	return func(x interface{}) interface{} {
		return fn(x.(int))
	}
}

func IntToValidatorFn(fn func(x int) Validator) BindFn {
	return func(x interface{}) Validator {
		return fn(x.(int))
	}
}

```
For more examples, see the tests