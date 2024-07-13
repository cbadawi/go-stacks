package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_RosettaAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test RosettaAPIService RosettaAccountBalance", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaAccountBalance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaBlock", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaBlock(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaBlockTransaction", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaBlockTransaction(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaConstructionCombine", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaConstructionCombine(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaConstructionDerive", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaConstructionDerive(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaConstructionHash", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaConstructionHash(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaConstructionMetadata", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaConstructionMetadata(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaConstructionParse", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaConstructionParse(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaConstructionPayloads", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaConstructionPayloads(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaConstructionPreprocess", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaConstructionPreprocess(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaConstructionSubmit", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaConstructionSubmit(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaMempool", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaMempool(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaMempoolTransaction", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaMempoolTransaction(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaNetworkList", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaNetworkList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaNetworkOptions", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaNetworkOptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RosettaAPIService RosettaNetworkStatus", func(t *testing.T) {
		resp, httpRes, err := apiClient.RosettaAPI.RosettaNetworkStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
