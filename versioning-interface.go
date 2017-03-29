package versioning

// Version struct
type Version struct {
	major, minor, patch int
	hash                string
	dirty               bool
}
