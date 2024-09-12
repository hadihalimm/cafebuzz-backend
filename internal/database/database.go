package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	Model(value interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Updates(value interface{}) *gorm.DB
	Delete(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB
	ScanRows(rows *sql.Rows, result interface{}) error
	DropTableIfExists(value interface{}) error
	AutoMigrate(value interface{}) error
	Transaction(fc func(tx Service) error) (err error)

	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, port)
	// db, err := sql.Open("pgx", connStr)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) Model(value interface{}) *gorm.DB {
	return s.db.Model(value)
}

func (s *service) Select(query interface{}, args ...interface{}) *gorm.DB {
	return s.db.Select(query, args...)
}

func (s *service) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return s.db.Find(dest, conds...)
}

func (s *service) Exec(sql string, values ...interface{}) *gorm.DB {
	return s.db.Exec(sql, values...)
}

func (s *service) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return s.db.First(dest, conds...)
}

func (s *service) Raw(sql string, values ...interface{}) *gorm.DB {
	return s.db.Raw(sql, values)
}

func (s *service) Create(value interface{}) *gorm.DB {
	return s.db.Create(value)
}

func (s *service) Save(value interface{}) *gorm.DB {
	return s.db.Save(value)
}

func (s *service) Updates(value interface{}) *gorm.DB {
	return s.db.Updates(value)
}

func (s *service) Delete(value interface{}) *gorm.DB {
	return s.db.Delete(value)
}

func (s *service) Where(query interface{}, args ...interface{}) *gorm.DB {
	return s.db.Where(query, args...)
}

func (s *service) Preload(column string, conditions ...interface{}) *gorm.DB {
	return s.db.Preload(column, conditions...)
}

func (s *service) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return s.db.Scopes(funcs...)
}

func (s *service) ScanRows(rows *sql.Rows, result interface{}) error {
	return s.db.ScanRows(rows, result)
}

func (s *service) DropTableIfExists(value interface{}) error {
	return s.db.Migrator().DropTable(value)
}

func (s *service) AutoMigrate(value interface{}) error {
	return s.db.AutoMigrate(value)
}

func (s *service) Transaction(fc func(tx Service) error) (err error) {
	panicked := true
	tx := s.db.Begin()
	defer func() {
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	txservice := &service{}
	txservice.db = tx
	err = fc(txservice)

	if err == nil {
		err = tx.Commit().Error
	}

	panicked = false
	return
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	sqlDb, _ := s.db.DB()
	return sqlDb.Close()
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	sqlDb, _ := s.db.DB()
	// Ping the database
	err := sqlDb.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := sqlDb.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}
