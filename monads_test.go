package monads

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Validator monads", func(t *testing.T) {
		t.Run("Map()", func(t *testing.T) {
			t.Run("it should run on a Some", func(t *testing.T) {
				Some(10).
					Map(IntToIntFn(double)).
					Map(func(x interface{}) interface{} {
						assert.Equal(t, 20, x)
						return 0
					})
			})

			t.Run("it should not run on a None", func(t *testing.T) {
				None().
					Map(func(x interface{}) interface{} {
						t.Error("Should not run Map on a None")
						return 0
					})
			})

			t.Run("it should not run on a Error", func(t *testing.T) {
				Error(errors.New("an error")).
					Map(func(x interface{}) interface{} {
						t.Error("Should not run Map on a Error")
						return 0
					})
			})
		})

		t.Run("MapError()", func(t *testing.T) {
			t.Run("it should handle an error condition", func(t *testing.T) {
				var ran bool
				Some(10).
					MapError(IntToIntErrorFn(doubleError)).
					CatchMap(func(err error) {
						assert.Equal(t, "double error", err.Error())
						ran = true
					})
				assert.True(t, ran)
			})

			t.Run("it should handle a no error condition", func(t *testing.T) {
				x, err := Some(10).
					Map(IntToIntFn(double)).
					MapError(func(x interface{}) (interface{}, error) {
						return x.(int) * 2, nil
					}).
					JoinError()
				assert.Equal(t, 40, x)
				assert.Nil(t, err)
			})

			t.Run("it should not run on a None", func(t *testing.T) {
				x, err := None().
					MapError(func(x interface{}) (interface{}, error) {
						t.Fail()
						return 0, nil
					}).
					JoinError()
				assert.Nil(t, x)
				assert.Nil(t, err)
			})
		})

		t.Run("Bind()", func(t *testing.T) {
			t.Run("it should run on a Some", func(t *testing.T) {
				x := Some(10).
					Bind(IntToValidatorFn(doubleBind)).
					Join()
				assert.Equal(t, 20, x)
			})

			t.Run("it should not run on a None", func(t *testing.T) {
				None().
					Bind(func(x interface{}) Validator {
						t.Error("Should not run Map on a None")
						return Some(10)
					})
			})

			t.Run("it should not run on a Error", func(t *testing.T) {
				Error(errors.New("an error")).
					Bind(func(x interface{}) Validator {
						t.Error("Should not run Map on a Error")
						return Some(0)
					})
			})
		})

		t.Run("CatchMap()", func(t *testing.T) {
			t.Run("it should run on an Error", func(t *testing.T) {
				var ran bool
				Error(errors.New("some error")).
					CatchMap(func(err error) {
						ran = true
				})
				assert.True(t, ran)
			})

			t.Run("it should not run on Some", func(t *testing.T) {
				Some(10).
					CatchMap(func(err error) {
						t.Error("Should not run CatchMap on Some")
				})
			})

			t.Run("it should not run on a None", func(t *testing.T) {
				None().
					CatchMap(func(err error) {
						t.Error("Should not run CatchMap on None")
				})
			})
		})
	})
}

func double(x int) int {
	return x * 2
}

func doubleError(x int) (int, error) {
	return 0, errors.New("double error")
}

func doubleBind(x int) Validator {
	return Some(x * 2)
}
