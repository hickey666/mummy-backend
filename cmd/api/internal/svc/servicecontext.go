package svc

import "mummy-storage/internal/environment"

type ServiceContext struct {
	Env *environment.ServiceContext
}

func NewServiceContext(env *environment.ServiceContext) *ServiceContext {
	return &ServiceContext{
		Env: env,
	}
}
