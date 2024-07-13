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
		resp, httpRes, err := apiClient.StackingRewardsAPI.GetBurnchainRewardList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StackingRewardsAPIService GetBurnchainRewardListByAddress", func(t *testing.T) {
		var address string

		resp, httpRes, err := apiClient.StackingRewardsAPI.GetBurnchainRewardListByAddress(context.Background(), address).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StackingRewardsAPIService GetBurnchainRewardSlotHolders", func(t *testing.T) {
		resp, httpRes, err := apiClient.StackingRewardsAPI.GetBurnchainRewardSlotHolders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StackingRewardsAPIService GetBurnchainRewardSlotHoldersByAddress", func(t *testing.T) {
		var address string

		resp, httpRes, err := apiClient.StackingRewardsAPI.GetBurnchainRewardSlotHoldersByAddress(context.Background(), address).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StackingRewardsAPIService GetBurnchainRewardsTotalByAddress", func(t *testing.T) {
		var address string

		resp, httpRes, err := apiClient.StackingRewardsAPI.GetBurnchainRewardsTotalByAddress(context.Background(), address).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
