package rambler

// Servicer is the interface implemented by the service.
type Servicer interface {
	Initialized() (bool, error)
	Initialize() error
	Available() ([]*Migration, error)
	Applied() ([]*Migration, error)
	Apply(*Migration) error
	Reverse(*Migration) error
}
