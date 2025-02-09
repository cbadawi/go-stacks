package api

import (
	"context"
	"testing"

	"github.com/cbadawi/go-stacks/stacks/config"
	"github.com/cbadawi/go-stacks/stacks/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_BlocksAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test BlocksAPIService GetAverageBlockTimes", func(t *testing.T) {
		resp, httpRes, _ := apiClient.BlocksAPI.GetAverageBlockTimes(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test BlocksAPIService GetBlock", func(t *testing.T) {
		value := int32(90000)
		heightOrHash := models.GetBurnBlockHeightOrHashParameter{Int32: &value}

		resp, httpRes, _ := apiClient.BlocksAPI.GetBlock(context.Background(), heightOrHash).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test BlocksAPIService GetBlockByBurnBlockHash", func(t *testing.T) {
		burnBlockHash := "000000000000000000000274d65f8b30011b263c5ec8feb20ff01eb579b19ed5"

		resp, httpRes, _ := apiClient.BlocksAPI.GetBlockByBurnBlockHash(context.Background(), burnBlockHash).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test BlocksAPIService GetBlockByHash", func(t *testing.T) {
		hash := "0xe8a2af8007a3550fe59e86c0f334a9b4089f895491ebda7ea345434227dcc8c6"

		resp, httpRes, _ := apiClient.BlocksAPI.GetBlockByHash(context.Background(), hash).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlockByHeight", func(t *testing.T) {
		height := float32(32)

		resp, httpRes, _ := apiClient.BlocksAPI.GetBlockByHeight(context.Background(), height).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test BlocksAPIService GetBlockList", func(t *testing.T) {
		resp, httpRes, _ := apiClient.BlocksAPI.GetBlockList(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlocksAPIService GetBlocks", func(t *testing.T) {
		resp, httpRes, _ := apiClient.BlocksAPI.GetBlocks(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
