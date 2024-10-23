package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPwd(t *testing.T) {
	pwd := RandomString(6)

	hashPwd, err := HashPwd(pwd)
	require.NoError(t, err)
	require.NotEmpty(t, hashPwd)

	err = CheckPwd(pwd, hashPwd)
	require.NoError(t, err)

	wrongPwd := RandomString(6)
	err = CheckPwd(wrongPwd, hashPwd)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
