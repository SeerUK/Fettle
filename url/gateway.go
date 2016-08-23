package url

// Gateway is the interface for find URL information in some data store.
type Gateway interface {
	FindByReference(string) (*URL, error)
	FindLatest() (*URL, error)
	Persist(URL) error
}
