package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"vutest/testutil/network"
	"vutest/testutil/nullify"
	"vutest/x/vutest/client/cli"
	"vutest/x/vutest/types"
)

func networkWithBaseDenomObjects(t *testing.T) (*network.Network, types.BaseDenom) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	baseDenom := &types.BaseDenom{}
	nullify.Fill(&baseDenom)
	state.BaseDenom = baseDenom
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.BaseDenom
}

func TestShowBaseDenom(t *testing.T) {
	net, obj := networkWithBaseDenomObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		args []string
		err  error
		obj  types.BaseDenom
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowBaseDenom(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetBaseDenomResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.BaseDenom)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.BaseDenom),
				)
			}
		})
	}
}
