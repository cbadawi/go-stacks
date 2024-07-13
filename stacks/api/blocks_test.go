package api

import (
	"context"
	"testing"

	"github.com/cbadawi/go-stacks/stacks/config"
	"github.com/cbadawi/go-stacks/stacks/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_BlocksAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test BlocksAPIService GetAverageBlockTimes", func(t *testing.T) {
		resp, httpRes, err := apiClient.BlocksAPI.GetAverageBlockTimes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlock", func(t *testing.T) {
		var heightOrHash models.GetBurnBlockHeightOrHashParameter

		resp, httpRes, err := apiClient.BlocksAPI.GetBlock(context.Background(), heightOrHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlockByBurnBlockHash", func(t *testing.T) {
		var burnBlockHash string

		resp, httpRes, err := apiClient.BlocksAPI.GetBlockByBurnBlockHash(context.Background(), burnBlockHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlockByBurnBlockHeight", func(t *testing.T) {
		var burnBlockHeight float32

		resp, httpRes, err := apiClient.BlocksAPI.GetBlockByBurnBlockHeight(context.Background(), burnBlockHeight).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlockByHash", func(t *testing.T) {
		var hash string

		resp, httpRes, err := apiClient.BlocksAPI.GetBlockByHash(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlockByHeight", func(t *testing.T) {
		var height float32

		resp, httpRes, err := apiClient.BlocksAPI.GetBlockByHeight(context.Background(), height).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlockList", func(t *testing.T) {
		resp, httpRes, err := apiClient.BlocksAPI.GetBlockList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlocks", func(t *testing.T) {
		resp, httpRes, err := apiClient.BlocksAPI.GetBlocks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlocksByBurnBlock", func(t *testing.T) {
		var heightOrHash models.GetBurnBlockHeightOrHashParameter

		resp, httpRes, err := apiClient.BlocksAPI.GetBlocksByBurnBlock(context.Background(), heightOrHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
