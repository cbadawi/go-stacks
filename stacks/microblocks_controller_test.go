package stacks

import (
	"context"
	"testing"

	"github.com/apimatic/go-core-runtime/testHelper"
)

// TestMicroblocksControllerTestGetMicroblockList tests the behavior of the MicroblocksController's
func TestMicroblocksControllerTestGetMicroblockList(t *testing.T) {
	ctx := context.Background()
	limit := 20
	offset := 42000
	apiResponse, err := microblocksController.GetMicroblockList(ctx, &limit, &offset)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json; charset=utf-8"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestMicroblocksControllerTestGetMicroblockByHash tests the behavior of the MicroblocksController's
func TestMicroblocksControllerTestGetMicroblockByHash(t *testing.T) {
	ctx := context.Background()
	hash := "0x3bfcdf84b3012adb544cf0f6df4835f93418c2269a3881885e27b3d58eb82d47"
	apiResponse, err := microblocksController.GetMicroblockByHash(ctx, hash)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json; charset=utf-8"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestMicroblocksControllerTestGetUnanchoredTxs tests the behavior of the MicroblocksController's
func TestMicroblocksControllerTestGetUnanchoredTxs(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := microblocksController.GetUnanchoredTxs(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json; charset=utf-8"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
