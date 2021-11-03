// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
)

// Review is the model entity for the Review schema.
type Review struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IsApproved holds the value of the "is_approved" field.
	IsApproved bool `json:"is_approved,omitempty"`
	// OperatorNote holds the value of the "operator_note" field.
	OperatorNote string `json:"operator_note,omitempty"`
	// CreatetimeUtc holds the value of the "createtime_utc" field.
	CreatetimeUtc int `json:"createtime_utc,omitempty"`
	// UpdatetimeUtc holds the value of the "updatetime_utc" field.
	UpdatetimeUtc int `json:"updatetime_utc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewQuery when eager-loading is set.
	Edges              ReviewEdges `json:"edges"`
	coin_info_reviews  *int
	transaction_review *int
}

// ReviewEdges holds the relations/edges for other nodes in the graph.
type ReviewEdges struct {
	// Transaction holds the value of the transaction edge.
	Transaction *Transaction `json:"transaction,omitempty"`
	// Coin holds the value of the coin edge.
	Coin *CoinInfo `json:"coin,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewEdges) TransactionOrErr() (*Transaction, error) {
	if e.loadedTypes[0] {
		if e.Transaction == nil {
			// The edge transaction was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: transaction.Label}
		}
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// CoinOrErr returns the Coin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewEdges) CoinOrErr() (*CoinInfo, error) {
	if e.loadedTypes[1] {
		if e.Coin == nil {
			// The edge coin was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: coininfo.Label}
		}
		return e.Coin, nil
	}
	return nil, &NotLoadedError{edge: "coin"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Review) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case review.FieldIsApproved:
			values[i] = new(sql.NullBool)
		case review.FieldID, review.FieldCreatetimeUtc, review.FieldUpdatetimeUtc:
			values[i] = new(sql.NullInt64)
		case review.FieldOperatorNote:
			values[i] = new(sql.NullString)
		case review.ForeignKeys[0]: // coin_info_reviews
			values[i] = new(sql.NullInt64)
		case review.ForeignKeys[1]: // transaction_review
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Review", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Review fields.
func (r *Review) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case review.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case review.FieldIsApproved:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_approved", values[i])
			} else if value.Valid {
				r.IsApproved = value.Bool
			}
		case review.FieldOperatorNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operator_note", values[i])
			} else if value.Valid {
				r.OperatorNote = value.String
			}
		case review.FieldCreatetimeUtc:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createtime_utc", values[i])
			} else if value.Valid {
				r.CreatetimeUtc = int(value.Int64)
			}
		case review.FieldUpdatetimeUtc:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updatetime_utc", values[i])
			} else if value.Valid {
				r.UpdatetimeUtc = int(value.Int64)
			}
		case review.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field coin_info_reviews", value)
			} else if value.Valid {
				r.coin_info_reviews = new(int)
				*r.coin_info_reviews = int(value.Int64)
			}
		case review.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field transaction_review", value)
			} else if value.Valid {
				r.transaction_review = new(int)
				*r.transaction_review = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTransaction queries the "transaction" edge of the Review entity.
func (r *Review) QueryTransaction() *TransactionQuery {
	return (&ReviewClient{config: r.config}).QueryTransaction(r)
}

// QueryCoin queries the "coin" edge of the Review entity.
func (r *Review) QueryCoin() *CoinInfoQuery {
	return (&ReviewClient{config: r.config}).QueryCoin(r)
}

// Update returns a builder for updating this Review.
// Note that you need to call Review.Unwrap() before calling this method if this Review
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Review) Update() *ReviewUpdateOne {
	return (&ReviewClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Review entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Review) Unwrap() *Review {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Review is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Review) String() string {
	var builder strings.Builder
	builder.WriteString("Review(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", is_approved=")
	builder.WriteString(fmt.Sprintf("%v", r.IsApproved))
	builder.WriteString(", operator_note=")
	builder.WriteString(r.OperatorNote)
	builder.WriteString(", createtime_utc=")
	builder.WriteString(fmt.Sprintf("%v", r.CreatetimeUtc))
	builder.WriteString(", updatetime_utc=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdatetimeUtc))
	builder.WriteByte(')')
	return builder.String()
}

// Reviews is a parsable slice of Review.
type Reviews []*Review

func (r Reviews) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
