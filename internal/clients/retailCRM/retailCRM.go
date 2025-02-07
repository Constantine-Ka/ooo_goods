package retailCRM

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"ooo_goods/config"
)

func GetOrders(cfg config.CFG, page int) ([]Order, int) {

	var orders []Order
	var response Response
	url := fmt.Sprintf("https://%s/api/v5/orders?limit=20&page=%d", cfg.Domain, page)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		cfg.Log.Error("Ошибка создания request", zap.Error(err))
		return orders, 0
	}
	req.Header.Add("X-API-KEY", cfg.Apikey)

	res, err := client.Do(req)
	if err != nil {
		cfg.Log.Error("Ошибка отправки запроса", zap.Error(err))
		return orders, 0
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		cfg.Log.Error("Ошибка чтения тела ответа", zap.Error(err))
		return nil, 0
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		cfg.Log.Error("Ошибка анмаршалинга JSON-ответа", zap.Error(err))
		return nil, 0
	}

	for _, order := range response.Orders {
		orders = append(orders, Order{
			OrderID:            order.Id,
			CustomerID:         order.Customer.Id,
			OrderNumber:        order.Number,
			FactData:           order.CustomFields.FaktData,
			PodtverjdenOtravka: order.CustomFields.PodtverzdenOtpravka,
			TotalSumm:          order.TotalSumm,
			PrepaySumm:         order.PrepaySum,
		})

	}
	cfg.Log.Info("Данные получены", zap.Int("Страница", page), zap.Int("Количество заказов", len(orders)))
	return orders, response.Pagination.TotalPageCount
}
