package controllers

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/cbadawi/go-stacks/stacks/errors"

	"github.com/cbadawi/go-stacks/stacks/models"
)

// AccountsController represents a controller struct.
type AccountsController struct {
	baseController
}

// NewAccountsController creates a new instance of AccountsController.
// It takes a baseController as a parameter and returns a pointer to the AccountsController.
func NewAccountsController(baseController baseController) *AccountsController {
	accountsController := AccountsController{baseController: baseController}
	return &accountsController
}

// GetAccountBalance takes context, principal, unanchored, untilBlock as parameters and
// returns an models.ApiResponse with models.AddressBalanceResponse data and
// an error if there was an issue with the request or response.
// Retrieves total account balance information for a given Address or Contract Identifier. This includes the balances of  STX Tokens, Fungible Tokens and Non-Fungible Tokens for the account.
func (a *AccountsController) GetAccountBalance(
	ctx context.Context,
	principal string,
	unanchored *bool,
	untilBlock *string) (
	models.ApiResponse[models.AddressBalanceResponse],
	error) {
	path := fmt.Sprintf("/extended/v1/address/%v/balances", principal)
	req := a.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Header("Content-Type", "application/json")
	req.Authenticate(true)

	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if untilBlock != nil {
		req.QueryParam("until_block", *untilBlock)
	}

	var result models.AddressBalanceResponse
	decoder, resp, err := a.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AddressBalanceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAccountStxBalance takes context, principal, unanchored, untilBlock as parameters and
// returns an models.ApiResponse with models.AddressStxBalanceResponse data and
// an error if there was an issue with the request or response.
// Retrieves STX token balance for a given Address or Contract Identifier.
func (a *AccountsController) GetAccountStxBalance(
	ctx context.Context,
	principal string,
	unanchored *bool,
	untilBlock *string) (
	models.ApiResponse[models.AddressStxBalanceResponse],
	error) {
	path := fmt.Sprintf("/extended/v1/address/%v/stx", principal)
	req := a.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if untilBlock != nil {
		req.QueryParam("until_block", *untilBlock)
	}

	var result models.AddressStxBalanceResponse
	decoder, resp, err := a.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AddressStxBalanceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAccountTransactions takes context, principal, limit, offset, height, unanchored, untilBlock as parameters and
// returns an models.ApiResponse with models.AddressTransactionsListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of all Transactions for a given Address or Contract Identifier. More information on Transaction types can be found [here](https://docs.stacks.co/understand-stacks/transactions#types).
// If you need to actively monitor new transactions for an address or contract id, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.
func (a *AccountsController) GetAccountTransactions(
	ctx context.Context,
	principal string,
	limit *int,
	offset *int,
	height *float64,
	unanchored *bool,
	untilBlock *string) (
	models.ApiResponse[models.AddressTransactionsListResponse],
	error) {
	path := fmt.Sprintf("/extended/v1/address/%v/transactions", principal)
	req := a.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if height != nil {
		req.QueryParam("height", *height)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if untilBlock != nil {
		req.QueryParam("until_block", *untilBlock)
	}

	var result models.AddressTransactionsListResponse
	decoder, resp, err := a.LogCallAsJSON(req, path)

	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AddressTransactionsListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetSingleTransactionWithTransfers takes context, principal, txId as parameters and
// returns an models.ApiResponse with models.AddressTransactionWithTransfers data and
// an error if there was an issue with the request or response.
// Retrieves transaction details for a given Transaction Id `tx_id`, for a given account or contract Identifier.
func (a *AccountsController) GetSingleTransactionWithTransfers(
	ctx context.Context,
	principal string,
	txId string) (
	models.ApiResponse[models.AddressTransactionWithTransfers],
	error) {
	path := fmt.Sprintf("/extended/v1/address/%v/%v/with_transfers", principal, txId)
	req := a.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)

	var result models.AddressTransactionWithTransfers
	decoder, resp, err := a.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AddressTransactionWithTransfers](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not found")
	}
	return models.NewApiResponse(result, resp), err
}

// GetAccountTransactionsWithTransfers takes context, principal, limit, offset, height, unanchored, untilBlock as parameters and
// returns an models.ApiResponse with models.AddressTransactionsWithTransfersListResponse data and
// an error if there was an issue with the request or response.
// Retrieve all transactions for an account or contract identifier including STX transfers for each transaction.
func (a *AccountsController) GetAccountTransactionsWithTransfers(
	ctx context.Context,
	principal string,
	limit *int,
	offset *int,
	height *float64,
	unanchored *bool,
	untilBlock *string) (
	models.ApiResponse[models.AddressTransactionsWithTransfersListResponse],
	error) {
	path := fmt.Sprintf("/extended/v1/address/%v/transactions_with_transfers", principal)
	req := a.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if height != nil {
		req.QueryParam("height", *height)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if untilBlock != nil {
		req.QueryParam("until_block", *untilBlock)
	}

	var result models.AddressTransactionsWithTransfersListResponse
	decoder, resp, err := a.LogCallAsJSON(req, path)

	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AddressTransactionsWithTransfersListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAccountNonces takes context, principal, blockHeight, blockHash as parameters and
// returns an models.ApiResponse with models.AddressNonces data and
// an error if there was an issue with the request or response.
// Retrieves the latest nonce values used by an account by inspecting the mempool, microblock transactions, and anchored transactions.
func (a *AccountsController) GetAccountNonces(
	ctx context.Context,
	principal string,
	blockHeight *float64,
	blockHash *string) (
	models.ApiResponse[models.AddressNonces],
	error) {
	path := fmt.Sprintf("/extended/v1/address/%v/nonces", principal)
	req := a.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if blockHeight != nil {
		req.QueryParam("block_height", *blockHeight)
	}
	if blockHash != nil {
		req.QueryParam("block_hash", *blockHash)
	}

	var result models.AddressNonces
	decoder, resp, err := a.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AddressNonces](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAccountAssets takes context, principal, limit, offset, unanchored, untilBlock as parameters and
// returns an models.ApiResponse with models.AddressAssetsListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of all assets events associated with an account or a Contract Identifier. This includes Transfers, Mints.
func (a *AccountsController) GetAccountAssets(
	ctx context.Context,
	principal string,
	limit *int,
	offset *int,
	unanchored *bool,
	untilBlock *string) (
	models.ApiResponse[models.AddressAssetsListResponse],
	error) {
	path := fmt.Sprintf("/extended/v1/address/%v/assets", principal)
	req := a.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if untilBlock != nil {
		req.QueryParam("until_block", *untilBlock)
	}

	var result models.AddressAssetsListResponse
	decoder, resp, err := a.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AddressAssetsListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAccountInbound takes context, principal, limit, offset, height, unanchored, untilBlock as parameters and
// returns an models.ApiResponse with models.AddressStxInboundListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of STX transfers with memos to the given principal. This includes regular transfers from a stx-transfer transaction type,
// and transfers from contract-call transactions a the `send-many-memo` bulk sending contract.
func (a *AccountsController) GetAccountInbound(
	ctx context.Context,
	principal string,
	limit *int,
	offset *int,
	height *float64,
	unanchored *bool,
	untilBlock *string) (
	models.ApiResponse[models.AddressStxInboundListResponse],
	error) {
	path := fmt.Sprintf("/extended/v1/address/%v/stx_inbound", principal)
	req := a.prepareRequest(
		ctx,
		"GET",
		path,
	)
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if height != nil {
		req.QueryParam("height", *height)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	if untilBlock != nil {
		req.QueryParam("until_block", *untilBlock)
	}

	var result models.AddressStxInboundListResponse
	decoder, resp, err := a.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AddressStxInboundListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAccountInfo takes context, principal, proof, tip as parameters and
// returns an models.ApiResponse with models.AccountDataResponse data and
// an error if there was an issue with the request or response.
// Retrieves the account data for a given Account or a Contract Identifier
// Where balance is the hex encoding of a unsigned 128-bit integer (big-endian), nonce is an unsigned 64-bit integer, and the proofs are provided as hex strings.
// For non-existent accounts, this does not return a 404 error, rather it returns an object with balance and nonce of 0.
func (a *AccountsController) GetAccountInfo(
	ctx context.Context,
	principal string,
	proof *int,
	tip *string) (
	models.ApiResponse[models.AccountDataResponse],
	error) {
	path := fmt.Sprintf("/v2/accounts/%v", principal)
	req := a.prepareRequest(ctx, "GET", path)
	req.Authenticate(true)
	if proof != nil {
		req.QueryParam("proof", *proof)
	}
	if tip != nil {
		req.QueryParam("tip", *tip)
	}

	var result models.AccountDataResponse
	decoder, resp, err := a.LogCallAsJSON(req, path)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountDataResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
