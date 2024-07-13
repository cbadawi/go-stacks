package api

import (
	"context"
	"fmt"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_FeesAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test FeesAPIService GetFeeTransfer", func(t *testing.T) {
		resp, httpRes, err := apiClient.FeesAPI.GetFeeTransfer(context.Background()).Execute()
		fmt.Println(err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeesAPIService PostFeeTransaction", func(t *testing.T) {
		t.Skip("TODO")
		resp, httpRes, _ := apiClient.FeesAPI.PostFeeTransaction(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
