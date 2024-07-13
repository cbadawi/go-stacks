package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_SmartContractsAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test SmartContractsAPIService CallReadOnlyFunction", func(t *testing.T) {
		var contractAddress string
		var contractName string
		var functionName string

		resp, httpRes, err := apiClient.SmartContractsAPI.CallReadOnlyFunction(context.Background(), contractAddress, contractName, functionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetContractById", func(t *testing.T) {
		var contractId string

		resp, httpRes, err := apiClient.SmartContractsAPI.GetContractById(context.Background(), contractId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetContractDataMapEntry", func(t *testing.T) {
		var contractAddress string
		var contractName string
		var mapName string

		resp, httpRes, err := apiClient.SmartContractsAPI.GetContractDataMapEntry(context.Background(), contractAddress, contractName, mapName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetContractEventsById", func(t *testing.T) {
		var contractId string

		resp, httpRes, err := apiClient.SmartContractsAPI.GetContractEventsById(context.Background(), contractId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetContractInterface", func(t *testing.T) {
		var contractAddress string
		var contractName string

		resp, httpRes, err := apiClient.SmartContractsAPI.GetContractInterface(context.Background(), contractAddress, contractName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetContractSource", func(t *testing.T) {
		var contractAddress string
		var contractName string

		resp, httpRes, err := apiClient.SmartContractsAPI.GetContractSource(context.Background(), contractAddress, contractName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetContractsByTrait", func(t *testing.T) {
		resp, httpRes, err := apiClient.SmartContractsAPI.GetContractsByTrait(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetSmartContractsStatus", func(t *testing.T) {
		resp, httpRes, err := apiClient.SmartContractsAPI.GetSmartContractsStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
