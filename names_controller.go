package stacksblockchainapi

import (
	"context"
	"fmt"
	"stacksblockchainapi/errors"
	"stacksblockchainapi/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// NamesController represents a controller struct.
type NamesController struct {
	baseController
}

// NewNamesController creates a new instance of NamesController.
// It takes a baseController as a parameter and returns a pointer to the NamesController.
func NewNamesController(baseController baseController) *NamesController {
	namesController := NamesController{baseController: baseController}
	return &namesController
}

// GetNamespacePrice takes context, tld as parameters and
// returns an models.ApiResponse with models.BnsGetNamespacePriceResponse data and
// an error if there was an issue with the request or response.
// Retrieves the price of a namespace. The `amount` given will be in the smallest possible units of the currency.
func (n *NamesController) GetNamespacePrice(
	ctx context.Context,
	tld string) (
	models.ApiResponse[models.BnsGetNamespacePriceResponse],
	error) {
	req := n.prepareRequest(ctx, "GET", fmt.Sprintf("/v2/prices/namespaces/%v", tld))
	req.Authenticate(true)

	var result models.BnsGetNamespacePriceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BnsGetNamespacePriceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetNamePrice takes context, name as parameters and
// returns an models.ApiResponse with models.BnsGetNamePriceResponse data and
// an error if there was an issue with the request or response.
// Retrieves the price of a name. The `amount` given will be in the smallest possible units of the currency.
func (n *NamesController) GetNamePrice(
	ctx context.Context,
	name string) (
	models.ApiResponse[models.BnsGetNamePriceResponse],
	error) {
	req := n.prepareRequest(ctx, "GET", fmt.Sprintf("/v2/prices/names/%v", name))
	req.Authenticate(true)

	var result models.BnsGetNamePriceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BnsGetNamePriceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAllNamespaces takes context as parameters and
// returns an models.ApiResponse with models.BnsGetAllNamespacesResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of all namespaces known to the node.
func (n *NamesController) GetAllNamespaces(ctx context.Context) (
	models.ApiResponse[models.BnsGetAllNamespacesResponse],
	error) {
	req := n.prepareRequest(ctx, "GET", "/v1/namespaces")
	req.Authenticate(true)
	var result models.BnsGetAllNamespacesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BnsGetAllNamespacesResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetNamespaceNames takes context, tld, page as parameters and
// returns an models.ApiResponse with []string data and
// an error if there was an issue with the request or response.
// Retrieves a list of names within a given namespace.
func (n *NamesController) GetNamespaceNames(
	ctx context.Context,
	tld string,
	page *int) (
	models.ApiResponse[[]string],
	error) {
	req := n.prepareRequest(ctx, "GET", fmt.Sprintf("/v1/namespaces/%v/names", tld))
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []string
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]string](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewBnsError(400, "Error")
	}
	if resp.StatusCode == 404 {
		err = errors.NewBnsError(404, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// GetAllNames takes context, page as parameters and
// returns an models.ApiResponse with []string data and
// an error if there was an issue with the request or response.
// Retrieves a list of all names known to the node.
func (n *NamesController) GetAllNames(
	ctx context.Context,
	page *int) (
	models.ApiResponse[[]string],
	error) {
	req := n.prepareRequest(ctx, "GET", "/v1/names")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	var result []string
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]string](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewBnsError(400, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// GetNameInfo takes context, name as parameters and
// returns an models.ApiResponse with models.BnsGetNameInfoResponse data and
// an error if there was an issue with the request or response.
// Retrieves details of a given name including the `address`, `status` and last transaction id - `last_txid`.
func (n *NamesController) GetNameInfo(
	ctx context.Context,
	name string) (
	models.ApiResponse[models.BnsGetNameInfoResponse],
	error) {
	req := n.prepareRequest(ctx, "GET", fmt.Sprintf("/v1/names/%v", name))
	req.Authenticate(true)

	var result models.BnsGetNameInfoResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BnsGetNameInfoResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewBnsError(400, "Error")
	}
	if resp.StatusCode == 404 {
		err = errors.NewBnsError(404, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// FetchSubdomainsListForName takes context, name as parameters and
// returns an models.ApiResponse with []string data and
// an error if there was an issue with the request or response.
// Retrieves the list of subdomains for a specific name
func (n *NamesController) FetchSubdomainsListForName(
	ctx context.Context,
	name string) (
	models.ApiResponse[[]string],
	error) {
	req := n.prepareRequest(ctx, "GET", fmt.Sprintf("/v1/names/%v/subdomains", name))
	req.Authenticate(true)

	var result []string
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]string](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// FetchZoneFile takes context, name as parameters and
// returns an models.ApiResponse with models.BnsFetchFileZoneResponse2 data and
// an error if there was an issue with the request or response.
// Retrieves a user’s raw zone file. This only works for RFC-compliant zone files. This method returns an error for names that have non-standard zone files.
func (n *NamesController) FetchZoneFile(
	ctx context.Context,
	name string) (
	models.ApiResponse[models.BnsFetchFileZoneResponse2],
	error) {
	req := n.prepareRequest(ctx, "GET", fmt.Sprintf("/v1/names/%v/zonefile", name))
	req.Authenticate(true)

	var result models.BnsFetchFileZoneResponse2
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BnsFetchFileZoneResponse2](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewBnsError(400, "Error")
	}
	if resp.StatusCode == 404 {
		err = errors.NewBnsError(404, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// GetHistoricalZoneFile takes context, name, zoneFileHash as parameters and
// returns an models.ApiResponse with models.BnsFetchHistoricalZoneFileResponse2 data and
// an error if there was an issue with the request or response.
// Retrieves the historical zonefile specified by the username and zone hash.
func (n *NamesController) GetHistoricalZoneFile(
	ctx context.Context,
	name string,
	zoneFileHash string) (
	models.ApiResponse[models.BnsFetchHistoricalZoneFileResponse2],
	error) {
	req := n.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/names/%v/zonefile/%v", name, zoneFileHash),
	)
	req.Authenticate(true)

	var result models.BnsFetchHistoricalZoneFileResponse2
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BnsFetchHistoricalZoneFileResponse2](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewBnsError(400, "Error")
	}
	if resp.StatusCode == 404 {
		err = errors.NewBnsError(404, "Error")
	}
	return models.NewApiResponse(result, resp), err
}

// GetNamesOwnedByAddress takes context, blockchain, address as parameters and
// returns an models.ApiResponse with models.BnsNamesOwnByAddressResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of names owned by the address provided.
func (n *NamesController) GetNamesOwnedByAddress(
	ctx context.Context,
	blockchain string,
	address string) (
	models.ApiResponse[models.BnsNamesOwnByAddressResponse],
	error) {
	req := n.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/addresses/%v/%v", blockchain, address),
	)
	req.Authenticate(true)

	var result models.BnsNamesOwnByAddressResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BnsNamesOwnByAddressResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewBnsError(404, "Error")
	}
	return models.NewApiResponse(result, resp), err
}
