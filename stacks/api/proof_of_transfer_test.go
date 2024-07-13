package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ProofOfTransferAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test ProofOfTransferAPIService GetPoxCycle", func(t *testing.T) {
		var cycleNumber int32

		resp, httpRes, err := apiClient.ProofOfTransferAPI.GetPoxCycle(context.Background(), cycleNumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProofOfTransferAPIService GetPoxCycleSigner", func(t *testing.T) {
		var cycleNumber int32
		var signerKey string

		resp, httpRes, err := apiClient.ProofOfTransferAPI.GetPoxCycleSigner(context.Background(), cycleNumber, signerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProofOfTransferAPIService GetPoxCycleSignerStackers", func(t *testing.T) {
		var cycleNumber int32
		var signerKey string

		resp, httpRes, err := apiClient.ProofOfTransferAPI.GetPoxCycleSignerStackers(context.Background(), cycleNumber, signerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProofOfTransferAPIService GetPoxCycleSigners", func(t *testing.T) {
		var cycleNumber int32

		resp, httpRes, err := apiClient.ProofOfTransferAPI.GetPoxCycleSigners(context.Background(), cycleNumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProofOfTransferAPIService GetPoxCycles", func(t *testing.T) {
		resp, httpRes, err := apiClient.ProofOfTransferAPI.GetPoxCycles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
