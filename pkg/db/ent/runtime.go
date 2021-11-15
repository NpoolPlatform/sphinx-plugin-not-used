// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/keystore"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/schema"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/transaction"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	coininfoFields := schema.CoinInfo{}.Fields()
	_ = coininfoFields
	// coininfoDescName is the schema descriptor for name field.
	coininfoDescName := coininfoFields[1].Descriptor()
	// coininfo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	coininfo.NameValidator = func() func(string) error {
		validators := coininfoDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// coininfoDescUnit is the schema descriptor for unit field.
	coininfoDescUnit := coininfoFields[2].Descriptor()
	// coininfo.UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	coininfo.UnitValidator = func() func(string) error {
		validators := coininfoDescUnit.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(unit string) error {
			for _, fn := range fns {
				if err := fn(unit); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// coininfoDescNeedSigninfo is the schema descriptor for need_signinfo field.
	coininfoDescNeedSigninfo := coininfoFields[3].Descriptor()
	// coininfo.DefaultNeedSigninfo holds the default value on creation for the need_signinfo field.
	coininfo.DefaultNeedSigninfo = coininfoDescNeedSigninfo.Default.(bool)
	keystoreFields := schema.KeyStore{}.Fields()
	_ = keystoreFields
	// keystoreDescAddress is the schema descriptor for address field.
	keystoreDescAddress := keystoreFields[1].Descriptor()
	// keystore.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	keystore.AddressValidator = func() func(string) error {
		validators := keystoreDescAddress.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(address string) error {
			for _, fn := range fns {
				if err := fn(address); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// keystoreDescPrivateKey is the schema descriptor for private_key field.
	keystoreDescPrivateKey := keystoreFields[2].Descriptor()
	// keystore.PrivateKeyValidator is a validator for the "private_key" field. It is called by the builders before save.
	keystore.PrivateKeyValidator = func() func(string) error {
		validators := keystoreDescPrivateKey.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(private_key string) error {
			for _, fn := range fns {
				if err := fn(private_key); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	reviewFields := schema.Review{}.Fields()
	_ = reviewFields
	// reviewDescIsApproved is the schema descriptor for is_approved field.
	reviewDescIsApproved := reviewFields[1].Descriptor()
	// review.DefaultIsApproved holds the default value on creation for the is_approved field.
	review.DefaultIsApproved = reviewDescIsApproved.Default.(bool)
	// reviewDescOperatorNote is the schema descriptor for operator_note field.
	reviewDescOperatorNote := reviewFields[2].Descriptor()
	// review.OperatorNoteValidator is a validator for the "operator_note" field. It is called by the builders before save.
	review.OperatorNoteValidator = reviewDescOperatorNote.Validators[0].(func(string) error)
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescAmountInt is the schema descriptor for amount_int field.
	transactionDescAmountInt := transactionFields[1].Descriptor()
	// transaction.AmountIntValidator is a validator for the "amount_int" field. It is called by the builders before save.
	transaction.AmountIntValidator = transactionDescAmountInt.Validators[0].(func(int) error)
	// transactionDescAmountDigits is the schema descriptor for amount_digits field.
	transactionDescAmountDigits := transactionFields[2].Descriptor()
	// transaction.DefaultAmountDigits holds the default value on creation for the amount_digits field.
	transaction.DefaultAmountDigits = transactionDescAmountDigits.Default.(int)
	// transaction.AmountDigitsValidator is a validator for the "amount_digits" field. It is called by the builders before save.
	transaction.AmountDigitsValidator = transactionDescAmountDigits.Validators[0].(func(int) error)
	// transactionDescAddressFrom is the schema descriptor for address_from field.
	transactionDescAddressFrom := transactionFields[3].Descriptor()
	// transaction.AddressFromValidator is a validator for the "address_from" field. It is called by the builders before save.
	transaction.AddressFromValidator = func() func(string) error {
		validators := transactionDescAddressFrom.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(address_from string) error {
			for _, fn := range fns {
				if err := fn(address_from); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// transactionDescAddressTo is the schema descriptor for address_to field.
	transactionDescAddressTo := transactionFields[4].Descriptor()
	// transaction.AddressToValidator is a validator for the "address_to" field. It is called by the builders before save.
	transaction.AddressToValidator = func() func(string) error {
		validators := transactionDescAddressTo.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(address_to string) error {
			for _, fn := range fns {
				if err := fn(address_to); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// transactionDescNeedManualReview is the schema descriptor for need_manual_review field.
	transactionDescNeedManualReview := transactionFields[5].Descriptor()
	// transaction.DefaultNeedManualReview holds the default value on creation for the need_manual_review field.
	transaction.DefaultNeedManualReview = transactionDescNeedManualReview.Default.(bool)
	// transactionDescTransactionIDInsite is the schema descriptor for transaction_id_insite field.
	transactionDescTransactionIDInsite := transactionFields[7].Descriptor()
	// transaction.TransactionIDInsiteValidator is a validator for the "transaction_id_insite" field. It is called by the builders before save.
	transaction.TransactionIDInsiteValidator = transactionDescTransactionIDInsite.Validators[0].(func(string) error)
	// transactionDescTransactionIDChain is the schema descriptor for transaction_id_chain field.
	transactionDescTransactionIDChain := transactionFields[8].Descriptor()
	// transaction.TransactionIDChainValidator is a validator for the "transaction_id_chain" field. It is called by the builders before save.
	transaction.TransactionIDChainValidator = transactionDescTransactionIDChain.Validators[0].(func(string) error)
	// transactionDescMutex is the schema descriptor for mutex field.
	transactionDescMutex := transactionFields[10].Descriptor()
	// transaction.DefaultMutex holds the default value on creation for the mutex field.
	transaction.DefaultMutex = transactionDescMutex.Default.(bool)
}
