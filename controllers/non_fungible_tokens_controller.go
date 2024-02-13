package controllers

import (
	"context"
	"fmt"

	"github.com/cbadawi/go-stacks/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// NonFungibleTokensController represents a controller struct.
type NonFungibleTokensController struct {
	baseController
}

// NewNonFungibleTokensController creates a new instance of NonFungibleTokensController.
// It takes a baseController as a parameter and returns a pointer to the NonFungibleTokensController.
func NewNonFungibleTokensController(baseController baseController) *NonFungibleTokensController {
	nonFungibleTokensController := NonFungibleTokensController{baseController: baseController}
	return &nonFungibleTokensController
}

// GetNftHoldings takes context, principal, assetIdentifiers, limit, offset, unanchored, txMetadata as parameters and
// returns an models.ApiResponse with models.NonFungibleTokenHoldingsList data and
// an error if there was an issue with the request or response.
// Retrieves the list of Non-Fungible Tokens owned by the given principal (STX address or Smart Contract ID).
// Results can be filtered by one or more asset identifiers and can include metadata about the transaction that made the principal own each token.
// More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).
func (n *NonFungibleTokensController) GetNftHoldings(
	ctx context.Context,
	principal string,
	assetIdentifiers []string,
	limit *int,
	offset *int,
	unanchored *bool,
	txMetadata *bool) (
	models.ApiResponse[models.NonFungibleTokenHoldingsList],
	error) {
	path := fmt.Sprintf("/extended/v1/tokens/nft/holdings")
	req := n.prepareRequest(ctx, "GET", path)
	req.Authenticate(true)
	req.QueryParam("principal", principal)
	if assetIdentifiers != nil {
		req.QueryParam("asset_identifiers", assetIdentifiers)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if txMetadata != nil {
		req.QueryParam("tx_metadata", *txMetadata)
	}
	var result models.NonFungibleTokenHoldingsList
	decoder, resp, err := n.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NonFungibleTokenHoldingsList](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetNftHistory takes context, assetIdentifier, value, limit, offset, unanchored, txMetadata as parameters and
// returns an models.ApiResponse with models.NonFungibleTokenHistoryEventList data and
// an error if there was an issue with the request or response.
// Retrieves all events relevant to a Non-Fungible Token. Useful to determine the ownership history of a particular asset.
// More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).
func (n *NonFungibleTokensController) GetNftHistory(
	ctx context.Context,
	assetIdentifier string,
	value string,
	limit *int,
	offset *int,
	unanchored *bool,
	txMetadata *bool) (
	models.ApiResponse[models.NonFungibleTokenHistoryEventList],
	error) {
	path := fmt.Sprintf("/extended/v1/tokens/nft/history")
	req := n.prepareRequest(ctx, "GET", path)
	req.Authenticate(true)
	req.QueryParam("asset_identifier", assetIdentifier)
	req.QueryParam("value", value)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if txMetadata != nil {
		req.QueryParam("tx_metadata", *txMetadata)
	}
	var result models.NonFungibleTokenHistoryEventList
	decoder, resp, err := n.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NonFungibleTokenHistoryEventList](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetNftMints takes context, assetIdentifier, limit, offset, unanchored, txMetadata as parameters and
// returns an models.ApiResponse with models.NonFungibleTokenMintList data and
// an error if there was an issue with the request or response.
// Retrieves all mint events for a Non-Fungible Token asset class. Useful to determine which NFTs of a particular collection have been claimed.
// More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).
func (n *NonFungibleTokensController) GetNftMints(
	ctx context.Context,
	assetIdentifier string,
	limit *int,
	offset *int,
	unanchored *bool,
	txMetadata *bool) (
	models.ApiResponse[models.NonFungibleTokenMintList],
	error) {
	path := fmt.Sprintf("/extended/v1/tokens/nft/mints")
	req := n.prepareRequest(ctx, "GET", path)
	req.Authenticate(true)
	req.QueryParam("asset_identifier", assetIdentifier)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if txMetadata != nil {
		req.QueryParam("tx_metadata", *txMetadata)
	}
	var result models.NonFungibleTokenMintList
	decoder, resp, err := n.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NonFungibleTokenMintList](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
