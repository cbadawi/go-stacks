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

	t.Run("Test SmartContractsAPIService GetContractById", func(t *testing.T) {
		contractId := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C.satoshibles"

		resp, httpRes, _ := apiClient.SmartContractsAPI.GetContractById(context.Background(), contractId).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetContractEventsById", func(t *testing.T) {
		contractId := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C.satoshibles"

		resp, httpRes, _ := apiClient.SmartContractsAPI.GetContractEventsById(context.Background(), contractId).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetContractInterface", func(t *testing.T) {
		contractAddress := "SPSCWDV3RKV5ZRN1FQD84YE1NQFEDJ9R1F4DYQ11"
		contractName := "newyorkcitycoin-core-v2"

		resp, httpRes, _ := apiClient.SmartContractsAPI.GetContractInterface(context.Background(), contractAddress, contractName).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SmartContractsAPIService GetContractSource", func(t *testing.T) {
		contractAddress := "SPSCWDV3RKV5ZRN1FQD84YE1NQFEDJ9R1F4DYQ11"
		contractName := "newyorkcitycoin-core-v2"

		resp, httpRes, _ := apiClient.SmartContractsAPI.GetContractSource(context.Background(), contractAddress, contractName).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartContractsAPIService GetSmartContractsStatus", func(t *testing.T) {
		req := apiClient.SmartContractsAPI.GetSmartContractsStatus(context.Background())
		contract := []string{"SPQZF23W7SEYBFG5JQ496NMY0G7379SRYEDREMSV.Candy"}
		req.contractId = &contract
		resp, httpRes, _ := req.Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
