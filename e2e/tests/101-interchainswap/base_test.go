package e2e

import (
	//"context"
	"context"
	"testing"

	//"time"

	//"github.com/strangelove-ventures/ibctest/broadcast"
	//"github.com/strangelove-ventures/ibctest/chain/cosmos"
	//"github.com/strangelove-ventures/ibctest/ibc"
	//"github.com/strangelove-ventures/ibctest/test"
	//"github.com/stretchr/testify/suite"

	"github.com/cosmos/ibc-go/e2e/testsuite"
	"github.com/cosmos/ibc-go/e2e/testvalues"
	"github.com/strangelove-ventures/ibctest/ibc"
	"github.com/strangelove-ventures/ibctest/test"
	"github.com/stretchr/testify/suite"

	types "github.com/sideprotocol/ibcswap/v4/modules/apps/101-interchain-swap/types"
)

func TestInterchainswapTestSuite(t *testing.T) {
	suite.Run(t, new(InterchainswapTestSuite))
}

type InterchainswapTestSuite struct {
	testsuite.E2ETestSuite
}

func (s *InterchainswapTestSuite) TestMsgCreatePool() {

	t := s.T()
	ctx := context.TODO()

	// setup relayers and connection-0 between two chains.
	relayer, channelA, _ := s.SetupChainsRelayerAndChannel(ctx, interchainswapChannelOptions())

	chainA, chainB := s.GetChains()

	chainADenom := chainA.Config().Denom

	// // create wallets for testing
	chainAWallet := s.CreateUserOnChainA(ctx, testvalues.StartingTokenAmount)
	chainAAddress := chainAWallet.Bech32Address("cosmos")

	// allocate tokens to the new account
	//initialBalances := sdk.NewCoins(sdk.NewCoin("marscoin", sdk.NewInt(1000)))

	//s.Require().NoError(err)

	//chainBAddress = s.CreateUserOnChainB(ctx, testvalues.StartingTokenAmount)

	// s.Require().NoError(test.WaitForBlocks(ctx, 1, chainA, chainB), "failed to wait for blocks")

	t.Run("start relayer", func(t *testing.T) {
		s.StartRelayer(relayer)
	})

	t.Run("send creat pool message", func(t *testing.T) {
		msg := types.NewMsgCreatePool(
			channelA.PortID,
			channelA.ChannelID,
			chainAAddress,
			"1:2",
			[]string{chainADenom, "marscoin"},
			[]uint32{10, 100},
		)

		resp, err := s.BroadcastMessages(ctx, chainA, chainAWallet, msg)
		s.AssertValidTxResponse(resp)
		s.Require().NoError(err)

		s.Require().NoError(test.WaitForBlocks(ctx, 20, chainA, chainB), "failed to wait for blocks")
		s.AssertPacketRelayed(
			ctx,
			chainA,
			types.ModuleName,
			channelA.ChannelID,
			1,
		)

		//poolId := types.GetPoolId([]string{chainADenom, "marscoin"})
		//_, err = s.QueryInterchainswapPool(ctx, chainA, poolId)
		//s.Require().NoError(err)
	})

	// t.Run("send deposit pool message", func(t *testing.T) {
	// 	poolId := types.GetPoolId([]string{chainADenom, "marscoin"})
	// 	msg := types.NewMsgDeposit(
	// 		poolId,
	// 		chainAAddress,
	// 		[]*sdk.Coin{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewInt(1000)}},
	// 	)
	// 	resp, err := s.BroadcastMessages(ctx, chainA, chainAWallet, msg)
	// 	fmt.Println(resp)
	// 	s.AssertValidTxResponse(resp)
	// 	s.Require().NoError(err)
	// })
}

// interchainswapChannelOptions configures both of the chains to have interchainswap enabled.
func interchainswapChannelOptions() func(options *ibc.CreateChannelOptions) {
	return func(opts *ibc.CreateChannelOptions) {
		opts.Version = "ics101-1"
		opts.Order = ibc.Unordered
		opts.DestPortName = types.ModuleName
		opts.SourcePortName = types.ModuleName
	}
}
