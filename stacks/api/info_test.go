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
		resp, httpRes, err := apiClient.InfoAPI.GetCoreApiInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetNetworkBlockTimeByNetwork", func(t *testing.T) {
		var network string

		resp, httpRes, err := apiClient.InfoAPI.GetNetworkBlockTimeByNetwork(context.Background(), network).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetNetworkBlockTimes", func(t *testing.T) {
		resp, httpRes, err := apiClient.InfoAPI.GetNetworkBlockTimes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetPoxInfo", func(t *testing.T) {
		resp, httpRes, err := apiClient.InfoAPI.GetPoxInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetStatus", func(t *testing.T) {
		resp, httpRes, err := apiClient.InfoAPI.GetStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetStxSupply", func(t *testing.T) {
		resp, httpRes, err := apiClient.InfoAPI.GetStxSupply(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetStxSupplyCirculatingPlain", func(t *testing.T) {
		resp, httpRes, err := apiClient.InfoAPI.GetStxSupplyCirculatingPlain(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetStxSupplyTotalSupplyPlain", func(t *testing.T) {
		resp, httpRes, err := apiClient.InfoAPI.GetStxSupplyTotalSupplyPlain(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InfoAPIService GetTotalStxSupplyLegacyFormat", func(t *testing.T) {
		resp, httpRes, err := apiClient.InfoAPI.GetTotalStxSupplyLegacyFormat(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
