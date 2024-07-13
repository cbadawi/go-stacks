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
		asset := "SPQZF23W7SEYBFG5JQ496NMY0G7379SRYEDREMSV.Candy::candy"
		value := "0x0100000000000000000000000000000803"
		req := apiClient.NonFungibleTokensAPI.GetNftHistory(context.Background())
		req.value = &value
		req.assetIdentifier = &asset
		resp, httpRes, _ := req.Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test NonFungibleTokensAPIService GetNftHoldings", func(t *testing.T) {
		principal := "SPNWZ5V2TPWGQGVDR6T7B6RQ4XMGZ4PXTEE0VQ0S.marketplace-v3"
		req := apiClient.NonFungibleTokensAPI.GetNftHoldings(context.Background())
		req.principal = &principal
		resp, httpRes, _ := req.Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test NonFungibleTokensAPIService GetNftMints", func(t *testing.T) {
		assetIdentifier := "SP2X0TZ59D5SZ8ACQ6YMCHHNR2ZN51Z32E2CJ173.the-explorer-guild::The-Explorer-Guild"
		req := apiClient.NonFungibleTokensAPI.GetNftMints(context.Background())
		req.assetIdentifier = &assetIdentifier
		resp, httpRes, _ := req.Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
