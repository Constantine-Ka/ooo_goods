package repository

import (
	"fmt"
	"go.uber.org/zap"
	"ooo_goods/internal/clients/retailCRM"
	"strings"
)

func (db DB) InsertOrder(data []retailCRM.Order) {
	var queryArr []string
	queryPrefix := fmt.Sprintf("REPLACE INTO %s (%s,%s,%s,%s,%s,%s,%s) values ", db.tablename, "order_id", "customer_id", "order_number", "fakt_data", "podtverzden_otpravka", "total_summ", "prepay_sum")
	for _, item := range data {
		queryArr = append(queryArr, fmt.Sprintf("(%d,%d,'%s',IF('%s'!='','%s',NULL),IF('%s'!='','%s',NULL),'%f','%f')", item.OrderID, item.CustomerID, strings.ReplaceAll(item.OrderNumber, "'", ""), item.FactData, item.FactData, item.PodtverjdenOtravka, item.PodtverjdenOtravka, item.TotalSumm, item.PrepaySumm))
	}
	query := queryPrefix + strings.Join(queryArr, ", ")
	_, err := db.DB.Exec(query)
	if err != nil {
		db.Log.Error("Запрос на добавление завершился с ошибкой", zap.String("query", query), zap.Error(err))
		return
	}
	db.Log.Info("Элементы добавлены в базу", zap.Int("Количество", len(data)))
}
