package db

import (
	"context"
	"testing"
	"time"

	"github.com/ericlamnguyen/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	// create a random password and hash it
	hashedPwd, err := util.HashPwd(util.RandomString(6))
	require.NoError(t, err)

	args := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPwd,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	retrieved_user, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, retrieved_user)
	require.Equal(t, user.Username, retrieved_user.Username)
	require.Equal(t, user.HashedPassword, retrieved_user.HashedPassword)
	require.Equal(t, user.FullName, retrieved_user.FullName)
	require.Equal(t, user.Email, retrieved_user.Email)
	require.WithinDuration(t, user.PasswordChangedAt, retrieved_user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, retrieved_user.CreatedAt, time.Second)
}
