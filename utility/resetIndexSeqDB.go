package utility

import (
	"fmt"

	"gorm.io/gorm"
)

func ResetAutoIncrement(db *gorm.DB, tableName string, columnName string) error {
	// Build the SQL query to reset the sequence
	query := fmt.Sprintf(
		"SELECT setval(pg_get_serial_sequence('%s', '%s'), COALESCE(MAX(%s), 1)) FROM %s",
		tableName, columnName, columnName, tableName,
	)

	// Execute the query
	return db.Exec(query).Error
}
