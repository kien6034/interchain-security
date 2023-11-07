package integration

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	consumertypes "github.com/cosmos/interchain-security/v3/x/ccv/consumer/types"
)

// Tests the tracking of historical info in the context of new blocks being committed
func (k CCVTestSuite) TestHistoricalInfo() { //nolint:govet // this is a test so we can copy locks
	fmt.Println("Normal Operations:: TestHistoricalInfo Init")
	consumerKeeper := k.consumerApp.GetConsumerKeeper()
	cCtx := k.consumerChain.GetContext

	// save init consumer valset length
	// initValsetLen := len(consumerKeeper.GetAllCCValidator(cCtx()))
	// save current block height
	initHeight := cCtx().BlockHeight()
	fmt.Println("Normal Operations::  AFTER SETUP ###")

	// define an utility function that creates a new cross-chain validator
	// and then call track historical info in the next block
	createVal := func(k CCVTestSuite) { //nolint:govet // this is a test so we can copy locks
		// add new validator to consumer states
		pk := ed25519.GenPrivKey().PubKey()
		cVal, err := consumertypes.NewCCValidator(pk.Address(), int64(1), pk)
		k.Require().NoError(err)

		consumerKeeper.SetCCValidator(k.consumerChain.GetContext(), cVal)

		// commit block in order to call TrackHistoricalInfo
		k.consumerChain.NextBlock()
	}

	// testsetup create 2 validators and then call track historical info with header block height
	// increased by HistoricalEntries in order to prune the historical info less or equal to the current block height
	// Note that historical info containing the created validators are stored during the next block BeginBlocker
	// and thus are indexed with the respective block heights InitHeight+1 and InitHeight+2
	testSetup := []func(CCVTestSuite){
		createVal,
		createVal,
		func(k CCVTestSuite) { //nolint:govet // this is a test so we can copy locks
			// historicalEntries := k.consumerApp.GetConsumerKeeper().GetHistoricalEntries(k.consumerCtx())
			newHeight := k.consumerChain.GetContext().BlockHeight()
			header := tmproto.Header{
				ChainID: "HelloChain",
				Height:  newHeight,
			}
			ctx := k.consumerChain.GetContext().WithBlockHeader(header)
			consumerKeeper.TrackHistoricalInfo(ctx)
		},
	}

	for _, ts := range testSetup {
		ts(k) //nolint:govet // this is a test so we can copy locks
	}
	fmt.Println("Normal Operations:: PREP CASES ###\n\n\n")

	// test cases verify that historical info entries are pruned when their height
	// is below CurrentHeight - HistoricalEntries, and check that their valset gets updated
	testCases := []struct {
		height int64
		err    error
		expLen int
	}{
		{
			height: initHeight + 1,
			err:    nil,
			expLen: 0,
		},
		// {
		// 	height: initHeight + 2,
		// 	err:    nil,
		// 	expLen: 0,
		// },
		// {
		// 	height: initHeight + ccvtypes.DefaultHistoricalEntries + 2,
		// 	err:    nil,
		// 	expLen: initValsetLen + 2,
		// },
	}

	for i, tc := range testCases {
		fmt.Printf("Normal Operations:: START TEST CASE %d ###\n", i)
		cCtx().WithBlockHeight(tc.height)
		hi, err := consumerKeeper.GetHistoricalInfo(cCtx().WithBlockHeight(tc.height), tc.height)
		fmt.Println(err)
		k.Require().Equal(tc.err, err)
		k.Require().Len(hi.Valset, tc.expLen)
		fmt.Printf("Normal Operations:: TEST CASE %d DONE ###\n", i)
	}
}
