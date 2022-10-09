package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/husamettinarabaci/tinyurl/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransform(t *testing.T) Transform {
	user := createRandomUser(t)

	arg := CreateTransformParams{
		Owner:    user.Username,
		LongUrl:  util.RandomLongUrl(),
		ShortUrl: util.RandomShortUrl(),
	}

	transform, err := testQueries.CreateTransform(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transform)

	require.Equal(t, arg.Owner, transform.Owner)
	require.Equal(t, arg.LongUrl, transform.LongUrl)
	require.Equal(t, arg.ShortUrl, transform.ShortUrl)

	require.NotZero(t, transform.ID)
	require.NotZero(t, transform.CreatedAt)

	return transform
}

func TestCreateTransform(t *testing.T) {
	createRandomTransform(t)
}

func TestGetTransform(t *testing.T) {
	transform1 := createRandomTransform(t)
	transform2, err := testQueries.GetTransform(context.Background(), transform1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transform2)

	require.Equal(t, transform1.ID, transform2.ID)
	require.Equal(t, transform1.Owner, transform2.Owner)
	require.Equal(t, transform1.LongUrl, transform2.LongUrl)
	require.Equal(t, transform1.ShortUrl, transform2.ShortUrl)
	require.WithinDuration(t, transform1.CreatedAt, transform2.CreatedAt, time.Second)
}

func TestUpdateTransform(t *testing.T) {
	transform1 := createRandomTransform(t)

	arg := UpdateTransformParams{
		ID:       transform1.ID,
		ShortUrl: util.RandomShortUrl(),
	}

	transform2, err := testQueries.UpdateTransform(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transform2)

	require.Equal(t, transform1.ID, transform2.ID)
	require.Equal(t, transform1.Owner, transform2.Owner)
	require.Equal(t, arg.ShortUrl, transform2.ShortUrl)
	require.Equal(t, transform1.LongUrl, transform2.LongUrl)
	require.WithinDuration(t, transform1.CreatedAt, transform2.CreatedAt, time.Second)
}

func TestDeleteTransform(t *testing.T) {
	transform1 := createRandomTransform(t)
	err := testQueries.DeleteTransform(context.Background(), transform1.ID)
	require.NoError(t, err)

	transform2, err := testQueries.GetTransform(context.Background(), transform1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transform2)
}

func TestListTransforms(t *testing.T) {
	var lastTransform Transform
	for i := 0; i < 10; i++ {
		lastTransform = createRandomTransform(t)
	}

	arg := ListTransformsParams{
		Owner:  lastTransform.Owner,
		Limit:  5,
		Offset: 0,
	}

	transforms, err := testQueries.ListTransforms(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transforms)

	for _, transform := range transforms {
		require.NotEmpty(t, transform)
		require.Equal(t, lastTransform.Owner, transform.Owner)
	}
}
