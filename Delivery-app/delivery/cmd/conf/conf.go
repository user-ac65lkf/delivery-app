package conf

import (
	"fmt"
	"github.com/Shemistan/uzum_delivery/internal/models"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const LocalRun = true

func NewConfig() (models.Config, error) {
	cfg := models.Config{}
	var err error

	if LocalRun {
		err = godotenv.Load("./dev/local.env")
		if err != nil {
			return models.Config{}, err
		}

		var cfg models.Config
		err = envconfig.Process("", &cfg)
		if err != nil {
			return models.Config{}, err
		}
		return cfg, nil
	}

	envconfig.MustProcess("", &cfg)
	return cfg, nil

}

func GetSqlConnectionString(cnf models.Config) string {
	sqlConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%v",
		cnf.DB.User, cnf.DB.Password, cnf.DB.Host, cnf.DB.Port, cnf.DB.Database, cnf.DB.SSLMode,
	)

	return sqlConnectionString
}
