package util_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/adisaputra10/codebasegolang/src/helper/util"
)

func TestToISODate(t *testing.T) {
	t.Run("should return error because of invalid location", func(t *testing.T) {
		loc := "test-unknown"
		time := time.Now()
		_, err := util.ToISODate(time, loc)
		assert.Error(t, err)
	})

	t.Run("should return success", func(t *testing.T) {
		expected := "2019-12-31T17:00:00.000Z"
		loc := "UTC"
		time, _ := time.Parse("2006-01-02T15:04:05.000Z07:00", expected)
		isoString, _ := util.ToISODate(time, loc)
		assert.Equal(t, expected, isoString)
	})
}

func TestGetRSAPublicKey(t *testing.T) {
	t.Run("should return nil pointer of public key because of file is not found", func(t *testing.T) {
		rsaPublicKey := util.GetRSAPublicKey("./no-file.pem")
		assert.Nil(t, rsaPublicKey)
	})

	t.Run("should return success", func(t *testing.T) {
		rsaPublicKey := util.GetRSAPublicKey("./../../../public.pem")
		assert.NotNil(t, rsaPublicKey)
	})
}

func TestGetRSAPrivateKey(t *testing.T) {
	t.Run("should return nil pointer of private key because of file is not found", func(t *testing.T) {
		rsaPrivateKey := util.GetRSAPrivateKey("./no-file.pem")
		assert.Nil(t, rsaPrivateKey)
	})

	t.Run("should return success", func(t *testing.T) {
		rsaPrivateKey := util.GetRSAPrivateKey("./../../../private.pem")
		assert.NotNil(t, rsaPrivateKey)
	})
}
