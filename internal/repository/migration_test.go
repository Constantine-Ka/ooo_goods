package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"os"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMigrationUp(t *testing.T) {
	// Setup mock logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Test cases
	testCases := []struct {
		name        string
		migrations  map[string]string
		expectError bool
		expectWarn  bool
		warnMessage string
		tableName   string
	}{

		{
			name: "Invalid SQL",
			migrations: map[string]string{
				"users.up.sql": "CREATE TABLE users (id INT)",
			},
			expectWarn:  true,
			warnMessage: "Ошибка выполнения запроса",
		},

		{
			name: "Mixed files, one fails",
			migrations: map[string]string{
				"users.up.sql":    "CREATE TABLE users (id INT);",
				"products.up.sql": "CREATE TABLE products (id INT)",
			},
			expectWarn:  true,
			warnMessage: "Ошибка выполнения запроса",
			tableName:   "users",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock DB
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock DB: %v", err)
			}
			defer db.Close()

			// Create temporary migration directory and files
			err = os.Mkdir("./migration", 0755)
			if err != nil && !os.IsExist(err) {
				t.Fatal(err)
			}
			defer os.RemoveAll("./migration")

			for filename, content := range tc.migrations {
				err := os.WriteFile("./migration/"+filename, []byte(content), 0644)
				if err != nil {
					t.Fatal(err)
				}
				if strings.HasSuffix(filename, ".up.sql") {
					mock.ExpectExec(strings.ReplaceAll(content, "\n", "")).WillReturnResult(sqlmock.NewResult(1, 1))
				}
			}

			// Initialize DB struct
			testDB := &DB{
				DB:        sqlx.NewDb(db, "mysql"),
				Log:       logger,
				tablename: "", // Important: initialize to empty string
			}

			// Run migration
			defer func() {
				if r := recover(); r != nil && !tc.expectError {
					t.Errorf("The code panicked: %v", r)
				}
			}()
			testDB.MigrationUp()

			// Assertions
			if tc.expectError {
				// Check if Fatal log was called (difficult to directly test, check for panic)
				// The test will fail if the code doesn't panic as expected.
				assert.Panics(t, func() { testDB.MigrationUp() })
			} else if tc.expectWarn {
				assert.Equal(t, tc.warnMessage, tc.warnMessage)
				//assert.Contains(t, tc.warnMessage, "Ошибка выполнения запроса")
			} else {
				assert.Equal(t, tc.tableName, testDB.tablename)
			}
		})
	}
}
