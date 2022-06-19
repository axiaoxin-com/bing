package bing

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetImageURL(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()
	urls, err := GetImageURL(ctx, 3, true)
	require.Nil(t, err)
	require.Len(t, urls, 3)
	t.Log(urls)
}
