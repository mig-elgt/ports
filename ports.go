package ports

// Port describes a port entinty
type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

// StorageService describes an interface to perfom CRUD operations.
type StorageService interface {
	// CreateOrUpdate creates a new port record or updates if the port is already exist.
	CreateOrUpdate(port *Port) error

	// Close closes a client connection.
	Close() error
}
