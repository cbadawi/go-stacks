package stacksblockchainapi

import (
	"context"
	"testing"

	"github.com/apimatic/go-core-runtime/testHelper"
)

// TestNamesControllerTestGetNamespacePrice tests the behavior of the NamesController's
func TestNamesControllerTestGetNamespacePrice(t *testing.T) {
	ctx := context.Background()
	tld := "id"
	apiResponse, err := NamesController.GetNamespacePrice(ctx, tld)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestNamesControllerTestGetNamePrice tests the behavior of the NamesController's
func TestNamesControllerTestGetNamePrice(t *testing.T) {
	ctx := context.Background()
	name := "muneeb.id"
	apiResponse, err := NamesController.GetNamePrice(ctx, name)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestNamesControllerTestFetchZoneFile tests the behavior of the NamesController's
func TestNamesControllerTestFetchZoneFile(t *testing.T) {
	ctx := context.Background()
	name := "bar.test"
	apiResponse, err := NamesController.FetchZoneFile(ctx, name)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"$ref":"./api/bns/manage-names/bns-fetch-zone-file-response.example.json"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestNamesControllerTestGetHistoricalZoneFile tests the behavior of the NamesController's
func TestNamesControllerTestGetHistoricalZoneFile(t *testing.T) {
	ctx := context.Background()
	name := "muneeb.id"
	zoneFileHash := "b100a68235244b012854a95f9114695679002af9"
	apiResponse, err := NamesController.GetHistoricalZoneFile(ctx, name, zoneFileHash)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"$ref":"./api/bns/name-querying/bns-get-historical-zone-file-response.example.json"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}
