package utility

import (
	"fmt"

	"gorm.io/gorm"
)

// ResetAutoIncrement resets the auto-increment sequence for a given table and column
func ResetAutoIncrement(db *gorm.DB, tableName string, columnName string) error {
	// Get the name of the sequence associated with the table and column
	var sequenceName string
	query := fmt.Sprintf(
		"SELECT pg_get_serial_sequence('%s', '%s')",
		tableName, columnName,
	)
	if err := db.Raw(query).Scan(&sequenceName).Error; err != nil {
		return err
	}

	// Build the SQL query to reset the sequence
	resetQuery := fmt.Sprintf(
		"SELECT setval('%s', COALESCE(MAX(%s), 1)) FROM %s",
		sequenceName, columnName, tableName,
	)

	// Execute the reset query
	return db.Exec(resetQuery).Error
}
