package url

import "sync"

// Referencer contains methods for referencing URLs.
type Referencer struct {
	sync.Mutex
	gateway Gateway
}

// Reference takes a URL, and returns a reference that can be used.
func (r *Referencer) Reference(url string) string {
	r.Lock()
	defer r.Unlock()

	_, err := r.gateway.FindLatest()
	if err != nil {
		// @todo: Handle
	}

	return ""
}
