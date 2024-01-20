package main

import (
	"context"
	"testing"

	"github.com/apimatic/go-core-runtime/testHelper"
)

// TestFeesControllerTestGetFeeTransfer tests the behavior of the FeesController's
func TestFeesControllerTestGetFeeTransfer(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := feesController.GetFeeTransfer(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"$ref":"./api/core-node/get-fee-transfer.example.json"}`
	testHelper.RawBodyMatcher(t, expected, apiResponse.Data)
}
