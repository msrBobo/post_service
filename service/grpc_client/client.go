package grpcClient

import (
	"POST_SERVICE/config"
	pbu "POST_SERVICE/genproto/user-service"
	"fmt"
	"google.golang.org/grpc"
)


// IServiceManager ...
type IServiceManager interface {
	UserService() pbu.UserServiceClient
}

type serviceManager struct {
	cfg         config.Config
	userService pbu.UserServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	usercoon, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port: %d",
			cfg.UserServiceHost, cfg.UserServicePort)
	}
	return &serviceManager{
		cfg:         cfg,
		userService: pbu.NewUserServiceClient(usercoon),
	}, nil
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}
