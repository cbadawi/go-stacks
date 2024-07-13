package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_AccountsAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test AccountsAPIService GetAccountAssets", func(t *testing.T) {
		var principal string

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountAssets(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountBalance", func(t *testing.T) {
		var principal string

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountBalance(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountInbound", func(t *testing.T) {
		var principal string

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountInbound(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountInfo", func(t *testing.T) {
		var principal string

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountInfo(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountNonces", func(t *testing.T) {
		var principal string

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountNonces(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountStxBalance", func(t *testing.T) {
		var principal string

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountStxBalance(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountTransactions", func(t *testing.T) {
		var principal string

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountTransactions(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountTransactionsWithTransfers", func(t *testing.T) {
		var principal string

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountTransactionsWithTransfers(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetSingleTransactionWithTransfers", func(t *testing.T) {
		var principal string
		var txId string

		resp, httpRes, err := apiClient.AccountsAPI.GetSingleTransactionWithTransfers(context.Background(), principal, txId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
