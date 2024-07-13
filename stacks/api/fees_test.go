package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_FeesAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test FeesAPIService FetchFeeRate", func(t *testing.T) {
		resp, httpRes, err := apiClient.FeesAPI.FetchFeeRate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeesAPIService GetFeeTransfer", func(t *testing.T) {
		resp, httpRes, err := apiClient.FeesAPI.GetFeeTransfer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeesAPIService PostFeeTransaction", func(t *testing.T) {
		resp, httpRes, err := apiClient.FeesAPI.PostFeeTransaction(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
