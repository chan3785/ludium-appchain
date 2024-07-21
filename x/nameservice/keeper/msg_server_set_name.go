package keeper

import (
	"context"

	"github.com/Jeongseup/ludiumapp/x/nameservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetName(goCtx context.Context, msg *types.MsgSetName) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetNameResponse{}, nil
}
