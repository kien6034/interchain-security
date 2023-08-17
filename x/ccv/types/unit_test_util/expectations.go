package ututil

import (
	time "time"

	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	ibctmtypes "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	"github.com/golang/mock/gomock"
	extra "github.com/oxyno-zeta/gomock-extra-matcher"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/interchain-security/v3/x/ccv/types"
)

//
// A file containing groups of commonly used mock expectations.
// Note: Each group of mock expectations is associated with a single method
// that may be called during unit tests.
//

// GetMocksForCreateConsumerClient returns mock expectations needed to call CreateConsumerClient().
func GetMocksForCreateConsumerClient(ctx sdk.Context, mocks *MockedKeepers,
	expectedChainID string, expectedLatestHeight clienttypes.Height,
) []*gomock.Call {
	// append MakeConsumerGenesis and CreateClient expectations
	expectations := GetMocksForMakeConsumerGenesis(ctx, mocks, time.Hour)
	createClientExp := mocks.MockClientKeeper.EXPECT().CreateClient(
		gomock.Any(),
		// Allows us to expect a match by field. These are the only two client state values
		// that are dependant on parameters passed to CreateConsumerClient.
		extra.StructMatcher().Field(
			"ChainId", expectedChainID).Field(
			"LatestHeight", expectedLatestHeight,
		),
		gomock.Any(),
	).Return("clientID", nil).Times(1)
	expectations = append(expectations, createClientExp)

	return expectations
}

// GetMocksForMakeConsumerGenesis returns mock expectations needed to call MakeConsumerGenesis().
func GetMocksForMakeConsumerGenesis(ctx sdk.Context, mocks *MockedKeepers,
	unbondingTimeToInject time.Duration,
) []*gomock.Call {
	return []*gomock.Call{
		mocks.MockStakingKeeper.EXPECT().UnbondingTime(gomock.Any()).Return(unbondingTimeToInject).Times(1),

		mocks.MockClientKeeper.EXPECT().GetSelfConsensusState(gomock.Any(),
			clienttypes.GetSelfHeight(ctx)).Return(&ibctmtypes.ConsensusState{}, nil).Times(1),

		mocks.MockStakingKeeper.EXPECT().IterateLastValidatorPowers(gomock.Any(), gomock.Any()).Times(1),
	}
}

// GetMocksForSetConsumerChain returns mock expectations needed to call SetConsumerChain().
func GetMocksForSetConsumerChain(ctx sdk.Context, mocks *MockedKeepers,
	chainIDToInject string,
) []*gomock.Call {
	return []*gomock.Call{
		mocks.MockChannelKeeper.EXPECT().GetChannel(ctx, types.ProviderPortID, gomock.Any()).Return(
			channeltypes.Channel{
				State:          channeltypes.OPEN,
				ConnectionHops: []string{"connectionID"},
			},
			true,
		).Times(1),
		mocks.MockConnectionKeeper.EXPECT().GetConnection(ctx, "connectionID").Return(
			conntypes.ConnectionEnd{ClientId: "clientID"}, true,
		).Times(1),
		mocks.MockClientKeeper.EXPECT().GetClientState(ctx, "clientID").Return(
			&ibctmtypes.ClientState{ChainId: chainIDToInject}, true,
		).Times(1),
	}
}

// GetMocksForStopConsumerChain returns mock expectations needed to call StopConsumerChain().
func GetMocksForStopConsumerChain(ctx sdk.Context, mocks *MockedKeepers) []*gomock.Call {
	dummyCap := &capabilitytypes.Capability{}
	return []*gomock.Call{
		mocks.MockChannelKeeper.EXPECT().GetChannel(gomock.Any(), types.ProviderPortID, "channelID").Return(
			channeltypes.Channel{State: channeltypes.OPEN}, true,
		).Times(1),
		mocks.MockScopedKeeper.EXPECT().GetCapability(gomock.Any(), gomock.Any()).Return(dummyCap, true).Times(1),
		mocks.MockChannelKeeper.EXPECT().ChanCloseInit(gomock.Any(), types.ProviderPortID, "channelID", dummyCap).Times(1),
	}
}

func GetMocksForHandleSlashPacket(ctx sdk.Context, mocks MockedKeepers,
	expectedProviderValConsAddr sdk.ConsAddress,
	valToReturn stakingtypes.Validator, expectJailing bool,
) []*gomock.Call {
	// These first two calls are always made.
	calls := []*gomock.Call{
		mocks.MockStakingKeeper.EXPECT().GetValidatorByConsAddr(
			ctx, expectedProviderValConsAddr).Return(
			valToReturn, true,
		).Times(1),

		mocks.MockSlashingKeeper.EXPECT().IsTombstoned(ctx,
			expectedProviderValConsAddr).Return(false).Times(1),
	}

	if expectJailing {
		calls = append(calls, mocks.MockStakingKeeper.EXPECT().Jail(
			gomock.Eq(ctx),
			gomock.Eq(expectedProviderValConsAddr),
		).Return())

		// JailUntil is set in this code path.
		calls = append(calls, mocks.MockSlashingKeeper.EXPECT().DowntimeJailDuration(ctx).Return(time.Hour).Times(1))
		calls = append(calls, mocks.MockSlashingKeeper.EXPECT().JailUntil(ctx,
			expectedProviderValConsAddr, gomock.Any()).Times(1))
	}

	return calls
}

func ExpectLatestConsensusStateMock(ctx sdk.Context, mocks MockedKeepers, clientID string, consState *ibctmtypes.ConsensusState) *gomock.Call {
	return mocks.MockClientKeeper.EXPECT().
		GetLatestClientConsensusState(ctx, clientID).Return(consState, true).Times(1)
}

func ExpectCreateClientMock(ctx sdk.Context, mocks MockedKeepers, clientID string, clientState *ibctmtypes.ClientState, consState *ibctmtypes.ConsensusState) *gomock.Call {
	return mocks.MockClientKeeper.EXPECT().CreateClient(ctx, clientState, consState).Return(clientID, nil).Times(1)
}

func ExpectGetCapabilityMock(ctx sdk.Context, mocks MockedKeepers, times int) *gomock.Call {
	return mocks.MockScopedKeeper.EXPECT().GetCapability(
		ctx, host.PortPath(types.ConsumerPortID),
	).Return(nil, true).Times(times)
}

func GetMocksForSendIBCPacket(ctx sdk.Context, mocks MockedKeepers, channelID string, times int) []*gomock.Call {
	return []*gomock.Call{
		mocks.MockChannelKeeper.EXPECT().GetChannel(ctx, types.ConsumerPortID,
			"consumerCCVChannelID").Return(channeltypes.Channel{}, true).Times(times),
		mocks.MockScopedKeeper.EXPECT().GetCapability(ctx,
			host.ChannelCapabilityPath(types.ConsumerPortID, "consumerCCVChannelID")).Return(
			capabilitytypes.NewCapability(1), true).Times(times),
		mocks.MockChannelKeeper.EXPECT().SendPacket(ctx,
			capabilitytypes.NewCapability(1),
			types.ConsumerPortID,
			"consumerCCVChannelID",
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(uint64(888), nil).Times(times),
	}
}
