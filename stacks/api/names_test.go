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
		var name string

		resp, httpRes, err := apiClient.NamesAPI.FetchSubdomainsListForName(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService FetchZoneFile", func(t *testing.T) {
		var name string

		resp, httpRes, err := apiClient.NamesAPI.FetchZoneFile(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetAllNames", func(t *testing.T) {
		resp, httpRes, err := apiClient.NamesAPI.GetAllNames(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetAllNamespaces", func(t *testing.T) {
		resp, httpRes, err := apiClient.NamesAPI.GetAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetHistoricalZoneFile", func(t *testing.T) {
		var name string
		var zoneFileHash string

		resp, httpRes, err := apiClient.NamesAPI.GetHistoricalZoneFile(context.Background(), name, zoneFileHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNameInfo", func(t *testing.T) {
		var name string

		resp, httpRes, err := apiClient.NamesAPI.GetNameInfo(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamePrice", func(t *testing.T) {
		var name string

		resp, httpRes, err := apiClient.NamesAPI.GetNamePrice(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamesOwnedByAddress", func(t *testing.T) {
		var blockchain string
		var address string

		resp, httpRes, err := apiClient.NamesAPI.GetNamesOwnedByAddress(context.Background(), blockchain, address).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamespaceNames", func(t *testing.T) {
		var tld string

		resp, httpRes, err := apiClient.NamesAPI.GetNamespaceNames(context.Background(), tld).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamespacePrice", func(t *testing.T) {
		var tld string

		resp, httpRes, err := apiClient.NamesAPI.GetNamespacePrice(context.Background(), tld).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
