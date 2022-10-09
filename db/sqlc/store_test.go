package db

import (
	"context"
	"testing"

	"github.com/husamettinarabaci/tinyurl/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransformTx(t *testing.T) {

	store := NewStore(testDB)

	user := createRandomUser(t)

	result, err := store.CreateTransformTx(context.Background(), CreateTransformParams{
		Owner:    user.Username,
		LongUrl:  util.RandomLongUrl(),
		ShortUrl: util.RandomShortUrl(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotZero(t, result.ID)

}

func TestDeleteransformTx(t *testing.T) {

	store := NewStore(testDB)

	user := createRandomUser(t)

	result, err := store.CreateTransformTx(context.Background(), CreateTransformParams{
		Owner:    user.Username,
		LongUrl:  util.RandomLongUrl(),
		ShortUrl: util.RandomShortUrl(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)

	err = store.DeleteTransformTx(context.Background(), result.ID)

	require.NoError(t, err)

	user2, err := store.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.Equal(t, user.UrlCount, user2.UrlCount)

}
