package monads

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Validator monads", func(t *testing.T) {
		t.Run("Map()", func(t *testing.T) {
			x := Some(10).
				Map(IntToIntFn(double)).
				Join()
			assert.Equal(t, 20, x)
		})

		t.Run("MapError()", func(t *testing.T) {
			t.Run("it should handle an error condition", func(t *testing.T) {
				_, err := Some(10).
					MapError(IntToIntErrorFn(doubleError)).
					Map(func(x interface{}) interface{} {
						t.Fail()
						return 0
					}).
					JoinError()
				assert.Equal(t, "double error", err.Error())
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
	})
}

func double(x int) int {
	return x * 2
}

func doubleError(x int) (int, error) {
	return 0, errors.New("double error")
}