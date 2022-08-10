package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type environment struct {
	dbName   string
	dbSchema string
	userName string
	host     string
	port     string
	password string
	charset  string
}

func ConnectionDB() (*gorm.DB, error) {
	envs := getDatabaseEnvs()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		envs.host, envs.userName, envs.password, envs.dbName, envs.port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(fmt.Sprintf("successful conected in schema: %s with user: %s and port: %s on database: %s",
		envs.dbSchema, envs.userName, envs.port, envs.dbName))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDatabaseEnvs() environment {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	return environment{
		dbName:   viper.GetString("DATABASE_NAME"),
		dbSchema: viper.GetString("DATABASE_SCHEMA"),
		userName: viper.GetString("DATABASE_USERNAME"),
		host:     viper.GetString("DATABASE_HOST"),
		port:     viper.GetString("DATABASE_PORT"),
		password: viper.GetString("DATABASE_PASSWORD"),
		charset:  viper.GetString("DATABASE_CHARSET"),
	}

}
