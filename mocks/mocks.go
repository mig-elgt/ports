package mocks

import "ports"

// StorageServiceMock describes a mock implementation of ports.StorageService interface.
type StorageServiceMock struct {
	CreateOrUpdateFn func(port *ports.Port) error
}

// CreateOrUpdate creates a new port record or updates if the port is already exist.
func (s *StorageServiceMock) CreateOrUpdate(port *ports.Port) error {
	return s.CreateOrUpdateFn(port)
}

// Close closes a client connection.
func (s *StorageServiceMock) Close() error {
	return nil
}
