package types

import (
	"math/big"

	"github.com/pokt-network/pocket/shared/crypto"
	"github.com/pokt-network/pocket/shared/modules"
	"github.com/pokt-network/pocket/utility/types"
)

// TODO (Research) is there anyway to not have to name these protobuf files uniquely?
//	not a fan of <module_name>_config/genesis.go would rather just config/genesis.go

func (x *PersistenceGenesisState) GetAccs() []modules.Account {
	return AccToAccInterface(x.GetAccounts())
}

func (x *PersistenceGenesisState) GetAccPools() []modules.Account {
	return AccToAccInterface(x.GetPools())
}

func (x *PersistenceGenesisState) GetApps() []modules.Actor {
	return ActorsToActorsInterface(x.GetApplications())
}

func (x *PersistenceGenesisState) GetVals() []modules.Actor {
	return ActorsToActorsInterface(x.GetValidators())
}

func (x *PersistenceGenesisState) GetFish() []modules.Actor {
	return ActorsToActorsInterface(x.GetFishermen())
}

func (x *PersistenceGenesisState) GetNodes() []modules.Actor {
	return ActorsToActorsInterface(x.GetServiceNodes())
}

func (x *PersistenceGenesisState) GetParameters() modules.Params {
	return x.GetParams()
}

// RESEARCH(olshansky): AFAIK this is the only way to convert slice of structs into interface - O(n)
// https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces
func ActorsToActorsInterface(a []*Actor) (actorI []modules.Actor) {
	for _, actor := range a {
		actorI = append(actorI, actor)
	}
	return
}

func AccToAccInterface(a []*Account) (accI []modules.Account) {
	for _, acc := range a {
		accI = append(accI, acc)
	}
	return
}

var (
	DefaultParamsOwner, _ = crypto.NewPrivateKey("ff538589deb7f28bbce1ba68b37d2efc0eaa03204b36513cf88422a875559e38d6cbe0430ddd85a5e48e0c99ef3dea47bf0d1a83c6e6ad1640f72201dc8a0120")
)

func DefaultParams() *Params { // TODO this is just a test / demo artifact and should be deprecated by genesis file
	return &Params{
		BlocksPerSession:                         4,
		AppMinimumStake:                          types.BigIntToString(big.NewInt(15000000000)),
		AppMaxChains:                             15,
		AppBaselineStakeRate:                     100,
		AppStakingAdjustment:                     0,
		AppUnstakingBlocks:                       2016,
		AppMinimumPauseBlocks:                    4,
		AppMaxPauseBlocks:                        672,
		ServiceNodeMinimumStake:                  types.BigIntToString(big.NewInt(15000000000)),
		ServiceNodeMaxChains:                     15,
		ServiceNodeUnstakingBlocks:               2016,
		ServiceNodeMinimumPauseBlocks:            4,
		ServiceNodeMaxPauseBlocks:                672,
		ServiceNodesPerSession:                   24,
		FishermanMinimumStake:                    types.BigIntToString(big.NewInt(15000000000)),
		FishermanMaxChains:                       15,
		FishermanUnstakingBlocks:                 2016,
		FishermanMinimumPauseBlocks:              4,
		FishermanMaxPauseBlocks:                  672,
		ValidatorMinimumStake:                    types.BigIntToString(big.NewInt(15000000000)),
		ValidatorUnstakingBlocks:                 2016,
		ValidatorMinimumPauseBlocks:              4,
		ValidatorMaxPauseBlocks:                  672,
		ValidatorMaximumMissedBlocks:             5,
		ValidatorMaxEvidenceAgeInBlocks:          8,
		ProposerPercentageOfFees:                 10,
		MissedBlocksBurnPercentage:               1,
		DoubleSignBurnPercentage:                 5,
		MessageDoubleSignFee:                     types.BigIntToString(big.NewInt(10000)),
		MessageSendFee:                           types.BigIntToString(big.NewInt(10000)),
		MessageStakeFishermanFee:                 types.BigIntToString(big.NewInt(10000)),
		MessageEditStakeFishermanFee:             types.BigIntToString(big.NewInt(10000)),
		MessageUnstakeFishermanFee:               types.BigIntToString(big.NewInt(10000)),
		MessagePauseFishermanFee:                 types.BigIntToString(big.NewInt(10000)),
		MessageUnpauseFishermanFee:               types.BigIntToString(big.NewInt(10000)),
		MessageFishermanPauseServiceNodeFee:      types.BigIntToString(big.NewInt(10000)),
		MessageTestScoreFee:                      types.BigIntToString(big.NewInt(10000)),
		MessageProveTestScoreFee:                 types.BigIntToString(big.NewInt(10000)),
		MessageStakeAppFee:                       types.BigIntToString(big.NewInt(10000)),
		MessageEditStakeAppFee:                   types.BigIntToString(big.NewInt(10000)),
		MessageUnstakeAppFee:                     types.BigIntToString(big.NewInt(10000)),
		MessagePauseAppFee:                       types.BigIntToString(big.NewInt(10000)),
		MessageUnpauseAppFee:                     types.BigIntToString(big.NewInt(10000)),
		MessageStakeValidatorFee:                 types.BigIntToString(big.NewInt(10000)),
		MessageEditStakeValidatorFee:             types.BigIntToString(big.NewInt(10000)),
		MessageUnstakeValidatorFee:               types.BigIntToString(big.NewInt(10000)),
		MessagePauseValidatorFee:                 types.BigIntToString(big.NewInt(10000)),
		MessageUnpauseValidatorFee:               types.BigIntToString(big.NewInt(10000)),
		MessageStakeServiceNodeFee:               types.BigIntToString(big.NewInt(10000)),
		MessageEditStakeServiceNodeFee:           types.BigIntToString(big.NewInt(10000)),
		MessageUnstakeServiceNodeFee:             types.BigIntToString(big.NewInt(10000)),
		MessagePauseServiceNodeFee:               types.BigIntToString(big.NewInt(10000)),
		MessageUnpauseServiceNodeFee:             types.BigIntToString(big.NewInt(10000)),
		MessageChangeParameterFee:                types.BigIntToString(big.NewInt(10000)),
		AclOwner:                                 DefaultParamsOwner.Address().String(),
		BlocksPerSessionOwner:                    DefaultParamsOwner.Address().String(),
		AppMinimumStakeOwner:                     DefaultParamsOwner.Address().String(),
		AppMaxChainsOwner:                        DefaultParamsOwner.Address().String(),
		AppBaselineStakeRateOwner:                DefaultParamsOwner.Address().String(),
		AppStakingAdjustmentOwner:                DefaultParamsOwner.Address().String(),
		AppUnstakingBlocksOwner:                  DefaultParamsOwner.Address().String(),
		AppMinimumPauseBlocksOwner:               DefaultParamsOwner.Address().String(),
		AppMaxPausedBlocksOwner:                  DefaultParamsOwner.Address().String(),
		ServiceNodeMinimumStakeOwner:             DefaultParamsOwner.Address().String(),
		ServiceNodeMaxChainsOwner:                DefaultParamsOwner.Address().String(),
		ServiceNodeUnstakingBlocksOwner:          DefaultParamsOwner.Address().String(),
		ServiceNodeMinimumPauseBlocksOwner:       DefaultParamsOwner.Address().String(),
		ServiceNodeMaxPausedBlocksOwner:          DefaultParamsOwner.Address().String(),
		ServiceNodesPerSessionOwner:              DefaultParamsOwner.Address().String(),
		FishermanMinimumStakeOwner:               DefaultParamsOwner.Address().String(),
		FishermanMaxChainsOwner:                  DefaultParamsOwner.Address().String(),
		FishermanUnstakingBlocksOwner:            DefaultParamsOwner.Address().String(),
		FishermanMinimumPauseBlocksOwner:         DefaultParamsOwner.Address().String(),
		FishermanMaxPausedBlocksOwner:            DefaultParamsOwner.Address().String(),
		ValidatorMinimumStakeOwner:               DefaultParamsOwner.Address().String(),
		ValidatorUnstakingBlocksOwner:            DefaultParamsOwner.Address().String(),
		ValidatorMinimumPauseBlocksOwner:         DefaultParamsOwner.Address().String(),
		ValidatorMaxPausedBlocksOwner:            DefaultParamsOwner.Address().String(),
		ValidatorMaximumMissedBlocksOwner:        DefaultParamsOwner.Address().String(),
		ValidatorMaxEvidenceAgeInBlocksOwner:     DefaultParamsOwner.Address().String(),
		ProposerPercentageOfFeesOwner:            DefaultParamsOwner.Address().String(),
		MissedBlocksBurnPercentageOwner:          DefaultParamsOwner.Address().String(),
		DoubleSignBurnPercentageOwner:            DefaultParamsOwner.Address().String(),
		MessageDoubleSignFeeOwner:                DefaultParamsOwner.Address().String(),
		MessageSendFeeOwner:                      DefaultParamsOwner.Address().String(),
		MessageStakeFishermanFeeOwner:            DefaultParamsOwner.Address().String(),
		MessageEditStakeFishermanFeeOwner:        DefaultParamsOwner.Address().String(),
		MessageUnstakeFishermanFeeOwner:          DefaultParamsOwner.Address().String(),
		MessagePauseFishermanFeeOwner:            DefaultParamsOwner.Address().String(),
		MessageUnpauseFishermanFeeOwner:          DefaultParamsOwner.Address().String(),
		MessageFishermanPauseServiceNodeFeeOwner: DefaultParamsOwner.Address().String(),
		MessageTestScoreFeeOwner:                 DefaultParamsOwner.Address().String(),
		MessageProveTestScoreFeeOwner:            DefaultParamsOwner.Address().String(),
		MessageStakeAppFeeOwner:                  DefaultParamsOwner.Address().String(),
		MessageEditStakeAppFeeOwner:              DefaultParamsOwner.Address().String(),
		MessageUnstakeAppFeeOwner:                DefaultParamsOwner.Address().String(),
		MessagePauseAppFeeOwner:                  DefaultParamsOwner.Address().String(),
		MessageUnpauseAppFeeOwner:                DefaultParamsOwner.Address().String(),
		MessageStakeValidatorFeeOwner:            DefaultParamsOwner.Address().String(),
		MessageEditStakeValidatorFeeOwner:        DefaultParamsOwner.Address().String(),
		MessageUnstakeValidatorFeeOwner:          DefaultParamsOwner.Address().String(),
		MessagePauseValidatorFeeOwner:            DefaultParamsOwner.Address().String(),
		MessageUnpauseValidatorFeeOwner:          DefaultParamsOwner.Address().String(),
		MessageStakeServiceNodeFeeOwner:          DefaultParamsOwner.Address().String(),
		MessageEditStakeServiceNodeFeeOwner:      DefaultParamsOwner.Address().String(),
		MessageUnstakeServiceNodeFeeOwner:        DefaultParamsOwner.Address().String(),
		MessagePauseServiceNodeFeeOwner:          DefaultParamsOwner.Address().String(),
		MessageUnpauseServiceNodeFeeOwner:        DefaultParamsOwner.Address().String(),
		MessageChangeParameterFeeOwner:           DefaultParamsOwner.Address().String(),
	}
}