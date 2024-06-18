package test

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"gorm.io/gorm"
	"server/config"
	"server/pkg/adapters/storage"
)

var (
	testDB *gorm.DB
	once   sync.Once
)

func TestMain(m *testing.M) {
	var err error
	once.Do(func() {
		testDB, err = setupDatabase()
		if err != nil {
			log.Fatalf("Failed to setup database: %v", err)
		}
	})

	// Run tests
	code := m.Run()

	os.Exit(code)
}

func setupDatabase() (*gorm.DB, error) {
	dbConfig := config.DB{
		Host:   "localhost",
		User:   "username",
		Pass:   "password",
		DBName: "testdb",
		Port:   5432,
	}

	db, err := storage.NewPostgresGormConnection(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = storage.Migrate(db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}

func TestDatabaseConnection(t *testing.T) {
	if testDB == nil {
		t.Fatalf("Database connection is nil")
	}

	sqlDB, err := testDB.DB()
	if err != nil {
		t.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
