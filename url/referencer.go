package url

// Referencer contains methods for referencing URLs.
type Referencer struct {
	gateway Gateway
}

// Reference takes a URL, and returns a reference that can be used.
func (r Referencer) Reference(url string) string {
	return r.gateway.FindLatest()
}
