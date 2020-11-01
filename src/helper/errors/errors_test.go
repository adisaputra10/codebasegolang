package errors_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	er "errors"

	e "github.com/adisaputra10/codebasegolang/src/helper/errors"
)

func TestGetHTTPStatusCodeByError(t *testing.T) {
	type spec struct {
		err  error
		code int
	}
	testTable := []spec{
		{e.ErrUnauthorized, http.StatusUnauthorized},
		{e.ErrNotFound, http.StatusNotFound},
		{e.ErrForbidden, http.StatusForbidden},
		{e.ErrInternalServer, http.StatusInternalServerError},
		{e.ErrUnprocessableEntitiy, http.StatusUnprocessableEntity},
		{e.ErrBadRequest, http.StatusBadRequest},
		{e.ErrMethodNotAllowed, http.StatusMethodNotAllowed},
		{e.ErrLocked, http.StatusLocked},
		{er.New("anyting"), http.StatusInternalServerError},
	}

	for _, spec := range testTable {
		expected := spec.code
		actual := e.GetHTTPStatusCodeByError(spec.err)

		assert.Equal(t, expected, actual)
	}
}
