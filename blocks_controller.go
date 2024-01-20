package main

import (
	"context"
	"fmt"

	"github.com/cbadawi/stacks-go-draft/errors"

	"github.com/cbadawi/stacks-go-draft/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// BlocksController represents a controller struct.
type BlocksController struct {
	baseController
}

// NewBlocksController creates a new instance of BlocksController.
// It takes a baseController as a parameter and returns a pointer to the BlocksController.
func NewBlocksController(baseController baseController) *BlocksController {
	blocksController := BlocksController{baseController: baseController}
	return &blocksController
}

// GetBlocksByBurnBlock takes context, heightOrHash, limit, offset as parameters and
// returns an models.ApiResponse with models.NakamotoBlockListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of blocks confirmed by a specific burn block
func (b *BlocksController) GetBlocksByBurnBlock(
	ctx context.Context,
	heightOrHash interface{},
	limit *int,
	offset *int) (
	models.ApiResponse[models.NakamotoBlockListResponse],
	error) {
	req := b.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v2/burn-blocks/%v/blocks", heightOrHash),
	)
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}

	var result models.NakamotoBlockListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NakamotoBlockListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBlocks takes context, limit, offset as parameters and
// returns an models.ApiResponse with models.NakamotoBlockListResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of recently mined blocks
func (b *BlocksController) GetBlocks(
	ctx context.Context,
	limit *int,
	offset *int) (
	models.ApiResponse[models.NakamotoBlockListResponse],
	error) {
	req := b.prepareRequest(ctx, "GET", "/extended/v2/blocks")
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	var result models.NakamotoBlockListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NakamotoBlockListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBlock takes context, heightOrHash as parameters and
// returns an models.ApiResponse with models.NakamotoBlock data and
// an error if there was an issue with the request or response.
// Retrieves a single block
func (b *BlocksController) GetBlock(
	ctx context.Context,
	heightOrHash interface{}) (
	models.ApiResponse[models.NakamotoBlock],
	error) {
	req := b.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v2/blocks/%v", heightOrHash),
	)
	req.Authenticate(true)

	var result models.NakamotoBlock
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NakamotoBlock](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBlockList takes context, limit, offset as parameters and
// returns an models.ApiResponse with models.BlockListResponse data and
// an error if there was an issue with the request or response.
// Deprecated: get_block_list is deprecated
// **NOTE:** This endpoint is deprecated in favor of [Get blocks](#operation/get_blocks).
// Retrieves a list of recently mined blocks
// If you need to actively monitor new blocks, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.
func (b *BlocksController) GetBlockList(
	ctx context.Context,
	limit *int,
	offset *int) (
	models.ApiResponse[models.BlockListResponse],
	error) {
	req := b.prepareRequest(ctx, "GET", "/extended/v1/block")
	req.Authenticate(true)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if offset != nil {
		req.QueryParam("offset", *offset)
	}
	var result models.BlockListResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BlockListResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBlockByHash takes context, hash as parameters and
// returns an models.ApiResponse with models.Block data and
// an error if there was an issue with the request or response.
// Deprecated: get_block_by_hash is deprecated
// **NOTE:** This endpoint is deprecated in favor of [Get block](#operation/get_block).
// Retrieves block details of a specific block for a given chain height. You can use the hash from your latest block ('get_block_list' API) to get your block details.
func (b *BlocksController) GetBlockByHash(
	ctx context.Context,
	hash string) (
	models.ApiResponse[models.Block],
	error) {
	req := b.prepareRequest(ctx, "GET", fmt.Sprintf("/extended/v1/block/%v", hash))
	req.Authenticate(true)

	var result models.Block
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Block](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Cannot find block with given ID")
	}
	return models.NewApiResponse(result, resp), err
}

// GetBlockByHeight takes context, height as parameters and
// returns an models.ApiResponse with models.Block data and
// an error if there was an issue with the request or response.
// Deprecated: get_block_by_height is deprecated
// **NOTE:** This endpoint is deprecated in favor of [Get block](#operation/get_block).
// Retrieves block details of a specific block at a given block height
func (b *BlocksController) GetBlockByHeight(
	ctx context.Context,
	height float64) (
	models.ApiResponse[models.Block],
	error) {
	req := b.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/block/by_height/%v", height),
	)
	req.Authenticate(true)

	var result models.Block
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Block](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Cannot find block with given height")
	}
	return models.NewApiResponse(result, resp), err
}

// GetBlockByBurnBlockHash takes context, burnBlockHash as parameters and
// returns an models.ApiResponse with models.Block data and
// an error if there was an issue with the request or response.
// Deprecated: get_block_by_burn_block_hash is deprecated
// **NOTE:** This endpoint is deprecated in favor of [Get blocks](#operation/get_blocks).
// Retrieves block details of a specific block for a given burnchain block hash
func (b *BlocksController) GetBlockByBurnBlockHash(
	ctx context.Context,
	burnBlockHash string) (
	models.ApiResponse[models.Block],
	error) {
	req := b.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/block/by_burn_block_hash/%v", burnBlockHash),
	)
	req.Authenticate(true)

	var result models.Block
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Block](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Cannot find block with given height")
	}
	return models.NewApiResponse(result, resp), err
}

// GetBlockByBurnBlockHeight takes context, burnBlockHeight as parameters and
// returns an models.ApiResponse with models.Block data and
// an error if there was an issue with the request or response.
// Deprecated: get_block_by_burn_block_height is deprecated
// **NOTE:** This endpoint is deprecated in favor of [Get blocks](#operation/get_blocks).
// Retrieves block details of a specific block for a given burn chain height
func (b *BlocksController) GetBlockByBurnBlockHeight(
	ctx context.Context,
	burnBlockHeight float64) (
	models.ApiResponse[models.Block],
	error) {
	req := b.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/extended/v1/block/by_burn_block_height/%v", burnBlockHeight),
	)
	req.Authenticate(true)

	var result models.Block
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Block](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Cannot find block with given height")
	}
	return models.NewApiResponse(result, resp), err
}
