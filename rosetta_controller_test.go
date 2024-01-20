package stacksblockchainapi

import (
	"context"
	"testing"

	"github.com/apimatic/go-core-runtime/testHelper"
)

// TestRosettaControllerTestRosettaNetworkList tests the behavior of the RosettaController's
func TestRosettaControllerTestRosettaNetworkList(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := rosettaController.RosettaNetworkList(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
