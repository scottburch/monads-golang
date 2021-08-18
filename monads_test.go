package monads

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Validation monads", func(t *testing.T) {
		t.Run("Map()", func(t *testing.T) {
			t.Run("it should run on a Success", func(t *testing.T) {
				Success(10).
					Map(IntToIntFn(double)).
					Map(func(x interface{}) interface{} {
						assert.Equal(t, 20, x)
						return 0
					})
			})

			t.Run("it should not run on a Fail", func(t *testing.T) {
				Fail(errors.New("an error")).
					Map(func(x interface{}) interface{} {
						t.Error("Should not run Map on a Fail")
						return 0
					})
			})
		})

		t.Run("MapError()", func(t *testing.T) {
			t.Run("it should handle an error condition", func(t *testing.T) {
				var ran bool
				Success(10).
					MapError(IntToIntErrorFn(doubleError)).
					CatchMap(func(err error) {
						assert.Equal(t, "double error", err.Error())
						ran = true
					})
				assert.True(t, ran)
			})

			t.Run("it should handle a no error condition", func(t *testing.T) {
				x, err := Success(10).
					Map(IntToIntFn(double)).
					MapError(func(x interface{}) (interface{}, error) {
						return x.(int) * 2, nil
					}).
					JoinError()
				assert.Equal(t, 40, x)
				assert.Nil(t, err)
			})

			t.Run("it should not run on a Fail", func(t *testing.T) {
				Fail(errors.New("some error")).
					MapError(func(x interface{}) (interface{}, error) {
						t.Error("MapError should not be running on a Fail")
						return 0, nil
					})
			})
		})

		t.Run("Bind()", func(t *testing.T) {
			t.Run("it should run on a Success", func(t *testing.T) {
				x := Success(10).
					Bind(IntToValidatorFn(doubleBind)).
					Join()
				assert.Equal(t, 20, x)
			})

			t.Run("it should not run on a Fail", func(t *testing.T) {
				Fail(errors.New("an error")).
					Bind(func(x interface{}) Validation {
						t.Error("Should not run Map on a Fail")
						return Success(0)
					})
			})
		})

		t.Run("CatchMap()", func(t *testing.T) {
			t.Run("it should run on an Fail", func(t *testing.T) {
				var ran bool
				Fail(errors.New("some error")).
					CatchMap(func(err error) {
						ran = true
					})
				assert.True(t, ran)
			})

			t.Run("it should not run on Success", func(t *testing.T) {
				Success(10).
					CatchMap(func(err error) {
						t.Error("Should not run CatchMap on Success")
					})
			})
		})
	})

	t.Run("types demo", func(t *testing.T) {
		Success(10).
			Map(IntToIntFn(double)).
			Bind(IntToValidatorFn(doubleBind)).
			Map(IntToIntFn(
				func(x int) int {
					return x + 10
				},
			)).
			Bind(IntToValidatorFn(
				func(x int) Validation {
					return Success(x * 10)
				},
			))
	})
}

func double(x int) int {
	return x * 2
}

func doubleError(x int) (int, error) {
	return 0, errors.New("double error")
}

func doubleBind(x int) Validation {
	return Success(x * 2)
}

