package url

import "sync"

// Referencer contains methods for referencing URLs.
type Referencer struct {
	sync.Mutex
	gateway Gateway
}

// Reference takes a URL, and creates a reference that can be used to retrieve the underlying URL.
func (r *Referencer) Reference(url string) URL {
	r.Lock()
	defer r.Unlock()

	_, err := r.gateway.FindLatest()
	if err != nil {
		// @todo: Handle
	}

	return URL{}
}
