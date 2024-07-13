package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_NamesAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test NamesAPIService FetchSubdomainsListForName", func(t *testing.T) {
		name := "stacks.btc"

		resp, httpRes, _ := apiClient.NamesAPI.FetchSubdomainsListForName(context.Background(), name).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService FetchZoneFile", func(t *testing.T) {
		name := "stacks.btc"

		resp, httpRes, _ := apiClient.NamesAPI.FetchZoneFile(context.Background(), name).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetAllNames", func(t *testing.T) {
		resp, httpRes, _ := apiClient.NamesAPI.GetAllNames(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetAllNamespaces", func(t *testing.T) {
		resp, httpRes, _ := apiClient.NamesAPI.GetAllNamespaces(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetHistoricalZoneFile", func(t *testing.T) {
		name := "stacks.btc"
		var zoneFileHash string

		resp, httpRes, _ := apiClient.NamesAPI.GetHistoricalZoneFile(context.Background(), name, zoneFileHash).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNameInfo", func(t *testing.T) {
		name := "stacks.btc"

		resp, httpRes, _ := apiClient.NamesAPI.GetNameInfo(context.Background(), name).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamePrice", func(t *testing.T) {
		name := "stacks.btc"

		resp, httpRes, _ := apiClient.NamesAPI.GetNamePrice(context.Background(), name).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamesOwnedByAddress", func(t *testing.T) {
		blockchain := "stacks"
		address := "SP000000000000000000002Q6VF78"

		resp, httpRes, _ := apiClient.NamesAPI.GetNamesOwnedByAddress(context.Background(), blockchain, address).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamespaceNames", func(t *testing.T) {
		tld := "id"

		resp, httpRes, _ := apiClient.NamesAPI.GetNamespaceNames(context.Background(), tld).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	// t.Run("Test NamesAPIService GetNamespacePrice", func(t *testing.T) {
	// 	var tld string

	// 	resp, httpRes, _ := apiClient.NamesAPI.GetNamespacePrice(context.Background(), tld).Execute()

	// 	require.NotNil(t, resp)
	// 	assert.Equal(t, 200, httpRes.StatusCode)

	// })

}
