package postgres

import (
	"fmt"
	"github.com/bahcasac/stock/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB() (*gorm.DB, error) {
	envs := config.LoadEnvVars()
	fmt.Println(envs)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		envs.Host, envs.UserName, envs.Password, envs.DBName, envs.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(fmt.Sprintf("successful conected in schema: %s with user: %s and port: %s on database: %s",
		envs.DBSchema, envs.UserName, envs.Port, envs.DBName))
	if err != nil {
		return nil, err
	}
	return db, nil
}
