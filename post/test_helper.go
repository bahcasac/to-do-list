// usage:
// testDB := testhelpers.NewTestDatabase(t)
// defer testDB.Close(t)
// println(testDB.ConnectionString(t))
package post

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	dbUser     = "postgres"
	dbPassword = "workshop"
	database   = "workshop"
)

type PostgresContainer struct {
	instance testcontainers.Container
	DSN      string
}

func SetupPostgres(ctx context.Context) (*PostgresContainer, error) {
	fmt.Println("comecou...")
	req := testcontainers.ContainerRequest{
		Image:        "postgres:11.6-alpine",
		ExposedPorts: []string{"5433/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": dbPassword,
			"POSTGRES_USER":     dbUser,
			"POSTGRES_DATABASE": database,
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
	}
	fmt.Println("comecou 2...")
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Println("erro ao gerar o container...")
		return nil, err
	}

	mappedPort, err := postgres.MappedPort(ctx, "5433")
	if err != nil {
		fmt.Println("erro para pegar a porta...")
		return nil, err
	}

	hostIP, err := postgres.Host(ctx)
	if err != nil {
		fmt.Println("erro para pegar o host")
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		hostIP, dbUser, dbPassword, database, mappedPort.Int())

	return &PostgresContainer{
		instance: postgres,
		DSN:      dsn,
	}, nil
}

func (db *PostgresContainer) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, db.instance.Terminate(ctx))
}
