package keeper

import (
	"context"

	"github.com/Jeongseup/ludiumapp/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) WhoisAll(goCtx context.Context, req *types.QueryAllWhoisRequest) (*types.QueryAllWhoisResponse, error) {
	// if req == nil {
	// 	return nil, status.Error(codes.InvalidArgument, "invalid request")
	// }

	// var whoiss []types.Whois
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// store := ctx.KVStore(k.storeKey)
	// whoisStore := prefix.NewStore(store, types.KeyPrefix(types.WhoisKeyPrefix))

	// pageRes, err := query.Paginate(whoisStore, req.Pagination, func(key []byte, value []byte) error {
	// 	var whois types.Whois
	// 	if err := k.cdc.Unmarshal(value, &whois); err != nil {
	// 		return err
	// 	}

	// 	whoiss = append(whoiss, whois)
	// 	return nil
	// })

	// if err != nil {
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }

	return &types.QueryAllWhoisResponse{Whois: []types.Whois{}, Pagination: &query.PageResponse{}}, nil
}

func (k Keeper) Whois(goCtx context.Context, req *types.QueryGetWhoisRequest) (*types.QueryGetWhoisResponse, error) {
	// if req == nil {
	// 	return nil, status.Error(codes.InvalidArgument, "invalid request")
	// }
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// val, found := k.GetWhois(
	// 	ctx,
	// 	req.Index,
	// )
	// if !found {
	// 	return nil, status.Error(codes.NotFound, "not found")
	// }

	return &types.QueryGetWhoisResponse{Whois: types.Whois{}}, nil
}
