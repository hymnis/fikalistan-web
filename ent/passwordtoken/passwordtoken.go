// Code generated by ent, DO NOT EDIT.

package passwordtoken

import (
	"time"
)

const (
	// Label holds the string label denoting the passwordtoken type in the database.
	Label = "password_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the passwordtoken in the database.
	Table = "password_tokens"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "password_tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "password_token_user"
)

// Columns holds all SQL columns for passwordtoken fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "password_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"password_token_user",
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
	// HashValidator is a validator for the "hash" field. It is called by the builders before save.
	HashValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
