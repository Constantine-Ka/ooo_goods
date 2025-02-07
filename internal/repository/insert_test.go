package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ooo_goods/internal/clients/retailCRM"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestDB_InsertOrder(t *testing.T) {
	//  Настраиваем тестовый логгер
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	testCases := []struct {
		name          string
		inputData     []retailCRM.Order
		expectedQuery string
		expectError   bool
	}{

		{
			name: "Ошибка выполнения запроса",
			inputData: []retailCRM.Order{
				{
					OrderID:            1,
					CustomerID:         2,
					OrderNumber:        "3",
					FactData:           "2024-01-01",
					PodtverjdenOtravka: "2024-01-02",
					TotalSumm:          100.50,
					PrepaySumm:         50.25,
				},
			},
			expectedQuery: "REPLACE INTO test_table (order_id,customer_id,order_number,fakt_data,podtverzden_otpravka,total_summ,prepay_sum) values (1,2,'3',IF('2024-01-01'!='','2024-01-01',NULL),IF('2024-01-02'!='','2024-01-02',NULL),'100.500000','50.250000')",
			expectError:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Не удалось создать mock DB: %v", err)
			}
			defer db.Close()

			testDB := DB{
				DB:        sqlx.NewDb(db, "sqlmock"),
				Log:       logger,
				tablename: "test_table",
			}

			if !tc.expectError {
				mock.ExpectExec(tc.expectedQuery).WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectExec(tc.expectedQuery).WillReturnError(fmt.Errorf("some error"))
			}

			testDB.InsertOrder(tc.inputData)

			if !tc.expectError {
				err = mock.ExpectationsWereMet()
				assert.NoError(t, err, "Все ожидания выполнены")
			}

		})
	}
}
