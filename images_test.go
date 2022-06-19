package bing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetImageURL(t *testing.T) {
	urls, err := GetImageURL(16, true)
	require.Nil(t, err)
	require.Len(t, urls, 16)
	t.Log(urls)
}
