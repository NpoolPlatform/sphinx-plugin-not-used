// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/keystore"
)

// KeyStore is the model entity for the KeyStore schema.
type KeyStore struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// PrivateKey holds the value of the "private_key" field.
	PrivateKey string `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KeyStoreQuery when eager-loading is set.
	Edges          KeyStoreEdges `json:"edges"`
	coin_info_keys *int32
}

// KeyStoreEdges holds the relations/edges for other nodes in the graph.
type KeyStoreEdges struct {
	// Coin holds the value of the coin edge.
	Coin *CoinInfo `json:"coin,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CoinOrErr returns the Coin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KeyStoreEdges) CoinOrErr() (*CoinInfo, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*KeyStore) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case keystore.FieldID:
			values[i] = new(sql.NullInt64)
		case keystore.FieldAddress, keystore.FieldPrivateKey:
			values[i] = new(sql.NullString)
		case keystore.ForeignKeys[0]: // coin_info_keys
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type KeyStore", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KeyStore fields.
func (ks *KeyStore) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case keystore.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ks.ID = int32(value.Int64)
		case keystore.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				ks.Address = value.String
			}
		case keystore.FieldPrivateKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_key", values[i])
			} else if value.Valid {
				ks.PrivateKey = value.String
			}
		case keystore.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field coin_info_keys", value)
			} else if value.Valid {
				ks.coin_info_keys = new(int32)
				*ks.coin_info_keys = int32(value.Int64)
			}
		}
	}
	return nil
}

// QueryCoin queries the "coin" edge of the KeyStore entity.
func (ks *KeyStore) QueryCoin() *CoinInfoQuery {
	return (&KeyStoreClient{config: ks.config}).QueryCoin(ks)
}

// Update returns a builder for updating this KeyStore.
// Note that you need to call KeyStore.Unwrap() before calling this method if this KeyStore
// was returned from a transaction, and the transaction was committed or rolled back.
func (ks *KeyStore) Update() *KeyStoreUpdateOne {
	return (&KeyStoreClient{config: ks.config}).UpdateOne(ks)
}

// Unwrap unwraps the KeyStore entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ks *KeyStore) Unwrap() *KeyStore {
	tx, ok := ks.config.driver.(*txDriver)
	if !ok {
		panic("ent: KeyStore is not a transactional entity")
	}
	ks.config.driver = tx.drv
	return ks
}

// String implements the fmt.Stringer.
func (ks *KeyStore) String() string {
	var builder strings.Builder
	builder.WriteString("KeyStore(")
	builder.WriteString(fmt.Sprintf("id=%v", ks.ID))
	builder.WriteString(", address=")
	builder.WriteString(ks.Address)
	builder.WriteString(", private_key=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// KeyStores is a parsable slice of KeyStore.
type KeyStores []*KeyStore

func (ks KeyStores) config(cfg config) {
	for _i := range ks {
		ks[_i].config = cfg
	}
}
