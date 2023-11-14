package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateBaseDenom = "create_base_denom"
	TypeMsgUpdateBaseDenom = "update_base_denom"
	TypeMsgDeleteBaseDenom = "delete_base_denom"
)

var _ sdk.Msg = &MsgCreateBaseDenom{}

func NewMsgCreateBaseDenom(creator string, denom string) *MsgCreateBaseDenom {
	return &MsgCreateBaseDenom{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgCreateBaseDenom) Route() string {
	return RouterKey
}

func (msg *MsgCreateBaseDenom) Type() string {
	return TypeMsgCreateBaseDenom
}

func (msg *MsgCreateBaseDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBaseDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBaseDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBaseDenom{}

func NewMsgUpdateBaseDenom(creator string, denom string) *MsgUpdateBaseDenom {
	return &MsgUpdateBaseDenom{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgUpdateBaseDenom) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBaseDenom) Type() string {
	return TypeMsgUpdateBaseDenom
}

func (msg *MsgUpdateBaseDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBaseDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBaseDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBaseDenom{}

func NewMsgDeleteBaseDenom(creator string) *MsgDeleteBaseDenom {
	return &MsgDeleteBaseDenom{
		Creator: creator,
	}
}
func (msg *MsgDeleteBaseDenom) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBaseDenom) Type() string {
	return TypeMsgDeleteBaseDenom
}

func (msg *MsgDeleteBaseDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBaseDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBaseDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
