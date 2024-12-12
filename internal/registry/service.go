package registry

import "context"

type ServiceRegistry struct {
	repo Repository
}

func NewServiceRegistry(repo Repository) *ServiceRegistry {
	return &ServiceRegistry{
		repo: repo,
	}
}

func (sr *ServiceRegistry) RegisterService(ctx context.Context, service *Service) error {
	return sr.repo.RegisterService(ctx, service)
}

func (sr *ServiceRegistry) GetServices(ctx context.Context) ([]*Service, error) {
	return sr.repo.GettServices(ctx)
}

func (sr *ServiceRegistry) DeregisterService(ctx context.Context, id string) error {
	return sr.repo.DeregisterService(ctx, id)
}
