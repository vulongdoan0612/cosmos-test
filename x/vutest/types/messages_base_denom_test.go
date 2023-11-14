package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"vutest/testutil/sample"
)

func TestMsgCreateBaseDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateBaseDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateBaseDenom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateBaseDenom{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateBaseDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateBaseDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateBaseDenom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateBaseDenom{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteBaseDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteBaseDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteBaseDenom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteBaseDenom{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
