// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AmountInt holds the value of the "amount_int" field.
	AmountInt int `json:"amount_int,omitempty"`
	// AmountDigits holds the value of the "amount_digits" field.
	AmountDigits int `json:"amount_digits,omitempty"`
	// AddressFrom holds the value of the "address_from" field.
	AddressFrom string `json:"address_from,omitempty"`
	// AddressTo holds the value of the "address_to" field.
	AddressTo string `json:"address_to,omitempty"`
	// NeedManualReview holds the value of the "need_manual_review" field.
	NeedManualReview bool `json:"need_manual_review,omitempty"`
	// Type holds the value of the "type" field.
	Type transaction.Type `json:"type,omitempty"`
	// TransactionIDInsite holds the value of the "transaction_id_insite" field.
	TransactionIDInsite string `json:"transaction_id_insite,omitempty"`
	// TransactionIDChain holds the value of the "transaction_id_chain" field.
	TransactionIDChain string `json:"transaction_id_chain,omitempty"`
	// Status holds the value of the "status" field.
	Status transaction.Status `json:"status,omitempty"`
	// Mutex holds the value of the "mutex" field.
	Mutex bool `json:"mutex,omitempty"`
	// CreatetimeUtc holds the value of the "createtime_utc" field.
	CreatetimeUtc int `json:"createtime_utc,omitempty"`
	// UpdatetimeUtc holds the value of the "updatetime_utc" field.
	UpdatetimeUtc int `json:"updatetime_utc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionQuery when eager-loading is set.
	Edges                  TransactionEdges `json:"edges"`
	coin_info_transactions *int
}

// TransactionEdges holds the relations/edges for other nodes in the graph.
type TransactionEdges struct {
	// Coin holds the value of the coin edge.
	Coin *CoinInfo `json:"coin,omitempty"`
	// Review holds the value of the review edge.
	Review []*Review `json:"review,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CoinOrErr returns the Coin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) CoinOrErr() (*CoinInfo, error) {
	if e.loadedTypes[0] {
		if e.Coin == nil {
			// The edge coin was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: coininfo.Label}
		}
		return e.Coin, nil
	}
	return nil, &NotLoadedError{edge: "coin"}
}

// ReviewOrErr returns the Review value or an error if the edge
// was not loaded in eager-loading.
func (e TransactionEdges) ReviewOrErr() ([]*Review, error) {
	if e.loadedTypes[1] {
		return e.Review, nil
	}
	return nil, &NotLoadedError{edge: "review"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldNeedManualReview, transaction.FieldMutex:
			values[i] = new(sql.NullBool)
		case transaction.FieldID, transaction.FieldAmountInt, transaction.FieldAmountDigits, transaction.FieldCreatetimeUtc, transaction.FieldUpdatetimeUtc:
			values[i] = new(sql.NullInt64)
		case transaction.FieldAddressFrom, transaction.FieldAddressTo, transaction.FieldType, transaction.FieldTransactionIDInsite, transaction.FieldTransactionIDChain, transaction.FieldStatus:
			values[i] = new(sql.NullString)
		case transaction.ForeignKeys[0]: // coin_info_transactions
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case transaction.FieldAmountInt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount_int", values[i])
			} else if value.Valid {
				t.AmountInt = int(value.Int64)
			}
		case transaction.FieldAmountDigits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount_digits", values[i])
			} else if value.Valid {
				t.AmountDigits = int(value.Int64)
			}
		case transaction.FieldAddressFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_from", values[i])
			} else if value.Valid {
				t.AddressFrom = value.String
			}
		case transaction.FieldAddressTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_to", values[i])
			} else if value.Valid {
				t.AddressTo = value.String
			}
		case transaction.FieldNeedManualReview:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field need_manual_review", values[i])
			} else if value.Valid {
				t.NeedManualReview = value.Bool
			}
		case transaction.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = transaction.Type(value.String)
			}
		case transaction.FieldTransactionIDInsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id_insite", values[i])
			} else if value.Valid {
				t.TransactionIDInsite = value.String
			}
		case transaction.FieldTransactionIDChain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id_chain", values[i])
			} else if value.Valid {
				t.TransactionIDChain = value.String
			}
		case transaction.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = transaction.Status(value.String)
			}
		case transaction.FieldMutex:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field mutex", values[i])
			} else if value.Valid {
				t.Mutex = value.Bool
			}
		case transaction.FieldCreatetimeUtc:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createtime_utc", values[i])
			} else if value.Valid {
				t.CreatetimeUtc = int(value.Int64)
			}
		case transaction.FieldUpdatetimeUtc:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updatetime_utc", values[i])
			} else if value.Valid {
				t.UpdatetimeUtc = int(value.Int64)
			}
		case transaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field coin_info_transactions", value)
			} else if value.Valid {
				t.coin_info_transactions = new(int)
				*t.coin_info_transactions = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCoin queries the "coin" edge of the Transaction entity.
func (t *Transaction) QueryCoin() *CoinInfoQuery {
	return (&TransactionClient{config: t.config}).QueryCoin(t)
}

// QueryReview queries the "review" edge of the Transaction entity.
func (t *Transaction) QueryReview() *ReviewQuery {
	return (&TransactionClient{config: t.config}).QueryReview(t)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return (&TransactionClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", amount_int=")
	builder.WriteString(fmt.Sprintf("%v", t.AmountInt))
	builder.WriteString(", amount_digits=")
	builder.WriteString(fmt.Sprintf("%v", t.AmountDigits))
	builder.WriteString(", address_from=")
	builder.WriteString(t.AddressFrom)
	builder.WriteString(", address_to=")
	builder.WriteString(t.AddressTo)
	builder.WriteString(", need_manual_review=")
	builder.WriteString(fmt.Sprintf("%v", t.NeedManualReview))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", transaction_id_insite=")
	builder.WriteString(t.TransactionIDInsite)
	builder.WriteString(", transaction_id_chain=")
	builder.WriteString(t.TransactionIDChain)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", mutex=")
	builder.WriteString(fmt.Sprintf("%v", t.Mutex))
	builder.WriteString(", createtime_utc=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatetimeUtc))
	builder.WriteString(", updatetime_utc=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdatetimeUtc))
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction

func (t Transactions) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
