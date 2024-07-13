package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/cbadawi/go-stacks/stacks/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_BurnBlocksAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test BurnBlocksAPIService GetBurnBlock", func(t *testing.T) {
		var heightOrHash models.GetBurnBlockHeightOrHashParameter

		resp, httpRes, err := apiClient.BurnBlocksAPI.GetBurnBlock(context.Background(), heightOrHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BurnBlocksAPIService GetBurnBlocks", func(t *testing.T) {
		resp, httpRes, err := apiClient.BurnBlocksAPI.GetBurnBlocks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
