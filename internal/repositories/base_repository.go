package repositories

import (
	"context"
	"database/sql"
)

// BaseRepository defines common database operations
type BaseRepository interface {
	// FindByID retrieves a record by its ID
	FindByID(ctx context.Context, id interface{}, dest interface{}) error
	
	// FindOne retrieves a single record that matches the given query and arguments
	FindOne(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	
	// FindAll retrieves all records that match the given query and arguments
	FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	
	// Create inserts a new record
	Create(ctx context.Context, entity interface{}) error
	
	// Update updates an existing record
	Update(ctx context.Context, entity interface{}) error
	
	// Delete removes a record by its ID
	Delete(ctx context.Context, id interface{}) error
	
	// ExecuteInTransaction executes the given function within a transaction
	ExecuteInTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error
}

// BaseRepositoryImpl is a base implementation of BaseRepository
type BaseRepositoryImpl struct {
	DB *sql.DB
}

// NewBaseRepository creates a new BaseRepositoryImpl
func NewBaseRepository(db *sql.DB) *BaseRepositoryImpl {
	return &BaseRepositoryImpl{
		DB: db,
	}
}

// ExecuteInTransaction executes the given function within a transaction
func (r *BaseRepositoryImpl) ExecuteInTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // re-throw panic after rollback
		}
	}()

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
