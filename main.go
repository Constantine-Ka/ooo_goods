package main

import (
	"go.uber.org/zap"
	"ooo_goods/config"
	"ooo_goods/internal/clients/retailCRM"
	"ooo_goods/internal/repository"
	"sync"
	"syscall"
	"time"
	//"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	cfg := config.Init()

	repo := repository.NewDB(cfg)
	repo.MigrationUp()
	page := 1
	cfg.Log.With(zap.Int("Поток", syscall.Getpid()), zap.Int("Страница", page))
	orders, maxPage := retailCRM.GetOrders(cfg, page)
	wg.Add(maxPage)
	go func(ords []retailCRM.Order) {
		cfg.Log.With(zap.Int("Поток", syscall.Getpid()), zap.Int("Страница", page))
		repo.InsertOrder(ords)
		wg.Done()
	}(orders)
	page++
	for ; page <= maxPage; page++ {
		p := page
		go func(pg int) {
			cfg.Log.With(zap.Int("Поток", syscall.Getpid()), zap.Int("Страница", page))
			orders, _ = retailCRM.GetOrders(cfg, pg)
			repo.InsertOrder(orders)

			wg.Done()
		}(p)
		//Согласно документации, существует ограничение на количество запросов с одного IP в 10 запросов в секунду
		if page%9 == 0 {
			time.Sleep(time.Second)
		}

	}
	wg.Wait()
	cfg.Log.Info("Приложение завершило свою работу", zap.Time("Начало", start), zap.Time("Окончание", time.Now()))
}
