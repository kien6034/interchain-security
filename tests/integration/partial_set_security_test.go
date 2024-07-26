package integration

import (
	"testing"

	"cosmossdk.io/math"
	ccv "github.com/cosmos/interchain-security/v5/x/ccv/types"
	"github.com/stretchr/testify/require"

	icstestingutils "github.com/cosmos/interchain-security/v5/testutil/ibc_testing"

	appConsumer "github.com/cosmos/interchain-security/v5/app/consumer"
	appProvider "github.com/cosmos/interchain-security/v5/app/provider"
)

// we need a stake multiplier because tokens do not directly correspond to voting power
// this is needed because 1000000 tokens = 1 voting power, so lower multipliers
// will be verbose and harder to read because small token numbers
// won't correspond to at least one voting power
const stake_multiplier = 1000000

// TestMinStake tests the min stake parameter.
// It starts a provider and single consumer chain,
// sets the initial powers according to the input, and then
// sets the min stake parameter according to the test case.
// Finally, it checks that the validator set on the consumer chain is as expected
// according to the min stake parameter.
func TestMinStake(t *testing.T) {
	testCases := []struct {
		name                string
		stakedTokens        []int64
		minStake            uint64
		expectedConsuValSet []int64
	}{
		{
			name: "disabled min stake",
			stakedTokens: []int64{
				1 * stake_multiplier,
				2 * stake_multiplier,
				3 * stake_multiplier,
				4 * stake_multiplier,
			},
			minStake: 0,
			expectedConsuValSet: []int64{
				1 * stake_multiplier,
				2 * stake_multiplier,
				3 * stake_multiplier,
				4 * stake_multiplier,
			},
		},
		{
			name: "stake multiplier - standard case",
			stakedTokens: []int64{
				1 * stake_multiplier,
				2 * stake_multiplier,
				3 * stake_multiplier,
				4 * stake_multiplier,
			},
			minStake: 3 * stake_multiplier,
			expectedConsuValSet: []int64{
				3 * stake_multiplier,
				4 * stake_multiplier,
			},
		},
		{
			name: "check min stake with multiple equal stakes",
			stakedTokens: []int64{
				1 * stake_multiplier,
				2 * stake_multiplier,
				2 * stake_multiplier,
				2 * stake_multiplier,
			},
			minStake: 2 * stake_multiplier,
			expectedConsuValSet: []int64{
				2 * stake_multiplier,
				2 * stake_multiplier,
				2 * stake_multiplier,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewCCVTestSuite[*appProvider.App, *appConsumer.App](
				// Pass in ibctesting.AppIniters for provider and consumer.
				icstestingutils.ProviderAppIniter, icstestingutils.ConsumerAppIniter, []string{})
			s.SetT(t)
			s.SetupTest()

			providerKeeper := s.providerApp.GetProviderKeeper()
			s.SetupCCVChannel(s.path)

			// set validator powers
			vals, err := providerKeeper.GetLastBondedValidators(s.providerChain.GetContext())
			s.Require().NoError(err)

			delegatorAccount := s.providerChain.SenderAccounts[0]

			for i, val := range vals {
				power := tc.stakedTokens[i]
				valAddr, err := providerKeeper.ValidatorAddressCodec().StringToBytes(val.GetOperator())
				s.Require().NoError(err)
				undelegate(s, delegatorAccount.SenderAccount.GetAddress(), valAddr, math.LegacyOneDec())

				// set validator power
				delegateByIdx(s, delegatorAccount.SenderAccount.GetAddress(), math.NewInt(power), i)
			}

			// end the epoch to apply the updates
			s.nextEpoch()

			// Relay 1 VSC packet from provider to consumer
			relayAllCommittedPackets(s, s.providerChain, s.path, ccv.ProviderPortID, s.path.EndpointB.ChannelID, 1)

			// end the block on the consumer to apply the updates
			s.consumerChain.NextBlock()

			// get the last bonded validators
			lastVals, err := providerKeeper.GetLastBondedValidators(s.providerChain.GetContext())
			s.Require().NoError(err)

			for i, val := range lastVals {
				// check that the intiial state was set correctly
				require.Equal(s.T(), math.NewInt(tc.stakedTokens[i]), val.Tokens)
			}

			// check the validator set on the consumer chain is the original one
			consuValSet := s.consumerChain.LastHeader.ValidatorSet
			s.Require().Equal(len(consuValSet.Validators), 4)

			// get just the powers of the consu val set
			consuValPowers := make([]int64, len(consuValSet.Validators))
			for i, consuVal := range consuValSet.Validators {
				// voting power corresponds to staked tokens at a 1:stake_multiplier ratio
				consuValPowers[i] = consuVal.VotingPower * stake_multiplier
			}

			s.Require().ElementsMatch(consuValPowers, tc.stakedTokens)

			// adjust parameters

			// set the minStake according to the test case
			providerKeeper.SetMinStake(s.providerChain.GetContext(), s.consumerChain.ChainID, tc.minStake)

			// undelegate and delegate to trigger a vscupdate
			delegateAndUndelegate(s, delegatorAccount.SenderAccount.GetAddress(), math.NewInt(1*stake_multiplier), 1)

			// end the epoch to apply the updates
			s.nextEpoch()

			// Relay 1 VSC packet from provider to consumer
			relayAllCommittedPackets(s, s.providerChain, s.path, ccv.ProviderPortID, s.path.EndpointB.ChannelID, 1)

			// end the block on the consumer to apply the updates
			s.consumerChain.NextBlock()

			// construct the new val powers
			newConsuValSet := s.consumerChain.LastHeader.ValidatorSet
			newConsuValPowers := make([]int64, len(newConsuValSet.Validators))
			for i, consuVal := range newConsuValSet.Validators {
				// voting power corresponds to staked tokens at a 1:stake_multiplier ratio
				newConsuValPowers[i] = consuVal.VotingPower * stake_multiplier
			}

			// check that the new validator set is as expected
			s.Require().ElementsMatch(newConsuValPowers, tc.expectedConsuValSet)
		})
	}
}
