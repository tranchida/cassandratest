package factory

import (
	"cassandratest/internal/cassandra"
	"cassandratest/internal/model"
	"cassandratest/internal/sqlite"
	"fmt"
)

func NewRepository(driver string) (model.Repository, error) {
	switch driver {
	case "cassandra":
		return cassandra.NewCassandraRepo( /* params ou via env */ )
	case "sqlite":
		return sqlite.NewSqliteRepo( /* params ou via env */ )
	default:
		return nil, fmt.Errorf("unknown driver: %s", driver)
	}
}
