package api

import (
	"context"
	"testing"

	config "github.com/cbadawi/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_SearchAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test SearchAPIService SearchById", func(t *testing.T) {
		id := "0xcf8b233f19f6c07d2dc1963302d2436efd36e9afac127bf6582824a13961c06d"

		resp, httpRes, _ := apiClient.SearchAPI.SearchById(context.Background(), id).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
