package bing

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetImages(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()
	imgs, err := GetImages(ctx)
	require.Nil(t, err)
	require.Len(t, imgs, 8)
	for i, img := range imgs {
		fmt.Println(i, " ", img.FullURL())
	}
}
