package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
)

type Config struct {
	IsProd     string `env:"IS_PROD"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     string `env:"DB_PORT"`
	UploadsDir string `env:"UPLOADS_DIR"`
	StaticDir  string `env:"STATIC_DIR"`
	AppPort    string `env:"APP_PORT"`
	AppHost    string `env:"APP_HOST"`
}

var (
	once    sync.Once
	Conf    Config
	tagName = "env"
)

func Get() *Config {
	once.Do(func() {
		var configBytes []byte
		var err error
		if os.Getenv("IS_PROD") != "" {
			if err := readEnv(); err != nil {
				log.Fatal(err)
			}

			configBytes, err = json.MarshalIndent(Conf, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

		} else {
			Conf = Config{
				IsProd:     "false",
				DbUser:     "admin",
				DbPassword: "admin",
				DbHost:     "localhost",
				DbName:     "postgres",
				DbPort:     "5432",
				UploadsDir: "../uploads",
				StaticDir:  "./static",
				AppPort:    "8000",
				AppHost:    "0.0.0.0",
			}
			configBytes, err = json.MarshalIndent(Conf, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("Configuration:", string(configBytes))
	})

	return &Conf
}

func readEnv() error {
	t := reflect.TypeOf(Conf)
	elem := reflect.ValueOf(&Conf).Elem()

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(tagName)

		if tag == "" {
			continue
		}

		envValue := os.Getenv(tag)
		if envValue == "" {
			message := fmt.Sprintf("%s environment variable is empty", tag)
			return errors.New(message)
		}

		if elem.Field(i).CanSet() {
			elem.Field(i).SetString(envValue)
		}
	}

	return nil
}
