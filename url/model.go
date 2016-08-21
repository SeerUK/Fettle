package url

// URL represents stored URL.
type URL struct {
	ID        uint64 `db:"id,omitempty"`
	Reference string `db:"reference"`
	URL       string `db:"url"`
}
