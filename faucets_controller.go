package stacksblockchainapi

import (
	"context"
	"stacksblockchainapi/errors"
	"stacksblockchainapi/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// FaucetsController represents a controller struct.
type FaucetsController struct {
	baseController
}

// NewFaucetsController creates a new instance of FaucetsController.
// It takes a baseController as a parameter and returns a pointer to the FaucetsController.
func NewFaucetsController(baseController baseController) *FaucetsController {
	faucetsController := FaucetsController{baseController: baseController}
	return &faucetsController
}

// RunFaucetStx takes context, address, stacking as parameters and
// returns an models.ApiResponse with models.RunFaucetResponse data and
// an error if there was an issue with the request or response.
// Add 500 STX tokens to the specified testnet address. Testnet STX addresses begin with `ST`. If the `stacking`
// parameter is set to `true`, the faucet will add the required number of tokens for individual stacking to the
// specified testnet address.
// The endpoint returns the transaction ID, which you can use to view the transaction in the
// [Stacks Explorer](https://explorer.hiro.so/?chain=testnet). The tokens are delivered once the transaction has
// been included in an anchor block.
// A common reason for failed faucet transactions is that the faucet has run out of tokens. If you are experiencing
// failed faucet transactions to a testnet address, you can get help in [Discord](https://stacks.chat).
// **Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.
func (f *FaucetsController) RunFaucetStx(
	ctx context.Context,
	address string,
	stacking *bool) (
	models.ApiResponse[models.RunFaucetResponse],
	error) {
	req := f.prepareRequest(ctx, "POST", "/extended/v1/faucets/stx")
	req.Authenticate(true)
	req.QueryParam("address", address)
	if stacking != nil {
		req.QueryParam("stacking", *stacking)
	}
	var result models.RunFaucetResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RunFaucetResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 500 {
		err = errors.NewApiError(500, "Faucet call failed")
	}
	return models.NewApiResponse(result, resp), err
}

// RunFaucetBtc takes context, address, body as parameters and
// returns an models.ApiResponse with models.RunFaucetResponse data and
// an error if there was an issue with the request or response.
// Add 1 BTC token to the specified testnet BTC address.
// The endpoint returns the transaction ID, which you can use to view the transaction in a testnet Bitcoin block
// explorer. The tokens are delivered once the transaction has been included in a block.
// **Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.
func (f *FaucetsController) RunFaucetBtc(
	ctx context.Context,
	address string,
	body *models.ExtendedV1FaucetsBtcRequest) (
	models.ApiResponse[models.RunFaucetResponse],
	error) {
	req := f.prepareRequest(ctx, "POST", "/extended/v1/faucets/btc")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	req.QueryParam("address", address)
	if body != nil {
		req.Json(*body)
	}
	var result models.RunFaucetResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RunFaucetResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 403 {
		err = errors.NewExtendedV1FaucetsBtc403Error(403, "BTC Faucet not fully configured")
	}
	if resp.StatusCode == 500 {
		err = errors.NewApiError(500, "Faucet call failed")
	}
	return models.NewApiResponse(result, resp), err
}
