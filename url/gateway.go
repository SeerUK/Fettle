package url

import "sync"

// -- URL Gateway Interface

// Gateway is the interface for find URL information in some data store.
type Gateway interface {
	FindByReference(string) (*URL, error)
	FindLatest() (*URL, error)
}

// -- Mock URL Gateway

// MockGateway provides an implementation of the Gateway interface that allows pre-settings values,
// for use in tests.
type MockGateway struct {
	mockURL   *URL
	mockError error
}

// FindByReference retrieves a pre-set mock URL instance.
func (g MockGateway) FindByReference(reference string) (*URL, error) {
	return g.mockURL, g.mockError
}

// FindLatest retrieves a pre-set mock URL instance.
func (g MockGateway) FindLatest() (*URL, error) {
	return g.mockURL, g.mockError
}

// -- MySQL URL Gateway

// MysqlGateway provides an implementation of the Gateway interface that uses Mysql as a data store.
type MysqlGateway struct {
	sync.RWMutex
}

// FindByReference finds a URL by it's reference.
func (g MysqlGateway) FindByReference(reference string) (*URL, error) {
	g.RLock()
	defer g.RUnlock()

	return &URL{}, nil
}

func (g MysqlGateway) findLatest() (*URL, error) {
	return &URL{}, nil
}

// FindLatest finds the latest URL in the data store.
func (g MysqlGateway) FindLatest() (*URL, error) {
	g.RLock()
	defer g.RUnlock()

	return g.findLatest()
}
