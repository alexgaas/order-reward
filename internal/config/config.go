package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type Flags struct {
	appAddress     string
	databaseDSN    string
	accrualAddress string
}

type Config struct {
	AppAddress     string
	DatabaseDSN    string
	AccrualAddress string
}

type environ struct {
	AppAddress     string `env:"RUN_ADDRESS" envDefault:"localhost:8000"`
	AccrualAddress string `env:"ACCRUAL_SYSTEM_ADDRESS" envDefault:"localhost:8080"`
	DatabaseDSN    string `env:"DATABASE_URI" envDefault:"order_reward.db"`
}

func GetAppFlags() Flags {
	flags := Flags{}
	flag.StringVar(&flags.appAddress, "a", "", "Address application, for example: http://localhost:8000")
	flag.StringVar(&flags.databaseDSN, "d", "", "Database connection string, for example: order_reward.db")
	flag.StringVar(&flags.accrualAddress, "r", "", "Accrual application, for example: http://localhost:8500")
	flag.Parse()
	return flags
}

func GetNewConfig(flags Flags) (*Config, error) {
	var err error
	var config Config
	var envs environ

	if err := env.Parse(&envs, env.Options{}); err != nil {
		return nil, err
	}

	if flags.appAddress == "" {
		config.AppAddress = envs.AppAddress
	} else {
		config.AppAddress = flags.appAddress
	}

	if flags.databaseDSN == "" {
		config.DatabaseDSN = envs.DatabaseDSN
	} else {
		config.DatabaseDSN = flags.databaseDSN
	}

	if flags.accrualAddress == "" {
		config.AccrualAddress = envs.AccrualAddress
	} else {
		config.AccrualAddress = flags.accrualAddress
	}

	return &config, err
}
