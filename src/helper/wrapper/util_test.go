package wrapper_test

import (
	"testing"

	"gitlab.playcourt.id/new-mypertamina/myptm-external-insurance-service/src/helper/errors"

	"github.com/stretchr/testify/assert"

	"gitlab.playcourt.id/new-mypertamina/myptm-external-insurance-service/src/helper/wrapper"
)

func TestFindOneErrorResult(t *testing.T) {
	t.Run("should return at least one error", func(t *testing.T) {
		errorResult := &wrapper.Result{Err: errors.ErrInternalServer}
		results := []*wrapper.Result{
			new(wrapper.Result),
			new(wrapper.Result),
			errorResult,
		}

		er := wrapper.FindOneErrorResult(results...)
		assert.NotNil(t, er)
		assert.Equal(t, errors.ErrInternalServer, er.Err)
	})

	t.Run("should return nil pointer for result caused by no error found", func(t *testing.T) {
		results := []*wrapper.Result{
			new(wrapper.Result),
			new(wrapper.Result),
		}

		er := wrapper.FindOneErrorResult(results...)
		assert.Nil(t, er)
	})
}
