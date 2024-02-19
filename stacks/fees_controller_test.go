package stacks

import (
	"context"
	"testing"

	"github.com/apimatic/go-core-runtime/testHelper"
)

// TODO BUG added an _ to the file name, now its ignored by the compiler
// TODO BUG	expected := `{"$ref":"./api/core-node/get-fee-transfer.example.json"}`
//

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
	expected := "1"
	testHelper.RawBodyMatcher(t, expected, apiResponse.Data)
}
