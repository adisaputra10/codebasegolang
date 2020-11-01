package wrapper_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.playcourt.id/new-mypertamina/myptm-external-insurance-service/src/helper/wrapper"
)

func TestHTTPResponseJSON(t *testing.T) {
	t.Run("should return http error result", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Err = errors.New("error")

		wrapper.HTTPResponseJSON(recorder, result, http.StatusInternalServerError, "Error")

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should return http success result", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Data = "test"

		wrapper.HTTPResponseJSON(recorder, result, http.StatusOK, "just for unit test")

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
