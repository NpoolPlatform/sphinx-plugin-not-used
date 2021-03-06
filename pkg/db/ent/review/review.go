// Code generated by entc, DO NOT EDIT.

package review

const (
	// Label holds the string label denoting the review type in the database.
	Label = "review"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsApproved holds the string denoting the is_approved field in the database.
	FieldIsApproved = "is_approved"
	// FieldOperatorNote holds the string denoting the operator_note field in the database.
	FieldOperatorNote = "operator_note"
	// FieldCreatetimeUtc holds the string denoting the createtime_utc field in the database.
	FieldCreatetimeUtc = "createtime_utc"
	// FieldUpdatetimeUtc holds the string denoting the updatetime_utc field in the database.
	FieldUpdatetimeUtc = "updatetime_utc"
	// EdgeTransaction holds the string denoting the transaction edge name in mutations.
	EdgeTransaction = "transaction"
	// EdgeCoin holds the string denoting the coin edge name in mutations.
	EdgeCoin = "coin"
	// Table holds the table name of the review in the database.
	Table = "reviews"
	// TransactionTable is the table that holds the transaction relation/edge.
	TransactionTable = "reviews"
	// TransactionInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionInverseTable = "transactions"
	// TransactionColumn is the table column denoting the transaction relation/edge.
	TransactionColumn = "transaction_review"
	// CoinTable is the table that holds the coin relation/edge.
	CoinTable = "reviews"
	// CoinInverseTable is the table name for the CoinInfo entity.
	// It exists in this package in order to avoid circular dependency with the "coininfo" package.
	CoinInverseTable = "coin_infos"
	// CoinColumn is the table column denoting the coin relation/edge.
	CoinColumn = "coin_info_reviews"
)

// Columns holds all SQL columns for review fields.
var Columns = []string{
	FieldID,
	FieldIsApproved,
	FieldOperatorNote,
	FieldCreatetimeUtc,
	FieldUpdatetimeUtc,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "reviews"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"coin_info_reviews",
	"transaction_review",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsApproved holds the default value on creation for the "is_approved" field.
	DefaultIsApproved bool
	// OperatorNoteValidator is a validator for the "operator_note" field. It is called by the builders before save.
	OperatorNoteValidator func(string) error
)
