package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"ooo_goods/internal/logging"
	"os"
)

func Init() CFG {
	var cfg CFG

	var ApikeyExist, DomainExist, DbHostExist, DbNameExist, DbLoginExist, DbTableExist bool
	vp := viper.New()
	cfg.Log = logging.CreateLogger()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		cfg.Log.Fatal("Error while reading config file,", zap.Any("err", err))
		return cfg
	}
	err = vp.Unmarshal(&cfg)
	if err != nil {
		cfg.Log.Fatal(fmt.Sprintf("Ошибка маршалинга %v", err))
	}
	flagApikey := flag.String("Apikey", cfg.Apikey, "apikey RetailCRM")
	flagDomain := flag.String("Domain", cfg.Domain, "domain RetailCRM")
	flagDbhost := flag.String("Dbhost", cfg.DB.Host, "Адрес БД ")
	flagDbname := flag.String("Dbname", cfg.DB.Name, "Название БД ")
	flagDblogin := flag.String("Dblogin", cfg.DB.Login, "Логин для БД ")
	flagDbpass := flag.String("DbPass", cfg.DB.Pass, "Пароль от БД")
	flagDBTable := flag.String("DbTable", cfg.DB.Table, "Таблица, куда будет производиться запись")
	flag.Parse()

	if *flagApikey == "" {
		cfg.Apikey, ApikeyExist = os.LookupEnv("Apikey")
		if !ApikeyExist {
			cfg.Log.Panic("Не указан флаг конфигурации.", zap.String("FIELD", "Apikey"))
		}
	} else {
		cfg.Apikey = *flagApikey
	}
	if *flagDomain == "" {
		cfg.Domain, DomainExist = os.LookupEnv("Domain")
		if !DomainExist {
			cfg.Log.Panic("Не указан флаг конфигурации.", zap.String("FIELD", "Domain"))
		}
	} else {
		cfg.Domain = *flagDomain
	}
	if *flagDbhost == "" {
		cfg.DB.Host, DbHostExist = os.LookupEnv("Dbhost")
		if !DbHostExist {
			cfg.Log.Panic("Не указан флаг конфигурации.", zap.String("FIELD", "Dbhost"))
		}
	} else {
		cfg.DB.Host = *flagDbhost
	}
	if *flagDbname == "" {
		cfg.DB.Name, DbNameExist = os.LookupEnv("Dbname")
		if !DbNameExist {
			cfg.Log.Panic("Не указан флаг конфигурации.", zap.String("FIELD", "Dbname"))
		}
	} else {
		cfg.DB.Name, *flagDbname = *flagDbname, *flagDbhost
	}
	if *flagDblogin == "" {
		cfg.DB.Login, DbLoginExist = os.LookupEnv("DbLogin")
		if !DbLoginExist {
			cfg.Log.Panic("Не указан флаг конфигурации.", zap.String("FIELD", "DbLogin"))
		}
	} else {
		cfg.DB.Login = *flagDblogin
	}
	if *flagDbpass == "" {
		// INFO Тут не создаю ошибку, тк авторизоваться можно с Пустым паролем
		cfg.DB.Pass, _ = os.LookupEnv("Dbpass")
	} else {
		cfg.DB.Pass = *flagDbpass
	}
	if *flagDBTable == "" {
		cfg.DB.Table, DbTableExist = os.LookupEnv("DBTable")
		if !DbTableExist {
			cfg.Log.Panic("Не указан флаг конфигурации.", zap.String("FIELD", "DBTable"))
		}
	} else {
		cfg.DB.Table = *flagDBTable
	}
	cfg.Log.Info("Поля конфигурации собраны.")
	return cfg

}
