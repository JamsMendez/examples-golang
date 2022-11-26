package postgres

import "context"

// Poo.Interface is an wraping to PgxPool to create test mocks
type PoolInterface interface {
  Conn()
}

// ConnectionDB is a simulation of  DB engine (For instance postgres)
type ConnectionDB struct {
}

func (c ConnectionDB) Conn() {
}

// GetConnection return connection pool from postgres drive
func GetConnection(context context.Context) *ConnectionDB {
  return &ConnectionDB{}
}

func RunMigration() {
  // Process for migrate DB in file
}
