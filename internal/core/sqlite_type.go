package core

import (
	"github.com/sqlc-dev/plugin-sdk-go/plugin"
	"github.com/sqlc-dev/plugin-sdk-go/sdk"
	"strings"
)

// sqliteType returns the SQLite type for a column
// https://www.sqlite.org/datatype3.html#determination_of_column_affinity
func sqliteType(col *plugin.Column) (string, bool) {
	columnType := strings.ToUpper(sdk.DataType(col.Type))
	switch {
	// 1. If the declared type contains the string "INT" then it is assigned INTEGER affinity.
	case strings.Contains(columnType, "INT"):
		return "Long", false
	// 2. If the declared type of the column contains any of the strings "CHAR", "CLOB", or "TEXT" then that column has TEXT affinity.
	// Notice that the type VARCHAR contains the string "CHAR" and is thus assigned TEXT affinity.
	case strings.Contains(columnType, "CHAR"),
		strings.Contains(columnType, "CLOB"),
		strings.Contains(columnType, "TEXT"):
		return "String", false
	// 3. If the declared type for a column contains the string "BLOB" or if no type is specified then the column has affinity BLOB.
	case strings.Contains(columnType, "BLOB"):
		return "ByteArray", false
	// 4. If the declared type for a column contains any of the strings "REAL", "FLOA", or "DOUB" then the column has REAL affinity.
	case strings.Contains(columnType, "REAL"),
		strings.Contains(columnType, "FLOA"),
		strings.Contains(columnType, "DOUB"):
		return "Double", false
	// 5. Otherwise, the affinity is NUMERIC.
	default:
		return "Any", false
	}
}
