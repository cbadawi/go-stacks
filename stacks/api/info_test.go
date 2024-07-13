package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_InfoAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test InfoAPIService GetCoreApiInfo", func(t *testing.T) {
		resp, httpRes, _ := apiClient.InfoAPI.GetCoreApiInfo(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetNetworkBlockTimeByNetwork", func(t *testing.T) {
		network := "mainnet"

		resp, httpRes, _ := apiClient.InfoAPI.GetNetworkBlockTimeByNetwork(context.Background(), network).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetNetworkBlockTimes", func(t *testing.T) {
		resp, httpRes, _ := apiClient.InfoAPI.GetNetworkBlockTimes(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetPoxInfo", func(t *testing.T) {
		resp, httpRes, _ := apiClient.InfoAPI.GetPoxInfo(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetStatus", func(t *testing.T) {
		resp, httpRes, _ := apiClient.InfoAPI.GetStatus(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetStxSupply", func(t *testing.T) {
		resp, httpRes, _ := apiClient.InfoAPI.GetStxSupply(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetStxSupplyCirculatingPlain", func(t *testing.T) {
		resp, httpRes, _ := apiClient.InfoAPI.GetStxSupplyCirculatingPlain(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetStxSupplyTotalSupplyPlain", func(t *testing.T) {
		resp, httpRes, _ := apiClient.InfoAPI.GetStxSupplyTotalSupplyPlain(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})
}
