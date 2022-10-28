package keeper_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/golang/mock/gomock"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	ibcsimapp "github.com/cosmos/ibc-go/v3/testing/simapp"
	"golang.org/x/exp/slices"

	testkeeper "github.com/cosmos/interchain-security/testutil/keeper"
	"github.com/cosmos/interchain-security/x/ccv/provider/types"
	ccv "github.com/cosmos/interchain-security/x/ccv/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmprotocrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/stretchr/testify/require"
)

// TestValsetUpdateBlockHeight tests the getter, setter, and deletion methods for valset updates mapped to block height
func TestValsetUpdateBlockHeight(t *testing.T) {
	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
	defer ctrl.Finish()

	blockHeight, found := providerKeeper.GetValsetUpdateBlockHeight(ctx, uint64(0))
	require.False(t, found)
	require.Zero(t, blockHeight)

	providerKeeper.SetValsetUpdateBlockHeight(ctx, uint64(1), uint64(2))
	blockHeight, found = providerKeeper.GetValsetUpdateBlockHeight(ctx, uint64(1))
	require.True(t, found)
	require.Equal(t, blockHeight, uint64(2))

	providerKeeper.DeleteValsetUpdateBlockHeight(ctx, uint64(1))
	blockHeight, found = providerKeeper.GetValsetUpdateBlockHeight(ctx, uint64(1))
	require.False(t, found)
	require.Zero(t, blockHeight)

	providerKeeper.SetValsetUpdateBlockHeight(ctx, uint64(1), uint64(2))
	providerKeeper.SetValsetUpdateBlockHeight(ctx, uint64(3), uint64(4))
	blockHeight, found = providerKeeper.GetValsetUpdateBlockHeight(ctx, uint64(3))
	require.True(t, found)
	require.Equal(t, blockHeight, uint64(4))
}

// TestSlashAcks tests the getter, setter, iteration, and deletion methods for stored slash acknowledgements
func TestSlashAcks(t *testing.T) {
	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
	defer ctrl.Finish()

	var chainsAcks [][]string

	penaltiesfN := func() (penalties []string) {
		providerKeeper.IterateSlashAcks(ctx, func(id string, acks []string) bool {
			chainsAcks = append(chainsAcks, acks)
			return true
		})
		return
	}

	chainID := "consumer"

	acks := providerKeeper.GetSlashAcks(ctx, chainID)
	require.Nil(t, acks)

	p := []string{"alice", "bob", "charlie"}
	providerKeeper.SetSlashAcks(ctx, chainID, p)

	acks = providerKeeper.GetSlashAcks(ctx, chainID)
	require.NotNil(t, acks)

	require.Len(t, acks, 3)
	slashAcks := providerKeeper.ConsumeSlashAcks(ctx, chainID)
	require.Len(t, slashAcks, 3)

	acks = providerKeeper.GetSlashAcks(ctx, chainID)
	require.Nil(t, acks)

	chains := []string{"c1", "c2", "c3"}

	for _, c := range chains {
		providerKeeper.SetSlashAcks(ctx, c, p)
	}

	penaltiesfN()
	require.Len(t, chainsAcks, len(chains))
}

// TestAppendSlashAck tests the append method for stored slash acknowledgements
func TestAppendSlashAck(t *testing.T) {
	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
	defer ctrl.Finish()

	p := []string{"alice", "bob", "charlie"}
	chains := []string{"c1", "c2"}
	providerKeeper.SetSlashAcks(ctx, chains[0], p)

	providerKeeper.AppendSlashAck(ctx, chains[0], p[0])
	acks := providerKeeper.GetSlashAcks(ctx, chains[0])
	require.NotNil(t, acks)
	require.Len(t, acks, len(p)+1)

	providerKeeper.AppendSlashAck(ctx, chains[1], p[0])
	acks = providerKeeper.GetSlashAcks(ctx, chains[1])
	require.NotNil(t, acks)
	require.Len(t, acks, 1)
}

// TestPendingVSCs tests the getter, appending, and deletion methods for stored pending VSCs
func TestPendingVSCs(t *testing.T) {
	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
	defer ctrl.Finish()

	chainID := "consumer"

	_, found := providerKeeper.GetPendingVSCs(ctx, chainID)
	require.False(t, found)

	pks := ibcsimapp.CreateTestPubKeys(4)
	var ppks [4]tmprotocrypto.PublicKey
	for i, pk := range pks {
		ppks[i], _ = cryptocodec.ToTmProtoPublicKey(pk)
	}

	packetList := []ccv.ValidatorSetChangePacketData{
		{
			ValidatorUpdates: []abci.ValidatorUpdate{
				{PubKey: ppks[0], Power: 1},
				{PubKey: ppks[1], Power: 2},
			},
			ValsetUpdateId: 1,
		},
		{
			ValidatorUpdates: []abci.ValidatorUpdate{
				{PubKey: ppks[2], Power: 3},
			},
			ValsetUpdateId: 2,
		},
	}
	for _, packet := range packetList {
		providerKeeper.AppendPendingVSC(ctx, chainID, packet)
	}

	packets, found := providerKeeper.GetPendingVSCs(ctx, chainID)
	require.True(t, found)
	require.Len(t, packets, 2)

	newPacket := ccv.ValidatorSetChangePacketData{
		ValidatorUpdates: []abci.ValidatorUpdate{
			{PubKey: ppks[3], Power: 4},
		},
		ValsetUpdateId: 3,
	}
	providerKeeper.AppendPendingVSC(ctx, chainID, newPacket)
	vscs := providerKeeper.ConsumePendingVSCs(ctx, chainID)
	require.Len(t, vscs, 3)
	require.True(t, vscs[len(vscs)-1].ValsetUpdateId == 3)
	require.True(t, vscs[len(vscs)-1].GetValidatorUpdates()[0].PubKey.String() == ppks[3].String())

	_, found = providerKeeper.GetPendingVSCs(ctx, chainID)
	require.False(t, found)
}

// TestInitHeight tests the getter and setter methods for the stored block heights (on provider) when a given consumer chain was started
func TestInitHeight(t *testing.T) {
	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
	defer ctrl.Finish()

	tc := []struct {
		chainID  string
		expected uint64
	}{
		{expected: 0, chainID: "chain"},
		{expected: 10, chainID: "chain1"},
		{expected: 12, chainID: "chain2"},
	}

	providerKeeper.SetInitChainHeight(ctx, tc[1].chainID, tc[1].expected)
	providerKeeper.SetInitChainHeight(ctx, tc[2].chainID, tc[2].expected)

	for _, tc := range tc {
		height, _ := providerKeeper.GetInitChainHeight(ctx, tc.chainID)
		require.Equal(t, tc.expected, height)
	}
}

// TestHandleSlashPacketDoubleSigning tests the handling of a double-signing related slash packet, with mocks and unit tests
func TestHandleSlashPacketDoubleSigning(t *testing.T) {

	chainId := "consumer"
	infractionHeight := int64(5)

	keeperParams := testkeeper.NewInMemKeeperParams(t)
	ctx := keeperParams.Ctx

	slashPacket := ccv.NewSlashPacketData(
		abci.Validator{Address: ed25519.GenPrivKey().PubKey().Address(),
			Power: int64(0)},
		uint64(0),
		stakingtypes.DoubleSign,
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mocks := testkeeper.NewMockedKeepers(ctrl)
	mockSlashingKeeper := mocks.MockSlashingKeeper
	mockStakingKeeper := mocks.MockStakingKeeper

	// Setup expected mock calls
	gomock.InOrder(

		mockStakingKeeper.EXPECT().GetValidatorByConsAddr(
			ctx, sdk.ConsAddress(slashPacket.Validator.Address)).Return(
			stakingtypes.Validator{Status: stakingtypes.Bonded}, true,
		).Times(1),

		mockSlashingKeeper.EXPECT().IsTombstoned(ctx, sdk.ConsAddress(slashPacket.Validator.Address)).Return(false).Times(1),

		mockSlashingKeeper.EXPECT().SlashFractionDoubleSign(ctx).Return(sdk.NewDec(1)).Times(1),

		mockSlashingKeeper.EXPECT().Tombstone(ctx, sdk.ConsAddress(slashPacket.Validator.Address)).Times(1),

		mockStakingKeeper.EXPECT().Slash(
			ctx,
			sdk.ConsAddress(slashPacket.Validator.Address),
			infractionHeight,
			int64(0),      // power
			sdk.NewDec(1), // Slash fraction
			stakingtypes.DoubleSign).Return().Times(1),

		mockStakingKeeper.EXPECT().Jail(
			gomock.Eq(ctx),
			gomock.Eq(sdk.ConsAddress(slashPacket.Validator.Address)),
		).Return(),

		mockSlashingKeeper.EXPECT().JailUntil(ctx, sdk.ConsAddress(slashPacket.Validator.Address),
			evidencetypes.DoubleSignJailEndTime).Times(1),
	)

	providerKeeper := testkeeper.NewInMemProviderKeeper(keeperParams, mocks)

	providerKeeper.SetInitChainHeight(ctx, chainId, uint64(infractionHeight))

	success, err := providerKeeper.HandleSlashPacket(ctx, chainId, slashPacket)
	require.NoError(t, err)
	require.True(t, success)
}

func TestIterateOverUnbondingOpIndex(t *testing.T) {

	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
	defer ctrl.Finish()

	chainID := "6"

	// mock an unbonding index
	unbondingOpIndex := []uint64{0, 1, 2, 3, 4, 5, 6}

	// set ubd ops by varying vsc ids and index slices
	for i := 1; i < len(unbondingOpIndex); i++ {
		providerKeeper.SetUnbondingOpIndex(ctx, chainID, uint64(i), unbondingOpIndex[:i])
	}

	// check iterator returns expected entries
	i := 1
	providerKeeper.IterateOverUnbondingOpIndex(ctx, chainID, func(vscID uint64, ubdIndex []uint64) bool {
		require.Equal(t, uint64(i), vscID)
		require.EqualValues(t, unbondingOpIndex[:i], ubdIndex)
		i++
		return true
	})
	require.Equal(t, len(unbondingOpIndex), i)
}

func TestMaturedUnbondingOps(t *testing.T) {

	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
	defer ctrl.Finish()

	ids, err := providerKeeper.GetMaturedUnbondingOps(ctx)
	require.NoError(t, err)
	require.Nil(t, ids)

	unbondingOpIds := []uint64{0, 1, 2, 3, 4, 5, 6}
	err = providerKeeper.AppendMaturedUnbondingOps(ctx, unbondingOpIds)
	require.NoError(t, err)

	ids, err = providerKeeper.ConsumeMaturedUnbondingOps(ctx)
	require.NoError(t, err)
	require.Equal(t, len(unbondingOpIds), len(ids))
	for i := 0; i < len(unbondingOpIds); i++ {
		require.Equal(t, unbondingOpIds[i], ids[i])
	}
}

// TestPendingSlashPacket tests the queue and iteration functions
// for pending slash packets with assertion of FIFO ordering
func TestPendingSlashPackets(t *testing.T) {

	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(
		t, testkeeper.NewInMemKeeperParams(t))
	defer ctrl.Finish()

	// Consistent time for "now"
	now := time.Now()

	// Queue 3 slash packets for chainIDs 0, 1, 2
	for i := 0; i < 3; i++ {
		packet := types.NewSlashPacket(now, "chain-"+fmt.Sprint(i), testkeeper.GetNewSlashPacketData())
		providerKeeper.QueuePendingSlashPacket(ctx, packet)
	}
	// Queue 3 slash packets for chainIDs 0, 1, 2 an hour later
	for i := 0; i < 3; i++ {
		packet := types.NewSlashPacket(now.Add(time.Hour), "chain-"+fmt.Sprint(i), testkeeper.GetNewSlashPacketData())
		providerKeeper.QueuePendingSlashPacket(ctx, packet)
	}

	// Retrieve packets from store
	packets := providerKeeper.GetAllPendingSlashPackets(ctx)

	// Assert that packets are obtained in FIFO order according to block time
	firstChainIdSet := []string{packets[0].ConsumerChainID, packets[1].ConsumerChainID, packets[2].ConsumerChainID}
	require.True(t, slices.Contains(firstChainIdSet, "chain-0"))
	require.True(t, slices.Contains(firstChainIdSet, "chain-1"))
	require.True(t, slices.Contains(firstChainIdSet, "chain-2"))
	secondChainIdSet := []string{packets[3].ConsumerChainID, packets[4].ConsumerChainID, packets[5].ConsumerChainID}
	require.True(t, slices.Contains(secondChainIdSet, "chain-0"))
	require.True(t, slices.Contains(secondChainIdSet, "chain-1"))
	require.True(t, slices.Contains(secondChainIdSet, "chain-2"))

	// Queue 3 slash packets for chainIDs 5, 6, 7 another hour later
	for i := 0; i < 3; i++ {
		packet := types.NewSlashPacket(now.Add(2*time.Hour), "chain-"+fmt.Sprint(i+5), testkeeper.GetNewSlashPacketData())
		providerKeeper.QueuePendingSlashPacket(ctx, packet)
	}

	// Retrieve packets from store
	packets = providerKeeper.GetAllPendingSlashPackets(ctx)

	// Assert that packets are obtained in FIFO order according to block time
	firstChainIdSet = []string{packets[0].ConsumerChainID, packets[1].ConsumerChainID, packets[2].ConsumerChainID}
	require.True(t, slices.Contains(firstChainIdSet, "chain-0"))
	require.True(t, slices.Contains(firstChainIdSet, "chain-1"))
	require.True(t, slices.Contains(firstChainIdSet, "chain-2"))
	secondChainIdSet = []string{packets[3].ConsumerChainID, packets[4].ConsumerChainID, packets[5].ConsumerChainID}
	require.True(t, slices.Contains(secondChainIdSet, "chain-0"))
	require.True(t, slices.Contains(secondChainIdSet, "chain-1"))
	require.True(t, slices.Contains(secondChainIdSet, "chain-2"))
	thirdChainIdSet := []string{packets[6].ConsumerChainID, packets[7].ConsumerChainID, packets[8].ConsumerChainID}
	require.True(t, slices.Contains(thirdChainIdSet, "chain-5"))
	require.True(t, slices.Contains(thirdChainIdSet, "chain-6"))
	require.True(t, slices.Contains(thirdChainIdSet, "chain-7"))

	// Test the callback break functionality of the iterator
	packets = []types.SlashPacket{}
	providerKeeper.IteratePendingSlashPackets(ctx, func(packet types.SlashPacket) bool {
		packets = append(packets, packet)
		// Break after any of the third set of packets is seen
		return slices.Contains(thirdChainIdSet, packet.ConsumerChainID)
	})
	// Expect first two sets of packets to be seen, and one packet from the third set
	require.Equal(t, 7, len(packets))
}

// TestPendingSlashPacketDeletion tests the deletion of pending slash packets with assertion of FIFO ordering
func TestPendingSlashPacketDeletion(t *testing.T) {

	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(
		t, testkeeper.NewInMemKeeperParams(t))
	defer ctrl.Finish()

	now := time.Now()

	packets := []types.SlashPacket{}
	packets = append(packets, types.NewSlashPacket(now, "chain-0", testkeeper.GetNewSlashPacketData()))
	packets = append(packets, types.NewSlashPacket(now.Add(time.Hour).UTC(), "chain-1", testkeeper.GetNewSlashPacketData()))
	packets = append(packets, types.NewSlashPacket(now.Add(2*time.Hour).Local(), "chain-2", testkeeper.GetNewSlashPacketData()))
	packets = append(packets, types.NewSlashPacket(now.Add(3*time.Hour).UTC(), "chain-3", testkeeper.GetNewSlashPacketData()))
	packets = append(packets, types.NewSlashPacket(now.Add(4*time.Hour).Local(), "chain-4", testkeeper.GetNewSlashPacketData()))
	packets = append(packets, types.NewSlashPacket(now.Add(5*time.Hour).UTC(), "chain-5", testkeeper.GetNewSlashPacketData()))
	packets = append(packets, types.NewSlashPacket(now.Add(6*time.Hour), "chain-6", testkeeper.GetNewSlashPacketData()))

	// Instantiate shuffled copy of above slice
	shuffledPackets := append([]types.SlashPacket{}, packets...)
	rand.Seed(now.UnixNano())
	rand.Shuffle(len(shuffledPackets), func(i, j int) {
		shuffledPackets[i], shuffledPackets[j] = shuffledPackets[j], shuffledPackets[i]
	})

	// Queue 7 slash packets with various block times in random order
	for _, packet := range shuffledPackets {
		providerKeeper.QueuePendingSlashPacket(ctx, packet)
	}

	// Assert obtained order is decided upon via block time, not insertion order
	gotPackets := providerKeeper.GetAllPendingSlashPackets(ctx)
	for i, gotPacket := range gotPackets {
		expectedPacket := packets[i]
		require.Equal(t, expectedPacket, gotPacket)
	}

	// Delete packets 1, 3, 5 (0-indexed)
	providerKeeper.DeletePendingSlashPackets(ctx, gotPackets[1], gotPackets[3], gotPackets[5])

	// Assert deletion and ordering
	gotPackets = providerKeeper.GetAllPendingSlashPackets(ctx)
	require.Equal(t, 4, len(gotPackets))
	require.Equal(t, "chain-0", gotPackets[0].ConsumerChainID)
	// Packet 1 was deleted
	require.Equal(t, "chain-2", gotPackets[1].ConsumerChainID)
	// Packet 3 was deleted
	require.Equal(t, "chain-4", gotPackets[2].ConsumerChainID)
	// Packet 5 was deleted
	require.Equal(t, "chain-6", gotPackets[3].ConsumerChainID)
}

// TestSlashGasMeter tests the getter and setter for the slash gas meter
func TestSlashGasMeter(t *testing.T) {

	testCases := []struct {
		meterValue  sdk.Int
		shouldPanic bool
	}{
		{meterValue: sdk.NewInt(-7999999999999999999), shouldPanic: true},
		{meterValue: sdk.NewInt(-tmtypes.MaxTotalVotingPower - 1), shouldPanic: true},
		{meterValue: sdk.NewInt(-tmtypes.MaxTotalVotingPower), shouldPanic: false},
		{meterValue: sdk.NewInt(-50000000078987), shouldPanic: false},
		{meterValue: sdk.NewInt(-4237), shouldPanic: false},
		{meterValue: sdk.NewInt(0), shouldPanic: false},
		{meterValue: sdk.NewInt(1), shouldPanic: false},
		{meterValue: sdk.NewInt(4237897), shouldPanic: false},
		{meterValue: sdk.NewInt(500078078987), shouldPanic: false},
		{meterValue: sdk.NewInt(tmtypes.MaxTotalVotingPower), shouldPanic: false},
		{meterValue: sdk.NewInt(tmtypes.MaxTotalVotingPower + 1), shouldPanic: true},
		{meterValue: sdk.NewInt(7999974823991111199), shouldPanic: true},
	}

	for _, tc := range testCases {
		providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(
			t, testkeeper.NewInMemKeeperParams(t))
		defer ctrl.Finish()

		if tc.shouldPanic {
			require.Panics(t, func() {
				providerKeeper.SetSlashGasMeter(ctx, tc.meterValue)
			})
		} else {
			providerKeeper.SetSlashGasMeter(ctx, tc.meterValue)
			gotMeterValue := providerKeeper.GetSlashGasMeter(ctx)
			require.Equal(t, tc.meterValue, gotMeterValue)
		}
	}
}

// TestLastSlashGasReplenishTime tests the getter and setter for the last slash gas replenish time
func TestLastSlashGasReplenishTime(t *testing.T) {

	testCases := []time.Time{
		time.Now(),
		time.Now().Add(1 * time.Hour).UTC(),
		time.Now().Add(2 * time.Hour).Local(),
		time.Now().Add(3 * time.Hour).In(time.FixedZone("UTC-8", -8*60*60)),
		time.Now().Add(4 * time.Hour).Local(),
		time.Now().Add(-1 * time.Hour).UTC(),
		time.Now().Add(-2 * time.Hour).Local(),
		time.Now().Add(-3 * time.Hour).UTC(),
		time.Now().Add(-4 * time.Hour).Local(),
	}

	for _, tc := range testCases {
		providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(
			t, testkeeper.NewInMemKeeperParams(t))
		defer ctrl.Finish()

		providerKeeper.SetLastSlashGasReplenishTime(ctx, tc)
		gotTime := providerKeeper.GetLastSlashGasReplenishTime(ctx)
		// Time should be returned in UTC
		require.Equal(t, tc.UTC(), gotTime)
	}
}
