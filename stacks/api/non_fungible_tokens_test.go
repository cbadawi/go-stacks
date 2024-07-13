package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_NonFungibleTokensAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test NonFungibleTokensAPIService GetNftHistory", func(t *testing.T) {
		resp, httpRes, err := apiClient.NonFungibleTokensAPI.GetNftHistory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonFungibleTokensAPIService GetNftHoldings", func(t *testing.T) {
		resp, httpRes, err := apiClient.NonFungibleTokensAPI.GetNftHoldings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonFungibleTokensAPIService GetNftMints", func(t *testing.T) {
		resp, httpRes, err := apiClient.NonFungibleTokensAPI.GetNftMints(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
