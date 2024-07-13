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
		cycleNumber := int32(2)

		resp, httpRes, _ := apiClient.ProofOfTransferAPI.GetPoxCycle(context.Background(), cycleNumber).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProofOfTransferAPIService GetPoxCycleSigner", func(t *testing.T) {
		cycleNumber := int32(6)
		signerKey := "0x038e3c4529395611be9abf6fa3b6987e81d402385e3d605a073f42f407565a4a3d"

		resp, httpRes, _ := apiClient.ProofOfTransferAPI.GetPoxCycleSigner(context.Background(), cycleNumber, signerKey).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProofOfTransferAPIService GetPoxCycleSignerStackers", func(t *testing.T) {
		cycleNumber := int32(6)
		signerKey := "0x038e3c4529395611be9abf6fa3b6987e81d402385e3d605a073f42f407565a4a3d"

		resp, httpRes, _ := apiClient.ProofOfTransferAPI.GetPoxCycleSignerStackers(context.Background(), cycleNumber, signerKey).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProofOfTransferAPIService GetPoxCycleSigners", func(t *testing.T) {
		cycleNumber := int32(6)

		resp, httpRes, _ := apiClient.ProofOfTransferAPI.GetPoxCycleSigners(context.Background(), cycleNumber).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProofOfTransferAPIService GetPoxCycles", func(t *testing.T) {
		resp, httpRes, _ := apiClient.ProofOfTransferAPI.GetPoxCycles(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
