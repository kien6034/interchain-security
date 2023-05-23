package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibctmtypes "github.com/cosmos/ibc-go/v4/modules/light-clients/07-tendermint/types"
	providertypes "github.com/cosmos/interchain-security/x/ccv/provider/types"
	ccvtypes "github.com/cosmos/interchain-security/x/ccv/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	ccvProviderKeeper     Keeper
	stakingKeeper         ccvtypes.StakingKeeper
	ccvProviderParamSpace paramtypes.Subspace
}

// NewMigrator returns a new Migrator.
func NewMigrator(ccvProviderKeeper Keeper, stakingKeeper ccvtypes.StakingKeeper,
	ccvProviderParamSpace paramtypes.Subspace) Migrator {
	return Migrator{ccvProviderKeeper: ccvProviderKeeper, ccvProviderParamSpace: ccvProviderParamSpace}
}

func (m Migrator) Migratev1p0To1p3(ctx sdk.Context) error {
	// Migrate params
	MigrateParamsv1p0To1p3(ctx,
		m.ccvProviderParamSpace,
		// See https://github.com/cosmos/interchain-security/blob/7861804cb311507ec6aebebbfad60ea42eb8ed4b/x/ccv/provider/keeper/params.go#L84
		// The v1.1.0-multiden version of ICS hardcodes this param as 10 of bond type: k.stakingKeeper.BondDenom(ctx).
		// Here we use the same starting value, but the param can now be changed through governance.
		sdk.NewCoin(m.stakingKeeper.BondDenom(ctx), sdk.NewInt(10000000)),
	)

	return nil
}

// MigrateParamsv1p0To1p3 migrates the provider CCV module params from v1.0.0 to v1.3.0,
// setting default values for new params.
func MigrateParamsv1p0To1p3(ctx sdk.Context, paramsSubspace paramtypes.Subspace, consumerRewardDenomRegistrationFee sdk.Coin) {
	// Get old params
	var templateClient ibctmtypes.ClientState
	paramsSubspace.Get(ctx, providertypes.KeyTemplateClient, &templateClient)
	var trustingPeriodFraction string
	paramsSubspace.Get(ctx, providertypes.KeyTrustingPeriodFraction, &trustingPeriodFraction)
	var ccvTimeoutPeriod time.Duration
	paramsSubspace.Get(ctx, ccvtypes.KeyCCVTimeoutPeriod, &ccvTimeoutPeriod)
	var initTimeoutPeriod time.Duration
	paramsSubspace.Get(ctx, providertypes.KeyInitTimeoutPeriod, &initTimeoutPeriod)
	var vscTimeoutPeriod time.Duration
	paramsSubspace.Get(ctx, providertypes.KeyVscTimeoutPeriod, &vscTimeoutPeriod)
	var slashMeterReplenishPeriod time.Duration
	paramsSubspace.Get(ctx, providertypes.KeySlashMeterReplenishPeriod, &slashMeterReplenishPeriod)
	var slashMeterReplenishFraction string
	paramsSubspace.Get(ctx, providertypes.KeySlashMeterReplenishFraction, &slashMeterReplenishFraction)
	var maxThrottledPackets int64
	paramsSubspace.Get(ctx, providertypes.KeyMaxThrottledPackets, &maxThrottledPackets)

	// Recycle old params, set new param to input value
	newParams := providertypes.NewParams(
		&templateClient,
		trustingPeriodFraction,
		ccvTimeoutPeriod,
		initTimeoutPeriod,
		vscTimeoutPeriod,
		slashMeterReplenishPeriod,
		slashMeterReplenishFraction,
		maxThrottledPackets,
		consumerRewardDenomRegistrationFee,
	)

	// Persist new params
	paramsSubspace.SetParamSet(ctx, &newParams)
}