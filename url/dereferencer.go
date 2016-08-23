package url

// Dereferencer contains methods for resolving a reference to a URL to it's underlying value.
type Dereferencer struct{}

// Dereference takes a reference to a URL, and returns the underlying URL.
func (d Dereferencer) Dereference(reference string) URL {
	return ""
}
