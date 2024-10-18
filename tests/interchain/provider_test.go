package interchain

import (
	"cosmos/interchain-security/tests/interchain/chainsuite"
	"testing"
	"time"

	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	providertypes "github.com/cosmos/interchain-security/v6/x/ccv/provider/types"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"

	"github.com/stretchr/testify/suite"
)

func TestProviderSuite(t *testing.T) {
	s := &ProviderSuite{}

	suite.Run(t, s)
}

func (s *ProviderSuite) TestConsumerAdditionProposalPositiveCases() {
	// submit message without setting all params
	chainNameRegistered := "consumerAdditionRegistered"
	proposalMsg := msgCreateConsumer(chainNameRegistered, nil, nil)
	s.Require().NoError(s.Provider.ExecuteProposalMsg(s.GetContext(), proposalMsg, chainsuite.GovModuleAddress, chainNameRegistered, cosmos.ProposalVoteYes, govv1.StatusPassed, false))
	consumerChain, err := s.Provider.GetConsumerChainByChainId(s.GetContext(), chainNameRegistered)
	s.Require().NoError(err)
	s.Require().Equal(providertypes.CONSUMER_PHASE_REGISTERED.String(), consumerChain.Phase)

	// submit message with future spawn time
	chainNameInitialized := "consumerAdditionInitialized"
	proposalMsg = msgCreateConsumer(chainNameInitialized, consumerInitParamsTemplate(time.Now().Add(time.Hour)), powerShapingParamsTemplate())
	s.Require().NoError(s.Provider.ExecuteProposalMsg(s.GetContext(), proposalMsg, chainsuite.GovModuleAddress, chainNameInitialized, cosmos.ProposalVoteYes, govv1.StatusPassed, false))
	consumerChain, err = s.Provider.GetConsumerChainByChainId(s.GetContext(), chainNameInitialized)
	s.Require().NoError(err)
	s.Require().Equal(providertypes.CONSUMER_PHASE_INITIALIZED.String(), consumerChain.Phase)

	// submit message and wait for spawn time
	chainNameLaunched := "consumerAdditionLaunched"
	spawnTime := 60 * time.Second
	consumerInitParams := consumerInitParamsTemplate(time.Now().Add(spawnTime))
	proposalMsg = msgCreateConsumer(chainNameLaunched, consumerInitParams, powerShapingParamsTemplate())
	s.Require().NoError(s.Provider.ExecuteProposalMsg(s.GetContext(), proposalMsg, chainsuite.GovModuleAddress, chainNameLaunched, cosmos.ProposalVoteYes, govv1.StatusPassed, false))
	consumerChain, err = s.Provider.GetConsumerChainByChainId(s.GetContext(), chainNameLaunched)
	s.Require().NoError(err)
	s.Require().NoError(s.Provider.OptIn(s.GetContext(), consumerChain.ConsumerID, 0))
	// wait for spawn time and then for one block for chain to be launched in begin blocker
	time.Sleep(spawnTime)
	s.Require().NoError(testutil.WaitForBlocks(s.GetContext(), 1, s.Provider))
	consumerChain, err = s.Provider.GetConsumerChainByChainId(s.GetContext(), chainNameLaunched)
	s.Require().NoError(err)
	s.Require().Equal(providertypes.CONSUMER_PHASE_LAUNCHED.String(), consumerChain.Phase)
	// get consumer genesis data
	consumerGenesis, err := s.Provider.GetConsumerGenesis(s.GetContext(), consumerChain.ConsumerID)
	s.Require().NoError(err)
	s.Require().Equal(consumerInitParams.ConsumerRedistributionFraction, consumerGenesis.Params.ConsumerRedistributionFraction)
}

func (s *ProviderSuite) TestConsumerAdditionProposalNegativeCases() {
	chainNameReject := "consumerAdditionReject"
	chainNameFailed := "consumerAdditionFailed"

	// reject consumer proposal and check that chain is not added
	proposalMsg := msgCreateConsumer(chainNameReject, nil, nil)
	s.Require().NoError(s.Provider.ExecuteProposalMsg(s.GetContext(), proposalMsg, chainsuite.GovModuleAddress, chainNameReject, cosmos.ProposalVoteNo, govv1.StatusRejected, false))
	_, err := s.Provider.GetConsumerChainByChainId(s.GetContext(), chainNameReject)
	s.Require().Error(err)

	// cannot create a Top N chain using the `MsgCreateConsumer` message
	proposalMsg = msgCreateConsumer(chainNameFailed, nil, powerShapingParamsTemplate())
	proposalMsg.PowerShapingParameters.Top_N = 100
	s.Require().Error(s.Provider.ExecuteProposalMsg(s.GetContext(), proposalMsg, chainsuite.GovModuleAddress, chainNameFailed, cosmos.ProposalVoteYes, govv1.StatusFailed, false))
	_, err = s.Provider.GetConsumerChainByChainId(s.GetContext(), chainNameFailed)
	s.Require().Error(err)

	// empty metadata
	proposalMsg = msgCreateConsumer(chainNameFailed, nil, nil)
	proposalMsg.Metadata = providertypes.ConsumerMetadata{}
	s.Require().Error(s.Provider.ExecuteProposalMsg(s.GetContext(), proposalMsg, chainsuite.GovModuleAddress, chainNameFailed, cosmos.ProposalVoteYes, govv1.StatusFailed, false))
	_, err = s.Provider.GetConsumerChainByChainId(s.GetContext(), chainNameFailed)
	s.Require().Error(err)
}

func (s *ProviderSuite) TestConsumerUpgradeProposal() {
	chainNameUpgrade := "consumerUpgrade"
	initParams := consumerInitParamsTemplate(time.Now().Add(time.Hour))
	powerShapingParams := powerShapingParamsTemplate()
	proposalMsg := msgCreateConsumer(chainNameUpgrade, initParams, powerShapingParams)
	s.Require().NoError(s.Provider.ExecuteProposalMsg(s.GetContext(), proposalMsg, chainsuite.GovModuleAddress, chainNameUpgrade, cosmos.ProposalVoteYes, govv1.StatusPassed, false))
	consumerChain, err := s.Provider.GetConsumerChainByChainId(s.GetContext(), chainNameUpgrade)
	s.Require().NoError(err)
	s.Require().Equal(0, consumerChain.TopN)

	powerShapingParams.Top_N = 80
	upgradeMsg := &providertypes.MsgUpdateConsumer{
		Owner:                    chainsuite.GovModuleAddress,
		ConsumerId:               consumerChain.ConsumerID,
		NewOwnerAddress:          chainsuite.GovModuleAddress,
		InitializationParameters: initParams,
		PowerShapingParameters:   powerShapingParams,
	}
	s.Require().NoError(s.Provider.ExecuteProposalMsg(s.GetContext(), upgradeMsg, chainsuite.GovModuleAddress, chainNameUpgrade, cosmos.ProposalVoteYes, govv1.StatusPassed, false))
	updatedChain, err := s.Provider.GetConsumerChain(s.GetContext(), consumerChain.ConsumerID)
	s.Require().NoError(err)
	s.Require().Equal(80, updatedChain.PowerShapingParams.TopN)
}

func (s *ProviderSuite) TestConsumerRemovalProposal() {
	// add consumer that will be later removed
	chainNameRemove := "consumerRemove"
	initParams := consumerInitParamsTemplate(time.Now().Add(time.Hour))
	powerShapingParams := powerShapingParamsTemplate()
	proposalMsg := msgCreateConsumer(chainNameRemove, initParams, powerShapingParams)
	s.Require().NoError(s.Provider.ExecuteProposalMsg(s.GetContext(), proposalMsg, chainsuite.GovModuleAddress, chainNameRemove, cosmos.ProposalVoteYes, govv1.StatusPassed, false))
	consumerChain, err := s.Provider.GetConsumerChainByChainId(s.GetContext(), chainNameRemove)
	s.Require().NoError(err)

	removeConsumerMsg := &providertypes.MsgRemoveConsumer{
		ConsumerId: consumerChain.ConsumerID,
		Owner:      chainsuite.GovModuleAddress,
	}
	// cannot be removed if not in phase CONSUMER_PHASE_LAUNCHED
	s.Require().Error(s.Provider.ExecuteProposalMsg(s.GetContext(), removeConsumerMsg, chainsuite.GovModuleAddress, chainNameRemove, cosmos.ProposalVoteYes, govv1.StatusPassed, false))

	// update spawn time
	spawTimeFromNow := 10 * time.Second
	initParams.SpawnTime = time.Now().Add(spawTimeFromNow)
	powerShapingParams.Top_N = 60
	upgradeMsg := &providertypes.MsgUpdateConsumer{
		Owner:                    chainsuite.GovModuleAddress,
		ConsumerId:               consumerChain.ConsumerID,
		NewOwnerAddress:          chainsuite.GovModuleAddress,
		InitializationParameters: initParams,
		PowerShapingParameters:   powerShapingParams,
	}
	s.Require().NoError(s.Provider.ExecuteProposalMsg(s.GetContext(), upgradeMsg, chainsuite.GovModuleAddress, chainNameRemove, cosmos.ProposalVoteYes, govv1.StatusPassed, false))
	time.Sleep(spawTimeFromNow)
	s.Require().NoError(testutil.WaitForBlocks(s.GetContext(), 1, s.Provider))
	chainToRemove, err := s.Provider.GetConsumerChain(s.GetContext(), consumerChain.ConsumerID)
	s.Require().NoError(err)
	s.Require().Equal(providertypes.CONSUMER_PHASE_LAUNCHED.String(), chainToRemove.Phase)

	// remove consumer successfully
	s.Require().NoError(s.Provider.ExecuteProposalMsg(s.GetContext(), removeConsumerMsg, chainsuite.GovModuleAddress, chainNameRemove, cosmos.ProposalVoteYes, govv1.StatusPassed, false))
	chainToRemove, err = s.Provider.GetConsumerChain(s.GetContext(), consumerChain.ConsumerID)
	s.Require().NoError(err)
	s.Require().Equal(providertypes.CONSUMER_PHASE_STOPPED.String(), chainToRemove.Phase)
	time.Sleep(chainsuite.ProviderUnbondingTime)
	s.Require().NoError(testutil.WaitForBlocks(s.GetContext(), 1, s.Provider))
	chainToRemove, err = s.Provider.GetConsumerChain(s.GetContext(), consumerChain.ConsumerID)
	s.Require().NoError(err)
	s.Require().Equal(providertypes.CONSUMER_PHASE_DELETED.String(), chainToRemove.Phase)
}
