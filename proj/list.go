package proj

// List represents project list
type List struct {
	Items []*Project `toml:"projects"`
}
