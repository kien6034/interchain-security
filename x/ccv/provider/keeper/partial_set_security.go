package keeper

import (
	errorsmod "cosmossdk.io/errors"
	abci "github.com/cometbft/cometbft/abci/types"
	tmprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/interchain-security/v4/x/ccv/provider/types"
	"sort"
)

type OptedInValidator struct {
	ProviderAddr types.ProviderConsAddress
	// block height the validator opted in at
	BlockHeight uint64
	// power the validator had when it opted in
	Power uint64
}

func (k Keeper) HandleOptIn(ctx sdk.Context, chainID string, providerAddr types.ProviderConsAddress, consumerKey *string) error {
	if !k.IsConsumerProposedOrRegistered(ctx, chainID) {
		return errorsmod.Wrapf(
			types.ErrUnknownConsumerChainId,
			"opting in to an unknown consumer chain, with id: %s", chainID)
	}

	validator, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, providerAddr.ToSdkConsAddr())
	if !found {
		return errorsmod.Wrapf(
			types.ErrNoValidatorProviderAddress,
			"could not find validator with consensus address: %s", providerAddr.ToSdkConsAddr().Bytes())
	}

	if k.IsToBeOptedOut(ctx, chainID, providerAddr) {
		// a validator to be opted in cancels out with a validator to be opted out
		k.DeleteToBeOptedOut(ctx, chainID, providerAddr)
	} else if !k.IsToBeOptedIn(ctx, chainID, providerAddr) && !k.IsOptedIn(ctx, chainID, providerAddr) {
		// a validator can only be set for opt in if it is not opted in and not already set for opt in
		k.SetToBeOptedIn(ctx, chainID, providerAddr)
	}

	if consumerKey != nil {
		consumerTMPublicKey, err := k.ParseConsumerKey(*consumerKey)
		if err != nil {
			return err
		}

		err = k.AssignConsumerKey(ctx, chainID, validator, consumerTMPublicKey)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) HandleOptOut(ctx sdk.Context, chainID string, providerAddr types.ProviderConsAddress) error {
	if _, found := k.GetConsumerClientId(ctx, chainID); !found {
		// A validator can only opt out from a running chain. We check this by checking the consumer client id, because
		// `SetConsumerClientId` is set when the chain starts in `CreateConsumerClientInCachedCtx` of `BeginBlockInit`.
		return errorsmod.Wrapf(
			types.ErrUnknownConsumerChainId,
			"opting out of an unknown or not running consumer chain, with id: %s", chainID)
	}

	if k.IsToBeOptedIn(ctx, chainID, providerAddr) {
		// a validator to be opted out cancels out a validator to be opted in
		k.DeleteToBeOptedIn(ctx, chainID, providerAddr)
	} else if !k.IsToBeOptedOut(ctx, chainID, providerAddr) && k.IsOptedIn(ctx, chainID, providerAddr) {
		// a validator can only be set for opt out if it is opted in and not already set for opt out
		k.SetToBeOptedOut(ctx, chainID, providerAddr)
	}

	return nil
}

// getValAddressAndPublicKey is a helper function that returns the `ValAddress` and the public key of
// the corresponding validator
func (k Keeper) getValAddressAndPublicKey(ctx sdk.Context, addr types.ProviderConsAddress,
) (sdk.ValAddress, tmprotocrypto.PublicKey, error) {
	validator, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, addr.ToSdkConsAddr())
	if !found {
		return sdk.ValAddress{}, tmprotocrypto.PublicKey{}, stakingtypes.ErrNoValidatorFound
	}

	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return sdk.ValAddress{}, tmprotocrypto.PublicKey{}, err
	}

	pubKey := tmprotocrypto.PublicKey{
		Sum: &tmprotocrypto.PublicKey_Ed25519{
			Ed25519: consAddr.Bytes(),
		},
	}
	return validator.GetOperator(), pubKey, nil
}

// ComputeNextValidators computes the next validator set that is responsible for validating on a consumer chain.
// The returned opted-in validators by `ComputeNextValidators` constitute the next `currentValidators`.
func (k Keeper) ComputeNextValidators(ctx sdk.Context,
	currentValidators []OptedInValidator,
	validatorAddressesToAdd []types.ProviderConsAddress,
	validatorAddressesToRemove []types.ProviderConsAddress,
) []OptedInValidator {
	isRemoved := make(map[string]bool)
	for _, val := range validatorAddressesToRemove {
		isRemoved[val.ToSdkConsAddr().String()] = true
	}

	var out []OptedInValidator
	for _, val := range currentValidators {
		if isRemoved[val.ProviderAddr.ToSdkConsAddr().String()] {
			continue
		}
		valAddress, _, err := k.getValAddressAndPublicKey(ctx, val.ProviderAddr)
		if err != nil {
			continue
		}

		val.Power = uint64(k.stakingKeeper.GetLastValidatorPower(ctx, valAddress))
		if val.Power == 0 {
			continue
		}
		out = append(out, val)
	}

	for _, addr := range validatorAddressesToAdd {
		valAddress, _, err := k.getValAddressAndPublicKey(ctx, addr)
		if err != nil {
			continue
		}

		validator, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, addr.ToSdkConsAddr())
		if !found {
			continue
		}
		if !validator.IsBonded() {
			continue
		}
		power := uint64(k.stakingKeeper.GetLastValidatorPower(ctx, valAddress))

		out = append(out, OptedInValidator{ProviderAddr: addr, BlockHeight: uint64(ctx.BlockHeight()), Power: power})
	}

	return out
}

// ComputeValidatorUpdates computes the validator updates needed to be sent to the consumer chain to capture
// the newly opted-in and opted-out validators, as well as validators that unbonded.
func (k Keeper) ComputeValidatorUpdates(ctx sdk.Context,
	currentValidators []OptedInValidator,
	validatorAddressesToAdd []types.ProviderConsAddress,
	validatorAddressesToRemove []types.ProviderConsAddress,
) []abci.ValidatorUpdate {
	var m = make(map[string]abci.ValidatorUpdate)

	for _, val := range currentValidators {
		valAddress, pubKey, err := k.getValAddressAndPublicKey(ctx, val.ProviderAddr)
		if err != nil {
			continue
		}

		validator, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, val.ProviderAddr.ToSdkConsAddr())
		if !found {
			continue
		}

		if val.Power == uint64(k.stakingKeeper.GetLastValidatorPower(ctx, validator.GetOperator())) {
			continue
		}

		// if `val` has unbonded, its `GetLastValidatorPower` power returns 0.
		m[pubKey.String()] = abci.ValidatorUpdate{
			PubKey: pubKey,
			Power:  k.stakingKeeper.GetLastValidatorPower(ctx, valAddress),
		}
	}

	for _, addr := range validatorAddressesToAdd {
		valAddress, pubKey, err := k.getValAddressAndPublicKey(ctx, addr)
		if err != nil {
			continue
		}

		validator, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, addr.ToSdkConsAddr())
		if !found {
			continue
		}

		// if a validator is in the active set, we do not add it
		if !validator.IsBonded() {
			continue
		}

		m[pubKey.String()] = abci.ValidatorUpdate{
			PubKey: pubKey,
			Power:  k.stakingKeeper.GetLastValidatorPower(ctx, valAddress),
		}
	}

	for _, addr := range validatorAddressesToRemove {
		_, pubKey, err := k.getValAddressAndPublicKey(ctx, addr)
		if err != nil {

		}

		m[pubKey.String()] = abci.ValidatorUpdate{
			PubKey: pubKey,
			Power:  0,
		}

	}

	var out []abci.ValidatorUpdate
	for _, update := range m {
		out = append(out, update)
	}

	// Similarly to `AccumulateChanges`, we sort validators for determinism.
	sort.Slice(out, func(i, j int) bool {
		if out[i].Power != out[j].Power {
			return out[i].Power > out[j].Power
		}
		return out[i].PubKey.String() > out[j].PubKey.String()
	})

	return out
}

// ResetCurrentValidators resets the opted-in validators with the newest set that was computed by
// `ComputePartialSetValidatorUpdates` and hence this method should only be called  after
// `ComputePartialSetValidatorUpdates` has complete. Also, clears all the `ToBeOptedIn` and `ToBeOptedOut` sets.
func (k Keeper) ResetCurrentValidators(ctx sdk.Context, chainID string, nextValidators []OptedInValidator) {
	k.DeleteAllOptedIn(ctx, chainID)
	for _, val := range nextValidators {
		k.SetOptedIn(ctx, chainID, val.ProviderAddr, val.BlockHeight, val.Power)
	}

	k.DeleteAllToBeOptedIn(ctx, chainID)
	k.DeleteAllToBeOptedOut(ctx, chainID)
}
