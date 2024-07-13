package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_MicroblocksAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test MicroblocksAPIService GetMicroblockByHash", func(t *testing.T) {
		var hash string

		resp, httpRes, _ := apiClient.MicroblocksAPI.GetMicroblockByHash(context.Background(), hash).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MicroblocksAPIService GetMicroblockList", func(t *testing.T) {
		resp, httpRes, _ := apiClient.MicroblocksAPI.GetMicroblockList(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MicroblocksAPIService GetUnanchoredTxs", func(t *testing.T) {
		resp, httpRes, _ := apiClient.MicroblocksAPI.GetUnanchoredTxs(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
