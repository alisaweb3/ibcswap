package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgMakeMultiAssetDeposit = "make_multi_asset_deposit"

var _ sdk.Msg = &MsgMakeMultiAssetDepositRequest{}

func NewMsgMakeMultiAssetDeposit(poolId string, senders []string, tokens []*sdk.Coin) *MsgMakeMultiAssetDepositRequest {
	return &MsgMakeMultiAssetDepositRequest{
		PoolId: poolId,
		Deposits: []*DepositAsset{
			{
				Sender:  senders[0],
				Balance: tokens[0],
			},
			{
				Sender:  senders[1],
				Balance: tokens[1],
			},
		},
	}
}

func (msg *MsgMakeMultiAssetDepositRequest) Route() string {
	return RouterKey
}

func (msg *MsgMakeMultiAssetDepositRequest) Type() string {
	return TypeMsgDeposit
}

func (msg *MsgMakeMultiAssetDepositRequest) GetSigners() []sdk.AccAddress {
	signers := []sdk.AccAddress{}

	creator, err := sdk.AccAddressFromBech32(msg.Deposits[0].Sender)
	if err != nil {
		panic(err)
	}
	signers = append(signers, creator)
	return signers
}

func (msg *MsgMakeMultiAssetDepositRequest) GetSignBytes() []byte {
	marshaledMsg := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(marshaledMsg)
}

func (msg *MsgMakeMultiAssetDepositRequest) ValidateBasic() error {
	// Check address
	for _, deposit := range msg.Deposits {
		_, err := sdk.AccAddressFromBech32(deposit.Sender)
		if err != nil {
			return ErrInvalidAddress
		}
		if deposit.Balance.Amount.Equal(sdk.NewInt(0)) {
			return ErrInvalidAmount
		}
	}
	return nil
}
