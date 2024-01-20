package stacksblockchainapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cbadawi/stacks-go-draft/errors"

	"github.com/cbadawi/stacks-go-draft/models"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
)

// TransactionsController represents a controller struct.
type TransactionsController struct {
	baseController
}

// NewTransactionsController creates a new instance of TransactionsController.
// It takes a baseController as a parameter and returns a pointer to the TransactionsController.
func NewTransactionsController(baseController baseController) *TransactionsController {
	transactionsController := TransactionsController{baseController: baseController}
	return &transactionsController
}

// GetTransactionList takes context, limit, offset, mType, unanchored as parameters and
// returns an models.ApiResponse with models.TransactionResults data and
// an error if there was an issue with the request or response.
// Retrieves all recently mined transactions
// If using TypeScript, import typings for this response from our types package:
// `import type { TransactionResults } from '@stacks/stacks-blockchain-api-types';`
func (t *TransactionsController) GetTransactionList(
	ctx context.Context,
	limit *int,
	offset *int,
	mType []models.TypeEnum,
	unanchored *bool) (
	models.ApiResponse[models.TransactionResults],
	error) {
	req := t.prepareRequest(ctx, "GET", "/extended/v1/tx")
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if mType != nil {
		req.QueryParam("type", mType)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	var result models.TransactionResults
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.TransactionResults](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetMempoolTransactionList takes context, senderAddress, recipientAddress, address, orderBy, order, limit, offset, unanchored as parameters and
// returns an models.ApiResponse with models.MempoolTransactionListResponse data and
// an error if there was an issue with the request or response.
// Retrieves all transactions that have been recently broadcast to the mempool. These are pending transactions awaiting confirmation.
// If you need to monitor new transactions, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.
func (t *TransactionsController) GetMempoolTransactionList(
	ctx context.Context,
	senderAddress *string,
	recipientAddress *string,
	address *string,
	orderBy *models.OrderByEnum,
	order *models.OrderEnum,
	limit *int,
	offset *int,
	unanchored *bool) (
	models.ApiResponse[models.MempoolTransactionListResponse],
	error) {
	req := t.prepareRequest(ctx, "GET", "/extended/v1/tx/mempool")
	req.Authenticate(true)
	if senderAddress != nil {
		req.QueryParam("sender_address", *senderAddress)
	}
	if recipientAddress != nil {
		req.QueryParam("recipient_address", *recipientAddress)
	}
	if address != nil {
		req.QueryParam("address", *address)
	}
	if orderBy != nil {
		req.QueryParam("order_by", *orderBy)
	}
	if order != nil {
		req.QueryParam("order", *order)
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
	var result models.MempoolTransactionListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MempoolTransactionListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetDroppedMempoolTransactionList takes context, limit, offset as parameters and
// returns an models.ApiResponse with models.MempoolTransactionListResponse data and
// an error if there was an issue with the request or response.
// Retrieves all recently-broadcast transactions that have been dropped from the mempool.
// Transactions are dropped from the mempool if:
// * they were stale and awaiting garbage collection or,
// * were expensive,  or
// * were replaced with a new fee
func (t *TransactionsController) GetDroppedMempoolTransactionList(
	ctx context.Context,
	limit *int,
	offset *int) (
	models.ApiResponse[models.MempoolTransactionListResponse],
	error) {
	req := t.prepareRequest(ctx, "GET", "/extended/v1/tx/mempool/dropped")
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	var result models.MempoolTransactionListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MempoolTransactionListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetMempoolTransactionStats takes context as parameters and
// returns an models.ApiResponse with models.MempoolTransactionStatsResponse data and
// an error if there was an issue with the request or response.
// Queries for transactions counts, age (by block height), fees (simple average), and size.
// All results broken down by transaction type and percentiles (p25, p50, p75, p95).
func (t *TransactionsController) GetMempoolTransactionStats(ctx context.Context) (
	models.ApiResponse[models.MempoolTransactionStatsResponse],
	error) {
	req := t.prepareRequest(ctx, "GET", "/extended/v1/tx/mempool/stats")
	req.Authenticate(true)
	var result models.MempoolTransactionStatsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MempoolTransactionStatsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetTxListDetails takes context, txId, eventOffset, eventLimit, unanchored as parameters and
// returns an models.ApiResponse with map[string]models.TransactionList data and
// an error if there was an issue with the request or response.
// Retrieves a list of transactions for a given list of transaction IDs
// If using TypeScript, import typings for this response from our types package:
// `import type { Transaction } from '@stacks/stacks-blockchain-api-types';`
func (t *TransactionsController) GetTxListDetails(
	ctx context.Context,
	txId []string,
	eventOffset *int,
	eventLimit *int,
	unanchored *bool) (
	models.ApiResponse[map[string]models.TransactionList],
	error) {
	req := t.prepareRequest(ctx, "GET", "/extended/v1/tx/multiple")
	req.Authenticate(true)
	req.QueryParam("tx_id", txId)
	if eventOffset != nil {
		req.QueryParam("event_offset", *eventOffset)
	}
	if eventLimit != nil {
		req.QueryParam("event_limit", *eventLimit)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}
	var result map[string]models.TransactionList
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[map[string]models.TransactionList](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Could not find any transaction by ID")
	}
	return models.NewApiResponse(result, resp), err
}

// GetTransactionById takes context, txId, eventOffset, eventLimit, unanchored as parameters and
// returns an models.ApiResponse with models.Transaction data and
// an error if there was an issue with the request or response.
// Retrieves transaction details for a given transaction ID
// `import type { Transaction } from '@stacks/stacks-blockchain-api-types';`
func (t *TransactionsController) GetTransactionById(
	ctx context.Context,
	txId string,
	eventOffset *int,
	eventLimit *int,
	unanchored *bool) (
	models.ApiResponse[models.Transaction],
	error) {
	req := t.prepareRequest(ctx, "GET", fmt.Sprintf("/extended/v1/tx/%v", txId))
	req.Authenticate(true)
	if eventOffset != nil {
		req.QueryParam("event_offset", *eventOffset)
	}
	if eventLimit != nil {
		req.QueryParam("event_limit", *eventLimit)
	}
	if unanchored != nil {
		req.QueryParam("unanchored", *unanchored)
	}

	var result models.Transaction
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Transaction](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Cannot find transaction for given ID")
	}
	return models.NewApiResponse(result, resp), err
}

// GetRawTransactionById takes context, txId as parameters and
// returns an models.ApiResponse with models.GetRawTransactionResult data and
// an error if there was an issue with the request or response.
// Retrieves a hex encoded serialized transaction for a given ID
func (t *TransactionsController) GetRawTransactionById(
	ctx context.Context,
	txId string) (
	models.ApiResponse[models.GetRawTransactionResult],
	error) {
	req := t.prepareRequest(ctx, "GET", fmt.Sprintf("/extended/v1/tx/%v/raw", txId))
	req.Authenticate(true)

	var result models.GetRawTransactionResult
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRawTransactionResult](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Cannot find transaction for given ID")
	}
	return models.NewApiResponse(result, resp), err
}

// PostCoreNodeTransactions takes context, body as parameters and
// returns an models.ApiResponse with string data and
// an error if there was an issue with the request or response.
// Broadcasts raw transactions on the network. You can use the [@stacks/transactions](https://github.com/blockstack/stacks.js) project to generate a raw transaction payload.
func (t *TransactionsController) PostCoreNodeTransactions(
	ctx context.Context,
	body *models.FileWrapper) (
	models.ApiResponse[string],
	error) {
	req := t.prepareRequest(ctx, "POST", "/v2/transactions")
	req.Authenticate(true)
	req.Header("Content-Type", "application/octet-stream")
	formFields := []https.FormParam{}
	if body != nil {
		bodyParam := https.FormParam{Key: "body", Value: *body, Headers: http.Header{}}
		formFields = append(formFields, bodyParam)
	}
	req.FormData(formFields)
	str, resp, err := req.CallAsText()
	var result string = str

	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	if resp.StatusCode == 400 {
		err = errors.NewPostCoreNodeTransactionsError(400, "Rejections result in a 400 error")
	}
	return models.NewApiResponse(result, resp), err
}

// GetTransactionsByBlock takes context, heightOrHash as parameters and
// returns an models.ApiResponse with models.TransactionResults data and
// an error if there was an issue with the request or response.
// Retrieves transactions confirmed in a single block
func (t *TransactionsController) GetTransactionsByBlock(
	ctx context.Context,
	heightOrHash interface{}) (
	models.ApiResponse[models.TransactionResults],
	error) {
	req := t.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v2/blocks/%v/transactions", heightOrHash),
	)
	req.Authenticate(true)

	var result models.TransactionResults
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.TransactionResults](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetTransactionsByBlockHash takes context, blockHash, limit, offset as parameters and
// returns an models.ApiResponse with models.TransactionResults data and
// an error if there was an issue with the request or response.
// Deprecated: get_transactions_by_block_hash is deprecated
// **NOTE:** This endpoint is deprecated in favor of [Get transactions by block](#operation/get_transactions_by_block).
// Retrieves a list of all transactions within a block for a given block hash.
func (t *TransactionsController) GetTransactionsByBlockHash(
	ctx context.Context,
	blockHash string,
	limit *int,
	offset *int) (
	models.ApiResponse[models.TransactionResults],
	error) {
	req := t.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/tx/block/%v", blockHash),
	)
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}

	var result models.TransactionResults
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.TransactionResults](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetTransactionsByBlockHeight takes context, height, limit, offset, unanchored as parameters and
// returns an models.ApiResponse with models.TransactionResults data and
// an error if there was an issue with the request or response.
// Deprecated: get_transactions_by_block_height is deprecated
// **NOTE:** This endpoint is deprecated in favor of [Get transactions by block](#operation/get_transactions_by_block).
// Retrieves all transactions within a block at a given height
func (t *TransactionsController) GetTransactionsByBlockHeight(
	ctx context.Context,
	height int,
	limit *int,
	offset *int,
	unanchored *bool) (
	models.ApiResponse[models.TransactionResults],
	error) {
	req := t.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/tx/block_height/%v", height),
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

	var result models.TransactionResults
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.TransactionResults](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAddressMempoolTransactions takes context, address, limit, offset, unanchored as parameters and
// returns an models.ApiResponse with models.MempoolTransactionListResponse data and
// an error if there was an issue with the request or response.
// Retrieves all transactions for a given address that are currently in mempool
func (t *TransactionsController) GetAddressMempoolTransactions(
	ctx context.Context,
	address string,
	limit *int,
	offset *int,
	unanchored *bool) (
	models.ApiResponse[models.MempoolTransactionListResponse],
	error) {
	req := t.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/address/%v/mempool", address),
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

	var result models.MempoolTransactionListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MempoolTransactionListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetFilteredEvents takes context, txId, address, limit, offset, mType as parameters and
// returns an models.ApiResponse with models.TransactionEventsResponse data and
// an error if there was an issue with the request or response.
// Retrieves the list of events filtered by principal (STX address or Smart Contract ID), transaction id or event types. The list of event types is ('smart_contract_log', 'stx_lock', 'stx_asset', 'fungible_token_asset', 'non_fungible_token_asset').
func (t *TransactionsController) GetFilteredEvents(
	ctx context.Context,
	txId *string,
	address *string,
	limit *int,
	offset *int,
	mType []models.Type2Enum) (
	models.ApiResponse[models.TransactionEventsResponse],
	error) {
	req := t.prepareRequest(ctx, "GET", "/extended/v1/tx/events")
	req.Authenticate(true)
	if txId != nil {
		req.QueryParam("tx_id", *txId)
	}
	if address != nil {
		req.QueryParam("address", *address)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	if mType != nil {
		req.QueryParam("type", mType)
	}
	var result models.TransactionEventsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.TransactionEventsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
