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
	configuration.Logger.Verbose = true
	apiClient := NewAPIClient(configuration)

	t.Run("Test TransactionsAPIService GetAddressMempoolTransactions", func(t *testing.T) {
		address := "SP000000000000000000002Q6VF78"

		resp, httpRes, _ := apiClient.TransactionsAPI.GetAddressMempoolTransactions(context.Background(), address).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetAddressTransactions", func(t *testing.T) {
		address := "SP000000000000000000002Q6VF78"

		resp, httpRes, _ := apiClient.TransactionsAPI.GetAddressTransactions(context.Background(), address).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetDroppedMempoolTransactionList", func(t *testing.T) {
		resp, httpRes, _ := apiClient.TransactionsAPI.GetDroppedMempoolTransactionList(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetFilteredEvents", func(t *testing.T) {
		adress := "SP11M3BVG5TGNWQ1FFE5EJM86TMV08YF3EK0JSMQ8"
		req := apiClient.TransactionsAPI.GetFilteredEvents(context.Background())
		req.address = &adress

		resp, httpRes, _ := req.Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TransactionsAPIService GetMempoolTransactionList", func(t *testing.T) {
		resp, httpRes, _ := apiClient.TransactionsAPI.GetMempoolTransactionList(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TransactionsAPIService GetMempoolTransactionStats", func(t *testing.T) {
		resp, httpRes, _ := apiClient.TransactionsAPI.GetMempoolTransactionStats(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetRawTransactionById", func(t *testing.T) {
		txId := "0xbc617b30c0932bbc11b21911a23324b50db1789e58e1c28b8d2da3b71c99808f"

		resp, httpRes, _ := apiClient.TransactionsAPI.GetRawTransactionById(context.Background(), txId).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TransactionsAPIService GetTransactionById", func(t *testing.T) {
		txId := "0xbc617b30c0932bbc11b21911a23324b50db1789e58e1c28b8d2da3b71c99808f"

		resp, httpRes, _ := apiClient.TransactionsAPI.GetTransactionById(context.Background(), txId).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TransactionsAPIService GetTransactionList", func(t *testing.T) {
		resp, httpRes, _ := apiClient.TransactionsAPI.GetTransactionList(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TransactionsAPIService GetTransactionsByBlock", func(t *testing.T) {
		value := int32(90000)
		heightOrHash := models.GetBurnBlockHeightOrHashParameter{Int32: &value}

		resp, httpRes, _ := apiClient.TransactionsAPI.GetTransactionsByBlock(context.Background(), heightOrHash).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionsByBlockHash", func(t *testing.T) {
		var blockHash string

		resp, httpRes, _ := apiClient.TransactionsAPI.GetTransactionsByBlockHash(context.Background(), blockHash).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionsByBlockHeight", func(t *testing.T) {
		height := int32(100000)

		resp, httpRes, _ := apiClient.TransactionsAPI.GetTransactionsByBlockHeight(context.Background(), height).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTxListDetails", func(t *testing.T) {
		req := apiClient.TransactionsAPI.GetTxListDetails(context.Background())
		req.txId = &[]string{"0xbc617b30c0932bbc11b21911a23324b50db1789e58e1c28b8d2da3b71c99808f"}
		resp, httpRes, _ := req.Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService PostCoreNodeTransactions", func(t *testing.T) {
		resp, httpRes, _ := apiClient.TransactionsAPI.PostCoreNodeTransactions(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
