// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/walletnode"
)

// WalletNode is the model entity for the WalletNode schema.
type WalletNode struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// HostVendor holds the value of the "host_vendor" field.
	HostVendor string `json:"host_vendor,omitempty"`
	// PublicIP holds the value of the "public_ip" field.
	PublicIP string `json:"public_ip,omitempty"`
	// LocalIP holds the value of the "local_ip" field.
	LocalIP string `json:"local_ip,omitempty"`
	// CreatetimeUtc holds the value of the "createtime_utc" field.
	CreatetimeUtc int `json:"createtime_utc,omitempty"`
	// LastOnlineTimeUtc holds the value of the "last_online_time_utc" field.
	LastOnlineTimeUtc int `json:"last_online_time_utc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WalletNodeQuery when eager-loading is set.
	Edges                  WalletNodeEdges `json:"edges"`
	coin_info_wallet_nodes *int32
}

// WalletNodeEdges holds the relations/edges for other nodes in the graph.
type WalletNodeEdges struct {
	// Coin holds the value of the coin edge.
	Coin *CoinInfo `json:"coin,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CoinOrErr returns the Coin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WalletNodeEdges) CoinOrErr() (*CoinInfo, error) {
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
func (*WalletNode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case walletnode.FieldID, walletnode.FieldCreatetimeUtc, walletnode.FieldLastOnlineTimeUtc:
			values[i] = new(sql.NullInt64)
		case walletnode.FieldUUID, walletnode.FieldLocation, walletnode.FieldHostVendor, walletnode.FieldPublicIP, walletnode.FieldLocalIP:
			values[i] = new(sql.NullString)
		case walletnode.ForeignKeys[0]: // coin_info_wallet_nodes
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WalletNode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WalletNode fields.
func (wn *WalletNode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case walletnode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wn.ID = int32(value.Int64)
		case walletnode.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				wn.UUID = value.String
			}
		case walletnode.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				wn.Location = value.String
			}
		case walletnode.FieldHostVendor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_vendor", values[i])
			} else if value.Valid {
				wn.HostVendor = value.String
			}
		case walletnode.FieldPublicIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_ip", values[i])
			} else if value.Valid {
				wn.PublicIP = value.String
			}
		case walletnode.FieldLocalIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field local_ip", values[i])
			} else if value.Valid {
				wn.LocalIP = value.String
			}
		case walletnode.FieldCreatetimeUtc:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createtime_utc", values[i])
			} else if value.Valid {
				wn.CreatetimeUtc = int(value.Int64)
			}
		case walletnode.FieldLastOnlineTimeUtc:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_online_time_utc", values[i])
			} else if value.Valid {
				wn.LastOnlineTimeUtc = int(value.Int64)
			}
		case walletnode.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field coin_info_wallet_nodes", value)
			} else if value.Valid {
				wn.coin_info_wallet_nodes = new(int32)
				*wn.coin_info_wallet_nodes = int32(value.Int64)
			}
		}
	}
	return nil
}

// QueryCoin queries the "coin" edge of the WalletNode entity.
func (wn *WalletNode) QueryCoin() *CoinInfoQuery {
	return (&WalletNodeClient{config: wn.config}).QueryCoin(wn)
}

// Update returns a builder for updating this WalletNode.
// Note that you need to call WalletNode.Unwrap() before calling this method if this WalletNode
// was returned from a transaction, and the transaction was committed or rolled back.
func (wn *WalletNode) Update() *WalletNodeUpdateOne {
	return (&WalletNodeClient{config: wn.config}).UpdateOne(wn)
}

// Unwrap unwraps the WalletNode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wn *WalletNode) Unwrap() *WalletNode {
	tx, ok := wn.config.driver.(*txDriver)
	if !ok {
		panic("ent: WalletNode is not a transactional entity")
	}
	wn.config.driver = tx.drv
	return wn
}

// String implements the fmt.Stringer.
func (wn *WalletNode) String() string {
	var builder strings.Builder
	builder.WriteString("WalletNode(")
	builder.WriteString(fmt.Sprintf("id=%v", wn.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(wn.UUID)
	builder.WriteString(", location=")
	builder.WriteString(wn.Location)
	builder.WriteString(", host_vendor=")
	builder.WriteString(wn.HostVendor)
	builder.WriteString(", public_ip=")
	builder.WriteString(wn.PublicIP)
	builder.WriteString(", local_ip=")
	builder.WriteString(wn.LocalIP)
	builder.WriteString(", createtime_utc=")
	builder.WriteString(fmt.Sprintf("%v", wn.CreatetimeUtc))
	builder.WriteString(", last_online_time_utc=")
	builder.WriteString(fmt.Sprintf("%v", wn.LastOnlineTimeUtc))
	builder.WriteByte(')')
	return builder.String()
}

// WalletNodes is a parsable slice of WalletNode.
type WalletNodes []*WalletNode

func (wn WalletNodes) config(cfg config) {
	for _i := range wn {
		wn[_i].config = cfg
	}
}
