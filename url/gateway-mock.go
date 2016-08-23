package url

// MockGateway provides an implementation of the Gateway interface that allows pre-settings values,
// for use in tests.
type MockGateway struct {
	mockURL   *URL
	mockError error
}

// FindByReference retrieves a pre-set mock URL instance, and mock error.
func (g MockGateway) FindByReference(reference string) (*URL, error) {
	return g.mockURL, g.mockError
}

// FindLatest retrieves a pre-set mock URL instance, and mock error.
func (g MockGateway) FindLatest() (*URL, error) {
	return g.mockURL, g.mockError
}

// Persist retrieves a pre-set mock error.
func (g MockGateway) Persist(url URL) error {
	return g.mockError
}
