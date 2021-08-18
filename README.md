## A monad library for GoLang

#### Features

* Works with types by using adapters to convert typed functions to untyped functions
* Works with an error return value through special mapping functions like MapError() and BindError()
* Cleaner more easily readable code

#### Validation Monad

The Validation monad has two types:
* Success = contains a regular value runs all .Map() .Bind()...
* Fail = contains an error value does not run .Map() .Bind()...

The Error type eliminates the need for 
```go
if x != nil {
	return errors.New(....)
}
```

#### Example
```go
Success(10).
	Map(IntToIntFn(double)).
	Bind(IntToValidationFn(doubleBind)).
	Bind(IntToValidationFn(
		func(x int) Validation {
			return Fail(errors.New("failing here"))
		},
	)).
    // The following function will not run due to the previous Fail
	Bind(IntToValidationFn(   
		func(x int) Validation {
			return Success(x * 10)
		},
	)).
    JoinError()   // returns x, error

// ***********************************************
func double(x int) int {
    return x * 2
}


func doubleBind(x int) Validation {
    return Success(x * 2)
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

func IntToValidationFn(fn func(x int) Validation) BindFn {
	return func(x interface{}) Validation {
		return fn(x.(int))
	}
}

```
For more examples, see the tests