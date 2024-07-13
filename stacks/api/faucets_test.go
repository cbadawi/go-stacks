package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_FaucetsAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test FaucetsAPIService RunFaucetBtc", func(t *testing.T) {
		resp, httpRes, err := apiClient.FaucetsAPI.RunFaucetBtc(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaucetsAPIService RunFaucetStx", func(t *testing.T) {
		resp, httpRes, err := apiClient.FaucetsAPI.RunFaucetStx(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
