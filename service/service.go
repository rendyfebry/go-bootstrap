package service

type (
	// SomeService interface
	SomeService interface {
		GetIndex() (string, error)
		GetHealth() (string, error)
	}

	// SomeSvc object implementation
	SomeSvc struct {
		DB     interface{}
		Config interface{}
	}
)

// NewService will create a new instance of SomeSvc
func NewService(cfg interface{}) SomeService {
	return &SomeSvc{
		Config: cfg,
	}
}

// GetIndex implementation
func (svc *SomeSvc) GetIndex() (string, error) {
	// Do Something
	return "index", nil
}

// GetHealth implementation
func (svc *SomeSvc) GetHealth() (string, error) {
	// Do Something
	return "health", nil
}
