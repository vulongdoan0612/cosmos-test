package keeper

import (
	"errors"
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// TxFeeChecker check if the provided fee is enough and returns the effective fee and tx priority,
// the effective fee should be deducted later, and the priority should be returned in abci response.

// DeductFeeDecorator deducts fees from the first signer of the tx
// If the first signer does not have the funds to pay for the fees, return with InsufficientFunds error
// Call next AnteHandler if fees successfully deducted
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator

type DeductFeeDecorator struct {
	accountKeeper  ante.AccountKeeper
	bankKeeper     types.BankKeeper
	feegrantKeeper ante.FeegrantKeeper
	txFeesChecker  Keeper
}

func NewDeductFeeDecorator_VuChain_On(ak ante.AccountKeeper, bk types.BankKeeper, fk ante.FeegrantKeeper, tfc Keeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		accountKeeper:  ak,
		bankKeeper:     bk,
		feegrantKeeper: fk,
		txFeesChecker:  tfc,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	feeTx, ok := tx.(sdk.FeeTx)

	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	var (
		priority int64
		// err      error
	)
	// feeCoins := feeTx.GetFee()

	baseDenom, found := dfd.txFeesChecker.GetBaseDenom(ctx)
	if !found {

		fmt.Print("ds")
	}
	fee := feeTx.GetFee()
	reqGas := feeTx.GetGas()
	reqGasInt := sdk.NewInt(int64(reqGas))
	minGasFee := reqGasInt.Mul(baseDenom.GetDenom().Amount)

	feeAmount := fee.AmountOf("stake").Uint64()

	feeAmountInt := sdk.NewInt(int64(feeAmount))

	if minGasFee.LT(feeAmountInt) {
		fmt.Print("  " ,"minGasFee:", minGasFee, "/", "feeTx.GetGas():", feeTx.GetGas(), "/", "feeAmount:", feeAmount, "/", "reqGas:", reqGas, "/", "baseDenom:", baseDenom.GetDenom().Amount)
	}

	if err := dfd.checkDeductFee(ctx, tx, fee, dfd.txFeesChecker); err != nil {
		return ctx, err
	}
	newCtx := ctx.WithPriority(priority)

	return next(newCtx, tx, simulate)
}

// DeductFees deducts fees from the given account.
func (dfd DeductFeeDecorator) checkDeductFee(ctx sdk.Context, sdkTx sdk.Tx, fee sdk.Coins, k Keeper) error {
	feeTx, ok := sdkTx.(sdk.FeeTx)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.accountKeeper.GetModuleAddress(types.FeeCollectorName); addr == nil {
		return fmt.Errorf("fee collector module account (%s) has not been set", types.FeeCollectorName)
	}
	// if fee[0].Denom != "ap" {
	// 	return sdkerrors.ErrInvalidCoins.Wrapf("fee must be AP, your fee is (%s). Try again", fee[0].Denom)
	// }

	feePayer := feeTx.FeePayer()
	feeGranter := feeTx.FeeGranter()
	deductFeesFrom := feePayer

	// if feegranter set deduct fee from feegranter account.
	// this works with only when feegrant enabled.
	if feeGranter != nil {
		if dfd.feegrantKeeper == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("fee grants are not enabled")
		} else if !feeGranter.Equals(feePayer) {
			err := dfd.feegrantKeeper.UseGrantedFees(ctx, feeGranter, feePayer, fee, sdkTx.GetMsgs())
			if err != nil {
				return sdkerrors.Wrapf(err, "%s does not allow to pay fees for %s", feeGranter, feePayer)
			}
		}

		deductFeesFrom = feeGranter
	}
	deductFeesFromAcc := dfd.accountKeeper.GetAccount(ctx, deductFeesFrom)
	if deductFeesFromAcc == nil {
		return sdkerrors.ErrUnknownAddress.Wrapf("fee payer address: %s does not exist", deductFeesFrom)
	}

	// deduct the fees
	if !fee.IsZero() {
		err := DeductFees(k, dfd.bankKeeper, ctx, deductFeesFromAcc, fee)
		if err != nil {
			return err
		}
	}

	events := sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeTx,
			sdk.NewAttribute(sdk.AttributeKeyFee, fee.String()),
			sdk.NewAttribute(sdk.AttributeKeyFeePayer, deductFeesFrom.String()),
		),
	}

	ctx.EventManager().EmitEvents(events)

	return nil
}

func (k Keeper) IsSufficientFee(ctx sdk.Context, minBaseGasPrice int, gasRequested uint64, feeCoin sdk.Coin) error {
	baseDenom, err := k.GetBaseDenom(ctx)
	if err {

	}
	// Determine the required fees by multiplying the required minimum gas
	// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
	minBaseGasPriceDec := sdk.NewDec(int64(minBaseGasPrice))
	gasRequestedDec := sdk.NewDec(int64(gasRequested))
	requiredFeeDec := minBaseGasPriceDec.Mul(gasRequestedDec)
	requiredBaseFee := sdk.NewCoin(baseDenom.String(), requiredFeeDec.Ceil().RoundInt())

	// check to ensure that the convertedFee should always be greater than or equal to the requireBaseFee
	errorsmod.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s which converts to NONE. required: %s", feeCoin, requiredBaseFee)

	return nil
}

// DeductFees deducts fees from the given account.
func DeductFees(txFeeKeeper Keeper, bankKeeper types.BankKeeper, ctx sdk.Context, acc types.AccountI, fees sdk.Coins,
) error {
	if !fees.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %s", fees)
	}
	baseDenom, found := txFeeKeeper.GetBaseDenom(ctx)

	if !found {
		return errors.New("base denom not found")
	}
	// checks if input fee is uOSMO (assumes only one fee token exists in the fees array (as per the check in mempoolFeeDecorator))
	if fees[0].Denom == baseDenom.GetDenom().Denom {

		// sends to FeeCollectorName module account, which sends to staking rewards
		err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.FeeCollectorName, fees)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
		}
	}

	return nil
}
