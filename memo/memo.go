package memo

import "ports"

// memo implements ports.StorageService interface
type memo struct {
	db map[string]*ports.Port
}

func New() *memo {
	return &memo{
		db: map[string]*ports.Port{},
	}
}

// CreateOrUpdate creates a new port record or updates if the port is already exist.
func (this *memo) CreateOrUpdate(port *ports.Port) error {
	this.db[port.Unlocs[0]] = port
	return nil
}

// Close closes a client connection.
func (this *memo) Close() error {
	return nil
}
