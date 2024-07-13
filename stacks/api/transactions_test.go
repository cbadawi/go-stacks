package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/cbadawi/go-stacks/stacks/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_TransactionsAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test TransactionsAPIService GetAddressMempoolTransactions", func(t *testing.T) {
		var address string

		resp, httpRes, err := apiClient.TransactionsAPI.GetAddressMempoolTransactions(context.Background(), address).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetAddressTransactionEvents", func(t *testing.T) {
		var address string
		var txId string

		resp, httpRes, err := apiClient.TransactionsAPI.GetAddressTransactionEvents(context.Background(), address, txId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetAddressTransactions", func(t *testing.T) {
		var address string

		resp, httpRes, err := apiClient.TransactionsAPI.GetAddressTransactions(context.Background(), address).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetDroppedMempoolTransactionList", func(t *testing.T) {
		resp, httpRes, err := apiClient.TransactionsAPI.GetDroppedMempoolTransactionList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetFilteredEvents", func(t *testing.T) {
		resp, httpRes, err := apiClient.TransactionsAPI.GetFilteredEvents(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetMempoolTransactionList", func(t *testing.T) {
		resp, httpRes, err := apiClient.TransactionsAPI.GetMempoolTransactionList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetMempoolTransactionStats", func(t *testing.T) {
		resp, httpRes, err := apiClient.TransactionsAPI.GetMempoolTransactionStats(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetRawTransactionById", func(t *testing.T) {
		var txId string

		resp, httpRes, err := apiClient.TransactionsAPI.GetRawTransactionById(context.Background(), txId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionById", func(t *testing.T) {
		var txId string

		resp, httpRes, err := apiClient.TransactionsAPI.GetTransactionById(context.Background(), txId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionList", func(t *testing.T) {
		resp, httpRes, err := apiClient.TransactionsAPI.GetTransactionList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionsByBlock", func(t *testing.T) {
		var heightOrHash models.GetBurnBlockHeightOrHashParameter

		resp, httpRes, err := apiClient.TransactionsAPI.GetTransactionsByBlock(context.Background(), heightOrHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionsByBlockHash", func(t *testing.T) {
		var blockHash string

		resp, httpRes, err := apiClient.TransactionsAPI.GetTransactionsByBlockHash(context.Background(), blockHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionsByBlockHeight", func(t *testing.T) {
		var height int32

		resp, httpRes, err := apiClient.TransactionsAPI.GetTransactionsByBlockHeight(context.Background(), height).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTxListDetails", func(t *testing.T) {
		resp, httpRes, err := apiClient.TransactionsAPI.GetTxListDetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService PostCoreNodeTransactions", func(t *testing.T) {
		resp, httpRes, err := apiClient.TransactionsAPI.PostCoreNodeTransactions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
