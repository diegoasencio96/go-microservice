package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type (
	Specification struct {
		ProjectName    string `default:"go-microservice"`
		ProjectVersion string `default:"0.0.0"`
		General        GeneralSpecification
		BaseRoutes     BaseRoutesSpecification
		Security       SecuritySpecification
		Database       DatabaseSpecification
		Redis          RedisSpecification
		Kafka          KafkaSpecification
	}

	GeneralSpecification struct {
		Environment string `envconfig:"ENVIRONMENT" default:"production"`
		Host        string `envconfig:"HOST" default:"0.0.0.0"`
		Port        string `envconfig:"PORT" default:"8080"`
		Language    string `envconfig:"LANGUAGE" default:"ES"`
		Timezone    string `envconfig:"TIMEZONE" default:"America/Bogota"`
		Country     string `envconfig:"COUNTRY" default:"CO"`
		BaseURL     string `envconfig:"BASE_URL" default:"127.0.0.1" required:"true"`
		Timeout     int    `envconfig:"TIMEOUT" default:"10"`
		APIKey      string `envconfig:"API_KEY"`
	}

	BaseRoutesSpecification struct {
		BaseRoute         string `default:"/"`
		StatusRoute       string `default:"/status"`
		OtherRoutes       []string
	}

	SecuritySpecification struct {
		ApplicationId string `envconfig:"x_application_id" default:"go-microservice"`
	}

	DatabaseSpecification struct {
		Name           string `envconfig:"DB_NAME" default:"db_name" required:"true"`
		Connection     string `envconfig:"DB_CONNECTION" required:"false"`
		MaxConnections int    `envconfig:"DB_MAX_CONNECTIONS" default:"5"`
		ConnectTimeout int    `envconfig:"DB_CONNECT_TIMEOUT" default:"30"`
		ReadTimeout    int    `envconfig:"DB_READ_TIMEOUT" default:"30000"`
		WriteTimeout   int    `envconfig:"DB_WRITE_TIMEOUT" default:"30000"`
	}

	RedisSpecification struct {
		Host        string `envconfig:"REDIS_HOST" required:"false"`
		Port        string `envconfig:"REDIS_PORT" required:"false"`
		DialTimeout int    `envconfig:"REDIS_DIAL_TIMEOUT" default:"10"`
		Timeout     int    `envconfig:"REDIS_TIMEOUT" default:"30"`
		PoolSize    int    `envconfig:"REDIS_POOL_SIZE" default:"10"`
	}

	KafkaSpecification struct {
		Host string `envconfig:"KAFKA_HOST" required:"false"`
		Port string `envconfig:"KAFKA_PORT" required:"false"`
	}
)

const (
	environmentKey = "ENVIRONMENT"
	developEnv     = "develop"
	productionEnv  = "production"
)

var Settings Specification

func init() {
	Settings.General.Environment = os.Getenv(environmentKey)
	if Settings.General.Environment == developEnv {

	}

	if err := envconfig.Process("", &Settings); err != nil {
		panic(err.Error())
	}

	Settings.BaseRoutes.BaseRoute = fmt.Sprintf("/api/%s", Settings.ProjectName)
	Settings.BaseRoutes.OtherRoutes = []string{
		Settings.BaseRoutes.BaseRoute + "",
	}
}
