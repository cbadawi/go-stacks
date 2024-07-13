package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_StackingRewardsAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test StackingRewardsAPIService GetBurnchainRewardList", func(t *testing.T) {
		resp, httpRes, _ := apiClient.StackingRewardsAPI.GetBurnchainRewardList(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StackingRewardsAPIService GetBurnchainRewardListByAddress", func(t *testing.T) {
		address := "SPQZF23W7SEYBFG5JQ496NMY0G7379SRYEDREMSV"

		resp, httpRes, _ := apiClient.StackingRewardsAPI.GetBurnchainRewardListByAddress(context.Background(), address).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StackingRewardsAPIService GetBurnchainRewardSlotHolders", func(t *testing.T) {
		resp, httpRes, _ := apiClient.StackingRewardsAPI.GetBurnchainRewardSlotHolders(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StackingRewardsAPIService GetBurnchainRewardSlotHoldersByAddress", func(t *testing.T) {
		address := "SPQZF23W7SEYBFG5JQ496NMY0G7379SRYEDREMSV"

		resp, httpRes, _ := apiClient.StackingRewardsAPI.GetBurnchainRewardSlotHoldersByAddress(context.Background(), address).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StackingRewardsAPIService GetBurnchainRewardsTotalByAddress", func(t *testing.T) {
		address := "SPQZF23W7SEYBFG5JQ496NMY0G7379SRYEDREMSV"

		resp, httpRes, _ := apiClient.StackingRewardsAPI.GetBurnchainRewardsTotalByAddress(context.Background(), address).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
