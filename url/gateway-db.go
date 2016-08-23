package url

import (
	"database/sql"
	"sync"
)

// MysqlGateway provides an implementation of the Gateway interface that uses Mysql as a data store.
type DBGateway struct {
	// RWMutex to prevent simultaneous access when persisting.
	sync.RWMutex
	// Pre-configured database connection.
	db *sql.DB
}

// NewDBGateway creates a new instance of DBGateway.
func NewDBGateway(db *sql.DB) *DBGateway {
	return &DBGateway{
		db: db,
	}
}

// FindByReference finds a URL by it's reference.
func (g *DBGateway) FindByReference(reference string) (*URL, error) {
	g.RLock()
	defer g.RUnlock()

	return &URL{}, nil
}

// FindLatest finds the latest URL in the data store.
func (g *DBGateway) FindLatest() (*URL, error) {
	g.RLock()
	defer g.RUnlock()

	return &URL{}, nil
}

// Persist takes a URL and saves it in the database.
func (g *DBGateway) Persist(url URL) error {
	g.Lock()
	defer g.Unlock()

	return nil
}
