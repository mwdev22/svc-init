package service

type Service struct {
	// Add your dependencies here, e.g., repositories, caches, etc.
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ExampleMethod() string {
	// implement your business logic here
	return "Hello from Service"
}
