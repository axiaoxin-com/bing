package bing

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetImageURL(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()
	urls, err := GetImageURL(ctx, ImageResolution1920x1080)
	require.Nil(t, err)
	require.Len(t, urls, 8)
	for i, url := range urls {
		fmt.Println(i, " ", url)
	}
}
