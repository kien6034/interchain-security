package integration

import (
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	providerkeeper "github.com/cosmos/interchain-security/v5/x/ccv/provider/keeper"
	ccv "github.com/cosmos/interchain-security/v5/x/ccv/types"
)

// TestUndelegationNormalOperation tests that undelegations complete after
// the unbonding period elapses on both the consumer and provider, without
// VSC packets timing out.
func (s *CCVTestSuite) TestUndelegationNormalOperation() {
	unbondConsumer := func(expectedPackets int) {
		// relay 1 VSC packet from provider to consumer
		relayAllCommittedPackets(s, s.providerChain, s.path, ccv.ProviderPortID, s.path.EndpointB.ChannelID, expectedPackets)
		// increment time so that the unbonding period ends on the consumer
		incrementTimeByUnbondingPeriod(s, Consumer)
		// relay 1 VSCMatured packet from consumer to provider
		relayAllCommittedPackets(s, s.consumerChain, s.path, ccv.ConsumerPortID, s.path.EndpointA.ChannelID, expectedPackets)
	}

	testCases := []struct {
		name     string
		shareDiv int64
		unbond   func(expBalance, balance math.Int)
	}{
		{
			"provider unbonding period elapses first", 2, func(expBalance, balance math.Int) {
				// increment time so that the unbonding period ends on the provider
				incrementTimeByUnbondingPeriod(s, Provider)

				// check that onHold is true
				checkStakingUnbondingOps(s, 1, true, true, "unbonding should be on hold")

				// check that the unbonding is not complete
				s.Require().Equal(expBalance, balance, "unexpected balance after provider unbonding")

				// undelegation complete on consumer
				unbondConsumer(1)
			},
		},
		{
			"consumer unbonding period elapses first", 2, func(expBalance, balance math.Int) {
				// undelegation complete on consumer
				unbondConsumer(1)

				// check that onHold is false
				checkStakingUnbondingOps(s, 1, true, false, "unbonding should be not be on hold")

				// check that the unbonding is not complete
				s.Require().Equal(expBalance, balance, "unexpected balance after consumer unbonding")

				// increment time so that the unbonding period ends on the provider
				incrementTimeByUnbondingPeriod(s, Provider)
			},
		},
		{
			"no valset changes", 1, func(expBalance, balance math.Int) {
				// undelegation complete on consumer
				unbondConsumer(1)

				// check that onHold is false
				checkStakingUnbondingOps(s, 1, true, false, "unbonding should be not be on hold")

				// check that the unbonding is not complete
				s.Require().Equal(expBalance, balance, "unexpected balance after consumer unbonding")

				// increment time so that the unbonding period ends on the provider
				incrementTimeByUnbondingPeriod(s, Provider)
			},
		},
	}

	for i, tc := range testCases {
		providerKeeper := s.providerApp.GetProviderKeeper()
		consumerKeeper := s.consumerApp.GetConsumerKeeper()
		stakingKeeper := s.providerApp.GetTestStakingKeeper()

		s.SetupCCVChannel(s.path)

		// set VSC timeout period to not trigger the removal of the consumer chain
		providerUnbondingPeriod, err := stakingKeeper.UnbondingTime(s.providerCtx())
		s.Require().NoError(err)
		consumerUnbondingPeriod := consumerKeeper.GetUnbondingPeriod(s.consumerCtx())
		providerKeeper.SetVscTimeoutPeriod(s.providerCtx(), providerUnbondingPeriod+consumerUnbondingPeriod+24*time.Hour)

		// delegate bondAmt and undelegate tc.shareDiv of it
		bondAmt := math.NewInt(10000000)
		delAddr := s.providerChain.SenderAccount.GetAddress()
		initBalance, valsetUpdateID := delegateAndUndelegate(s, delAddr, bondAmt, tc.shareDiv)
		// - check that staking unbonding op was created and onHold is true
		checkStakingUnbondingOps(s, 1, true, true, "test: "+tc.name)
		// - check that CCV unbonding op was created
		checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, true, "test: "+tc.name)

		// call NextBlock on the provider (which increments the height)
		s.nextEpoch()

		// unbond both on provider and consumer and check that
		// the balance remains unchanged in between
		tc.unbond(initBalance.Sub(bondAmt), getBalance(s, s.providerCtx(), delAddr))

		// check that the unbonding operation completed
		// - check that ccv unbonding op has been deleted
		checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, false, "test: "+tc.name)
		// - check that staking unbonding op has been deleted
		checkStakingUnbondingOps(s, valsetUpdateID, false, false, "test: "+tc.name)
		// - check that necessary delegated coins have been returned
		unbondAmt := bondAmt.Sub(bondAmt.Quo(math.NewInt(tc.shareDiv)))
		s.Require().Equal(
			initBalance.Sub(unbondAmt),
			getBalance(s, s.providerCtx(), delAddr),
			"unexpected initial balance after unbonding; test: %s", tc.name,
		)

		if i+1 < len(testCases) {
			// reset suite to reset provider client
			s.SetupTest()
		}
	}
}

// TestUndelegationVscTimeout tests that an undelegation
// completes after vscTimeoutPeriod even if it does not
// reach maturity on the consumer chain. In this case,
// the consumer chain is removed.
func (s *CCVTestSuite) TestUndelegationVscTimeout() {
	providerKeeper := s.providerApp.GetProviderKeeper()

	s.SetupCCVChannel(s.path)

	// set VSC timeout period to trigger the removal of the consumer chain
	vscTimeout := providerKeeper.GetVscTimeoutPeriod(s.providerCtx())

	// delegate bondAmt and undelegate 1/2 of it
	bondAmt := math.NewInt(10000000)
	delAddr := s.providerChain.SenderAccount.GetAddress()
	initBalance, valsetUpdateID := delegateAndUndelegate(s, delAddr, bondAmt, 2)
	// - check that staking unbonding op was created and onHold is true
	checkStakingUnbondingOps(s, 1, true, true)
	// - check that CCV unbonding op was created
	checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, true)

	// call NextBlock on the provider (which increments the height)
	s.providerChain.NextBlock()

	// increment time so that the unbonding period ends on the provider
	incrementTimeByUnbondingPeriod(s, Provider)

	// check that onHold is true
	checkStakingUnbondingOps(s, 1, true, true, "unbonding should be on hold")

	// check that the unbonding is not complete
	s.Require().Equal(
		initBalance.Sub(bondAmt),
		getBalance(s, s.providerCtx(), delAddr),
		"unexpected balance after provider unbonding")

	// increment time
	incrementTime(s, vscTimeout)

	// check whether the chain was removed
	chainID := s.consumerChain.ChainID
	_, found := providerKeeper.GetConsumerClientId(s.providerCtx(), chainID)
	s.Require().Equal(false, found, "consumer chain was not removed")

	// check if the chain was properly removed
	s.checkConsumerChainIsRemoved(chainID, true)

	// check that the unbonding operation completed
	// - check that ccv unbonding op has been deleted
	checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, false)
	// - check that staking unbonding op has been deleted
	checkStakingUnbondingOps(s, valsetUpdateID, false, false)
	// - check that necessary delegated coins have been returned
	unbondAmt := bondAmt.Sub(bondAmt.Quo(math.NewInt(2)))
	s.Require().Equal(
		initBalance.Sub(unbondAmt),
		getBalance(s, s.providerCtx(), delAddr),
		"unexpected initial balance after VSC timeout",
	)
}

// TestUndelegationDuringInit checks that before the CCV channel is established
//   - no undelegations can complete, even if the provider unbonding period elapses
//   - all the VSC packets are stored in state as pending
//   - if the channel handshake times out, then the undelegation completes
func (s *CCVTestSuite) TestUndelegationDuringInit() {
	testCases := []struct {
		name                       string
		updateInitTimeoutTimestamp func(*providerkeeper.Keeper, time.Duration)
		removed                    bool
	}{
		{
			"channel handshake completes after unbonding period", func(pk *providerkeeper.Keeper, pUnbondingPeriod time.Duration) {
				// change the init timeout timestamp for this consumer chain
				// to make sure the chain is not removed before the unbonding period elapses
				ts := s.providerCtx().BlockTime().Add(pUnbondingPeriod + 24*time.Hour)
				pk.SetInitTimeoutTimestamp(s.providerCtx(), s.consumerChain.ChainID, uint64(ts.UnixNano()))
			}, false,
		},
		{
			"channel handshake times out before unbonding period", func(pk *providerkeeper.Keeper, pUnbondingPeriod time.Duration) {
				// change the init timeout timestamp for this consumer chain
				// to make sure the chain is removed before the unbonding period elapses
				ts := s.providerCtx().BlockTime().Add(pUnbondingPeriod - 24*time.Hour)
				pk.SetInitTimeoutTimestamp(s.providerCtx(), s.consumerChain.ChainID, uint64(ts.UnixNano()))
			}, true,
		},
	}

	for i, tc := range testCases {
		providerKeeper := s.providerApp.GetProviderKeeper()
		stakingKeeper := s.providerApp.GetTestStakingKeeper()

		// delegate bondAmt and undelegate 1/2 of it
		bondAmt := math.NewInt(10000000)
		delAddr := s.providerChain.SenderAccount.GetAddress()
		initBalance, valsetUpdateID := delegateAndUndelegate(s, delAddr, bondAmt, 2)
		// - check that staking unbonding op was created and onHold is true
		checkStakingUnbondingOps(s, 1, true, true, "test: "+tc.name)
		// - check that CCV unbonding op was created
		checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, true, "test: "+tc.name)

		// get provider unbonding period
		providerUnbondingPeriod, err := stakingKeeper.UnbondingTime(s.providerCtx())
		s.Require().NoError(err)
		// update init timeout timestamp
		tc.updateInitTimeoutTimestamp(&providerKeeper, providerUnbondingPeriod)

		s.nextEpoch()

		// check that the VSC packet is stored in state as pending
		pendingVSCs := providerKeeper.GetPendingVSCPackets(s.providerCtx(), s.consumerChain.ChainID)
		s.Require().Lenf(pendingVSCs, 1, "no pending VSC packet found; test: %s", tc.name)

		// delegate again to create another VSC packet
		delegate(s, delAddr, bondAmt)

		s.nextEpoch()

		// check that the VSC packet is stored in state as pending
		pendingVSCs = providerKeeper.GetPendingVSCPackets(s.providerCtx(), s.consumerChain.ChainID)
		s.Require().Lenf(pendingVSCs, 2, "only one pending VSC packet found; test: %s", tc.name)

		// increment time so that the unbonding period ends on the provider
		incrementTimeByUnbondingPeriod(s, Provider)

		// check whether the unbonding op is still there and onHold is true
		checkStakingUnbondingOps(s, 1, !tc.removed, true, "test: "+tc.name)

		if !tc.removed {
			// check that unbonding has not yet completed, i.e., the initBalance
			// is still lower by the bond amount, because it has been taken out of
			// the delegator's account
			s.Require().Equal(
				initBalance.Sub(bondAmt).Sub(bondAmt),
				getBalance(s, s.providerCtx(), delAddr),
				"unexpected initial balance before unbonding; test: %s", tc.name,
			)

			// complete CCV channel setup
			s.SetupCCVChannel(s.path)
			s.nextEpoch()

			// relay VSC packets from provider to consumer
			relayAllCommittedPackets(s, s.providerChain, s.path, ccv.ProviderPortID, s.path.EndpointB.ChannelID, 2)

			// increment time so that the unbonding period ends on the consumer
			incrementTimeByUnbondingPeriod(s, Consumer)

			// relay VSCMatured packets from consumer to provider
			relayAllCommittedPackets(s, s.consumerChain, s.path, ccv.ConsumerPortID, s.path.EndpointA.ChannelID, 2)

			// check that the unbonding operation completed
			// - check that ccv unbonding op has been deleted
			checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, false, "test: "+tc.name)
			// - check that staking unbonding op has been deleted
			checkStakingUnbondingOps(s, valsetUpdateID, false, false, "test: "+tc.name)
			// - check that one quarter the delegated coins have been returned
			s.Require().Equal(
				initBalance.Sub(bondAmt).Sub(bondAmt.Quo(math.NewInt(2))),
				getBalance(s, s.providerCtx(), delAddr),
				"unexpected initial balance after unbonding; test: %s", tc.name,
			)
		}

		if i+1 < len(testCases) {
			// reset suite to reset provider client
			s.SetupTest()
		}
	}
}

// Bond some tokens on provider
// Unbond them to create unbonding op
// Check unbonding ops on both sides
// Advance time so that provider's unbonding op completes
// Check that unbonding has completed in provider staking
func (s *CCVTestSuite) TestUnbondingNoConsumer() {
	providerKeeper := s.providerApp.GetProviderKeeper()
	providerStakingKeeper := s.providerApp.GetTestStakingKeeper()

	// remove all consumer chains, which were already started during setup
	for chainID := range s.consumerBundles {
		err := providerKeeper.StopConsumerChain(s.providerCtx(), chainID, true)
		s.Require().NoError(err)
	}

	// delegate bondAmt and undelegate 1/2 of it
	bondAmt := math.NewInt(10000000)
	delAddr := s.providerChain.SenderAccount.GetAddress()
	initBalance, valsetUpdateID := delegateAndUndelegate(s, delAddr, bondAmt, 2)
	// - check that staking unbonding op was created and onHold is FALSE
	checkStakingUnbondingOps(s, 1, true, false)
	// - check that CCV unbonding op was NOT created
	checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, false)

	// increment time so that the unbonding period ends on the provider;
	// cannot use incrementTimeByUnbondingPeriod() since it tries
	// to also update the provider's client on the consumer
	providerUnbondingPeriod, err := providerStakingKeeper.UnbondingTime(s.providerCtx())
	s.Require().NoError(err)
	s.coordinator.IncrementTimeBy(providerUnbondingPeriod + time.Hour)

	// call NextBlock on the provider (which increments the height)
	s.providerChain.NextBlock()

	// check that the unbonding operation completed
	// - check that staking unbonding op has been deleted
	checkStakingUnbondingOps(s, valsetUpdateID, false, false)
	// - check that half the coins have been returned
	s.Require().True(getBalance(s, s.providerCtx(), delAddr).Equal(initBalance.Sub(bondAmt.Quo(math.NewInt(2)))))
}

// TestRedelegationNoConsumer tests a redelegate transaction
// submitted on a provider chain with no consumers
func (s *CCVTestSuite) TestRedelegationNoConsumer() {
	providerKeeper := s.providerApp.GetProviderKeeper()
	stakingKeeper := s.providerApp.GetTestStakingKeeper()

	// stop the consumer chain, which was already started during setup
	err := providerKeeper.StopConsumerChain(s.providerCtx(), s.consumerChain.ChainID, true)
	s.Require().NoError(err)

	// Setup delegator, bond amount, and src/dst validators
	bondAmt := math.NewInt(10000000)
	delAddr := s.providerChain.SenderAccount.GetAddress()
	_, srcVal := s.getValByIdx(0)
	_, dstVal := s.getValByIdx(1)

	delegateAndRedelegate(
		s,
		delAddr,
		srcVal,
		dstVal,
		bondAmt,
	)

	// 1 redelegation record should exist for original delegator
	redelegations := checkRedelegations(s, delAddr, 1)

	// Check that the only entry has appropriate maturation time, the unbonding period from now
	unbondingTime, err := stakingKeeper.UnbondingTime(s.providerCtx())
	s.Require().NoError(err)
	checkRedelegationEntryCompletionTime(
		s,
		redelegations[0].Entries[0],
		s.providerCtx().BlockTime().Add(unbondingTime),
	)

	// required before call to incrementTimeByUnbondingPeriod or else a panic
	// occurs in ibc-go because trusted validators don't match last trusted.
	s.providerChain.NextBlock()

	// Increment time so that the unbonding period passes on the provider
	incrementTimeByUnbondingPeriod(s, Provider)

	// Call NextBlock on the provider (which increments the height)
	s.providerChain.NextBlock()

	// No redelegation records should exist for original delegator anymore
	checkRedelegations(s, delAddr, 0)
}

// TestRedelegationWithConsumer tests a redelegate transaction submitted on a provider chain
// when the unbonding period elapses first on the provider chain
func (s *CCVTestSuite) TestRedelegationProviderFirst() {
	s.SetupCCVChannel(s.path)
	s.SetupTransferChannel()

	providerKeeper := s.providerApp.GetProviderKeeper()
	consumerKeeper := s.consumerApp.GetConsumerKeeper()
	stakingKeeper := s.providerApp.GetTestStakingKeeper()

	// set VSC timeout period to not trigger the removal of the consumer chain
	providerUnbondingPeriod, err := stakingKeeper.UnbondingTime(s.providerCtx())
	s.Require().NoError(err)
	consumerUnbondingPeriod := consumerKeeper.GetUnbondingPeriod(s.consumerCtx())
	providerKeeper.SetVscTimeoutPeriod(s.providerCtx(), providerUnbondingPeriod+consumerUnbondingPeriod+24*time.Hour)

	// Setup delegator, bond amount, and src/dst validators
	bondAmt := math.NewInt(10000000)
	delAddr := s.providerChain.SenderAccount.GetAddress()
	_, srcVal := s.getValByIdx(0)
	_, dstVal := s.getValByIdx(1)

	delegateAndRedelegate(
		s,
		delAddr,
		srcVal,
		dstVal,
		bondAmt,
	)

	// 1 redelegation record should exist for original delegator
	redelegations := checkRedelegations(s, delAddr, 1)

	// Check that the only entry has appropriate maturation time, the unbonding period from now
	unbondingTime, err := stakingKeeper.UnbondingTime(s.providerCtx())
	s.Require().NoError(err)
	checkRedelegationEntryCompletionTime(
		s,
		redelegations[0].Entries[0],
		s.providerCtx().BlockTime().Add(unbondingTime),
	)

	// Save the current valset update ID
	valsetUpdateID := providerKeeper.GetValidatorSetUpdateId(s.providerCtx())

	// Check that CCV unbonding op was created from AfterUnbondingInitiated hook
	checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, true)

	// move forward by an epoch to be able to relay VSC packets
	s.nextEpoch()

	// Relay 2 VSC packets from provider to consumer (original delegation, and redelegation)
	relayAllCommittedPackets(s, s.providerChain, s.path,
		ccv.ProviderPortID, s.path.EndpointB.ChannelID, 2)

	// Increment time so that the unbonding period ends on the provider
	incrementTimeByUnbondingPeriod(s, Provider)

	// 1 redelegation record should still exist for original delegator on provider
	checkRedelegations(s, delAddr, 1)

	// CCV unbonding op should also still exist
	checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, true)

	// Increment time so that the unbonding period ends on the consumer
	incrementTimeByUnbondingPeriod(s, Consumer)

	// Relay 2 VSCMatured packets from consumer to provider (original delegation and redelegation)
	relayAllCommittedPackets(s, s.consumerChain,
		s.path, ccv.ConsumerPortID, s.path.EndpointA.ChannelID, 2)

	//
	// Check that the redelegation operation has now completed on provider
	//

	// Redelegation record should be deleted for original delegator
	checkRedelegations(s, delAddr, 0)

	// Check that ccv unbonding op has been deleted
	checkCCVUnbondingOp(s, s.providerCtx(), s.consumerChain.ChainID, valsetUpdateID, false)
}

// This test reproduces a fixed bug when an inactive validator enters back into the active set.
// It used to cause a panic in the provider module hook called by AfterUnbondingInitiated
// during the staking module EndBlock.
func (s *CCVTestSuite) TestTooManyLastValidators() {
	sk := s.providerApp.GetTestStakingKeeper()

	getLastValsFn := func(ctx sdk.Context) []stakingtypes.Validator {
		lastVals, err := sk.GetLastValidators(s.providerCtx())
		s.Require().NoError(err)
		return lastVals
	}

	// get current staking params
	p, err := sk.GetParams(s.providerCtx())
	s.Require().NoError(err)

	// get validators, which are all active at the moment
	vals, err := sk.GetAllValidators(s.providerCtx())
	s.Require().NoError(err)

	s.Require().Equal(len(vals), len(getLastValsFn(s.providerCtx())))

	// jail a validator
	val := vals[0]
	consAddr, err := val.GetConsAddr()
	s.Require().NoError(err)
	sk.Jail(s.providerCtx(), consAddr)

	// save the current number of bonded vals
	lastVals := getLastValsFn(s.providerCtx())

	// pass one block to apply the validator set changes
	// (calls ApplyAndReturnValidatorSetUpdates in the the staking module EndBlock)
	s.providerChain.NextBlock()

	// verify that the number of bonded validators is decreased by one
	s.Require().Equal(len(lastVals)-1, len(getLastValsFn(s.providerCtx())))

	// update maximum validator to equal the number of bonded validators
	p.MaxValidators = uint32(len(getLastValsFn(s.providerCtx())))
	sk.SetParams(s.providerCtx(), p)

	// pass one block to apply validator set changes
	s.providerChain.NextBlock()

	// unjail validator
	// Note that since validators are sorted in descending order, the unjailed validator
	// enters the active set again since it's ranked first by voting power.
	sk.Unjail(s.providerCtx(), consAddr)

	// pass another block to update the validator set
	// which causes a panic due to a GetLastValidator call in
	// ApplyAndReturnValidatorSetUpdates where the staking module has a inconsistent state
	s.Require().NotPanics(s.providerChain.NextBlock)
	s.Require().NotPanics(func() { sk.ApplyAndReturnValidatorSetUpdates(s.providerCtx()) })
	s.Require().NotPanics(func() { getLastValsFn(s.providerCtx()) })
}
